package service

import "time"

type DataCalculateRevenue struct {
	// Transactions []Transaction
	TransactionData   []Transaction
	Config Config
}
type Config struct{
	SkipCorporateAction string
	FinancialYear       string
	StartFinancialMonth string
	EndFinancialMonth   string
}

type Transaction struct {
	Market    string
	Direction string
	Cost      float32
	Price     float32
	Quantity  int
	Date      time.Time
	Activity  string
	UnitPrice float32
}