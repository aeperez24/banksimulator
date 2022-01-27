package handler

import (
	"aeperez24/banksimulator/dto"
	"aeperez24/banksimulator/port"
	"encoding/json"
	"net/http"
)

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(dto.ResponseDto{Data: payload})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

type HandlerConfig struct {
	AccountHandler port.AccountHandler
	UserHandler    port.UserHandler
}
