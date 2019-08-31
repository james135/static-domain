package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func findIPAddress() (string, error) {
	response, err := http.Get("https://api.ipify.org")
	if err != nil {
		return "", err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalln(err)
	}
	return string(body), nil
}
