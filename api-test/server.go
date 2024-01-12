package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", Books)
	http.HandleFunc("/envs", Envs)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func Books(w http.ResponseWriter, r *http.Request) {
	books := []Book{
		{Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams", ISBN: "0345391802"},
		{Title: "The Hobbit", Author: "J.R.R. Tolkien", ISBN: "0345391802"},
		{Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", ISBN: "0345391802"},
		{Title: "The Silmarillion", Author: "J.R.R. Tolkien", ISBN: "0345391802"},
	}
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(books)
}

type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	ISBN   string `json:"isbn"`
}

func Envs(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(map[string]string{"name": name, "age": age})

}
