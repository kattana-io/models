package models

type CustomEvent struct {
	Tx       string `json:"tx,omitempty"`
	Contract string `json:"contract,omitempty"`
	Wallet   string `json:"wallet,omitempty"`
	Entity   string `json:"entity,omitempty"`
	Klass    string `json:"klass,omitempty"`
	Name     string `json:"name,omitempty"`
	Data     string `json:"data,omitempty"`
}
