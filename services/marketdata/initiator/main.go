package main

import (
	"fmt"
	"os"
	"os/signal"
	"path"

	"github.com/ilovelili/sumoproto/services/marketdata/initiator/core"
	marketdata "github.com/ilovelili/sumoproto/services/marketdata/shared"
	"github.com/quickfixgo/enum"
	"github.com/quickfixgo/field"
	"github.com/quickfixgo/quickfix"
)

var (
	config      *marketdata.Config
	appSettings *quickfix.Settings
)

//TradeClient implements the quickfix.Application interface
type TradeClient struct{}

// OnCreate is called when quickfix creates a new session.
// A session comes into and remains in existence for the life of the application.
// Sessions exist whether or not a counter party is connected to it.
// As soon as a session is created, you can begin sending messages to it.
// If no one is logged on, the messages will be sent at the time a connection is established with the counterparty.
func (c TradeClient) OnCreate(sessionID quickfix.SessionID) {
	fmt.Println("On session create")
	return
}

// OnLogon notifies you when a valid logon has been established with a counter party.
// This is called when a connection has been established and the FIX logon process has completed with both parties exchanging valid logon messages.
func (c TradeClient) OnLogon(sessionID quickfix.SessionID) {
	fmt.Println("Logging on")
	if err := core.QueryMarketDataRequest(appSettings); err != nil {
		panic(err)
	}

	return
}

// OnLogout notifies you when an FIX session is no longer online.
// This could happen during a normal logout exchange or because of a forced termination or a loss of network connection.
func (c TradeClient) OnLogout(sessionID quickfix.SessionID) {
	fmt.Println("Logging out")
	return
}

// FromAdmin notifies you when an administrative message is sent from a counterparty to your FIX engine.
// This can be useful for doing extra validation on logon messages such as for checking passwords.
func (c TradeClient) FromAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	return
}

// ToAdmin provides you with a peak at the administrative messages that are being sent from your FIX engine to the counter party.
// This is normally not useful for an application however it is provided for any logging you may wish to do.
// Notice that the Message is not const. This allows you to add fields before an administrative message is sent out.
func (c TradeClient) ToAdmin(msg *quickfix.Message, sessionID quickfix.SessionID) {
	if msg.IsMsgTypeOf(string(enum.MsgType_LOGON)) {
		// this is important since username / password in plaintext is required by currenex
		msg.Body.Set(field.NewUsername(config.UserName))
		msg.Body.Set(field.NewPassword(config.Password))

		// manually set reset seq number flag to true
		msg.Body.Set(field.NewResetSeqNumFlag(true))

	} else if msg.IsMsgTypeOf(string(enum.MsgType_LOGOUT)) {
		// manually set reset seq number flag to true
		msg.Body.Set(field.NewResetSeqNumFlag(true))
	}

	return
}

// ToApp notifies you of application messages that you are being sent to a counterparty.
// Notice that the Message is not const. This allows you to add fields before an application message before it is sent out.
func (c TradeClient) ToApp(msg *quickfix.Message, sessionID quickfix.SessionID) (err error) {
	fmt.Println()
	fmt.Println("========================= Message To Currenex =========================")
	fmt.Println(msg.String())
	fmt.Println()

	return
}

// FromApp is one of the core entry points for your FIX application.
// Every application level request will come through here.
// If, for example, your application is a sell-side OMS, this is where you will get your new order requests.
// If you were a buy side, you would get your execution reports here.
func (c TradeClient) FromApp(msg *quickfix.Message, sessionID quickfix.SessionID) (reject quickfix.MessageRejectError) {
	fmt.Println("========================= Message From Currenex =========================")
	fmt.Println()
	fmt.Println(msg.String())
	fmt.Println()

	// Todo: handle snapshot here
	if msg.IsMsgTypeOf(string(enum.MsgType_MARKET_DATA_SNAPSHOT_FULL_REFRESH)) {

		fmt.Println("Received a full market data snapshot")
	} else if msg.IsMsgTypeOf(string(enum.MsgType_MARKET_DATA_INCREMENTAL_REFRESH)) {
		fmt.Println("Received an incremental market data update")
	}

	return
}

func init() {
	if conf, err := marketdata.GetConfig(); err != nil {
		// failed to get config.
		panic(err)
	} else {
		config = conf
	}
}

func main() {
	cwd, _ := os.Getwd()
	cfgFileName := path.Join(cwd, "config", "currenex.cfg")

	cfg, err := os.Open(cfgFileName)
	defer cfg.Close()

	if err != nil {
		fmt.Printf("Error opening %v, %v\n", cfgFileName, err)
		return
	}

	settings, err := quickfix.ParseSettings(cfg)
	if err != nil {
		fmt.Println("Error reading cfg,", err)
		return
	}
	appSettings = settings

	app := TradeClient{}

	logFactory, err := quickfix.NewFileLogFactory(appSettings)
	if err != nil {
		fmt.Println("Error creating file log factory,", err)
		return
	}

	storeFactory := marketdata.NewSimpleFileStoreFactory()

	initiator, err := quickfix.NewInitiator(app, storeFactory, appSettings, logFactory)
	if err != nil {
		fmt.Printf("Unable to create Initiator: %s\n", err)
		return
	}

	initiator.Start()

	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt, os.Kill)
	<-interrupt

	initiator.Stop()

	fmt.Println("Bye!")
}
