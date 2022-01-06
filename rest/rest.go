package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

func GetRequest(ctx context.Context, url string, result any, f ...func(req *http.Request)) error {
	return request(ctx, http.MethodGet, url, nil, result, f...)
}

func PostRequest(ctx context.Context, url string, reqBody any, result any, f ...func(req *http.Request)) error {
	return request(ctx, http.MethodPost, url, reqBody, result, f...)
}

func PutRequest(ctx context.Context, url string, reqBody io.Reader, result any, f ...func(req *http.Request)) error {
	return request(ctx, http.MethodPut, url, reqBody, result, f...)
}

func DeleteRequest(ctx context.Context, url string, reqBody io.Reader, result any, f ...func(req *http.Request)) error {
	return request(ctx, http.MethodDelete, url, reqBody, result, f...)
}

func request(ctx context.Context, method string, url string, reqBody any, result any, f ...func(req *http.Request)) error {
	b, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	r := bytes.NewReader(b)
	req, err := http.NewRequestWithContext(ctx, method, url, r)
	for _, fun := range f {
		fun(req)
	}
	if err != nil {
		return err
	}
	req.Header.Set("accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != 200 {
		return fmt.Errorf("status code: %v", res.StatusCode)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	if result == nil {
		return nil
	}
	return json.Unmarshal(body, result)
}
