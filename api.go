package main

import (
		"io/ioutil"
		"encoding/json"
		"fmt"
    "net/http"
)

func Get(endpoint string, apiKey string) map[string]interface{} {
	baseURL := "http://api.abyiss.com"

	url := baseURL + endpoint

	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println(err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("apiKey", apiKey)

	res, err := httpClient.Do(req)
	if err != nil {
		fmt.Println(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(body), &result)
	return result
}
