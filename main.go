package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
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

	port := port()

	if err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil); err != nil {
		panic(err)
	}
}

func port() string {
	port := os.Getenv("GO_SAMPLE_HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("port: %s", port)
	return port
}
