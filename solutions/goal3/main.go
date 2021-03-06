package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var serverURL = "http://localhost:8080/"

var database map[string]int

// EncodeRequest represents a request to encode
type EncodeRequest struct {
	URL string
}

// EncodeResponse represents the result of encoding
type EncodeResponse struct {
	ShortURL string
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusNotImplemented)
		fmt.Fprintf(w, "NOT IMPLEMENTED")

	case "POST":
		jsonBody := json.NewDecoder(r.Body)
		var encReq EncodeRequest
		jsonBody.Decode(&encReq)

		if len(encReq.URL) > 0 {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")

			var res EncodeResponse
			res.ShortURL = encodeURL(encReq.URL)

			jsonRes, err := json.Marshal(res)
			if err != nil {
				fmt.Println("error:", err)
			}

			fmt.Fprintf(w, "%s", jsonRes)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, " { \"error\": \"No URL found\" }")
	}

}

func encodeURL(URL string) string {
	if database == nil {
		database = make(map[string]int)
	}

	nextInt := len(database) + 1
	database[URL] = nextInt
	return fmt.Sprintf("%s%d", serverURL, nextInt)
}
