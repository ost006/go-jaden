package adapter

import "math/big"

// 블록
type Block struct {
	Id				BlockId			`json:"id"`
	PrevId			BlockId			`json:"prev-id"`
	Transactions	[]Transaction	`json:"transactions"`
}

// 블록 Identifier
type BlockId struct {
	Hash	string	`json:"hash"`
	Index	uint64	`json:"index"`
}

// 트랜잭션
type Transaction struct {
	Id			string		`json:"id"`
	Status		string		`json:"status"`
	Operations	[]Operation	`json:"operations"`
}

// 트랜잭션 상세 오퍼레이션
type Operation struct {
	Id					OperationId		`json:"id"`
	RelatedOperationId	*OperationId	`json:"related-id"`
	Type				string			`json:"type"`
	Account				Account			`json:"account"`
	Amount				Amount			`json:"amount"`
	Metadata			string			`json:"meta"`
}

// 오퍼레이션 Identifier
type OperationId struct {
	TransactionId	string	`json:"id"`
	Index			uint8	`json:"index"`
}

// 계정 주소
type Account struct {
	Id			AccountId	`json:"id"`
	SubAddress	string		`json:"sub-address"`
}

// 계정 Identifier
type AccountId	struct {
	Address		string	`json:"address"`
}

// 금액 정보
type Amount struct {
	Value		big.Int		`json:"value"`
	Currency	Currency	`json:"currency"`
}

// 화폐 정보
type Currency struct {
	Symbol		string	`json:"symbol"`
	Decimals	uint8	`json:"decimals"`
	Unit		string	`json:"unit"`
}

// 계정 상세 정보
type AccountDetail struct {
	Id			AccountId	`json:"id"`
	Balance		[]Amount	`json:"balances"`
	Metadata	string		`json:"meta"`
}