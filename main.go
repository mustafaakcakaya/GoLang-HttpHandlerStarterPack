package main

import (
	"encoding/json" //to produce json
	"fmt"           //format library
	"io/ioutil"     //input/output to read data
	"log"           //log. library
	"net/http"      //to use huge http library
)

//object example
type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

//a data structure type
type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	//to log if endpoint has been hit
	fmt.Println("Enpoint Hit: All Articles Endpoint")
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello Http"},
		Article{Title: "Test Title2", Desc: "Test Description2", Content: "Hello Http"},
	}
	json.NewEncoder(w).Encode(articles)
}
func sendData(wr http.ResponseWriter, r *http.Request) {

	log.Println("Your Data Sent")

	//error coming as well, we can handle it with this 'd, err :='
	d, _ := ioutil.ReadAll(r.Body)

	log.Printf("data: %s\n", d)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "homepage Endpoint Hit")
}
func handleRequests() {
	//simple default with forward slash
	http.HandleFunc("/", homePage)
	//endpoint examples
	http.HandleFunc("/articles", allArticles)
	http.HandleFunc("/senddata", sendData)

	//to listen and serve on a port. For simple way to bind specific adress '192.141.241.19:25001' (ip is example)
	//nill means i want a basic service(on the other hand, we can do things such as how to deploy app, how to test them, link them together, building caching etc. )
	log.Fatal(http.ListenAndServe(":25001", nil))
}
func main() {
	//main to start program
	handleRequests()
}
