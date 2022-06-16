package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphaelmb/go-web-dev/controllers"
	"github.com/raphaelmb/go-web-dev/templates"
	"github.com/raphaelmb/go-web-dev/views"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.StaticHandlder(views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))))

	r.Get("/contact", controllers.StaticHandlder(views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))))

	r.Get("/faq", controllers.FAQ(views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	
	fmt.Println("Starting server on port 3000.")
	http.ListenAndServe(":3000", r)
}
