package core

import (
	"context"
	"strconv"
	"time"

	"github.com/ilovelili/fixdecoder"
	"github.com/micro/go-micro/client"

	proto "github.com/ilovelili/sumoproto/services/marketdata/proto"
	marketdata "github.com/ilovelili/sumoproto/services/marketdata/shared"
)

// PublishMarketData publish market data got frm fix engine and decoded by fix decoder
func PublishMarketData(ctx context.Context, fixdata fixdecoder.DecodedFields) error {
	msg := client.NewPublication(marketdata.MarketDataServiceName, &proto.Fix{
		Msg:  fixdata.String(),
		Time: time.Now().UnixNano(),
	})

	return client.Publish(ctx, msg)
}

func resolveFieldID(fieldID string) int32 {
	result, _ := strconv.Atoi(fieldID)
	return int32(result)
}
