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

	if config.ServeHttp {
		wait.Add(1)
		go server.Serve(":"+config.PortHttp, &wait)
	} else {
		wait.Add(1)
		go server.RedirectToHttps(config.PortHttp, config.PortHttps, &wait)
	}

	if config.ServeHttps {
		wait.Add(1)
		go server.ServeTLS(":"+config.PortHttps, config.CertFile, config.KeyFile, &wait)
	}

	wait.Wait()
}
