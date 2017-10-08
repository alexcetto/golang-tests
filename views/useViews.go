package views

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"html/template"
)

func PopulateTemplates() {
	var allFiles [] string
	templatesDir := "./templates/"
	files, err := ioutil.ReadDir(templatesDir)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasSuffix(filename,".html") {
			allFiles = append(allFiles, templatesDir+filename)
		}
	}
	templates, err = template.ParseFiles(allFiles...)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	homeTemplate = templates.Lookup("home.html")
}