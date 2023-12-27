package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/maickmachado/cestas-do-bem/internal/config"
	"github.com/maickmachado/cestas-do-bem/internal/handlers"
	"github.com/maickmachado/cestas-do-bem/internal/render"
	"github.com/maickmachado/cestas-do-bem/internal/routes"
)

const portNumber = ":8080"

func main() {
	//var app config.AppConfig
	app := &config.App
	//change it to true when in production
	app.InProduction = false

	// Initialize a new session manager and configure the session lifetime.
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction
	// add the session to the AppConfig
	app.Session = sessionManager

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error creating template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepository(app)
	handlers.NewHandlers(repo)
	render.NewTemplates(app)

	fmt.Printf("Staring application on port %s\n", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes.Routes(app),
	}
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
