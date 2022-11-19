FROM golang:1.19-bullseye AS builder

WORKDIR /src
ENV GOPROXY=https://goproxy.io
COPY go.* .
RUN go mod download
COPY . .

ENV CGO_ENABLED=1
RUN --mount=type=cache,target=/root/go/ \
    --mount=type=cache,target=/root/.cache/go-build \
    go build -o /app/bin/ -trimpath -ldflags "-s -w" github.com/wenerme/go-wecom/cmd/wwfinance-poller \
    && mkdir -p /app/libs/ \
    && cp WeWorkFinanceSDK/libs/libWeWorkFinanceSdk_C.so /app/libs/ \
    && cp WeWorkFinanceSDK/libs/md5sum.txt /app/libs/ \
    && cp WeWorkFinanceSDK/libs/version.txt /app/libs/

FROM debian:bullseye-slim
WORKDIR /app
COPY --from=builder /app .
ENV LD_LIBRARY_PATH=/app/libs
CMD [ "/app/bin/wwfinance-poller" ]
