package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func reverseWords(input string) string {
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	output := string(rune)
	return output
}

func findWordsInText(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for key, value := range r.Form {
		fmt.Println("key", key)
		fmt.Println("val:", strings.Join(value, ""))
	}
	fmt.Fprintf(w, "Hello & Welcome To My First GoLang Server")
}

func text(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	var result string
	if r.Method == "GET" {
		t, _ := template.ParseFiles("form.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		result = reverseWords(r.FormValue("text"))
		fmt.Println("Top Ten Words:", reverseWords(r.FormValue("text")))
	}
	w.Write([]byte(result))
}

func main() {
	http.HandleFunc("/", findWordsInText)
	http.HandleFunc("/text", text)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
