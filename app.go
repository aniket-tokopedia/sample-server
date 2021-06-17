package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func main() {

	http.HandleFunc("/v1/sample/api", func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Set("Content-Type", "application/json")
		response := struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
		}{
			FirstName: "Aniket",
			LastName:  "Chadha",
		}
		resJson, _ := json.Marshal(response)
		writer.WriteHeader(http.StatusOK)
		writer.Write(resJson)
	})

	server := &http.Server{
		ReadHeaderTimeout: 10*time.Second,
		WriteTimeout: 10*time.Second,
		Addr: ":8086",
	}
	server.ListenAndServe()
}
