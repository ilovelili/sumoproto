package core

import (
	"context"

	"github.com/ilovelili/fixdecoder"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	"github.com/quickfixgo/quickfix"

	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/nats"
	_ "github.com/micro/go-plugins/transport/nats"
)

var (
	fd *fixdecoder.FixDecoder
)

func init() {
	fd = fixdecoder.NewFixDecoder()
	cmd.Init()
}

type marketDataLog struct {
	title string
}

func (l marketDataLog) OnIncoming(msg []byte) {
	fix := fd.Decode(string(msg))
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "marketdatapublisher",
	})

	PublishMarketData(ctx, fix)
}

func (l marketDataLog) OnOutgoing(msg []byte) {
}

func (l marketDataLog) OnEvent(msg string) {
	// Maybe we can add some event watcher here
}

func (l marketDataLog) OnEventf(format string, v ...interface{}) {
}

type marketDataLogFactory struct{}

func (marketDataLogFactory) Create() (quickfix.Log, error) {
	log := marketDataLog{"GLOBAL"}
	return log, nil
}

func (marketDataLogFactory) CreateSessionLog(sessionID quickfix.SessionID) (quickfix.Log, error) {
	log := marketDataLog{sessionID.String()}
	return log, nil
}

// NewMarketDataLogFactory creates an instance of LogFactory that publishes messages and events to NATS server.
func NewMarketDataLogFactory() quickfix.LogFactory {
	return marketDataLogFactory{}
}
