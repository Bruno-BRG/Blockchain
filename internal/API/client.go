package API

import (
	"Blockchain/internal/server"
	"net/http"
)

func Run() {
	http.HandleFunc("/blocks", server.GetBlocks)
	http.ListenAndServe(":8080", nil)
}
