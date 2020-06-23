package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, err := http.Get("https://api.covid19india.org/data.json")
	if err != nil {
		fmt.Println(err)
	} else {
		rp, _ := ioutil.ReadAll(res.Body)

		type covid struct {
			CasesTimeSeries []struct {
				Dailyconfirmed string `json:"dailyconfirmed"`
				Dailydeceased  string `json:"dailydeceased"`
				Dailyrecovered string `json:"dailyrecovered"`
				Date           string `json:"date"`
				Totalconfirmed string `json:"totalconfirmed"`
				Totaldeceased  string `json:"totaldeceased"`
				Totalrecovered string `json:"totalrecovered"`
			} `json:"cases_time_series"`
			Statewise []struct {
				Active          string `json:"active"`
				Confirmed       string `json:"confirmed"`
				Deaths          string `json:"deaths"`
				Deltaconfirmed  string `json:"deltaconfirmed"`
				Deltadeaths     string `json:"deltadeaths"`
				Deltarecovered  string `json:"deltarecovered"`
				Lastupdatedtime string `json:"lastupdatedtime"`
				Recovered       string `json:"recovered"`
				State           string `json:"state"`
				Statecode       string `json:"statecode"`
				Statenotes      string `json:"statenotes"`
			} `json:"statewise"`
			Tested []struct {
				Individualstestedperconfirmedcase string `json:"individualstestedperconfirmedcase"`
				Positivecasesfromsamplesreported  string `json:"positivecasesfromsamplesreported"`
				Samplereportedtoday               string `json:"samplereportedtoday"`
				Source                            string `json:"source"`
				Testpositivityrate                string `json:"testpositivityrate"`
				Testsconductedbyprivatelabs       string `json:"testsconductedbyprivatelabs"`
				Testsperconfirmedcase             string `json:"testsperconfirmedcase"`
				Testspermillion                   string `json:"testspermillion"`
				Totalindividualstested            string `json:"totalindividualstested"`
				Totalpositivecases                string `json:"totalpositivecases"`
				Totalsamplestested                string `json:"totalsamplestested"`
				Updatetimestamp                   string `json:"updatetimestamp"`
			} `json:"tested"`
		}
		c := &covid{}
		json.Unmarshal([]byte(rp), &c)
		for _, v := range c.Statewise {
			fmt.Println("Active: ", v.Active, "\nConfirmed: ", v.Confirmed, "\nDeaths: ", v.Deaths, "\nstate: ", v.State)
			fmt.Println("-----------------------------------------------------------------------------------------------")
		}
	}
}
