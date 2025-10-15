package goglib

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HttpRequest(ctx *gin.Context, payload []byte, method string, baseurl string) ([]byte, error, int) {
	url := fmt.Sprintf("%s/saga/order", baseurl)

	// req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBuffer(payload))
	req, err := http.NewRequest(method, url, bytes.NewBuffer(payload))
	if err != nil {

		return nil, err, http.StatusNotFound
	}

	// 헤더 복사 (JWT, Trace-Id, Correlation-Id 등)
	for key, values := range ctx.Request.Header {
		for _, v := range values {
			// logger.Log.Print(2, "header key : %v, value : %v", key, v)
			req.Header.Add(key, v)
		}
	}

	req.Header.Set("Content-Type", "application/json")

	var client *http.Client
	client = &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
		DisableKeepAlives: true,
	}}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Http Request 호출 실패: %w", err), http.StatusNotFound
	}
	defer resp.Body.Close()

	rBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err, resp.StatusCode
	}

	if resp.StatusCode != http.StatusOK {
		return rBody, fmt.Errorf("saga request fail.. [%d]: %s", resp.StatusCode, string(rBody)), resp.StatusCode
	}

	return rBody, nil, resp.StatusCode
}
