package asset

type Asset struct {
	Asset  uint64 `json:"asset"`
	Symbol string `json:"symbol"`
	Active bool   `json:"active"`
}
