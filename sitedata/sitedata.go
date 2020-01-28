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

func readTitle(file *os.File) (string,error) {

	reader     := bufio.NewReader(file)
	title, err := reader.ReadString('\n') // read first line

	if err != nil {
		return "", err
	}

	return title, nil
}

func (page *Page) addMenuEntry(file *os.File, path string) {

	title, err := readTitle(file)

	if err != nil {
		log.Print("Reading title from "+path+": "+err.Error())
		return
	}

	entry := MenuEntry {
		Title: strings.TrimRight(title, "\n"),
		Dest:  strings.TrimSuffix( filepath.Base(path), ".md" ),
	}

	page.Menu = append(page.Menu, entry)
}

func Get() Page {

	page := Page {
		Title:   "404 Not found",
		Content: "<h1>404: Page not found</h1>",
		Menu:    []MenuEntry{},
	}

	err := filepath.Walk("website", func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() || !strings.HasSuffix(path, ".md") {
			return nil
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}

		page.addMenuEntry(file, path)

		file.Close()

		return nil
	})

	if err != nil {
		log.Print(err)
	}

	return page
}
