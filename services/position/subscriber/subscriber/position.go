package subscriber

import (
	"log"

	proto "github.com/ilovelili/sumoproto/services/position/proto"
	"github.com/ilovelili/sumoproto/services/position/subscriber/publisher"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

// Position is a struct that contains position data handlers
type Position struct {
	Client client.Client
}

// Handle will respond to relevant messages on the topic it is registered
func (m *Position) Handle(ctx context.Context, msg *proto.Fix) error {
	log.Print("Handler received position. Publishing...")
	publisher.PublishPosition(msg)
	return nil
}
