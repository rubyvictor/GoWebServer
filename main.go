package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"sort"
	s "strings"
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

func filterPairs(allPairs PairList) PairList {
	return allPairs[0:10]
}

func rankByWordCount(wordFrequencies map[string]int) PairList {
	pl := make(PairList, len(wordFrequencies))
	i := 0
	for k, v := range wordFrequencies {
		pl[i] = Pair{k, v}
		i++
	}
	sort.Sort(pl)
	return filterPairs(pl)
}

type Pair struct {
	Key   string
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

func displayPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	fmt.Fprintf(w, "Hello & Welcome To My First GoLang Server")
}

func text(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("form.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// fmt.Println(r.Form)
		// fmt.Println("path", r.URL.Path)
		// fmt.Println("scheme", r.URL.Scheme)
		// fmt.Println(r.Form["text"])
		for key, value := range r.Form {
			fmt.Println("key", key)
			fmt.Println("val:", s.Join(value, ""))
			fmt.Fprintf(w, "Top 10 Words:%v", rankByWordCount(countWords(getWordsFrom(s.Join(value, "")))))
		}
	}

}

func main() {
	http.HandleFunc("/", displayPage)
	http.HandleFunc("/text", text)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
