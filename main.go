package main

import (
	"sync"
	"github.com/GOKOP/markdown-website/server"
	"github.com/GOKOP/markdown-website/siteconfig"
)

func main() {

	config := siteconfig.Read("config.yaml")

	server.HandlerSetup(config.Files)

	var wait sync.WaitGroup
	wait.Add(2)

	if config.ServeHttp {
		go server.Serve(":"+config.PortHttp, &wait)
	}

	if config.ServeHttps {
		go server.ServeTLS(":"+config.PortHttps, config.CertFile, config.KeyFile, &wait)
	}

	wait.Wait()
}
