package adapters

import "adapter/classes"

func RESTAdapter(responseBody string, responseStatus int) classes.IResponse {
	response := &classes.Response{}
	headers := []string{"content-Type: application/json"} // REST headers goes here
	response.SetHeaders(headers)
	response.SetResponseBody(responseBody)
	response.SetStatusCode(responseStatus)
	return response
}
