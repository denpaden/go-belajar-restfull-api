package web

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func ToWebResponse(code int, status string, data interface{}) WebResponse {
	return WebResponse{
		Code:   code,
		Status: status,
		Data:   data,
	}
}
func ToWebResponseNoData(code int, status string) WebResponse {
	return WebResponse{
		Code:   code,
		Status: status,
	}
}
