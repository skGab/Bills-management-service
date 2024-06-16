package controllers

import (
	"net/http"

	"github.com/skGab/Bills-microservices/Bills-apigateway-service/app/services"

	"github.com/gin-gonic/gin"
)

type GateWayPipe struct {
	Services *services.Handlers
}

// ServeHTTP implements http.Handler.
func (gp *GateWayPipe) ServeHTTP(http.ResponseWriter, *http.Request) {
	panic("unimplemented")
}

// THE WAY THE USER COULD PASS DATA WITHIN THE REQUEST, SHOULD
// WILL GET THE REQUEST FROM THE USER AND EXTRACT THE ROUTING KEY
// THEN WOULD CALL THE CORREPONDING SERVICE BASED ON THE ROUTING KEY
func (gp *GateWayPipe) Run(gin *gin.Context) {
	payload := gp.Services.HandleRequest(gin)

	switch payload.routingKey {
	case "database":

		result, err := gp.Services.DatabaseHandle(*body, *params)

		if err != nil {
			gin.JSON(http.StatusInternalServerError, err)
		}

		gin.JSON(http.StatusOK, result)
	default:

		gin.JSON(http.StatusBadRequest, struct{ msg string }{msg: "Algo deu errado durante a requisição"})
	}

}
