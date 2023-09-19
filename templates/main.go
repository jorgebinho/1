package main

import (
	"html/template"
	"os"
	"strings"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// t := template.Must(template.New("content.html").ParseFiles(templates...))
	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
	t = template.Must(t.ParseFiles(templates...))
	err := t.Execute(os.Stdout, Cursos{
		{"Go", 40},
		{"JavaScript", 100},
		{"Python", 20},
	})
	if err != nil {
		panic(err)
	}
	// })
	// http.ListenAndServe(":8282", nil)
}
