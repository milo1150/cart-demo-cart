package nats

import (
	"log"
	"os"

	"github.com/nats-io/nats.go"
	"go.uber.org/zap"
)

func ConnectNATS() *nats.Conn {
	url := os.Getenv("NATS_URL")
	token := os.Getenv("NATS_TOKEN")

	nc, err := nats.Connect(url, nats.Token(token))
	if err != nil {
		log.Fatalf("Failed to connect NATS server")
	}

	return nc
}

func SubscribeToUserService(nc *nats.Conn, log *zap.Logger) {
	go SubscribeCreateUserEvent(nc, log)
}
