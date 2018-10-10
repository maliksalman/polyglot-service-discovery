package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"
)

var colorClient = &http.Client{Timeout: 10 * time.Second}

type ColorValue struct {
	Color string
}

func createResponse(w http.ResponseWriter, color string) {
	w.Header().Add("Content-Type", "application/json")
	response := map[string]string{"name": "frontend", "color": color}
	json.NewEncoder(w).Encode(response)
}

func handle(w http.ResponseWriter, r *http.Request) {

	colorServiceUrl := getEnvWithDefault("COLOR_SVC", "http://localhost:8080")
	req, _ := http.NewRequest("GET", colorServiceUrl+"/color", nil)
	req.Header.Set("State", r.Header.Get("State"))

	res, err := colorClient.Do(req)
	if err != nil {
		log.Printf("Unable to connect to %s", colorServiceUrl)
		createResponse(w, "unavailable")
	} else {
		defer res.Body.Close()
		var color ColorValue
		json.NewDecoder(res.Body).Decode(&color)
		createResponse(w, color.Color)
	}
}

func getEnvWithDefault(key string, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	} else {
		return defaultVal
	}
}

func main() {
	http.HandleFunc("/", handle)
	port := getEnvWithDefault("SERVER_PORT", "8080")
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
