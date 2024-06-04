package nats

import (
	"log"
)

func (c *NATSClient) Publish(subject string, msg []byte) {
	if err := c.Conn.Publish(subject, msg); err != nil {
		log.Printf("Error publishing message to %s: %v", subject, err)
	}
}
