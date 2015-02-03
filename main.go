package main

import (
	"net/http"
  	"os"
)

func main() {

	router := NewRouter()

  	http.ListenAndServe(":" + os.Getenv("PORT"), router)
}