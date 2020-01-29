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
		templ := template.Parse("<!DOCTYPE html><html>
		                         <head><title>{{.Title}}</title></head>
		                         <body><h1>Default template!</h1>
		                         {{range .Menu}}<a href='{{.Dest}}'>{{.Title}}</a><br/>{{end}}<br/>
		                         {{.Content}}</body>
		                         </html>"
	}

	templ.Execute(w, page)
}
