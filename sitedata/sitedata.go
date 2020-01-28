package sitedata

import (
	"log"
	"os"
	"path/filepath"
	"bufio"
	"io/ioutil"
	"strings"
	"bytes"
	"gopkg.in/russross/blackfriday.v2"
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

func (page *Page) addMenuEntry(title string, path string) {

	entry := MenuEntry {
		Title: strings.TrimRight(title, "\n"),
		Dest:  strings.TrimSuffix( filepath.Base(path), ".md" ),
	}

	page.Menu = append(page.Menu, entry)
}

func (page *Page) setData(title string, path string) {

	markdown, err := readContent(path)

	if err != nil {
		page.setError()
	}

	html := blackfriday.Run(markdown)

	page.Title   = title
	page.Content = string(html)
}

func (page *Page) setError() {
	page.Title   = "Error"
	page.Content = "<h1>An error ocurred while retreiving contents of this page</h1>"
}

func readTitle(path string) (string, error) {

	file, err := os.Open(path)

	if err != nil {
		log.Print("Opening "+path+": "+err.Error())
		return "", err
	}

	reader     := bufio.NewReader(file)
	title, err := reader.ReadString('\n') // read first line

	file.Close()

	if err != nil {
		log.Print("Reading title from "+path+": "+err.Error())
		return "", err
	}

	return title, nil
}

func readContent(path string) ([]byte, error) {

	content, err := ioutil.ReadFile(path)

	if err != nil {
		log.Print("Reading content of "+path+": "+err.Error())
		return nil, err
	}

	endOfTitle := bytes.IndexByte(content, '\n')
	content    =  content[ endOfTitle+1 :]

	return content, nil
}

func Get(name string) Page {

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

		targetPage := false
		if strings.TrimSuffix( filepath.Base(path), ".md" ) == name {
			targetPage = true
		}

		title, err := readTitle(path)

		if err == nil {
			page.addMenuEntry(title, path)

			if targetPage {
				page.setData(title, path)
			}

		} else if targetPage {
			page.setError()
		}

		return nil
	})

	if err != nil {
		log.Print(err)
	}

	return page
}
