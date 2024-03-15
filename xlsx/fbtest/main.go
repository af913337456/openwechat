package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	FB_ACCESS_TOKEN = "EAAGNO4a7r2wBO2Ac5A6RvMEEggW3QfzLil5KCM9IJvNYrR78hIElZBHyqszGsBnxZCCiZBV3uhVBMQeQmEG9dT9DAri3bvJhLEZB8ayUzgcKZCMZCF4BuE254OVqgYzLDe8iSTGjwAaJcFVzjnBV8EGIBfJM93AqTWbtvtX9uN9AGjUbAqX30kFlEbgwZDZD"
	FB_API_VERSION  = "v3.1"
)

// Struct generated using https://mholt.github.io/json-to-go/
type GraphAPIResponse struct {
	Data []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"data"`
}

func main() {
	printMessage("Getting accounts...")
	getAccounts()
}

// Get accounts
func getAccounts() {
	// Make /me/adaccounts call to graphAPI
	url := "https://graph.facebook.com/" + FB_API_VERSION + "/me/adaccounts?limit=500&fields=id,name&access_token=" + FB_ACCESS_TOKEN

	response, err := http.Get(url)

	fmt.Println(response.StatusCode)

	checkErr(err)

	defer response.Body.Close()

	items := GraphAPIResponse{}

	body, err := ioutil.ReadAll(response.Body)

	checkErr(err)

	err = json.Unmarshal(body, &items)

	checkErr(err)

	// Accounts have been retrieved - save them to DB
	printMessage("Accounts have been retrieved successfully!")

	// For each item
	for _, item := range items.Data {
		accountID := item.ID
		accountName := item.Name

		fmt.Println("Account ID: " + accountID + ", account name: " + accountName)
	}
}

// Helper function for handling messages
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// Helper function for handling errors
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
