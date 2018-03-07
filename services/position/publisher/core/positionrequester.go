package core

import (
	"errors"
	"fmt"
	"time"

	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	fix44rfp "github.com/quickfixgo/fix44/requestforpositions"
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

func getSenderCompID(settings *quickfix.Settings) string {
	if senderCompID, err := getSettings(config.SenderCompID, settings); err != nil {
		panic(err)
	} else {
		if senderCompID == "" {
			panic(errors.New("Oops no SenderCompID found"))
		}

		return senderCompID
	}
}

func getTargetCompID(settings *quickfix.Settings) string {
	if targetCompID, err := getSettings(config.TargetCompID, settings); err != nil {
		panic(err)
	} else {
		if targetCompID == "" {
			panic(errors.New("Oops no TargetCompID found"))
		}

		return targetCompID
	}
}

type header interface {
	Set(f quickfix.FieldWriter) *quickfix.FieldMap
}

func queryHeader(h header, settings *quickfix.Settings) {
	h.Set(field.NewSenderCompID(getSenderCompID(settings)))
	h.Set(field.NewTargetCompID(getTargetCompID(settings)))
}

func queryPositionRequest(settings *quickfix.Settings) fix44rfp.RequestForPositions {
	request := fix44rfp.New(
		field.NewPosReqID("POS"+time.Now().Format("20060102150405")),
		field.NewPosReqType(enum.PosReqType_POSITIONS),
		field.NewAccount(getSenderCompID(settings)),
		field.NewAccountType(enum.AccountType_ACCOUNT_IS_CARRIED_ON_CUSTOMER_SIDE_OF_THE_BOOKS),
		field.NewClearingBusinessDate(time.Now().Format("20060102")),
		field.NewTransactTime(time.Now()),
	)

	noPartyIDs := fix44rfp.NewNoPartyIDsRepeatingGroup()
	noPartyIDs.Add().SetPartyID(getSenderCompID(settings))
	noPartyIDs.Add().SetPartyRole(enum.PartyRole_CLIENT_ID)
	request.SetNoPartyIDs(noPartyIDs)
	// this line is important. Otherwise CNX will send IncorrectGroupNo error resposne
	request.Set(field.NewNoPartyIDs(1))

	request.SetSubscriptionRequestType(enum.SubscriptionRequestType_SNAPSHOT_PLUS_UPDATES)
	// set sending time
	request.SetSendingTime(time.Now())

	queryHeader(request.Header, settings)
	return request
}

// QueryPositionRequest position data req
func QueryPositionRequest(settings *quickfix.Settings) (err error) {
	beginstring, err := getSettings(config.BeginString, settings)
	if err != nil {
		return err
	}

	fmt.Println("Protocol Version: ", beginstring)
	req := queryPositionRequest(settings)

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
