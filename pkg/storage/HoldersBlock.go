package models

type HoldersBlock struct {
	Block     uint64    `json:"block,omitempty"`
	Timestamp int64     `json:"timestamp,omitempty"`
	Chain     string    `json:"chain,omitempty"`
	Notify    bool      `json:"notify,omitempty"`
	Holders   []*Holder `json:"holders,omitempty"`
}
