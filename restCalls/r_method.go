package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	fmt.Println("Get Method")
	resget, err := http.Get("https://httpbin.org/ip")
	if err != nil {
		fmt.Println("Error in Get method")
	} else {
		rg, _ := ioutil.ReadAll(resget.Body)
		type Get struct {
			Origin string `json:"origin"`
		}
		g1 := &Get{}
		json.Unmarshal([]byte(rg), g1)
		fmt.Println(*g1)

	}

	jsonData := map[string]string{"firstname": "Nic", "lastname": "Raboy"}
	jsonValue, _ := json.Marshal(jsonData)
	respost, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		fmt.Println("Error in Post method")
	} else {

		rp, _ := ioutil.ReadAll(respost.Body)
		type Post struct {
			Args struct {
			} `json:"args"`
			Data  string `json:"data"`
			Files struct {
			} `json:"files"`
			Form struct {
			} `json:"form"`
			Headers struct {
				AcceptEncoding string `json:"Accept-Encoding"`
				ContentLength  string `json:"Content-Length"`
				ContentType    string `json:"Content-Type"`
				Host           string `json:"Host"`
				UserAgent      string `json:"User-Agent"`
				XAmznTraceID   string `json:"X-Amzn-Trace-Id"`
			} `json:"headers"`
			JSON struct {
				Firstname string `json:"firstname"`
				Lastname  string `json:"lastname"`
			} `json:"json"`
			Origin string `json:"origin"`
			URL    string `json:"url"`
		}

		p1 := &Post{}
		json.Unmarshal([]byte(rp), p1)
		fmt.Println(p1.Data)

	}
}
