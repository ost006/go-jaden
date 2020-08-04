package puller

import (
	"encoding/json"
	"github.com/ost006/go-jaden/adapter"
	"log"
)

type Stub struct {
	chTxs		chan []adapter.Transaction
	chReady		chan struct{}
}

func NewStub() *Stub {
	return &Stub{}
}

func (o *Stub) GetChTxs() chan []adapter.Transaction {
	return o.chTxs
}

func (o *Stub) SetChTxs(chTxs chan []adapter.Transaction) {
	o.chTxs = chTxs
}

func (o *Stub) GetChReady() chan struct{} {
	return o.chReady
}

func (o *Stub) SetChReady(chReady chan struct{}) {
	o.chReady = chReady
}


func (o *Stub) Run() {
	for {
		o.chReady <- struct{}{}
		txs := <- o.chTxs

		j, err := json.Marshal(txs)
		if err != nil {
			log.Printf("error on json.Marshal: %v", string(j))
			return
		}

		log.Printf("%v", string(j))
	}
}