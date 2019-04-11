package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func httpGet(url string) (result string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result = string(body)
	return
}

func httpPost(url string, data url.Values) (result string, err error) {
	resp, err := http.PostForm(url, data)
	if err != nil {
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	result = string(body)
	return
}

func TestCounterAdd(t *testing.T) {
	go main()

	v := url.Values{}
	v.Set("name", "Tom")
	result, err := httpPost("http://localhost:8323/api/furyCount.Add", v)
	if err != nil {
		t.Error(err)
	}

	if result != "" {
		t.Error("Add Fail:" + result)
	}

	result, err = httpPost("http://localhost:8323/api/furyCount.Get", v)
	if err != nil {
		t.Error(err)
	}

	if result != "0" {
		t.Error("Add Fail:" + result)
	}
}
