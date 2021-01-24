package models

import (
	"bytes"
	"https://github.com/thehaohcm/simple-onedrive/enums"
)

type HttpRequest struct {
	HttpMethod enums.HttpRequestMethod
	Url        string
	Body       *bytes.Buffer
	Headers    []*HttpHeader
}

func InitHttpRequest(httpMethod enums.HttpRequestMethod, url string, body []byte, headers []*HttpHeader) *HttpRequest {
	var httpRequest HttpRequest
	httpRequest.HttpMethod = httpMethod
	httpRequest.Url = url
	//old version
	httpRequest.Body = bytes.NewBuffer(body)
	httpRequest.Headers = headers
	return &httpRequest
}
