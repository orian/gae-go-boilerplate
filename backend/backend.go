package backend

import (
	"backend/handlers"
	"net/http"
)

func init() {
  http.HandleFunc("/", handlers.Main)
	http.HandleFunc("/helloWorld", handlers.HelloWorld)
}
