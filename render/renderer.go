package render

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func Template(w http.ResponseWriter, tn string, data any) error {
	paths := []string{
		"render/src/+components.html",
		"render/src/+layout.html",
		fmt.Sprintf("render/src/%v.html", tn),
	}
	parseMe := ""
	for _, p := range paths {
		file, err := os.ReadFile(p)
		if err != nil {
			return err
		}

		parseMe += string(file[:]) + "\n"
	}

	templ, err := template.New(tn).Funcs(funcMap).Parse(parseMe)
	if err != nil {
		return err
	}
	err = templ.Execute(w, data)
	if err != nil {
		return err
	}
	return nil
}

var funcMap = template.FuncMap{
	"safeURL": func(u string) template.URL { return template.URL(u) },
}
