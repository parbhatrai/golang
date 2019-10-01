package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	    filename := p.Title + ".txt"
	        return ioutil.WriteFile(filename, p.Body, 0600)
	}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Welcome</h1>")
	fmt.Fprint(w, "To get in touch, please send an email to <a href=\"maileto:support@wiki.com\">support@wiki.com</a>.")
}

func main() {
	http.HandleFunc("/", handlerFunc)
	http.ListenAndServe("localhost:3000", nil)
}






