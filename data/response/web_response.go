package response

type Response struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"`
	Message string      `json:"error"`
	Data    interface{} `json:"data,omitempty"`
}
