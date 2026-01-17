package main

import (
	// "fmt"
	"embed"
	readingfiles "go-playground/17_reading_files"
	"html/template"
	"io"
	"io/fs"
	"os"
	"path/filepath"
)

// var (
// 	// postTemplate = `<h1>{{.Title}}</h1><p>{{.Description}}</p>Tags: <ul>{{range .Tags}}<li>{{.}}</li>{{end}}</ul>`
// postTemplates embed.FS
// )

// goembed comment is important , maintain syntax
var (
	//go:embed "templates/*"
	postTemplates embed.FS
)

func main() {
	// entry point
	filepath.WalkDir(".", func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		println("path:", path)
		return nil
	})

}
func Render(w io.Writer, p readingfiles.Post) error {
	/*
		_, err := fmt.Fprintf(w, "<h1>%s</h1>\n<p>%s</p>\n", p.Title, p.Description)
		if err != nil {
			return err
		}

		_, err = fmt.Fprintf(w, "Tags: <ul>")
		if err != nil {
			return err
		}
		for _, t := range p.Tags {
			_, err = fmt.Fprintf(w, "<li>%s</li>", t)
			if err != nil {
				return err
			}

		}
		_, err = fmt.Fprintf(w, "</ul>")
		if err != nil {
			return err
		}
		return err

	*/

	// templ, err := template.New("blog").Parse(postTemplate)

	// 		fs.WalkDir(postTemplates, ".", func(path string, d fs.DirEntry, err error) error {
	// 			println("path below=========")
	// 	println(path)
	// 	return nil
	// })

	templ, err := template.ParseFS(postTemplates, "templates/*.gohtml")
	if err != nil {
		return err
	}

	if err := templ.Execute(w, p); err != nil {
		return err
	}

	return nil
}
