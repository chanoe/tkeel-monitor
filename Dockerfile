FROM golang:1.17 AS builder

COPY . /src
WORKDIR /src

RUN GOPROXY="https://goproxy.cn,direct" make build

FROM alpine:3.13


ENV PLUGIN_ID=tkeel-monitor
COPY --from=builder /src/bin /app

WORKDIR /app

CMD ["./monitor"]
