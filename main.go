package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

//build server
// Listen --> Accept -(conn)--> READ or WRITE
// Think about personal interaction with cell phone

type person int

func (p person) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method string
		URL *url.URL
		Submissions map[string][]string
		Header http.Header
		Host string
		ContentLength int64
	}{
		req.Method,
		req.URL,
		req.Form,
		req.Header,
		req.Host,
		req.ContentLength,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var p person
	http.ListenAndServe(":8080", p)
}

//type handler interface {
//	ServeHTTP(ResponseWriter, *request)
//}