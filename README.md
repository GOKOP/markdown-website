# markdown-website

markdown-website is a simple webserver serving a very minimalistic website based on markdown files.
The website is generated on request from a template so modifying it boils down to editing .md files.

## Obtaining and building
You need properly configured Go installation.
```
go get gopkg.in/russross/blackfriday.v2
go get gopkg.in/yaml.v3
go get github.com/GOKOP/markdown-website
go build
```

## Running
Just launch the executable

## Files
* `template.html` - Go template from which the site is generated. 
	Every page of the site is built from the same template.
	Default one is perfectly usable but you can change it as you want.
	You should at least change the copyright notice in the footer.
* `config.yaml` - configuration
* `website/` - your markdown files go here.
	Every file in one subpage and one entry in navigation menu.
	Main page should be called `index.md`.
	All files should have the `.md` extension.
	First line of the file is page's tab title (the <title></title> one, you know) and doesn't appear in the page's content.
* `files/` - folder for all the other files you may need to serve, eg. CSS stylesheets.
	You also need to add an entry in the config for each one of them.
	This is for the purpose of not serving a filesystem to avoid vulnerabilites.

All other files (except for the executable of course) are code so you can get rid of them after building if you want to.
