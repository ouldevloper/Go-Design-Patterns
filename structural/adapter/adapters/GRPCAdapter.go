package adapters

import "adapter/classes"

func GRPCAdapter(responseBody string, responseStatus int) classes.IResponse {
	response := &classes.Response{}
	headers := []string{"content-Type: stream"} // GRPC headers goes here
	response.SetHeaders(headers)
	response.SetResponseBody(responseBody)
	response.SetStatusCode(responseStatus)
	return response
}
