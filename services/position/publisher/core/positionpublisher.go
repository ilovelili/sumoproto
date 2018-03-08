package core

import (
	"context"
	"time"

	"github.com/ilovelili/fixdecoder"

	"github.com/micro/go-micro/client"

	proto "github.com/ilovelili/sumoproto/services/position/proto"
	position "github.com/ilovelili/sumoproto/services/position/shared"

	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/nats"
	_ "github.com/micro/go-plugins/transport/nats"
)

// PublishPosition publish market data got frm fix engine and decoded by fix decoder
func PublishPosition(ctx context.Context, fix fixdecoder.DecodedFields) error {
	msg := client.NewPublication(position.PositionServiceName, &proto.Fix{
		Msg:  fix.String(),
		Time: time.Now().UnixNano(),
	})

	return client.Publish(ctx, msg)
}
