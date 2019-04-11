package main

import (
	"io/ioutil"
	"net/http"
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

func TestAbs(t *testing.T) {
	go main()

	result, err := httpGet("http://localhost:8323/")
	if err != nil {
		t.Error(err)
	}

	if result != "Hello, World!" {
		t.Error("Invali body:" + result)
	}
}
