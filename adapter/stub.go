package adapter

import "fmt"

type Stub struct {

}

var a Adapter = Stub{}

func (o Stub) GetLatestBlockId() (*BlockId, error) {
	return &BlockId{"test", 0}, nil
}

func (o Stub) GetBlock(id BlockId) (*Block, error) {
	return &Block{}, nil
}

func (o Stub) GetTransaction(id string) (*Transaction, error) {
	return &Transaction{}, nil
}

func (o Stub) GetAccount(id AccountId) (*AccountDetail, error) {
	return &AccountDetail{}, nil
}

func (o Stub) GetTransactions(id BlockId, txs *[]Transaction) (*BlockId, error) {
	*txs = append(*txs, Transaction{Id: fmt.Sprintf("%s-%6d", id.Hash, id.Index)})

	return &BlockId{id.Hash, id.Index+1}, nil
}
