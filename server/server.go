package main

import (
	"fmt"
	"log"
	"net/http"
	_ "time"
	_ "math"
)

func main() {
	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
			return
		}
		request := r.URL.Query().Get("request")
		if request == "" {
			http.Error(w, "Missing 'request' parameter", http.StatusBadRequest)
			return
		}

		temp := 3
		for i := 0; i < 10000; i++ {
			temp = temp*temp%1000
		}

		response := fmt.Sprintf("Processed request '%s'\n", request)
		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		if _, err := w.Write([]byte(response)); err != nil {
			log.Printf("Error writing response: %v", err)
		}
		fmt.Println("Send")
	})

	log.Println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
