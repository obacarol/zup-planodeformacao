package models

import (
	"gopkg.in/guregu/null.v4"
)

type Account struct {
	ID              int     `json:"id"`
	Name            string  `json:"name"`
	Cpf             int     `json:"cpf"`
	Creation_date   string  `json:"creation_date"`
	Balance_account float64 `json:"balance_account"`
}

type Records struct {
	ID                *int64      `json:"id"`
	ID_account_from   int64       `json:"id_account_from"`
	ID_account_to     null.Int    `json:"id_account_to"`
	Transaction_date  string      `json:"transaction_date"`
	Transaction_type  Transaction `json:"transaction_type"`
	Transaction_value float64     `json:"transaction_value"`
}

type Transaction string

const (
	Pay_In     = "Pay in"
	Withdrawal = "Withdrawal"
	Transfer   = "Transfer"
)
