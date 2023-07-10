package model

type BlockDump struct {
	BlockNumber int64    `json:"block_number"`
	Dump        []string `json:"dump"`
}
