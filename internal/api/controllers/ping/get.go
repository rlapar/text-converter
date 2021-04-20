package ping

import (
	"encoding/json"
	"net/http"
)

type PingResponse struct {
	Pong string `json:"pong"`
}

func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(PingResponse{Pong: "ok"})
	if err != nil {
		return
	}
}