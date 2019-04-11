package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func httpGet(url string) (result string) {
	resp, err := http.Get(url)
	if err != nil {
		return err.Error()
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}
	return string(body)
}

func httpPost(url string, data url.Values) (result string) {
	resp, err := http.PostForm(url, data)
	if err != nil {
		return err.Error()
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err.Error()
	}

	return string(body)
}

func TestCounterAdd(t *testing.T) {
	go main()

	v := url.Values{}
	v.Set("name", "Tom")
	result := httpPost("http://localhost:8323/api/furyCount.Add", v)
	if result != "" {
		t.Error("Add Fail:" + result)
	}

	result = httpPost("http://localhost:8323/api/furyCount.Add", nil)
	if result != "Must provide counter name" {
		t.Error("Add Fail:" + result)
	}

	result = httpPost("http://localhost:8323/api/furyCount.Add", v)
	if result != "counter exist" {
		t.Error("Add Fail:" + result)
	}

	result = httpPost("http://localhost:8323/api/furyCount.Get", v)
	if result != "0" {
		t.Error("Get Fail:" + result)
	}

	v.Set("name", "Kim")
	result = httpPost("http://localhost:8323/api/furyCount.Get", v)
	if result != "counter not found" {
		t.Error("Get Fail fi:" + result)
	}

	result = httpPost("http://localhost:8323/api/furyCount.Get", nil)
	if result != "Must provide counter name" {
		t.Error("Get Fail fi:" + result)
	}

	result = httpGet("http://localhost:8323/")
	if result != "Hello, World!" {
		t.Error("index:" + result)
	}
}
