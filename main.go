package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Tweet struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Body string `json:"body"`
}

func main() {
	bytes, err := ioutil.ReadFile("sample.json")
	if err != nil {
		log.Fatal(err)
	}

	var tweets []Tweet
	if err := json.Unmarshal(bytes, &tweets); err != nil {
		log.Fatal(err)
	}

	for _, v := range tweets {
		fmt.Printf("%d: %s\n", v.Id, v.Name)
	}

}
