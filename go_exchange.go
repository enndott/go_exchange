// go_exchange... simple command line access to exchange rates
//
// by: nicholas ward
// http://github.com/enndott
//
// note: mostly just to familiarize myself w/ how to deal with JSON in go.
// definitely not stable :)

package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func fetchExchangeRatesJSON(apikey string) []byte {

	//gets the latest values from openexchangerates
	res, err := http.Get("http://openexchangerates.org/api/latest.json?app_id=" + apikey)
	if err != nil {
		panic(err)
	}

	//read the response into a variable, jsonData
	jsonData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	//close it
	res.Body.Close()
	return jsonData
}

func parseExchangeRatesJSON(jsonData []byte) map[string]interface{} {

	//parses the exchange rate JSON downloaded in fetchExchangeRatesJSON
	//returns a map of [currency code]:[exchange rate] ex: "USD:1"
	response := map[string]interface{}{}
	json.Unmarshal(jsonData, &response)

	//check for the error response in the JSON
	//echo and quit with details if found.
	if response["error"] != nil {
		fmt.Println("There was an error: " + response["message"].(string))
		fmt.Println("Details: " + response["description"].(string))
		fmt.Println("")
		os.Exit(1)
	}

	//skip to the exchange rate map and return
	return response["rates"].(map[string]interface{})

}

func main() {

	//set the API key
	//to get an API key go to: https://openexchangerates.org/signup
	var apikey string = ""
	if apikey == "" {
		fmt.Println("There was an error: no api key defined.")
		fmt.Println("Please add an API key to line 64.")
		os.Exit(1)
	}

	//get the latest values and parse
	jsonData := fetchExchangeRatesJSON(apikey)
	rates := parseExchangeRatesJSON(jsonData)

	//check to see of an argument was passed. If so,
	//parse the list
	if len(os.Args) > 1 {

		//split it by comma
		currencyList := strings.Split(os.Args[1], ",")
		for _, value := range currencyList {
			if rates[value] != nil {
				fmt.Println(value, ": ", rates[value])
			} else {
				fmt.Println(value, ":  currency not found.")
			}

		}

	} else {
		fmt.Println("Printing all available exchange rates.")
		for key, value := range rates {
			fmt.Println(key, ": ", value)
		}
	}

}
