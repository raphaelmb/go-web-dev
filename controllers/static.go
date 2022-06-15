package controllers

import (
	"net/http"

	"github.com/raphaelmb/go-web-dev/views"
)

func StaticHandlder(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}