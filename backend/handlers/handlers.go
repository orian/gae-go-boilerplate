package handlers

import (
  // "appengine"
  // "backend/models"
  "fmt"
  "net/http"
  "html/template"
)

var (
  indexTemplate = template.Must(template.ParseFiles("templates/index.html"))
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
  fmt.Fprint(w, "Hello, world!")
}

func Main(w http.ResponseWriter, r *http.Request) {
  indexTemplate.Execute(w, nil)
}
