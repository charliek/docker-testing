package main

import (
	"html/template"
	"net/http"
	"os"
)

var base_html = `
<html>
	<head>
		<title>{{template "title" .}}</title>
		<style>
			body {
				background-color: red;
				color: white;
			}
		</style>
	</head>
	<body>
		{{template "body" .}}
	</body>
</html>
`

var index_html = `
{{define "title"}}index{{end}}

{{define "body"}}
	<h1>Container Environment</h1>
	<ul>
	{{range .}}
		<li>{{.}}</li>
	{{end}}
	</ul>
{{end}}
`

func main() {

	base := template.Must(template.New("base").Parse(base_html))
	index := template.Must(base.New("index").Parse(index_html))

	rootHandler := func(w http.ResponseWriter, r *http.Request) {
		env := os.Environ()
		index.ExecuteTemplate(w, "base", env)
	}

	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":9090", nil)
}
