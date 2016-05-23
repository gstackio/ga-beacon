package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	initBeacon()
	err := http.ListenAndServe(fmt.Sprintf(":%v", getPort()), nil)
	if err != nil {
		panic(err)
	}
}

func getPort() string {
	if configuredPort := os.Getenv("PORT"); configuredPort == "" {
		return "8080"
	} else {
		return configuredPort
	}
}
