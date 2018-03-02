package publisher

import (
	"context"
	"strconv"
	"time"

	"github.com/ilovelili/fixdecoder"
	"github.com/micro/go-micro/client"

	marketdata "github.com/ilovelili/sumoproto/services/marketdata/const"
	proto "github.com/ilovelili/sumoproto/services/marketdata/proto"
)

// PublishMarketData publish market data got frm fix engine and decoded by fix decoder
func PublishMarketData(ctx context.Context, fixdata *fixdecoder.DecodedField) error {
	msg := client.NewPublication(marketdata.MarketDataServiceName, &proto.Fix{
		Fieldid:      resolveFieldID(fixdata.FieldID),
		Value:        fixdata.Value,
		Name:         fixdata.Field.Name,
		Type:         fixdata.Field.Type,
		Decodedvalue: fixdata.DecodedValue,
		Time:         time.Now().UnixNano(),
	})

	return client.Publish(ctx, msg)
}

func resolveFieldID(fieldID string) int32 {
	result, _ := strconv.Atoi(fieldID)
	return int32(result)
}
