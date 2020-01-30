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

	if pageName == "" {
		pageName = "index"
	}

	page := sitedata.Get(pageName)

	templ, err := template.ParseFiles("template.html")

	if err != nil {
		log.Print("Reading html template: "+err.Error())
		fmt.Fprintf(w, "<h1>Error while reading html template</h1>")
		return
	}

	templ.Execute(w, page)
}

func createFileHandlers(files []string) {

	for _,file := range files {
		http.HandleFunc("/"+file, func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "files/"+file)
		})
	}
}

func Serve(port string, files []string) {

	http.HandleFunc("/", mainHandler)
	createFileHandlers(files)
	log.Fatal( http.ListenAndServe(port, nil) )
}
