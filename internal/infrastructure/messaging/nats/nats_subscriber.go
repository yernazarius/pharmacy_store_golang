package nats

import (
	"log"

	"github.com/nats-io/nats.go"
)

func (c *NATSClient) Subscribe(subject string, handler nats.MsgHandler) {
	_, err := c.Conn.Subscribe(subject, handler)
	if err != nil {
		log.Printf("Error subscribing to %s: %v", subject, err)
	}
}
