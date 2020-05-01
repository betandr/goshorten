package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var serverURL = "http://localhost:8080/"

var lookup map[string]int = make(map[string]int)
var redirect map[int]string = make(map[int]string)

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
		uriParts := strings.Split(r.URL.String(), "/")

		i, err := strconv.Atoi(uriParts[len(uriParts)-1])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if to, ok := redirect[i]; ok {
			http.Redirect(w, r, to, http.StatusMovedPermanently)
			return
		}

		w.WriteHeader(http.StatusNotFound)

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
	if val, ok := lookup[URL]; ok {
		return fmt.Sprintf("%s%d", serverURL, val)
	}

	nextInt := len(lookup) + 1
	lookup[URL] = nextInt
	shortURL := fmt.Sprintf("%s%d", serverURL, nextInt)
	redirect[nextInt] = URL
	return shortURL
}
