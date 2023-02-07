package models

type Holder struct {
	Token  string `json:"token"`
	From   string `json:"from"`
	To     string `json:"to"`
	Tx     string `json:"tx"`
	Amount int64  `json:"amount,omitempty"`
}
