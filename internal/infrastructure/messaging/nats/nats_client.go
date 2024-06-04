package nats

import (
	"log"

	"github.com/nats-io/nats.go"
	"github.com/yernazarius/pharmacy_store_golang/configs"
)

type NATSClient struct {
	Conn *nats.Conn
}

func NewNATSClient(config *configs.Config) *NATSClient {
	nc, err := nats.Connect(config.NATSUrl)
	if err != nil {
		log.Fatalf("Error connecting to NATS: %v", err)
	}

	return &NATSClient{Conn: nc}
}

func (c *NATSClient) Close() {
	c.Conn.Close()
}
