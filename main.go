package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	o := flag.Bool("o", false, "output hello world")
	p := flag.Bool("p", false, "make panic")
	h := flag.Bool("h", false, "run http server localhost:8080")
	flag.Parse()

	if *o {
		fmt.Println("hellow world")
	}

	if *p {
		panic("Panic!")
	}

	if *h {
		httpServ()
	}
}

func httpServ() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})
	http.ListenAndServe(":8080", nil)
}
