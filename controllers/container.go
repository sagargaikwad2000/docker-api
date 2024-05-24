package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Ashu23042000/docker-api/constants"
	"github.com/Ashu23042000/docker-api/models"
	"github.com/Ashu23042000/docker-api/services"
	"github.com/gin-gonic/gin"
)

func Get(ctx *gin.Context) {

	response := models.Response{}
	res, err := services.Get(constants.ContainerList)

	if err != nil {
		response.StatusCode = http.StatusInternalServerError
		response.Msg = "Failed to get list of containers running"
		response.Err = err.Error()
	}

	response.StatusCode = http.StatusOK
	response.Msg = "Success"

	response.Result = res

	jsonResponse, _ := json.MarshalIndent(response, "", " ")

	ctx.JSON(200, gin.H{"response": string(jsonResponse)})

}
