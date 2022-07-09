package controllers

import (
	"encoding/json"
	"net/http"
)

func (s *Server) Home(w http.ResponseWriter, r *http.Request) {
	responseWithJSON(w, 200, "Hello world")
}


func responseWithJSON(response http.ResponseWriter, statusCode int, data interface{}) {
	result, _ := json.Marshal(data)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(statusCode)
	response.Write(result)
	}