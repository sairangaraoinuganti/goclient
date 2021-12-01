package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

//response struct
type Response struct {
	UID string `json:"uid"`
}

//api struct
type API struct {
	Client  *http.Client
	baseURL string
	path    string
}

// send request capture response and vaildate the response
func (api *API) sendRequest() ([]byte, error) {
	resp, err := api.Client.Get(api.baseURL + api.path)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error Occured. %+v", err)
	}
	var response Response
	unmarshalError := json.Unmarshal(body, &response)
	if unmarshalError != nil {
		log.Fatalf("Couldn't parse response body. %+v", unmarshalError)
	}
	if !isValidUID(response.UID) {
		log.Fatalf("invalid UUID")
	}
	// handling error and doing stuff with body that needs to be unit tested
	return body, err
}

func main() {

	c := http.Client{Timeout: time.Duration(1) * time.Second}
	path := os.Args[1]
	api := API{&c, "https://ionaapp.com/assignment-magic/dk", path}
	pathArray := strings.Split(api.path, "/")
	if pathArray[1] == "short" {
		if !isvaildparamforshort(pathArray[2]) {
			log.Fatalf("invalid path: " + path)
		}
	} else if pathArray[1] == "long" {
		if !isvaildparamforlong(pathArray[2]) {
			log.Fatalf("invalid path: " + path)
		}
	}
	body, _ := api.sendRequest()
	log.Println("Response Body:", string(body))
}

//check if UID is 32 digit alphanumeric characters
func isValidUID(uuid string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{32}$")
	return r.MatchString(uuid)
}

//check if param is 2 digit alphanumeric characters when endpiont is short
func isvaildparamforshort(pathparam string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{2}$")
	return r.MatchString(pathparam)
}

//check if param is 3 digit alphanumeric characters when endpiont is long
func isvaildparamforlong(pathparam string) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{3}$")
	return r.MatchString(pathparam)
}
