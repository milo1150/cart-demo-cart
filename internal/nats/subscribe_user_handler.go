package nats

import (
	"strconv"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

func HandlerCreateUserEvent(msg *nats.Msg, log *zap.Logger) {
	data, err := strconv.Atoi(string(msg.Data))
	if err != nil {
		log.Error("Failed to parse msg", zap.Error(err))
	}

	log.Info("HandlerCreateUserEvent",
		zap.Any("data:", data),
	)

	// TODO: create Cart and bind User ID
}
