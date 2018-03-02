package log

import (
	"context"
	"fmt"
	"sync"

	"github.com/ilovelili/fixdecoder"
	"github.com/ilovelili/sumoproto/services/marketdata/pub/publisher"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	"github.com/quickfixgo/quickfix"

	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/nats"
	_ "github.com/micro/go-plugins/transport/nats"
)

var (
	fd *fixdecoder.FixDecoder
	wg sync.WaitGroup
)

func init() {
	fd = fixdecoder.NewFixDecoder()
	cmd.Init()
}

type marketDataLog struct {
	title string
}

func (p marketDataLog) OnIncoming(msg []byte) {
	fix := fd.Decode(string(msg))

	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "marketdatapublisher",
	})

	// Todo: parse necessay data from raw and publish it (like tick). So actually this loop is not needed
	for _, fixdataline := range fix {
		// filter out incremental market data
		if fixdataline.Decoded {
			wg.Add(1)
			fmt.Printf("Publishing %v\n", fixdataline)
			go publisher.PublishMarketData(ctx, fixdataline)
		}
	}

	wg.Wait()
}

func (p marketDataLog) OnOutgoing(msg []byte) {
	// TBD
}

func (p marketDataLog) OnEvent(msg string) {
	// TBD
}

func (p marketDataLog) OnEventf(format string, v ...interface{}) {
	// TBD
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
