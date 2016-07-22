package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	o := flag.Bool("o", false, "output hello world")
	p := flag.Bool("p", false, "make panic")
	s := flag.Bool("s", false, "run http server localhost:8080")
	flag.Parse()

	if *o {
		fmt.Println("Hello, World")
	}

	if *p {
		panic("Hello, World from panic")
	}

	if *s {
		httpServ()
	}
}

func httpServ() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})
	http.ListenAndServe(":8080", nil)
}
