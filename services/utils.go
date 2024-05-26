package services

import (
	"io"
	"net/http"

	"github.com/Ashu23042000/logger/logger"
)

func Get(url string) (string, error) {

	client := http.Client{}
	log := logger.New(nil)

	resp, err := client.Get(url)
	if err != nil {
		log.Info("error while making http request")
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Info("error while reading response body")
		return "", err
	}

	return string(body), nil
}

func Post(url string, contentType string, reqbody io.Reader) (string, error) {

	client := http.Client{}
	log := logger.New(nil)

	resp, err := client.Post(url, contentType, reqbody)
	if err != nil {
		log.Info("error while making http request")
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Info("error while reading response body")
		return "", err
	}

	return string(body), nil
}
