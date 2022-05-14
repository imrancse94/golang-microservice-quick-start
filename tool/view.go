package tool

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

// View Render specific view
func View(path string, w http.ResponseWriter, content *interface{}) {
	viewPath := GetDynamicPath(fmt.Sprintf("assets/view/%s", path))
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	t, err := template.ParseFiles(viewPath)
	if err != nil {
		log.Fatal(err)
	}

	_ = t.Execute(w, content)
}
