package main

import (
	"encoding/json"
	"fmt"
	"io"
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
	fmt.Println("hello")

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	log.Println("Listing for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
