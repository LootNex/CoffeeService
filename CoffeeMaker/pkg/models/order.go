package models

type Order struct {
	ID    string `json:"id"`
	Buyer string `json:"buyer"`
	Item  string `json:"item"`
}
