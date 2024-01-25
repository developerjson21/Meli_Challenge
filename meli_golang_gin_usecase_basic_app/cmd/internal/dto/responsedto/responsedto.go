package responsedto


type ResponseBody struct {
	StatusCode int `json:"status_code" `
	Body any `json:"body" `
}
