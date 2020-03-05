package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHttpRequest(t *testing.T){
	handler:= func(w http.ResponseWriter, r * http.Request) {
		io.WriteString(w, "{ \"status\": \"good\" }")
	}

	req:=httptest.NewRequest("GET","https://www.google.com",nil)
	w:=httptest.NewRecorder()
	handler(w,req)

	resp:=w.Result()
	body,_:=ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	if 200!=resp.StatusCode {
		t.Fatal("Status code not OK")
	}
}
