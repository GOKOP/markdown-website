package server

import (
	"net"
	"net/http"
	"strings"
	"text/template"
	"log"
	"fmt"
	"sync"
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

func HandlerSetup(files []string) {

	http.HandleFunc("/", mainHandler)
	createFileHandlers(files)
}

type HttpsRedirect struct {
	HttpsPort string
}

func (red HttpsRedirect) handler(w http.ResponseWriter, r *http.Request) {

	host := r.Host

	if strings.Contains(r.Host, ":") {
		splitHost, _, err := net.SplitHostPort(r.Host)
		host = splitHost

		if err != nil {
			log.Print("HTTPS redirection: Failed to separate host from port;", err.Error())
			return
		}
	}

	if red.HttpsPort == "443" {
		red.HttpsPort = ""
	}

	target := "https://" + host + ":" + red.HttpsPort + r.URL.Path
	http.Redirect(w, r, target, http.StatusPermanentRedirect)
}

func RedirectToHttps(httpPort string, httpsPort string, wait *sync.WaitGroup) {

	defer wait.Done()

	redirect := HttpsRedirect {
		HttpsPort: httpsPort,
	}

	log.Print("Redirecting HTTP on port ", httpPort, " to HTTPS on port ", httpsPort)
	http.ListenAndServe(":"+httpPort, http.HandlerFunc(redirect.handler))

}

func Serve(port string, wait *sync.WaitGroup) {

	defer wait.Done()

	log.Print("Starting HTTP server on port "+port)
	log.Fatal( http.ListenAndServe(":"+port, nil) )
}

func ServeTLS(port string, cert string, key string, wait *sync.WaitGroup) {

	defer wait.Done()

	log.Print("Starting HTTPS server on port "+port)
	log.Fatal( http.ListenAndServeTLS(":"+port, cert, key, nil) )
}
