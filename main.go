package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonRequest struct {
	Message string `json:"message"`
}

type JsonResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	http.HandleFunc("/json-endpoint", handleJSONRequest)
	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", nil)
}

func handleJSONRequest(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	var req map[string]interface{}
	if err := decoder.Decode(&req); err != nil {
		res := JsonResponse{
			Status:  "400",
			Message: "Invalid JSON message",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(res); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		return
	}

	message, ok := req["message"].(string)
	if !ok || message == "" {

		res := JsonResponse{
			Status:  "400",
			Message: "Invalid JSON message",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)

		encoder := json.NewEncoder(w)
		if err := encoder.Encode(res); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
		return
	}

	fmt.Println("Received message:", message)

	res := JsonResponse{
		Status:  "success",
		Message: "Data successfully received",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(res); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

