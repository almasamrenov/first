package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rs/cors"
)

type PersonRequest struct {
	Person struct {
		Name     string `json:"name"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	} `json:"person"`
}


type PersonResponse struct {
	Status string `json:"status"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/person-endpoint", handlePersonRequest)

	handler := cors.Default().Handler(mux)

	fmt.Println("Server listening on :8080")
	http.ListenAndServe(":8080", handler)
}

func handlePersonRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		handlePersonPost(w, r)
	default:
		sendErrorResponse(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func handlePersonPost(w http.ResponseWriter, r *http.Request) {
	var req PersonRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&req); err != nil {
		sendErrorResponse(w, "Invalid JSON message", http.StatusBadRequest)
		return
	}

	fmt.Printf("Received POST request with Person data: %+v\n", req.Person)

	res := PersonResponse{
		Status: "success",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(res); err != nil {
		sendErrorResponse(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}


func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	errRes := ErrorResponse{
		Status:  fmt.Sprintf("%d", statusCode),
		Message: message,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	encoder := json.NewEncoder(w)
	if err := encoder.Encode(errRes); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}
