package server

import (
	"net/http"
	"strings"
	"text/template"
	"log"
	"fmt"
	"github.com/GOKOP/markdown-website/sitedata"
)

func mainHandler(w http.ResponseWriter, r *http.Request) {

	pageName := strings.Trim( r.URL.Path, "/" )
	page     := sitedata.Get(pageName)

	templ, err := template.ParseFiles("template/template.html")

	if err != nil {
		log.Print("Reading html template: "+err.Error())
		fmt.Fprintf(w, "<h1>Error while reading html template</h1>")
		return
	}

	templ.Execute(w, page)
}

func Serve(port string) {

	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(port, nil)
}
