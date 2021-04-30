package ping

import (
	"encoding/json"
	"net/http"
)

type pingResponse struct {
	Pong string `json:"pong"`
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(pingResponse{Pong: "ok"})
	if err != nil {
		return
	}
}