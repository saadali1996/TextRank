package main

import (
	"fmt"

	"github.com/DavidBelicza/TextRank"
	"runtime"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"encoding/json"
)
func generate_data(rawText string)[]string{
	res:=[]string{}
	tr := textrank.NewTextRank()
	// Default Rule for parsing.
	rule := textrank.NewDefaultRule()
	// Default Language for filtering stop words.
	language := textrank.NewDefaultLanguage()
	// Default algorithm for ranking text.
	algorithmDef := textrank.NewChainAlgorithm()

	// Add text.
	tr.Populate(rawText, language, rule)
	// Run the ranking.
	tr.Ranking(algorithmDef)

	// Get the most important 10 sentences. Importance by phrase weights.
	sentences_realtional := textrank.FindSentencesByRelationWeight(tr, 20)
	sentences_wordweight := textrank.FindSentencesByWordQtyWeight(tr, 20)
	// Found sentences
	for i, _ := range sentences_realtional {
		res=append(res, sentences_realtional[i].Value)
		res=append(res,sentences_wordweight[i].Value)


	}
	runtime.GC()
	return res

}
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("200 - OK"))
}
func getsummary_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	r.ParseForm()
	url := r.Form["text"]
	title:=r.Form["title"]
	out := generate_data(url[0]+title[0])
	json.NewEncoder(w).Encode(out)

}
func main() {
	fmt.Println("WeLCOME TO GET SENTENCES API")
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/getsummary", getsummary_handler)
	router.HandleFunc("/", ServeHTTP)
	log.Fatal(http.ListenAndServe(":8080", router))

}