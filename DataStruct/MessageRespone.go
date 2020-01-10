package datastruct

// MessageRespone is ok
type MessageRespone struct {
	Message string      `json:"message"`
	Body    interface{} `json:"body"`
}
