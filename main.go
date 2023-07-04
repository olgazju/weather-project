package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Dataset struct {
	UID          string  `json:"uid"`
	MindDate     string  `json:"mindate"`
	MaxDate      string  `json:"maxdate"`
	Name         string  `json:"name"`
	Datacoverage float32 `json:"datacoverage"`
	Id           string  `json:"id"`
}

type Datasets struct {
	Results []Dataset `json:"results"`
}

func main() {
	req, err := http.NewRequest("GET", "https://www.ncdc.noaa.gov/cdo-web/api/v2/datasets", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Replace "Your-Token" with your actual token.
	req.Header.Set("token", "Your-Token")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	var datasets Datasets
	err = json.Unmarshal(body, &datasets)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, dataset := range datasets.Results {
		fmt.Println(dataset.Name)
	}
}
