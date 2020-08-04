package adapter

type Adapter interface {
	GetLatestBlockId() (*BlockId, error)
	GetBlock(id BlockId) (*Block, error)
	GetTransaction(id string) (*Transaction, error)
	GetAccount(id AccountId) (*AccountDetail, error)

	// @param id: BlockId
	// @param txs: result of Transactions reference
	// @return *BlockId: next BlockId
	GetTransactions(id BlockId, txs *[]Transaction) (*BlockId, error)
}