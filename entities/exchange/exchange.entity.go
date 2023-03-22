package exchange

type Exchange struct {
	Exchange uint64 `json:"exchange"`
	Name     string `json:"name"`
	Active   bool   `json:"active"`
}
