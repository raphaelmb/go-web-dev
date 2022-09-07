package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/raphaelmb/go-web-dev/pkg/config"
	"github.com/raphaelmb/go-web-dev/pkg/handlers"
	"github.com/raphaelmb/go-web-dev/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf("Starting server at port %s\n", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
