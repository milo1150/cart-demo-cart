package nats

import (
	"cart-service/internal/repositories"
	"fmt"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"gorm.io/gorm"

	cartpkg "github.com/milo1150/cart-demo-pkg/pkg"
)

func HandlerCreateUserEvent(msg *nats.Msg, log *zap.Logger, db *gorm.DB) {
	data, err := cartpkg.BytesToUint(msg.Data)
	if err != nil {
		log.Error("Failed to parse CreateUserEvent data", zap.Error(err))
	}

	log.Info("HandlerCreateUserEvent",
		zap.Any("UserID:", data),
	)

	// Create new Cart and binding UserID
	if err := repositories.CreateCart(db, data); err != nil {
		log.Error(fmt.Sprintf("Failed to create Cart with UserID: %v", data))
	}
}
