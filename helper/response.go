package helper

import (
	"encoding/json"
	"net/http"
)

// ResponseJSON sends a JSON response with proper formatting
func ResponseJSON(w http.ResponseWriter, statusCode int, data interface{}) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)

    // Indent the JSON response for better readability
    encodedResponse, err := json.MarshalIndent(data, "", "    ")
    if err != nil {
        // Handle encoding error if needed
        w.Write([]byte(`{"error": "Error encoding JSON"}`))
        return
    }

    // Write the indented JSON response
    w.Write(encodedResponse)
}
