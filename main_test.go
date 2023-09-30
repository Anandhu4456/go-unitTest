package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDemoHandler(t *testing.T) {
	server:=httptest.NewServer(http.HandlerFunc(DemoHandler))
	response,err:=http.Get(server.URL)
	if err!=nil{
		t.Error(err)
	}
	if response.StatusCode != http.StatusOK{
		t.Errorf("expected 200 got %d",response.StatusCode)
	}
	defer response.Body.Close()
	expected:="hey"
	bytes,err:= ioutil.ReadAll(response.Body)
	if err !=nil{
		t.Error(err)
	}
	if string(bytes) != expected{
		t.Errorf("expected %s and got %s",expected,bytes)
	}
}
