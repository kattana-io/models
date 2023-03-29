package models

import (
	"math/big"
)

type Block struct {
	Type      string   `json:"type"`
	Network   string   `json:"network"`
	Number    *big.Int `json:"number"`
	Node      string   `json:"node"`
	Notify    bool     `json:"notify"`
	Timestamp uint64   `json:"timestamp"`

	// block trace
	DiscoveredAt       int64 `json:"discovered_at"`
	ReceivedInParserAt int64 `json:"received_in_parser_at"`
	ParsedAt           int64 `json:"parsed_at"`
	ConsumedAt         int64 `json:"consumed_at"`
	LiveSendAt         int64 `json:"live_send_at"`
}
