package main

import (
	"github.com/GOKOP/markdown-website/server"
	"github.com/GOKOP/markdown-website/siteconfig"
)

func main() {

	config := siteconfig.Read("config.yaml")
	server.Serve(":"+config.Port, config.Files)
}
