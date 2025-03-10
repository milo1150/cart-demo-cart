package nats

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func SubscribeCreateUserEvent(nc *nats.Conn, log *zap.Logger, db *gorm.DB) {
	subject := "user.created"

	nc.Subscribe(subject, func(msg *nats.Msg) {
		HandlerCreateUserEvent(msg, log, db)
	})
}
