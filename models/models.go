package models

type AccountInput struct {
	AccId int `json:"accId"`
}

type AccountDetails struct {
	AccId          int `json:"accId"`
	DocumentNumber int `json:"documentNumber"`
}
