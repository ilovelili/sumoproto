package subscriber

import (
	"log"

	proto "github.com/ilovelili/sumoproto/services/marketdata/proto"
	"github.com/ilovelili/sumoproto/services/marketdata/subscriber/publisher"
	"github.com/micro/go-micro/client"
	"golang.org/x/net/context"
)

// MarketData is a struct that contains market data handlers
type MarketData struct {
	Client client.Client
}

// Handle will respond to relevant messages on the topic it is registered
func (m *MarketData) Handle(ctx context.Context, msg *proto.Fix) error {
	log.Print("Handler received market data. Publishing...")
	publisher.PublishMarketData(msg)
	return nil
}
