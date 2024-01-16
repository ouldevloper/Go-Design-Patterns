package adapters

import "adapter/classes"

func QRAPHQLAdapter(responseBody string, responseStatus int) classes.IResponse {

	response := &classes.Response{}
	headers := []string{"content-Type: stream"} // QRAPHQLAdapter headers goes here
	response.SetHeaders(headers)
	response.SetResponseBody(responseBody)
	response.SetStatusCode(responseStatus)
	return response
}
