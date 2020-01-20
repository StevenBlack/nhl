package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func scoring() {
	url := urls["scoring"]

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Print(err)
	}

	req.Header.Set("User-Agent", "nhl-stats-api")

	res, getErr := client.Do(req)
	if getErr != nil {
		fmt.Print(getErr)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Print(err)
	}

	// var stats nhlStandings
	// err = json.Unmarshal(data, &stats)
	// if err != nil {
	// 	fmt.Println("error:", err)
	// }

	// create and load standings
	fmt.Println(data)

}
