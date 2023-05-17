package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var (
	lastIP = ""
)

func GetMyIP() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	newIP, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(newIP), nil
}

func updateDNS() {
	newIP, err := GetMyIP()
	if err != nil {
		log.Printf("failed to get my IP address. %s\n", err)
		return
	}
	if newIP != lastIP {
		log.Printf("my new IP address is %s\n", newIP)
		lastIP = newIP

		updtUrl := fmt.Sprintf("%s%s", *dyndns2, newIP)
		resp, err := http.Get(updtUrl)
		if err != nil {
			log.Printf("failed to update DNS. %s\n", err)
			lastIP = ""
			return
		}
		defer resp.Body.Close()

		text, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("failed to read body of response to my update DNS request. Assuming it worked!")
			return
		}
		log.Printf("DNS updated.\n%s\n", text)
	}
}

func continuouslyUpdateIP() {
	updateDNS()
	ticker := time.NewTicker(time.Minute * 5)
	for range ticker.C {
		updateDNS()
	}
}
