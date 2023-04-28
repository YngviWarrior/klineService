package rabbitmqstructs

type InputAvgMessageDto struct {
	Asset            uint64  `json:"asset"`
	AssetSymbol      string  `json:"asset_symbol"`
	AssetQuote       uint64  `json:"asset_quote"`
	AssetQuoteSymbol string  `json:"asset_quote_symbol"`
	Exchange         uint64  `json:"exchange"`
	Avg              float64 `json:"avg"`
	Period           string  `json:"period"`
	Roc              float64 `json:"roc"`
}
