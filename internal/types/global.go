package types

import (
	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AppState struct {
	DB   *gorm.DB
	NATS *nats.Conn
	Log  *zap.Logger
}
