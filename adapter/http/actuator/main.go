package actuator

import (
	"encoding/json"
	"net/http"
)

// HealthBody type to describe the health of the application
type HealthBody struct {
	Status string `json:"status"`
}

// Health function to get status of the application
func Health(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	var health = HealthBody{"alive"}

	_ = json.NewEncoder(w).Encode(health)

}
