package render

import (
	"bytes"
	"log"
	"net/http"
	"path/filepath"
	"text/template"

	"github.com/vkhoa145/golang-bookings/pkg/config"
	"github.com/vkhoa145/golang-bookings/pkg/models"
)

var functions = template.FuncMap{}

var app *config.AppConfig

// NewTemplates stes th config for the template package
func NewTemplates(a *config.AppConfig) {
	app = a
}

func AddDefaultData(td *models.TemplateData) *models.TemplateData {
	return td
}

func RenderTemplate(w http.ResponseWriter, html string, td *models.TemplateData) {
	var tc map[string]*template.Template

	if app.UseCache {
		// get template cache from app config
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemaplateCache()

	}

	// create a template cache
	// tc, err := CreateTemaplateCache()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// get requested template from cache
	t, ok := tc[html]
	if !ok {
		log.Fatal("could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	td = AddDefaultData(td)

	_ = t.Execute(buf, td)

	// render the template
	_, err := buf.WriteTo(w)
	if err != nil {
		log.Fatal(err)
	}
}

func CreateTemaplateCache() (map[string]*template.Template, error) {
	// myCache := make(map[string]*template.Template)
	myCache := map[string]*template.Template{}

	// get all of the files named *page.html from ./templates
	pages, err := filepath.Glob("./templates/*.page.html")

	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.html

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).ParseFiles(page)

		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}

// func RenderTemplate(w http.ResponseWriter, html string) {
// 	parsedTemplate, _ := template.ParseFiles("./templates/"+html, "./templates/base.layout.html")
// 	// parsedTemplate, _ := template.ParseFiles("./templates/base.layout.html", "./templates/"+html)
// 	err := parsedTemplate.Execute(w, nil)

// 	if err != nil {
// 		fmt.Println("error parsing template:", err)
// 		return
// 	}
// }

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var html *template.Template
// 	var err error

// 	// check to see if we already have the template in our cache
// 	_, inMap := tc[t]
// 	if !inMap {
// 		// need to create the template
// 		log.Println("creating template and adding to cache")
// 		err = createTemaplateCache(t)

// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {
// 		// we have the template in the cache
// 		log.Println("using cached template")
// 	}

// 	html = tc[t]

// 	err = html.Execute(w, nil)

// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemaplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.html",
// 		"./templates/product.layout.html",
// 	}

// 	// parse the template
// 	html, err := template.ParseFiles(templates...)

// 	if err != nil {
// 		return err
// 	}

// 	// add template to cache (map)

// 	tc[t] = html

// 	return nil
// }
