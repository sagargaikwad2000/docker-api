package services

import (
	"fmt"

	"github.com/Ashu23042000/docker-api/constants"
)

func StartContainer(id string) (string, error) {
	url := fmt.Sprintf("http://127.0.0.1:2375/v1.45/containers/%s/start", id)
	return Post(url, constants.ContentType, nil)
}
