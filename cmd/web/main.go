package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/maickmachado/cestas-do-bem/internal/config"
	"github.com/maickmachado/cestas-do-bem/internal/handlers"
	"github.com/maickmachado/cestas-do-bem/internal/models"
	"github.com/maickmachado/cestas-do-bem/internal/render"
	"github.com/maickmachado/cestas-do-bem/internal/routes"
)

const portNumber = ":8080"

// var app config.AppConfig
var app = &config.App
var sessionManager *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {

	err := run()
	if err != nil {
		log.Fatal(err)
	}

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

func run() error {
	//what am I going to put in session
	gob.Register(models.Register{})

	//change it to true when in production
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	// Initialize a new session manager and configure the session lifetime.
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction
	// add the session to the AppConfig
	app.Session = sessionManager

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error creating template cache")
		return err
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepository(app)
	handlers.NewHandlers(repo)
	render.NewTemplates(app)

	return nil
}
