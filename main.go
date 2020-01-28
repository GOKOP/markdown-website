package main

import (
	"fmt"
	"github.com/GOKOP/markdown-website/sitedata"
)

func main() {
	page := sitedata.Get("index")

	fmt.Println(page.Title)
	fmt.Println(page.Content)

	for i,entry := range page.Menu {
		fmt.Println(i)
		fmt.Println("  "+entry.Title)
		fmt.Println("  "+entry.Dest)
	}
}
