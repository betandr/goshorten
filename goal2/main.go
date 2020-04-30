package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// EncodeRequest represents a request to encode
type EncodeRequest struct {
	URL string
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "OK")

	case "POST":
		jsonBody := json.NewDecoder(r.Body)
		var encReq EncodeRequest
		jsonBody.Decode(&encReq)

		if len(encReq.URL) > 0 {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")

			fmt.Fprintf(w, "GOT: %s", encReq.URL)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, " { \"error\": \"No URL found\" }")
	}

}
