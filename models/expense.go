package models

//Expense test
type Expense struct {
	ID     int    `json:id`
	PaidAt string `json:paidAt`
	Title  string `json:title`
	Amount int    `json:amount`
}
