package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	Path := r.URL.Path[1:]

	//Main path for the website
	if Path == "" {
		dat, err := box.FindString("index.html")
		check(err)
		fmt.Fprintf(w, dat)
	}

	if Path == "add" {
		dat, err := box.FindString("add.html")
		check(err)
		fmt.Fprintf(w, dat)
	}

	if Path == "tab" {
		dat, err := box.FindString("tab.html")
		check(err)
		fmt.Fprintf(w, dat)
	}
}
