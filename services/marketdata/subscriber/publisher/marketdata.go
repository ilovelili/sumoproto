package publisher

import (
	"fmt"
	"log"

	proto "github.com/ilovelili/sumoproto/services/marketdata/proto"
	nats "github.com/nats-io/nats"
)

const (
	natshost = "nats://nats:4222"
)

// PublishMarketData sends out single Fix data to telegraf to save to InfluxDB
func PublishMarketData(f *proto.Fix) {
	tag := f.Fieldid
	value := f.Value
	name := f.Name
	decodedvalue := resolveDecodedValue(f.Decodedvalue)
	time := f.Time

	nc, err := nats.Connect(natshost)
	if err != nil {
		log.Println(err)
	}
	defer nc.Close()

	msg := fmt.Sprintf(`marketdata,tag=%d name="%s",value="%s",decodedvalue="%s" %d`, tag, name, value, decodedvalue, time)
	if err := nc.Publish("invast.sumo.srv.telegraf", []byte(msg)); err != nil {
		log.Println(err)
	}
}

func resolveDecodedValue(decodedvalue string) string {
	if decodedvalue == "" {
		return "NA"
	}

	return decodedvalue
}
