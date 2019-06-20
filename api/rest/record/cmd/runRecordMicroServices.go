package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"github.com/justinhuang159/ustart_tutorial/api/rest/record"
	//prof "github.com/sea350/ustart_tutorial/customer"
	//"github.com/sea350/ustart_tutorial/customer/storage"
	//elasticstore "github.com/sea350/ustart_tutorial/customer/storage/elastic"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	log.Println("Dialing up...")

	var config customer.Config

	//Importing configuration from json
	file, err := os.Open("config.json")
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		panic(err)
	}

	//Generating api object
	restAPI, err := customer.New(&config)
	if err != nil {
		panic(err)
	}

	//Assigning the handler functions to a url
	http.HandleFunc("/", nil)
	http.HandleFunc("/pull", restAPI.Insert)
	http.HandleFunc("/register", restAPI.History)
	http.HandleFunc("/search", Pay) 

	//Hear and handle
	http.ListenAndServe(":"+strconv.Itoa(config.Port), nil)
}
