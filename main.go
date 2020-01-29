package main

import (
	"github.com/GOKOP/markdown-website/server"
)

func main() {
	files := []string{"style.css"}
	server.Serve(":8080", files)
}
