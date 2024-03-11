package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templateCache = make(map[string]*template.Template)

func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	_, inMap := templateCache[t]

	if !inMap {
		// Need to add the template to the cache,
		log.Println("Creating the template and adding it to the cache", t)
		err = createTemplateCache(t)
		if err != nil {
			log.Println("Error creating template cache", err)
		}
	} else {
		// We have the template, so get it from the cache
		log.Println("Using the template from the cache", t)
	}

	tmpl = templateCache[t]

	err = tmpl.Execute(w, nil)

	if err != nil {
		log.Println("Error executing template", err)
		return
	}
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf(("./templates/%s"), t),
		"./templates/base.layout.tmpl",
	}

	tmpl, err := template.ParseFiles(templates...)

	if err != nil {
		return err
	}

	templateCache[t] = tmpl

	return nil
}
