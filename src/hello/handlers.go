package main

import (
	"encoding/json"
	"net/http"
)

func GetStatusEndpoint(w http.ResponseWriter, r *http.Request) {
	var status Status
	status.Msg = "Hello World"
	status.Version = "1.0.1"

	json.NewEncoder(w).Encode(status)
}
