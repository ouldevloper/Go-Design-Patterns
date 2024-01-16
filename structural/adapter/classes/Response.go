package classes

type Response struct {
	headers      []string
	responseBody string
	statusCode   int
}

func (obj *Response) SetStatusCode(status int) {
	obj.statusCode = status
}

func (obj *Response) GetStatusCode() int {
	return obj.statusCode
}

func (obj *Response) SetHeaders(headers []string) {
	obj.headers = headers
}

func (obj *Response) GetHeaders() []string {
	return obj.headers
}

func (obj *Response) SetResponseBody(body string) {
	obj.responseBody = body
}

func (obj *Response) GetResponseBody() string {
	return obj.responseBody
}
