package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type client struct {
	apiKey string
	client *http.Client
}

var lock = &sync.Mutex{}

var clientInstance *client

func getApiKey() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv("API_KEY")
	return apiKey
}
func getHttpClient() *client {
	if clientInstance == nil {
		
		lock.Lock()
		defer lock.Unlock()

		if clientInstance == nil {
			clientInstance = &client{
				apiKey: getApiKey(),
				client: &http.Client{}}

			log.Default().Println("New HTTP Client instantiated")
		} else {
			log.Default().Println("HTTP Client already instantiated")
		}
	} else {
		log.Default().Println("HTTP Client already instantiated")
	}
	return clientInstance
}
