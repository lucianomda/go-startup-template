package httptests

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
)

func setHeaders(c *gin.Context, headers map[string]string) {
	for k, v := range headers {
		c.Request.Header.Set(k, v)
	}
}

func GetGinContextStub(body, method, requestURL string, queryParams *url.Values,
	params *[]gin.Param, headers map[string]string) (*gin.Context, *httptest.ResponseRecorder) {

	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	if queryParams != nil {
		requestURL = requestURL + "?" + queryParams.Encode()
	}
	if params != nil {
		c.Params = *params
	}
	c.Request, _ = http.NewRequest(method, requestURL, strings.NewReader(body))
	if headers != nil {
		setHeaders(c, headers)
	}
	return c, rr
}
