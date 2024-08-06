package main

import (
	"log"
	"net/http"
	"os"
)

func main() {

	port := func() string {
		if len(os.Args) == 2 {
			return os.Args[1]
		} else {
			return "8080"
		}
	}()

	dir, err_wd := os.Getwd()
	if err_wd != nil {
		log.Fatal(err_wd)
	}

	fs := http.FileServer(http.Dir(dir))

	http.Handle("/", fs)

	log.Print("server started on port ", port)
	err := http.ListenAndServe(":"+port, nil)

	if err != nil {
		log.Fatal(err)
	}
}
