package web

type DefaultResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Data   any    `json:"data"`
}

type ErrorResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
	Errors any    `json:"errors"`
}
