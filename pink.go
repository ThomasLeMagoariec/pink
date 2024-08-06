package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	dir, err_wd := os.Getwd()
	if err_wd != nil {
		log.Fatal(err_wd)
	}

	fs := http.FileServer(http.Dir(dir))

	http.Handle("/", fs)

	log.Print("listening on port 8000")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal(err)
	}
}
