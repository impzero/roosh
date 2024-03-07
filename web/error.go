package web

type errorResponse struct {
	ErrorMessage string `json:"message"`
	ErrorCode    string `json:"status_code"`
}
