package nats

import (
	"context"
	"time"

	"github.com/nats-io/nats.go/jetstream"
	"go.uber.org/zap"
)

func PublishCreateCheckoutEvent(js jetstream.JetStream, log *zap.Logger) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	js.CreateStream(ctx, jetstream.StreamConfig{
		Name:      "CHECKOUT",
		Subjects:  []string{"checkout.*"},
		Retention: jetstream.WorkQueuePolicy, // acknowledges the message will be removed.
	})
}
