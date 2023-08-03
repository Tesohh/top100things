package render

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func Template(w http.ResponseWriter, tn string, data any) error {
	headFile, err := os.ReadFile("render/layouts/head.html")
	if err != nil {
		return err
	}
	stylesFile, err := os.ReadFile("render/layouts/styles.html")
	if err != nil {
		return err
	}

	path := fmt.Sprintf("render/src/%v.html", tn)
	templFile, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	templ, err := template.New(path).Parse(fmt.Sprintf("%s\n%s\n%s", headFile, stylesFile, templFile))
	if err != nil {
		return err
	}
	templ.Execute(w, data)
	return nil
}
