package web

type WebResponse struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type WebResponseAll struct {
	Code         int         `json:"code"`
	Status       bool        `json:"status"`
	Message      string      `json:"message"`
	TotalAllData int         `json:"total_all_data"`
	Data         interface{} `json:"data"`
}
