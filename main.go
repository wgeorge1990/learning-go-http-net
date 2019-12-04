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
	w.Header().Set("firstname", "william")
	w.Header().Set("lastname", "george")
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
	var personOne person
	http.ListenAndServe(":8080", personOne)
}

//type handler interface {
//	ServeHTTP(ResponseWriter, *request)
//}

// func (r *Request) formValue(key string) string {
// pass in form key and get back the value. This is a helper method when dealing with forms.
//}

//  type responseWriter interface {
//  	has below three methods:
//  	1.	Header() Header
//  	2.	write([]byte) (int, err)
//  	3.	writeHeader(int)
//  }