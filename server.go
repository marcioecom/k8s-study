package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/configmap", configMap)
	http.ListenAndServe(":3000", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := os.Getenv("NAME")
	age := os.Getenv("AGE")

	fmt.Fprintf(w, "Hello, I'm %s and I'm %s years old", name, age)
}

func configMap(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("myskills/skills.txt")
	if err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	fmt.Fprintf(w, "My skills: %s", string(data))
}
