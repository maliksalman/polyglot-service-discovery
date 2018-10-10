package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func getEnvWithDefault(key string, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	} else {
		return defaultVal
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	color := getEnvWithDefault("MY_COLOR", "unknown")
	w.Header().Add("Content-Type", "application/json")
	response := map[string]string{"color": color}
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/color", handle)
	port := getEnvWithDefault("SERVER_PORT", "8080")
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
