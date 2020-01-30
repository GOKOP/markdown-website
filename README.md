# markdown-website

*current version is 1.1*

markdown-website is a simple webserver serving a very minimalistic website based on markdown files.
The website is generated on request from a template so modifying it boils down to editing .md files.

## Features

* **Simplistic website markdown** - Each markdown file is one page and one entry in the navigation menu.
	Each page is generated from the same template, in the name of simplicity.
* **Serving files** - If your website needs some CSS or background image, you can do that.
	Currently however each file is served through a separate handler for the purpose of not serving a filesystem 
	(to avoid unnecessary vulnerabilites)
	so this program is not a very good option if you want image galeries or something like that.
* **HTTPS** - you can serve your website on HTTP, HTTPS, both or neither.
	If you choose to only serve on HTTPS, HTTP port will be used to redirect connections to HTTPS.
	If you choose neither then the program will exit.

## Planned features

* blog

## Obtaining and building
You need properly configured Go installation.
```
go get github.com/GOKOP/markdown-website
(in the project folder)
go build -i
```

You can also download precompiled binaries from [here](https://github.com/GOKOP/markdown-website/releases)

## Running
Just launch the executable (you may want to redirect its output to some file in case of some failure)

## Files
* `template.html` - Go template from which the site is generated. 
	Every page of the site is built from the same template.
	Default one is perfectly usable but you can change it as you want.
	(you should at least change the copyright notice in the footer)
* `config.yaml` - configuration
* `website/` - your markdown files go here.
	Every file is one subpage and one entry in navigation menu.
	Main page should be called `index.md`.
	All files should have the `.md` extension.
	First line of the file is page's tab title (the <title></title> one, you know) and doesn't appear in the page's content.
	All md files must have endlines formatted in Unix (LF) format, otherwise markdown may render incorrectly.
* `files/` - folder for all the other files you may need to serve, eg. CSS stylesheets.
	You also need to add an entry in the config for each one of them.
	This is for the purpose of not serving a filesystem to avoid vulnerabilites.

All other files (except for the executable of course) are code so you can get rid of them after building if you want to.
