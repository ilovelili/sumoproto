package subscriber

import (
	"log"

	proto "github.com/ilovelili/sumoproto/services/tickrecorder/proto"
	"github.com/ilovelili/sumoproto/services/tickrecorder/publisher"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

// Tick is a struct that contains Tick handlers
type Tick struct {
	Client client.Client
}

// Handle will respond to relevant messages on the topic it is registered
func (e *Tick) Handle(ctx context.Context, msg *proto.Tick) error {
	log.Print("Handler received tick data. Publishing...")
	publisher.PublishTick(msg)
	return nil
}
