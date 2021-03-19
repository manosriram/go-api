package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

var functions = template.FuncMap{}

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	tc, err := CreateTemplateCache(w, tmpl)
	if err != nil {
		log.Fatal()
	}

	t, ok := tc[tmpl]
	if !ok {
		log.Fatal()
	}

	buf := new(bytes.Buffer)
	_ = t.Execute(buf, nil)
	_, err = buf.WriteTo(w)
	if err != nil {
		fmt.Println("Error writing template to browser: ", err)
	}

	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err = parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func CreateTemplateCache(w http.ResponseWriter, tmpl string) (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.page.tmpl")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		fmt.Println("Page is currently: ", name)
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)

		if err != nil {
			return cache, err
		}
		matches, err := filepath.Glob("./templates/*.layout.tmpl")

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.tmpl")
			if err != nil {
				return cache, err
			}
		}
		cache[name] = ts
	}
	return cache, nil
}
