package main

import (
	"html/template"
	"log"
	"net/http"
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
	tpl.ExecuteTemplate(w, "index.gohtml", req.Form)
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