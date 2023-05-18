package api

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Payload interface{} `json:"payload"`
}

func CreateNewResponse(w http.ResponseWriter, response *Response) error {
	//Set the response code on the api
	w.WriteHeader(response.Status)
	//Convert the struct to a JSON body
	err := json.NewEncoder(w).Encode(response)
	return err
}
