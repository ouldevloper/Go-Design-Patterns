package adapters

import "adapter/classes"

func SAOPAdapter(responseBody string, responseStatus int) classes.IResponse {
	response := &classes.Response{}
	headers := []string{"content-Type: application/xml"} // SAOP headers goes here
	response.SetHeaders(headers)
	response.SetResponseBody(responseBody)
	response.SetStatusCode(responseStatus)
	return response
}
