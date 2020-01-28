package sitedata

import (
	"log"
	"os"
	"path/filepath"
)

type MenuEntry struct {
	Title string
	Dest  string
}

type Page struct {
	Title   string
	Content string
	Menu []MenuEntry
}

func Get() Page {

	page := Page {
		Title:   "404 Not found",
		Content: "<h1>404: Page not found</h1",
		Menu:    nil,
	}

	err := filepath.Walk("website", func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		// to do
		return nil
	})

	if err != nil {
		log.Print("filepath.Walk: "+err.Error())
	}

	return page
}
