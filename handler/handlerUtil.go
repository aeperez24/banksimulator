package handler

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/port"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(dto.ResponseDto{Data: payload})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithJSONGin(c *gin.Context, code int, payload interface{}) {
	c.JSON(code, dto.ResponseDto{Data: payload})
}

type HandlerConfig struct {
	AccountHandler        port.AccountHandler
	UserHandler           port.UserHandler
	AuthenticationHandler port.AuthenticationHandler
}

type HandlerConfigGin struct {
	AccountHandler        GinAccountHandlerImpl
	AuthenticationHandler GinAutenticationHandlerImpl
	UserHandler           GinUserHandlerImpl
}
