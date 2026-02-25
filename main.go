package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "salve")
}

func main() {
	http.HandleFunc("/", defaultHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Erro ao iniciar servidor: %v", err)
	}
}
