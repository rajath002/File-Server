package render

import (
	"bytes"
	"errors"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/rajath002/File-Server/internal/models"
)

var pathToTemplates = "./templates"

func CreateTemplateDynamicCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all the files names *.page.tmpl from ./templates
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))

	if err != nil {
		log.Println("Something went wrong! 98")
		log.Println(err, pages)
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		// ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			log.Println("Something went wrong!108")
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))

		if err != nil {
			log.Println("Something went wrong!115")
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				log.Println("Something went wrong!122")
				return myCache, err
			}
		}

		myCache[name] = ts
	}
	return myCache, nil
}

// Template Renders html/template
func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {
	// Create a template cache
	var tc map[string]*template.Template

	tc, _ = CreateTemplateDynamicCache()

	// get a requested template cache
	t, ok := tc[tmpl]
	if !ok {
		return errors.New("could not get the template from the template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, td)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
