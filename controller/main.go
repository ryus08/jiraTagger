package controller

type RequestBody struct {
	Content   string `json:"content"`
	Key       string `json:"key"`
	Token     string `json:"token"`
	Challenge string `json:"challenge"`
}

type ResponseBody struct {
	Message string

	*RequestBody
}

type Receive struct {
}

func (receive *Receive) Handler(req *RequestBody) *ResponseBody {
	response := &ResponseBody{Message: "success", RequestBody: req}
	return response
}
