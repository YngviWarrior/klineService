package bybitstructs

type GetKlinesParams struct {
	Test      bool
	Symbol    string `json:"symbol"`
	Interval  string `json:"interval"`
	Limit     int64  `json:"limit"`
	StartTime int64  `json:"startTime"`
	EndTime   int64  `json:"endTime"`
}
