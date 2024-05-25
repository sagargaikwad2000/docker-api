package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Ashu23042000/docker-api/constants"
	"github.com/Ashu23042000/docker-api/models"
	"github.com/Ashu23042000/docker-api/services"
	"github.com/gin-gonic/gin"
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

	var containers []models.Container

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
