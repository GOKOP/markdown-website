package main

import (
	"fmt"
	"github.com/GOKOP/markdown-website/sitedata"
)

func main() {
	page := sitedata.Get()

	fmt.Println(page.Title)
	fmt.Println(page.Content)
}
