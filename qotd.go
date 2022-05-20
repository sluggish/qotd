package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/json"
)

type Response struct {
	Quote  string `json:"quote"`
	Author string `json:"author"`
}

func main() {
	res, err := http.Get("https://free-quotes-api.herokuapp.com/")
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)

	var result Response
    if err := json.Unmarshal(body, &result); err != nil {
        fmt.Println(err)
    }
	fmt.Println(result.Quote + " by ", result.Author)
}
