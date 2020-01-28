package sitedata

import (
	"log"
	"os"
	"path/filepath"
	"bufio"
	"strings"
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

func (page *Page) addMenuEntry(path string) {

	file, err := os.Open(path)

	if err != nil {
		log.Print("Opening "+path+": "+err.Error())
		return
	}

	reader     := bufio.NewReader(file)
	title, err := reader.ReadString('\n') // read first line

	file.Close()

	if err != nil {
		log.Print("Reading title from "+path+": "+err.Error())
	}

	entry := MenuEntry {
		Title: strings.TrimRight(title, "\n"),
		Dest:  strings.TrimSuffix( filepath.Base(path), ".md" ),
	}

	page.Menu = append(page.Menu, entry)
}

func Get(pageName string) Page {

	page := Page {
		Title:   "404 Not found",
		Content: "<h1>404: Page not found</h1>",
		Menu:    []MenuEntry{},
	}

	err := filepath.Walk("website", func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		page.addMenuEntry(path)
		return nil
	})

	if err != nil {
		log.Print("Traversing files: "+err.Error())
	}

	return page
}
