package marketdata

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path"
	"sync"
	"time"

	"github.com/quickfixgo/quickfix"
)

// simple file storage factory without cache writing feature
type simpleFileStoreFactory struct{}

//NewSimpleFileStoreFactory returns a MessageStoreFactory instance that created simple file MessageStores
func NewSimpleFileStoreFactory() quickfix.MessageStoreFactory { return simpleFileStoreFactory{} }

var (
	filesource = path.Join(Getwd(), "simplefilestore.json") // make it simple
	mu         sync.RWMutex
)

// Create create simplefilestore
func (f simpleFileStoreFactory) Create(sessionID quickfix.SessionID) (quickfix.MessageStore, error) {
	simplefilestore := new(simpleFileStore)
	simplefilestore.Reset()
	return simplefilestore, nil
}

type simpleFileStore struct {
	SenderMsgSeqNum int       `json:"senderMsgSeqNum"`
	TargetMsgSeqNum int       `json:"targetMsgSeqNum"`
	Time            time.Time `json:"creationTime"`
	messageMap      map[int][]byte
}

func (store *simpleFileStore) Reset() (err error) {
	if mystore, err := unmarshalFromFile(); err != nil {
		store.SenderMsgSeqNum = 0
		store.TargetMsgSeqNum = 0
		store.Time = time.Now()
	} else {
		store.SenderMsgSeqNum = mystore.SenderMsgSeqNum
		store.TargetMsgSeqNum = mystore.TargetMsgSeqNum
		store.Time = mystore.Time
	}

	// clear message map always
	store.messageMap = nil
	return nil
}

func (store *simpleFileStore) Refresh() error {
	//nop, nothing to refresh
	return nil
}

func (store *simpleFileStore) Close() error {
	//nop, nothing to close
	return nil
}

func (store *simpleFileStore) CreationTime() time.Time {
	return store.Time
}

func (store *simpleFileStore) GetMessages(beginSeqNum, endSeqNum int) ([][]byte, error) {
	var msgs [][]byte
	for seqNum := beginSeqNum; seqNum <= endSeqNum; seqNum++ {
		if m, ok := store.messageMap[seqNum]; ok {
			msgs = append(msgs, m)
		}
	}

	return msgs, nil
}

func (store *simpleFileStore) SaveMessage(seqNum int, msg []byte) error {
	if store.messageMap == nil {
		store.messageMap = make(map[int][]byte)
	}

	store.messageMap[seqNum] = msg
	marshalToFile(store)
	return nil
}

func (store *simpleFileStore) NextSenderMsgSeqNum() int {
	return store.SenderMsgSeqNum + 1
}

func (store *simpleFileStore) NextTargetMsgSeqNum() int {
	return store.TargetMsgSeqNum + 1
}

func (store *simpleFileStore) IncrNextSenderMsgSeqNum() error {
	store.SenderMsgSeqNum++
	return nil
}

func (store *simpleFileStore) IncrNextTargetMsgSeqNum() error {
	store.TargetMsgSeqNum++
	return nil
}

func (store *simpleFileStore) SetNextSenderMsgSeqNum(nextSeqNum int) error {
	store.SenderMsgSeqNum = nextSeqNum - 1
	return nil
}

func (store *simpleFileStore) SetNextTargetMsgSeqNum(nextSeqNum int) error {
	store.TargetMsgSeqNum = nextSeqNum - 1
	return nil
}

func unmarshalFromFile() (store *simpleFileStore, err error) {
	file, err := os.Open(filesource)
	defer file.Close()
	// maybe file not found
	if err != nil {
		return nil, err
	}

	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(&store)
	return
}

func marshalToFile(store *simpleFileStore) (err error) {
	stream, err := json.Marshal(store)
	if err != nil {
		return err
	}

	mu.Lock()
	defer mu.Unlock()
	return ioutil.WriteFile(filesource, stream, os.ModePerm)
}
