package classes

type IResponse interface {
	SetStatusCode(status int)
	GetStatusCode() int
	SetHeaders(headers []string)
	GetHeaders() []string
	SetResponseBody(body string)
	GetResponseBody() string
}
