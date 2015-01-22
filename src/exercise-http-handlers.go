package main

import (
	"fmt"
	"log"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s+"\n")
}

type Struct struct {
	Greeting string
	Punct    string
	Who      string
}

func (s Struct) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, fmt.Sprintf("%s%s %s\n", s.Greeting, s.Punct, s.Who))
}

func main() {
	http.Handle("/", &Struct{"Hello", ":", "use /string or /struct resource!"})
	http.Handle("/string", String("I'm a frayed knot."))
	http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
	log.Fatal(http.ListenAndServe("localhost:40000", nil))
}
