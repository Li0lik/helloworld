package main

import (
	"healthcheck/bean"
	"log"
	"net/http"
)

const REST_HEALTH_API_NAME = "/api/v1/healthcheck"
const PORT_NAME = "8888"

func main() {

	httpresolver := bean.HttpResolver{}

	http.HandleFunc(REST_HEALTH_API_NAME, httpresolver.Resolver)

	log.Println("[INFO] Server running on port " + PORT_NAME)
	err := http.ListenAndServe(":"+PORT_NAME, nil)
	if err != nil {
		log.Println("[ERROR ]Server failed " + err.Error())
	}
}
