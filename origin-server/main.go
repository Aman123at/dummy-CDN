package main

import (
	"log"
	"net/http"
)

var staticPath = "/Users/amantiwari/Desktop/masterclass-excersizes/dummy-cdn/origin-server/static"

func handleAllRequests(w http.ResponseWriter, r *http.Request) {
	reqPath := r.URL.Path
	log.Printf("Received a request in origin server from path : %s", reqPath)
	// w.Write([]byte(reqPath + "/mynewresponsefromorigin"))
	w.Header().Set("Content-Type", "image/jpg")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	http.ServeFile(w, r, staticPath+reqPath)
}

func main() {
	log.Println("Welcome to origin server")
	http.HandleFunc("/", handleAllRequests)
	log.Println("Origin server running at 9000...")
	log.Fatal(http.ListenAndServe(":9000", nil))
}
