package core

import (
	"errors"
	"fmt"
	"time"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	fix44mdr "github.com/quickfixgo/fix44/marketdatarequest"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/config"
)

// getSettings get session settings by key
func getSettings(key string, setting *quickfix.Settings) (value string, err error) {
	for _, sessionSettings := range setting.SessionSettings() {
		value, err = sessionSettings.Setting(key)
		if err != nil {
			return
		}
	}

	return
}

func getSenderCompID(settings *quickfix.Settings) field.SenderCompIDField {
	if senderCompID, err := getSettings(config.SenderCompID, settings); err != nil {
		panic(err)
	} else {
		if senderCompID == "" {
			panic(errors.New("Oops no SenderCompID found"))
		}

		return field.NewSenderCompID(senderCompID)
	}
}

func getTargetCompID(settings *quickfix.Settings) field.TargetCompIDField {
	if targetCompID, err := getSettings(config.TargetCompID, settings); err != nil {
		panic(err)
	} else {
		if targetCompID == "" {
			panic(errors.New("Oops no TargetCompID found"))
		}

		return field.NewTargetCompID(targetCompID)
	}
}

type header interface {
	Set(f quickfix.FieldWriter) *quickfix.FieldMap
}

func queryHeader(h header, settings *quickfix.Settings) {
	h.Set(getSenderCompID(settings))
	h.Set(getTargetCompID(settings))
}

func queryMarketDataRequest(settings *quickfix.Settings) fix44mdr.MarketDataRequest {
	request := fix44mdr.New(
		field.NewMDReqID("EUR/USD"+time.Now().Format("20060102150405")),
		field.NewSubscriptionRequestType(enum.SubscriptionRequestType_SNAPSHOT_PLUS_UPDATES),
		// set 0 as tag 264 (full book)
		field.NewMarketDepth(0),
	)

	// set 0 & 1 as tag 269
	entryTypes := fix44mdr.NewNoMDEntryTypesRepeatingGroup()
	entryTypes.Add().SetMDEntryType(enum.MDEntryType_BID)
	request.SetNoMDEntryTypes(entryTypes)
	entryTypes.Add().SetMDEntryType(enum.MDEntryType_OFFER)
	request.SetNoMDEntryTypes(entryTypes)

	relatedSym := fix44mdr.NewNoRelatedSymRepeatingGroup()
	relatedSym.Add().SetSymbol("EUR/USD")
	request.SetNoRelatedSym(relatedSym)

	// set 1 as tag 265
	request.SetMDUpdateType(enum.MDUpdateType_INCREMENTAL_REFRESH)

	// set N as tag 266
	request.SetAggregatedBook(false)

	// set sending time
	request.SetSendingTime(time.Now())

	// currenex customized tag
	SetNewAttributedPrices(request, true)

	queryHeader(request.Header, settings)
	return request
}

// QueryMarketDataRequest Market data req
func QueryMarketDataRequest(settings *quickfix.Settings) (err error) {
	beginstring, err := getSettings(config.BeginString, settings)
	if err != nil {
		return err
	}

	fmt.Println("Protocol Version: ", beginstring)
	req := queryMarketDataRequest(settings)

SendingWithRetry:
	for i := 1; i <= 10; i++ {
		if err = quickfix.Send(req); err == nil {
			break SendingWithRetry
		}
		// set a interval
		time.Sleep(3 * time.Second)
		fmt.Println("Sending request... Retry count: ", i)
	}

	return err
}
