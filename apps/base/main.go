package base

import (
    "html/template"
    "net/http"
)

var index = template.Must(template.ParseFiles(
  "static/templates/base.html",
  "static/templates/index.html",
))

var page404 = template.Must(template.ParseFiles(
  "static/templates/base.html",
  "static/templates/static_pages/404.html",
))

func init() {
    http.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
    if err := index.Execute(w, nil); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        if err := page404.Execute(w, nil); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
    }
}
