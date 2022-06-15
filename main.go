package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/raphaelmb/go-web-dev/controllers"
	"github.com/raphaelmb/go-web-dev/views"
)

func main() {
	r := chi.NewRouter()
	
	r.Get("/", controllers.StaticHandlder(views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))))

	r.Get("/contact", controllers.StaticHandlder(views.Must(views.Parse(filepath.Join("templates", "contact.gohtml"))))) 

	r.Get("/faq", controllers.StaticHandlder(views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound) 
	})
	fmt.Println("Starting server on port 3000.")
	http.ListenAndServe(":3000", r)
}