package utils

import (
	"net/http"
	"https://github.com/thehaohcm/simple-onedrive/config"
	"https://github.com/thehaohcm/simple-onedrive/enums"
	"https://github.com/thehaohcm/simple-onedrive/models"
	"https://github.com/thehaohcm/simple-onedrive/token"
	"strconv"
	"strings"
)

func HandleHttpRequestForUploading(httpRequest *models.HttpRequest) (*http.Response, error) {
	bodyBytes := httpRequest.Body
	request, _ := http.NewRequest(string(httpRequest.HttpMethod), httpRequest.Url, bodyBytes)
	for _, header := range httpRequest.Headers {
		request.Header.Add(header.Key, header.Value)
	}
	request.Header.Add("Authorization", config.TokenType+" "+config.SavedToken.AccessToken)
	httpClient := &http.Client{}
	return httpClient.Do(request)
}

func SendInitUploadRequest(fileName string, uploadFolderPath string) models.UploadSessionResponse {
	sessionUrL := strings.Replace(config.UploadAPIEndPoint, "{FILE_NAME}", fileName, 1)
	//Create an Upload Session
	var payload = []byte(strings.Replace(config.UploadBodyJson, "{FILE_NAME}", fileName, 1))
	var httpHeaders [](*models.HttpHeader)
	httpHeaders = append(httpHeaders, models.InitHttpHeader("Content-Type", "application/json"))
	httpRequest := models.InitHttpRequest(enums.POST, sessionUrL, payload, httpHeaders)

	var uploadJSONResult models.UploadSessionResponse
	resp, err := HandleHttpRequestForUploading(httpRequest)
	if err != nil {
		panic(err)
	}
	parseData(resp, &uploadJSONResult)
	return uploadJSONResult
}

func SendUploadBlockRequest(url string, blockOrder int, maxBlockRange int, byteBlock []byte) (interface{}, error) {
	sizeByteBlock := len(byteBlock)
	var err error
	var payload = []byte(byteBlock)
	param := "bytes " + strconv.Itoa(blockOrder*fragSize) + "-" + strconv.Itoa(maxBlockRange-1) + "/" + strconv.Itoa(fileSize)
	var httpHeaders [](*models.HttpHeader)
	httpHeaders = append(httpHeaders, models.InitHttpHeader("Content-Length", strconv.Itoa(sizeByteBlock)))
	httpHeaders = append(httpHeaders, models.InitHttpHeader("Content-Range", param))
	httpRequest := models.InitHttpRequest(enums.PUT, url, payload, httpHeaders)

	var jsonResponse interface{}
	if blockOrder < (blockSize - 1) {
		var uploadBlockResponse models.UploadBlockResponse
		resp, err := HandleHttpRequestForUploading(httpRequest)
		if err != nil {
			panic(err)
		}
		parseData(resp, &uploadBlockResponse)
		jsonResponse = uploadBlockResponse
	} else {
		var uploadFinishedResponse models.UploadFinishedResponse
		resp, err := HandleHttpRequestForUploading(httpRequest)
		if err != nil {
			panic(err)
		}
		parseData(resp, &uploadFinishedResponse)
		jsonResponse = uploadFinishedResponse
	}
	return jsonResponse, err
}

func ShareLinkFunc(uploadFinishedResponse *models.UploadFinishedResponse) string {
	token.RefreshToken()
	if uploadFinishedResponse != nil && uploadFinishedResponse.Id != "" {
		//share the item's link
		sharedLinkAPIEndpoint := strings.Replace(config.ShareAPIEndPoint, "{UPLOADED_FILE_ID}", uploadFinishedResponse.Id, 1)
		payload := []byte(config.ShareBodyJson)
		var httpHeaders [](*models.HttpHeader)
		httpHeaders = append(httpHeaders, models.InitHttpHeader("Content-Type", "application/json"))
		httpRequest := models.InitHttpRequest(enums.POST, sharedLinkAPIEndpoint, payload, httpHeaders)

		var sharedLinkResponse models.SharedLinkResponse
		resp, err := HandleHttpRequestForUploading(httpRequest)
		if err != nil {
			panic(err)
		}
		parseData(resp, &sharedLinkResponse)

		return sharedLinkResponse.Link.WebUrl
	}
	return ""
}
