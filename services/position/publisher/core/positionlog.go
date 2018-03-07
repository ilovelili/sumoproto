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

type positionLog struct {
	title string
}

func (l positionLog) OnIncoming(msg []byte) {
	fix := fd.Decode(string(msg))
	ctx := metadata.NewContext(context.Background(), map[string]string{
		"X-User-Id": "positionpublisher",
	})

	PublishPosition(ctx, fix)
}

func (l positionLog) OnOutgoing(msg []byte) {
	// TBD
}

func (l positionLog) OnEvent(msg string) {
	// Maybe we can add some event watcher here
}

func (l positionLog) OnEventf(format string, v ...interface{}) {
	// TBD
}

type positionLogFactory struct{}

func (positionLogFactory) Create() (quickfix.Log, error) {
	log := positionLog{"GLOBAL"}
	return log, nil
}

func (positionLogFactory) CreateSessionLog(sessionID quickfix.SessionID) (quickfix.Log, error) {
	log := positionLog{sessionID.String()}
	return log, nil
}

// NewPositionLogFactory creates an instance of LogFactory that publishes messages and events to NATS server.
func NewPositionLogFactory() quickfix.LogFactory {
	return positionLogFactory{}
}
