package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Ashu23042000/docker-api/constants"
	"github.com/Ashu23042000/docker-api/models"
	"github.com/Ashu23042000/docker-api/models/container"
	"github.com/Ashu23042000/docker-api/services"
	"github.com/Ashu23042000/logger/logger"
	"github.com/gin-gonic/gin"
)

var (
	log = logger.New(nil)
)

func GetContainers(ctx *gin.Context) {
	response := models.Response{}
	res, err := services.Get(constants.ContainerList)

	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Msg = "Failed"
		response.Err = err.Error()
		ctx.JSON(response.StatusCode, gin.H{"response": response})
	}

	var containers []container.Container

	err = json.Unmarshal([]byte(res), &containers)
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Msg = "Failed"
		response.Err = err.Error()
		ctx.JSON(response.StatusCode, gin.H{"response": response})
	}

	response.StatusCode = http.StatusOK
	response.Msg = "Success"
	response.Result = containers

	ctx.JSON(response.StatusCode, gin.H{"response": response})
}

func CreateContainer(ctx *gin.Context) {
	response := models.Response{}
	res, err := services.Post(constants.CreateContainer, constants.ContentType, ctx.Request.Body)

	log.Debug(res)

	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Msg = "Failed"
		response.Err = err.Error()
		ctx.JSON(response.StatusCode, gin.H{"response": response})
	}

	type ResultStruct struct {
		Id       string   `json:"id,omitempty"`
		Warnings []string `json:"warnings,omitempty"`
	}

	var result = ResultStruct{}

	err = json.Unmarshal([]byte(res), &result)
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Msg = "Failed"
		response.Err = err.Error()
		ctx.JSON(response.StatusCode, gin.H{"response": response})
	}

	resp, err := services.StartContainer(result.Id)
	if err != nil {
		log.Debugf(err.Error())
		response.StatusCode = http.StatusInternalServerError
		response.Msg = "Failed"
		response.Err = err.Error()
		ctx.JSON(response.StatusCode, gin.H{"response": resp})
	}

	response.StatusCode = http.StatusOK
	response.Msg = "Success"
	response.Result = result

	ctx.JSON(response.StatusCode, gin.H{"response": response})

}

func InspectContainer(ctx *gin.Context) {
	response := models.Response{}
	id := ctx.Param("id")
	url := fmt.Sprintf("http://127.0.0.1:2375/v1.45/containers/%s/json", id)

	res, err := services.Get(url)

	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Msg = "Failed"
		response.Err = err.Error()
		ctx.JSON(response.StatusCode, gin.H{"response": response})
	}

	var containers []container.InspectContainer

	err = json.Unmarshal([]byte(res), &containers)
	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Msg = "Failed"
		response.Err = err.Error()
		ctx.JSON(response.StatusCode, gin.H{"response": response})
	}

	response.StatusCode = http.StatusOK
	response.Msg = "Success"
	response.Result = containers

	ctx.JSON(response.StatusCode, gin.H{"response": response})

}
