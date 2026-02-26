package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type EditablePage struct {
	Title string
	Data  string
}

var editableTemplate = template.Must(template.ParseFiles("templates/editable.html"))

func editableTemplateHandler(w http.ResponseWriter, r *http.Request) {
	fileData, err := os.ReadFile("data.txt")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	pageData := &EditablePage{
		Title: "Vault",
		Data:  string(fileData),
	}

	err = editableTemplate.Execute(w, pageData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	http.HandleFunc("/", editableTemplateHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Erro ao iniciar servidor: %v", err)
	}
}
