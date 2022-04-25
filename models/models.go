package models

type AccountInput struct {
	AccId int `json:"accId"`
}

type AccountDetails struct {
	AccId          int `json:"accId"`
	DocumentNumber int `json:"documentNumber"`
}

type Transactions struct {
	AccountId       int     `json:"accountId"`
	OperationTypeId int     `json:"operationTypeId"`
	Amount          float64 `json:"amount"`
}

type AccountResponse struct {
	AccountId int `json:"accountId"`
}

type TransactionResponse struct {
	TransactionId int `json:"transactionId"`
}
