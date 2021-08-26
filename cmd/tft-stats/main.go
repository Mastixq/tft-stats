package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type apiKey struct {
	KeyType string `json:"key_type"`
	Key     string `json:"key"`
}

func main() {

	jsonFile, err := os.Open("cmd/resources/api-key.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var key apiKey
	json.Unmarshal([]byte(byteValue), &key)

	fmt.Println(key)

	var client http.Client
	request_string := fmt.Sprintf("https://eun1.api.riotgames.com/tft/league/v1/challenger?api_key=%s", key.Key)
	fmt.Println(request_string)
	resp, err := client.Get(request_string)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp)
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
	}
}
