package okk_source

type HistoryMarkPriceCandle struct {
	Code int        `json:"code"`
	Msg  string     `json:"msg"`
	Data [][]string `json:"data"`
}
