package response

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func JSON(
	w http.ResponseWriter,
	statusCode int,
	success bool,
	message string,
	data any,
) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := APIResponse{
		Success: success,
		Message: message,
		Data:    data,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "failed to encode response", http.StatusInternalServerError)
	}
}

func Success(w http.ResponseWriter, statusCode int, message string, data any) {
	JSON(w, statusCode, true, message, data)
}

func Error(w http.ResponseWriter, statusCode int, message string, data any) {
	JSON(w, statusCode, false, message, data)
}
