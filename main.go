package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zet4/go-travis-docker-test/lib"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			return
		}
		host, err := os.Hostname()
		if err != nil {
			w.Write([]byte("Encountered error while trying to get hostname..."))
			return
		}
		fmt.Fprintf(w, "Hello there from %s, this instance has seen %d hits.", host, lib.IncrementAndReturn())
	})

	http.HandleFunc("/reset", func(w http.ResponseWriter, r *http.Request) {
		host, err := os.Hostname()
		if err != nil {
			w.Write([]byte("Encountered error while trying to get hostname..."))
			return
		}
		lib.Reset()
		fmt.Fprintf(w, "Hello there from %s, this instance's counter has been reset.", host)
	})

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		log.Fatal(err)
	}
}
