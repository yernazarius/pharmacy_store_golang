package nats

import (
	"github.com/nats-io/nats.go"
)

func NewNatsClient(url string) (*nats.Conn, error) {
	return nats.Connect(url)
}
