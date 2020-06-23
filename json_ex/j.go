package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// small key name wont , if you want use `json:name`
	type Person struct {
		Name    string   `json:"name"`
		Age     int      `json:"age"`
		Depart  string   `json:"depart"`
		Friends []string `json:"friends"`
	}

	p := &Person{Name: "kalai", Age: 27, Depart: "EEE", Friends: []string{"lokesh", "gowtham"}}
	data, _ := json.Marshal(p)
	fmt.Println(string(data))

	dd := `{"Name": "kalai", "Age": 27, "Depart": "EEE", "Friends":["lokesh","gowtham"]}`
	p2 := &Person{}
	json.Unmarshal([]byte(dd), p2)
	fmt.Println(*p2)

	dm := map[string]string{"Name": "Kalai", "Age": "27", "Depart": "EEE"}
	fmt.Println(dm)
	da, _ := json.Marshal(dm)
	fmt.Println(string(da))

}
