package adapters

import "adapter/classes"

func ViewAdapter(responseBody string, responseStatus int) classes.IResponse {
	response := &classes.Response{}
	headers := []string{"content-Type: application/html"} // View headers goes here
	response.SetHeaders(headers)
	response.SetResponseBody(responseBody)
	response.SetStatusCode(responseStatus)
	return response
}
