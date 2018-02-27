package publisher

import (
	"fmt"
	"log"

	nats "github.com/nats-io/nats"

	proto "github.com/ilovelili/sumoproto/services/tickrecorder/proto"
)

const (
	natshost = "nats://nats:4222"
)

// PublishTick sends out single tick to telegraf to save to InfluxDB
func PublishTick(t *proto.Tick) {
	broker := t.Broker
	last := t.Last
	ask := t.Ask
	bid := t.Bid
	pair := t.Pair
	time := t.Time

	nc, err := nats.Connect(natshost)
	if err != nil {
		log.Println(err)
	}
	defer nc.Close()

	msg := fmt.Sprintf("tick,broker=%s,pair=%s ask=%f,bid=%f,last=%f %d", broker, pair, ask, bid, last, time)
	if err := nc.Publish("go.micro.telegraf", []byte(msg)); err != nil {
		log.Println(err)
	}
}
