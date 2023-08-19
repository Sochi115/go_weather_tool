package userprompt

import "net/http"

type Prompt interface {
	GetCoordinates(httpClient *http.Client, apiKey string) (float64, float64, error)
}