package models

import (
	"math/big"
)

type DexeEvents struct {
	Number       *big.Int       `json:"number,omitempty"`
	Timestamp    int64          `json:"timestamp,omitempty"`
	Network      string         `json:"network,omitempty"`
	CustomEvents []*CustomEvent `json:"customEvents,omitempty"`
}
