package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	req, err := http.NewRequest("GET", "http://web-1.service.consul:8090/hello", nil)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	c := http.Client{}

	res, err := c.Do(req)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	bRes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Write(bRes)

}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Println(err)
	}
}
