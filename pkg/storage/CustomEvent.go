package models

import "time"

type CustomEvent struct {
	BlockNumber uint64    `json:"blocknumber"`
	Date        time.Time `json:"date"`
	Tx          string    `json:"tx"`
	Contract    string    `json:"contract"`
	Chain       string    `json:"chain"`
	Wallet      string    `json:"wallet"`
	Name        string    `json:"name"`
	Data        string    `json:"data"`
}
