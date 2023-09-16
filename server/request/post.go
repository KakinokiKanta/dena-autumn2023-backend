package request

import (
	"bytes"
	"io"
	"net/http"
)

func Post(url string, contentType string, body []byte) ([]byte, int, error) {
	resp, err := http.Post(url, contentType, bytes.NewBuffer(body))
	if err != nil {
		return nil, resp.StatusCode, err
	}

	defer resp.Body.Close()

	byteArray, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, 500, err
	}

	return byteArray, resp.StatusCode, nil
}
