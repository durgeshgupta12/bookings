package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/durgeshgupta12/bookings/internal/config"
	"github.com/durgeshgupta12/bookings/internal/models"
	"github.com/justinas/nosurf"
)

var functions = template.FuncMap{}

var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func RenderTemplate(w http.ResponseWriter, r *http.Request, hbs string, td *models.TemplateData) {
	//create template cache
	var tc map[string]*template.Template
	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	//get requested template from cache
	t, ok := tc[hbs]
	if !ok {
		log.Fatal("could not get template from template cache")
	}
	buf := new(bytes.Buffer)

	td = AddDefaultData(td, r)

	_ = t.Execute(buf, td)

	//render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := make(map[string]*template.Template)
	//get all the files named *.hbs from ./templates
	pages, err := filepath.Glob("./templates/*.hbs")

	if err != nil {
		return myCache, err
	}
	//range through all files ending with *.hbs
	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.hbs")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.hbs")
			if err != nil {
				return myCache, err
			}
		}
		myCache[name] = ts
	}

	return myCache, nil
}
