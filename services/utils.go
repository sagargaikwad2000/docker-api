package services

import (
	"io"
	"log"
	"net/http"
)

func Get(url string) (string, error) {

	client := http.Client{}

	resp, err := client.Get(url)
	if err != nil {
		log.Println("error while making http request")
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("error while reading response body")
		return "", err
	}

	return string(body), nil
}
