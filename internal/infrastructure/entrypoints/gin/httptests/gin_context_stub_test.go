package httptests

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const (
	headerKey0     = "header-key-0"
	headerKey1     = "header-key-1"
	headerValue0   = "header_Value_0"
	headerValue1   = "header_Value_0"
	ginParamKey0   = "ginParamKey0"
	ginParamKey1   = "ginParamKey1"
	ginParamValue0 = "ginParamValue0"
	ginParamValue1 = "ginParamValue1"
	queryParam0    = "queryParam"
	queryParam1    = "queryParam1"
	qpValue0       = "qpValue0"
	qpValue1A      = "qpValue1A"
	qpValue1B      = "qpValue1B"
	expectedBody   = `{"id":"123"}`
	expectedScheme = "http"
	expectedHost   = "localhost"
	expectedPort   = "2020"
	expectedPath   = "/shang/tsung"
)

func assertExpectedBody(t *testing.T, ctx *gin.Context) {
	bodyMap := make(map[string]interface{}, 1)
	err := ctx.ShouldBindJSON(&bodyMap)
	assert.NoError(t, err)
	assert.Equal(t, "123", bodyMap["id"])
}

func assertExpectedRequestURL(t *testing.T, ctx *gin.Context) {
	assert.Equal(t, expectedScheme, ctx.Request.URL.Scheme)
	assert.Equal(t, expectedHost, ctx.Request.URL.Hostname())
	assert.Equal(t, expectedPort, ctx.Request.URL.Port())
	assert.Equal(t, expectedPath, ctx.Request.URL.Path)
}

func assertContainsAllQueryParams(t *testing.T, ctx *gin.Context, expectedQueryParams url.Values) {
	for key, values := range expectedQueryParams {
		queryArray, queryArrayFound := ctx.GetQueryArray(key)
		assert.True(t, queryArrayFound)
		assert.Len(t, queryArray, len(values))
		for _, value := range values {
			assert.Contains(t, queryArray, value)
		}
	}
}

func assertContainsAllGinParams(t *testing.T, ctx *gin.Context, expectedGinParams []gin.Param) {
	for _, expectedGinParam := range expectedGinParams {
		paramValue, paramFound := ctx.Params.Get(expectedGinParam.Key)
		assert.True(t, paramFound)
		assert.Equal(t, expectedGinParam.Value, paramValue)
	}
}

func assertContainsAllHeaders(t *testing.T, ctx *gin.Context, expectedHeaders map[string]string) {
	for k, v := range expectedHeaders {
		assert.Equal(t, v, ctx.GetHeader(k))
	}
}

func TestGetGinContextStub(t *testing.T) {
	expectedRequestURL := fmt.Sprintf("%s://%s:%s%s", expectedScheme, expectedHost, expectedPort, expectedPath)
	expectedMethod := http.MethodPost
	expectedQueryParams := url.Values{queryParam0: {qpValue0}, queryParam1: {qpValue1A, qpValue1B}}
	expectedGinParams := []gin.Param{{Key: ginParamKey0, Value: ginParamValue0}, {Key: ginParamKey1, Value: ginParamValue1}}
	expectedHeaders := map[string]string{headerKey0: headerValue0, headerKey1: headerValue1}

	ctx, _ := GetGinContextStub(
		expectedBody, expectedMethod, expectedRequestURL, &expectedQueryParams, &expectedGinParams, expectedHeaders)

	assertExpectedBody(t, ctx)
	assertExpectedRequestURL(t, ctx)
	assert.Equal(t, expectedMethod, ctx.Request.Method)
	assertContainsAllQueryParams(t, ctx, expectedQueryParams)
	assertContainsAllGinParams(t, ctx, expectedGinParams)
	assertContainsAllHeaders(t, ctx, expectedHeaders)
}
