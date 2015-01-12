// handlers
package main

import (
	//"fmt"
	"net/http"
	"text/template"
)

func index(w http.ResponseWriter, r *http.Request) {
	t1, _ := template.New("shit").Parse(`<a>{{.}}</a>`)
	for _, h := range hosts {
		t1.Execute(w, h.Ip)
	}
}
