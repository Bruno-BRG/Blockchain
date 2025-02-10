package server

import (
	"Blockchain/internal/service"
	"encoding/json"
	"net/http"
)

func GetBlocks(w http.ResponseWriter, r *http.Request) {
	// Marshal the blocks from the global blockchain instance
	jsonData, err := json.Marshal(service.BlockchainInstance.Blocks)
	if err != nil {
		http.Error(w, "Error generating JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
