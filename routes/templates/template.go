package templates

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
)

func RenderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	cwd, _ := os.Getwd()
	fmt.Println("cwd", cwd)
	fmt.Println("filepath.Join( cwd, ./routes/"+tmpl+"/ + tmpl + .html)", filepath.Join(cwd, "./routes/"+tmpl+"/"+tmpl+".html"))
	t, err := template.ParseFiles(filepath.Join(cwd, "./routes/"+tmpl+"/"+tmpl+".html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
