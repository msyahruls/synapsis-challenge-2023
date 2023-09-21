package domain

type DataWS struct {
	Method string   `json:"method"`
	Params ParamsWS `json:"params"`
}

type ParamsWS struct {
	Channel string      `json:"channel"`
	Data    interface{} `json:"data"`
}
