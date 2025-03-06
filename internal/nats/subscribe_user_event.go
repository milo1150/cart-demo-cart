package nats

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

func SubscribeCreateUserEvent(nc *nats.Conn, log *zap.Logger) {
	subject := "user.created"

	nc.Subscribe(subject, func(msg *nats.Msg) {
		HandlerCreateUserEvent(msg, log)
	})
}
