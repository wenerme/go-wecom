package wecom

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"sync"
	"sync/atomic"
	"time"

	"github.com/gorilla/websocket"
)

const (
	DefaultWSEndpoint   = "wss://openws.work.weixin.qq.com"
	DefaultPingInterval = 30 * time.Second
	DefaultWriteTimeout = 10 * time.Second
	DefaultReadTimeout  = 60 * time.Second
)

// BotWSClient is a WebSocket client for WeChat Work intelligent bot (智能机器人长连接).
//
// It handles:
//   - Connection establishment and subscription
//   - Heartbeat (ping/pong every 30s)
//   - Automatic reconnection
//   - Message/event callback dispatching
//   - Stream response support
//
// See https://developer.work.weixin.qq.com/document/path/101463
type BotWSClient struct {
	// BotID is the bot identifier
	BotID string
	// Secret is the bot secret
	Secret string
	// Endpoint override (default: wss://openws.work.weixin.qq.com)
	Endpoint string
	// PingInterval override (default: 30s)
	PingInterval time.Duration

	// OnMessage is called when a user message is received (aibot_msg_callback)
	OnMessage func(ctx context.Context, msg *BotMessage) error
	// OnEvent is called when an event is received (aibot_event_callback)
	OnEvent func(ctx context.Context, event *BotEvent) error
	// OnError is called on connection errors
	OnError func(err error)
	// Logger override (defaults to slog.Default())
	Logger *slog.Logger

	conn    *websocket.Conn
	mu      sync.Mutex
	reqSeq  atomic.Int64
	pending sync.Map // req_id -> chan *WSResponse
	done    chan struct{}
}

// WSMessage is the base WebSocket message format
type WSMessage struct {
	Cmd     string          `json:"cmd"`
	Headers WSHeaders       `json:"headers,omitempty"`
	Body    json.RawMessage `json:"body,omitempty"`
	ErrCode int             `json:"errcode,omitempty"`
	ErrMsg  string          `json:"errmsg,omitempty"`
}

// WSHeaders contains message metadata
type WSHeaders struct {
	ReqID string `json:"req_id,omitempty"`
}

// WSResponse is a generic response
type WSResponse struct {
	Headers WSHeaders `json:"headers"`
	ErrCode int       `json:"errcode"`
	ErrMsg  string    `json:"errmsg"`
}

// BotMessage represents a message callback from a user
type BotMessage struct {
	MsgID    string          `json:"msgid"`
	AiBotID  string          `json:"aibotid"`
	ChatID   string          `json:"chatid,omitempty"`
	ChatType string          `json:"chattype"` // "single" or "group"
	From     BotUser         `json:"from"`
	MsgType  string          `json:"msgtype"` // text, image, voice, file, video, mixed
	Text     *BotTextContent `json:"text,omitempty"`
	Image    *BotMediaRef    `json:"image,omitempty"`
	Voice    *BotVoiceRef    `json:"voice,omitempty"`
	File     *BotMediaRef    `json:"file,omitempty"`
	Video    *BotMediaRef    `json:"video,omitempty"`
	Mixed    *BotMixedMsg    `json:"mixed,omitempty"`

	// ReqID for replying
	ReqID string `json:"-"`
	// client reference for reply helpers
	client *BotWSClient
}

// BotEvent represents an event callback
type BotEvent struct {
	MsgID      string         `json:"msgid"`
	CreateTime int64          `json:"create_time"`
	AiBotID    string         `json:"aibotid"`
	ChatID     string         `json:"chatid,omitempty"`
	ChatType   string         `json:"chattype,omitempty"`
	From       BotUser        `json:"from"`
	MsgType    string         `json:"msgtype"` // "event"
	Event      BotEventDetail `json:"event"`

	ReqID  string `json:"-"`
	client *BotWSClient
}

type BotEventDetail struct {
	EventType string `json:"eventtype"` // enter_chat, template_card_event, feedback_event, disconnected_event
}

type BotUser struct {
	UserID string `json:"userid"`
}

type BotTextContent struct {
	Content string `json:"content"`
}

type BotMediaRef struct {
	MediaID string `json:"media_id"`
}

type BotVoiceRef struct {
	MediaID     string `json:"media_id"`
	TextContent string `json:"text_content,omitempty"` // voice-to-text
}

type BotMixedMsg struct {
	MsgItem []json.RawMessage `json:"msg_item"`
}

// BotStreamReply is used for streaming responses (typing effect)
type BotStreamReply struct {
	ID      string `json:"id"`
	Finish  bool   `json:"finish"`
	Content string `json:"content"`
}

// Connect establishes the WebSocket connection and starts the message loop.
// Blocks until ctx is cancelled or an unrecoverable error occurs.
func (c *BotWSClient) Connect(ctx context.Context) error {
	if c.Endpoint == "" {
		c.Endpoint = DefaultWSEndpoint
	}
	if c.PingInterval == 0 {
		c.PingInterval = DefaultPingInterval
	}
	if c.Logger == nil {
		c.Logger = slog.Default()
	}

	c.done = make(chan struct{})

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		err := c.connectOnce(ctx)
		if err != nil {
			if ctx.Err() != nil {
				return ctx.Err()
			}
			c.Logger.Warn("connection lost, reconnecting in 5s...", "error", err)
			if c.OnError != nil {
				c.OnError(err)
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(5 * time.Second):
			}
		}
	}
}

func (c *BotWSClient) connectOnce(ctx context.Context) error {
	c.Logger.Info("connecting...")

	conn, _, err := websocket.DefaultDialer.DialContext(ctx, c.Endpoint, nil)
	if err != nil {
		return fmt.Errorf("dial: %w", err)
	}

	c.mu.Lock()
	c.conn = conn
	c.mu.Unlock()

	defer func() {
		c.mu.Lock()
		c.conn = nil
		c.mu.Unlock()
		_ = conn.Close()
	}()

	// Subscribe
	if err := c.subscribe(ctx); err != nil {
		return fmt.Errorf("subscribe: %w", err)
	}

	c.Logger.Info("subscribed, listening for messages...")

	// Start ping loop
	pingCtx, pingCancel := context.WithCancel(ctx)
	defer pingCancel()
	go c.pingLoop(pingCtx)

	// Read loop
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		_, data, err := conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("read: %w", err)
		}

		var msg WSMessage
		if err := json.Unmarshal(data, &msg); err != nil {
			c.Logger.Warn("failed to parse message", "error", err)
			continue
		}

		go c.handleMessage(ctx, &msg)
	}
}

func (c *BotWSClient) subscribe(ctx context.Context) error {
	reqID := c.nextReqID()

	body, _ := json.Marshal(map[string]string{
		"bot_id": c.BotID,
		"secret": c.Secret,
	})

	msg := WSMessage{
		Cmd:     "aibot_subscribe",
		Headers: WSHeaders{ReqID: reqID},
		Body:    body,
	}

	resp, err := c.sendAndWait(ctx, &msg, 10*time.Second)
	if err != nil {
		return err
	}
	if resp.ErrCode != 0 {
		return fmt.Errorf("subscribe failed: %d %s", resp.ErrCode, resp.ErrMsg)
	}
	return nil
}

func (c *BotWSClient) pingLoop(ctx context.Context) {
	ticker := time.NewTicker(c.PingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			reqID := c.nextReqID()
			msg := WSMessage{
				Cmd:     "ping",
				Headers: WSHeaders{ReqID: reqID},
			}
			if _, err := c.sendAndWait(ctx, &msg, 10*time.Second); err != nil {
				c.Logger.Warn("ping failed", "error", err)
			}
		}
	}
}

func (c *BotWSClient) handleMessage(ctx context.Context, msg *WSMessage) {
	// Check if it's a response to a pending request
	if msg.Cmd == "" && msg.Headers.ReqID != "" {
		if ch, ok := c.pending.LoadAndDelete(msg.Headers.ReqID); ok {
			resp := &WSResponse{
				Headers: msg.Headers,
				ErrCode: msg.ErrCode,
				ErrMsg:  msg.ErrMsg,
			}
			ch.(chan *WSResponse) <- resp
			return
		}
	}

	switch msg.Cmd {
	case "aibot_msg_callback":
		if c.OnMessage == nil {
			return
		}
		var body BotMessage
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			c.Logger.Warn("failed to parse message callback", "error", err)
			return
		}
		body.ReqID = msg.Headers.ReqID
		body.client = c
		if err := c.OnMessage(ctx, &body); err != nil {
			c.Logger.Warn("message handler error", "error", err)
		}

	case "aibot_event_callback":
		if c.OnEvent == nil {
			return
		}
		var body BotEvent
		if err := json.Unmarshal(msg.Body, &body); err != nil {
			c.Logger.Warn("failed to parse event callback", "error", err)
			return
		}
		body.ReqID = msg.Headers.ReqID
		body.client = c
		if err := c.OnEvent(ctx, &body); err != nil {
			c.Logger.Warn("event handler error", "error", err)
		}
	}
}

// ─── Reply helpers ────────────────────────────────────────────────────

// ReplyText sends a text reply to the message
func (m *BotMessage) ReplyText(_ context.Context, content string) error {
	body, _ := json.Marshal(map[string]any{
		"msgtype": "text",
		"text":    map[string]string{"content": content},
	})
	return m.client.send(&WSMessage{
		Cmd:     "aibot_respond_msg",
		Headers: WSHeaders{ReqID: m.ReqID},
		Body:    body,
	})
}

// ReplyMarkdown sends a markdown reply
func (m *BotMessage) ReplyMarkdown(_ context.Context, content string) error {
	body, _ := json.Marshal(map[string]any{
		"msgtype":  "markdown",
		"markdown": map[string]string{"content": content},
	})
	return m.client.send(&WSMessage{
		Cmd:     "aibot_respond_msg",
		Headers: WSHeaders{ReqID: m.ReqID},
		Body:    body,
	})
}

// ReplyStream sends a streaming response chunk. Use same streamID for continuation.
// Set finish=true for the final chunk.
func (m *BotMessage) ReplyStream(_ context.Context, streamID string, content string, finish bool) error {
	body, _ := json.Marshal(map[string]any{
		"msgtype": "stream",
		"stream": BotStreamReply{
			ID:      streamID,
			Finish:  finish,
			Content: content,
		},
	})
	return m.client.send(&WSMessage{
		Cmd:     "aibot_respond_msg",
		Headers: WSHeaders{ReqID: m.ReqID},
		Body:    body,
	})
}

// ReplyWelcome sends a welcome message (only for enter_chat events, within 5s)
func (e *BotEvent) ReplyWelcome(_ context.Context, content string) error {
	body, _ := json.Marshal(map[string]any{
		"msgtype": "text",
		"text":    map[string]string{"content": content},
	})
	return e.client.send(&WSMessage{
		Cmd:     "aibot_respond_welcome_msg",
		Headers: WSHeaders{ReqID: e.ReqID},
		Body:    body,
	})
}

// SendMessage sends an active message (not in response to a callback)
func (c *BotWSClient) SendMessage(ctx context.Context, chatID string, chatType int, msgType string, content any) error {
	body, _ := json.Marshal(map[string]any{
		"chatid":    chatID,
		"chat_type": chatType,
		"msgtype":   msgType,
		msgType:     content,
	})
	_, err := c.sendAndWait(ctx, &WSMessage{
		Cmd:     "aibot_send_msg",
		Headers: WSHeaders{ReqID: c.nextReqID()},
		Body:    body,
	}, 10*time.Second)
	return err
}

// ─── Internal ─────────────────────────────────────────────────────────

func (c *BotWSClient) nextReqID() string {
	return fmt.Sprintf("req_%d", c.reqSeq.Add(1))
}

func (c *BotWSClient) send(msg *WSMessage) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.conn == nil {
		return fmt.Errorf("not connected")
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	return c.conn.WriteMessage(websocket.TextMessage, data)
}

func (c *BotWSClient) sendAndWait(ctx context.Context, msg *WSMessage, timeout time.Duration) (*WSResponse, error) {
	ch := make(chan *WSResponse, 1)
	c.pending.Store(msg.Headers.ReqID, ch)
	defer c.pending.Delete(msg.Headers.ReqID)

	if err := c.send(msg); err != nil {
		return nil, err
	}

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-time.After(timeout):
		return nil, fmt.Errorf("timeout waiting for response to %s", msg.Cmd)
	case resp := <-ch:
		if resp.ErrCode != 0 {
			return resp, fmt.Errorf("%d: %s", resp.ErrCode, resp.ErrMsg)
		}
		return resp, nil
	}
}

// Close gracefully closes the WebSocket connection
func (c *BotWSClient) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}
