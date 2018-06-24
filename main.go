package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	s "strings"
	"sort"
	"regexp"
)

func getWordsFrom(text string) []string {
	words := regexp.MustCompile("\\w+")
	return words.FindAllString(s.ToLower(text), -1)
}

func countWords(words []string) map[string]int {
	wordCounts := make(map[string]int)
	for _, word := range words {
		wordCounts[word]++
	}
	return wordCounts
}

func consoleOut(orderedWordCounts PairList) {
	for _, wordCount := range orderedWordCounts {
		fmt.Printf("%v\n", wordCount)
	}
}

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(sort.Reverse(pl))
	slicedPl := pl[0:10]
	return slicedPl
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }


func findWordsInText(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for key, value := range r.Form {
		fmt.Println("key", key)
		fmt.Println("val:", s.Join(value, ""))
	}
	fmt.Fprintf(w, "Hello & Welcome To My First GoLang Server")
}

func text(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	// var result PairList
	if r.Method == "GET" {
		t, _ := template.ParseFiles("form.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// result = reverseWords(r.FormValue("text"))
		// result = rankByWordCount(countWords(getWordsFrom(r.FormValue("text"))))
		fmt.Println("Top Ten Words:", rankByWordCount(countWords(getWordsFrom(r.FormValue("text")))))
	}
	
}

func main() {
	http.HandleFunc("/", findWordsInText)
	http.HandleFunc("/text", text)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
