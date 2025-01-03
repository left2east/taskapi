package util

import (
	"bytes"
	"io"
	"net/http"
)

func HttpRequest(url string, method string, headers map[string]string, body []byte) ([]byte, error) {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	responseBody, err1 := io.ReadAll(resp.Body)
	if err1 != nil {
		return nil, err1
	}

	return responseBody, err
}
