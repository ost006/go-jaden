package puller

import (
	"fmt"
	"github.com/ost006/go-jaden/adapter"
	"log"
	"time"
)

/*
블록 조회
 */
type Puller struct {
	Config
	adapter.Adapter

	chTxs		chan []adapter.Transaction
	chReady		chan struct{}

	running	bool
}

type Config struct {
	StartIndex		uint64	`json:"start-index" yaml:"start-index"`
	EndIndex		uint64	`json:"end-index" yaml:"end-index"`

	IntervalSeconds	uint8	`json:"interval-sec"`
	SleepSeconds	uint8	`json:"sleep-sec"`
}

func New(config Config) *Puller {
	o := &Puller{
		Config: config,
		chTxs: make(chan []adapter.Transaction, 0),
		chReady: make(chan struct{}),
	}

	return o
}

func (o *Puller) GetChTxs() chan []adapter.Transaction {
	return o.chTxs
}

func (o *Puller) SetChTxs(chTxs chan []adapter.Transaction) {
	o.chTxs = chTxs
}

func (o *Puller) GetChReady() chan struct{} {
	return o.chReady
}

func (o *Puller) SetChReady(chReady chan struct{}) {
	o.chReady = chReady
}

// run daemon process
func (o *Puller) Run(adt *adapter.Adapter) error {
	if o.running {
		return fmt.Errorf("already running: Puller")
	}

	o.Adapter = *adt

	blockId, err := o.GetLatestBlockId()
	for err != nil {
		log.Printf("error on GetLatestBlockId: %v", err)
		time.Sleep(time.Duration(o.SleepSeconds) * time.Second)
		blockId, err = o.GetLatestBlockId()
	}

	log.Printf("check latest block id: %v / %v", blockId.Index, blockId.Hash)

	nextBlockId := &adapter.BlockId{
		Index: o.StartIndex,
		Hash: blockId.Hash,
	}

	o.running = true
	for o.running {
		time.Sleep(time.Duration(o.IntervalSeconds) * time.Second)
		txs := make([]adapter.Transaction, 0)
		nextBlockId, err = o.GetTransactions(*nextBlockId, &txs)
		<- o.chReady
		o.chTxs <- txs
	}
	return nil
}

// stop daemon process
func (o *Puller) Stop() error {
	if !o.running {
		return fmt.Errorf("already stopped: Puller")
	}

	o.running = false
	return nil
}

