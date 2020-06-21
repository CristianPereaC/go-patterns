package http_utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"strings"
)

func BuildGetHttpRequestContext(url string, headers map[string][]string) (*gin.Context, *httptest.ResponseRecorder) {
	return buildRequestContext("GET", url, headers, "")
}


func BuildPostHttpRequestContext(url string, headers map[string][]string, body string) (*gin.Context, *httptest.ResponseRecorder) {
	return buildRequestContext("POST", url, headers, body)
}

func BuildPutHttpRequestContext(url string, headers map[string][]string, body string) (*gin.Context, *httptest.ResponseRecorder) {
	return buildRequestContext("PUT", url, headers, body)
}

func BuildDeleteHttpRequestContext(url string, headers map[string][]string) (*gin.Context, *httptest.ResponseRecorder) {
	return buildRequestContext("DELETE", url, headers, "")
}

func buildRequestContext(method string, url string, headers map[string][]string, body string) (*gin.Context, *httptest.ResponseRecorder)  {

	responseRecorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(responseRecorder)

	request, _ := http.NewRequest(method, url, strings.NewReader(body))
	request.Header = headers

	context.Request = request

	return context, responseRecorder

}

