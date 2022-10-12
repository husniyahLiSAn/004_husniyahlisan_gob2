package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	data := map[string]interface{}{
		"title":  "Airel",
		"body":   "Jordan",
		"userId": 1,
	}

	requestJson, err := json.Marshal(data)

	client := &http.Client{}
	if err != nil {
		log.Fatalln(err)
	}

	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", bytes.NewBuffer(requestJson))

	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		log.Fatalln(err)
	}

	res, err := client.Do(req)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()

	sb := string(body)
	fmt.Println(sb)
}
