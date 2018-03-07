package publisher

import (
	"fmt"
	"log"
	"strings"

	proto "github.com/ilovelili/sumoproto/services/position/proto"
	nats "github.com/nats-io/nats"
)

const (
	natshost = "nats://nats:4222"
)

// PublishPosition sends out single Fix data to telegraf to save to InfluxDB
func PublishPosition(f *proto.Fix) {
	nc, err := nats.Connect(natshost)
	if err != nil {
		log.Println(err)
	}
	defer nc.Close()

	santized := sanitizeFIXMessage(f.Msg)
	msg := fmt.Sprintf(`position message="%s" %d`, santized, f.Time)
	if err := nc.Publish("invast.sumo.srv.telegraf", []byte(msg)); err != nil {
		log.Println(err)
	}
}

// sanitizeFIXMessage escape the json to see if it works
func sanitizeFIXMessage(decodeFixMsg string) string {
	replacer := strings.NewReplacer("\"", "", ",", "|")
	result := replacer.Replace(decodeFixMsg)
	fmt.Println("FIX message sanitized: ", result)
	return result
}
