package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

var localCache = map[string][]byte{}

const originServerURL = "http://localhost:9000"

func handleGetRequest(w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	w.Header().Set("Content-Type", "image/jpg")
	reqPath := r.URL.Path
	if localCache[reqPath] == nil {
		// if request file not present in localCache, get it from origin server and cache it
		log.Printf("Getting from origin server : %s", reqPath)
		completeURL := fmt.Sprintf(originServerURL + reqPath)

		// call origin server
		resp, err := http.Get(completeURL)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		defer resp.Body.Close()
		bytes, convErr := io.ReadAll(resp.Body)
		if convErr != nil {
			log.Printf("Convert error : %v", convErr)
		}

		// set response in localCache
		localCache[reqPath] = bytes
		w.Write(bytes)
		log.Printf("Time taken by : %s is %s", reqPath, time.Since(startTime))
		return
	} else {
		// if requested file is present in localCache, return it
		w.Write(localCache[reqPath])
		log.Printf("Time taken by : %s is %s", reqPath, time.Since(startTime))
		return
	}
}

func main() {
	log.Println("Welcome to dummy CDN.")
	http.HandleFunc("/", handleGetRequest)
	log.Println("Starting CDN at 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))

}
