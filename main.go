package main

import (
	"fmt"
	"net/http"
	"os"
)

func defaultHandler(w http.ResponseWriter, _ *http.Request) {
	data, err := os.ReadFile("data.txt")
	if err != nil {
		http.Error(w, "failed to read file", http.StatusInternalServerError)
		return
	}
	w.Write(data)
}

func main() {
	http.HandleFunc("/", defaultHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Erro ao iniciar servidor: %v", err)
	}
}
