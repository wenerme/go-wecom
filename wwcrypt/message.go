package wwcrypt

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"errors"
)

type ReceiveMessage struct {
	ToUsername string          `xml:"ToUserName"`
	Encrypt    string          `xml:"Encrypt"`
	AgentID    string          `xml:"AgentID"`
	Content    *ReceiveContent `xml:"-"`
}

type ReceiveContent struct {
	ReceiverID string
	Content    []byte
}

func (rc *ReceiveContent) MarshalBinary() (data []byte, err error) {
	buf := &bytes.Buffer{}
	r := make([]byte, 16)
	_, err = rand.Read(r)
	if err != nil {
		return nil, err
	}
	_, _ = buf.Write(r)

	binary.BigEndian.PutUint32(r, uint32(len(rc.Content)))
	_, _ = buf.Write(r[:4])
	_, _ = buf.Write(rc.Content)
	_, _ = buf.Write([]byte(rc.ReceiverID))
	return buf.Bytes(), nil
}

func (rc *ReceiveContent) UnmarshalBinary(data []byte) (err error) {
	var size uint32
	if len(data) < 20 {
		goto UNDERFLOW
	}
	// random
	data = data[16:]
	size = binary.BigEndian.Uint32(data)
	// size
	data = data[4:]
	if len(data) < int(size) {
		goto UNDERFLOW
	}
	rc.Content = data[:size]
	rc.ReceiverID = string(data[size:])
	return
UNDERFLOW:
	err = errors.New("invalid message data")
	return
}

func (rc *ReceiveContent) VerifyReceiverID(r string) bool {
	return rc.ReceiverID == r
}
