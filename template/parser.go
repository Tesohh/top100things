package template

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func Parse(w http.ResponseWriter, tn string, data any) error {
	path := fmt.Sprintf("template/src/%v.html", tn)
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	templ, err := template.New(path).Parse(string(file))
	if err != nil {
		return err
	}
	templ.Execute(w, data)
	return nil
}
