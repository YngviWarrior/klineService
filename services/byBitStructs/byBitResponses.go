package bybitstructs

type GetServerTimestamp struct {
	RetCode int64  `json:"ret_code"`
	RetMsg  string `json:"ret_msg"`
	ExtCode string `json:"ext_code"`
	ExtInfo string `json:"ext_info"`
	Result  struct {
		TimeNow int64 `json:"serverTime"`
	} `json:"result"`
}

type Kline struct {
	Open_time string `json:"0"`
	Open      string `json:"1"`
	High      string `json:"2"`
	Low       string `json:"3"`
	Close     string `json:"4"`
	Volume    string `json:"5"`
	TurnOver  string `json:"6"`
}

type KlineResponse struct {
	RetCode int64  `json:"retCode"`
	RetMsg  string `json:"retMsg"`
	Result  struct {
		Symbol   string  `json:"symbol"`
		Category string  `json:"category"`
		List     [][]any `json:"list"`
	} `json:"result"`
	RetExtInfo any   `json:"retExtInfo"`
	Time       int64 `json:"time"`
}
