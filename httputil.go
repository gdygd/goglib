package goglib

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

var httpClient = &http.Client{Transport: &http.Transport{
	TLSClientConfig: &tls.Config{
		InsecureSkipVerify: true,
	},
	DisableKeepAlives: false,
}}

func HttpRequest(ctx *gin.Context, payload []byte, method string, url string) (int, []byte, error) {

	// req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(payload))
	req, err := http.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(payload))
	if err != nil {

		return http.StatusNotFound, nil, err
	}

	// 헤더 복사 (JWT, Trace-Id, Correlation-Id 등)
	for key, values := range ctx.Request.Header {
		for _, v := range values {
			// logger.Log.Print(2, "header key : %v, value : %v", key, v)
			req.Header.Add(key, v)
		}
	}

	req.Header.Set("Content-Type", "application/json")

	// var client *http.Client
	// client = &http.Client{Transport: &http.Transport{
	// 	TLSClientConfig: &tls.Config{
	// 		InsecureSkipVerify: true,
	// 	},
	// 	DisableKeepAlives: false,
	// }}

	resp, err := httpClient.Do(req)
	if err != nil {
		return http.StatusNotFound, nil, fmt.Errorf("Http Request 호출 실패: %w", err)
	}
	defer resp.Body.Close()

	rBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, rBody, fmt.Errorf("saga request fail.. [%d]: %s", resp.StatusCode, string(rBody))
	}

	return resp.StatusCode, rBody, nil
}
