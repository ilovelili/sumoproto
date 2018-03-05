package publisher

import (
	"fmt"
	"log"
	"strings"

	proto "github.com/ilovelili/sumoproto/services/marketdata/proto"
	nats "github.com/nats-io/nats"
)

const (
	natshost = "nats://nats:4222"
)

// PublishMarketData sends out single Fix data to telegraf to save to InfluxDB
func PublishMarketData(f *proto.Fix) {
	nc, err := nats.Connect(natshost)
	if err != nil {
		log.Println(err)
	}
	defer nc.Close()

	sanitizedMsg := sanitizeFIXMessage(f.Msg)
	msg := fmt.Sprintf(`marketdata,type=rawmsg, message="%s" %d`, sanitizedMsg, f.Time)
	// msg := fmt.Sprintf(`marketdata,tag=%d name="%s",value="%s",decodedvalue="%s" %d`, tag, name, value, decodedvalue, time)
	if err := nc.Publish("invast.sumo.srv.telegraf", []byte(msg)); err != nil {
		log.Println(err)
	}
}

// sanitizeFIXMessage escape the json to see if it works
func sanitizeFIXMessage(decodeFixMsg string) string {
	replacer := strings.NewReplacer("\"", "", ":", " is ", "\n", "  ", "{", "|", "}", "|", ",", "")
	result := replacer.Replace(decodeFixMsg)
	fmt.Println("FIX message sanitized: ", result)
	return result
}
