FROM golang:1.24-alpine3.22 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o tgbot .

FROM alpine:3.22

RUN apk --no-cache add ca-certificates

RUN addgroup -g 1001 -S appgroup && \
    adduser -u 1001 -S appuser -G appgroup

WORKDIR /app

COPY --from=builder /app/tgbot .

RUN chown appuser:appgroup /app/tgbot

USER appuser

ENV TG_BOT_TOKEN=""

CMD ["./tgbot"]