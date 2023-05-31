package helper

// Response is used for static shape json return
type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
}

func APIResponse(message string, code int, status string, data interface{}) Response {
	jsonResponse := Response{
		Message: message,
		Code:    code,
		Data:    data,
	}

	return jsonResponse
}
