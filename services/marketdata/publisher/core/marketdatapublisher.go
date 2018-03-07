package core

import (
	"context"
	"time"

	"github.com/ilovelili/fixdecoder"

	"github.com/micro/go-micro/client"

	proto "github.com/ilovelili/sumoproto/services/marketdata/proto"
	marketdata "github.com/ilovelili/sumoproto/services/marketdata/shared"
)

// PublishMarketData publish market data got frm fix engine and decoded by fix decoder. Subscriber/main.go handles service discovery
func PublishMarketData(ctx context.Context, fix fixdecoder.DecodedFields) error {
	msg := client.NewPublication(marketdata.MarketDataServiceName, &proto.Fix{
		Msg:  fix.String(),
		Time: time.Now().UnixNano(),
	})

	return client.Publish(ctx, msg)
}
