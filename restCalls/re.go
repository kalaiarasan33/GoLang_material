package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	type rate struct {
		Base  string `json:"base"`
		Rates struct {
			GBP float64 `json:"GBP"`
			HKD float64 `json:"HKD"`
			IDR float64 `json:"IDR"`
			ILS float64 `json:"ILS"`
			DKK float64 `json:"DKK"`
			INR float64 `json:"INR"`
			CHF float64 `json:"CHF"`
			MXN float64 `json:"MXN"`
			CZK float64 `json:"CZK"`
			SGD float64 `json:"SGD"`
			THB float64 `json:"THB"`
			HRK float64 `json:"HRK"`
			MYR float64 `json:"MYR"`
			NOK float64 `json:"NOK"`
			CNY float64 `json:"CNY"`
			BGN float64 `json:"BGN"`
			PHP float64 `json:"PHP"`
			SEK float64 `json:"SEK"`
			PLN float64 `json:"PLN"`
			ZAR float64 `json:"ZAR"`
			CAD float64 `json:"CAD"`
			ISK float64 `json:"ISK"`
			BRL float64 `json:"BRL"`
			RON float64 `json:"RON"`
			NZD float64 `json:"NZD"`
			TRY float64 `json:"TRY"`
			JPY float64 `json:"JPY"`
			RUB float64 `json:"RUB"`
			KRW float64 `json:"KRW"`
			USD float64 `json:"USD"`
			HUF float64 `json:"HUF"`
			AUD float64 `json:"AUD"`
		} `json:"rates"`
		Date string `json:"date"`
	}
	start := time.Now()

	resp, _ := http.Get("https://api.ratesapi.io/api/latest")

	rb, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(rb))
	rt := &rate{}
	json.Unmarshal([]byte(rb), rt)
	rates := *rt

	fmt.Println(rates.Rates.GBP)

	t := time.Now()
	elapsed := t.Sub(start)

	fmt.Println(elapsed)
}
