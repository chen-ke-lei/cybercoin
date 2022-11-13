package net

import (
	"context"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

func GetWithUrl(ctx context.Context, url string) (map[string]interface{}, error) {
	return GetWithParam(ctx, url, nil)

}

func GetWithParam(ctx context.Context, url string, param map[string]string) (map[string]interface{}, error) {
	return Get(ctx, url, param, nil)
}
func Get(ctx context.Context, url string, param map[string]string, header map[string]string) (map[string]interface{}, error) {
	request, err := BuildGetRequest(url, param, header)
	if err != nil {
		return nil, err
	}
	req, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	return ReadResponse(ctx, req)
}

func BuildGetRequest(url string, param map[string]string, header map[string]string) (*http.Request, error) {
	if header == nil {
		header = map[string]string{}
	}
	if _, ok := header["user-agent"]; !ok {
		header["user-agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/107.0.0.0 Safari/537.36"
	}
	url += "?"
	for k, v := range param {
		url += k + "=" + v + "&"
	}
	url = url[0 : len(url)-1]
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		request.Header.Set(k, v)
	}
	return request, nil
}

func ReadResponse(ctx context.Context, resp *http.Response) (map[string]interface{}, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	res := make(map[string]interface{})
	err = jsoniter.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}
	return res, nil
}
