package handlers

import (
	"encoding/gob"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/maickmachado/cestas-do-bem/internal/config"
	"github.com/maickmachado/cestas-do-bem/internal/middlewares"
	"github.com/maickmachado/cestas-do-bem/internal/models"
	"github.com/maickmachado/cestas-do-bem/internal/render"
)

var functions = template.FuncMap{}

// var app config.AppConfig
var app = &config.App
var sessionManager *scs.SessionManager
var pathToTemplates = "./../../templates"

// TODO: change to TestMain
func getRoutes() http.Handler {
	//what am I going to put in session
	gob.Register(models.Register{})

	//change it to true when in production
	app.InProduction = false

	// Initialize a new session manager and configure the session lifetime.
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction
	// add the session to the AppConfig
	app.Session = sessionManager

	tc, err := CreateTemplateCache()
	if err != nil {
		log.Fatal("error creating template cache")
	}

	app.TemplateCache = tc
	app.UseCache = true

	repo := NewRepository(app)
	NewHandlers(repo)
	render.NewTemplates(app)

	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	//mux.Use(middlewares.NoSurf)
	mux.Use(middlewares.SessionLoad)

	mux.Get("/", Repo.Home)

	mux.Get("/about", Repo.About)
	mux.Get("/general", Repo.General)
	mux.Get("/major", Repo.Major)
	mux.Get("/contact", Repo.Contact)

	mux.Get("/beneficiary-register", Repo.BeneficiaryRegister)
	mux.Post("/beneficiary-register", Repo.PostBeneficiaryRegister)
	mux.Get("/beneficiary-sumary", Repo.BeneficiarySumary)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	myCache := map[string]*template.Template{}

	// get all of the files named *.page.tmpl from ./templates
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.tmpl", pathToTemplates))
	if err != nil {
		return myCache, err
	}

	// range through all files ending with *.page.tmpl
	for _, page := range pages {
		//Base retorna somente a ultima parte do path ou seja o nome do template
		name := filepath.Base(page)
		//associa o path 'page' com o template 'name'
		//porque faço isso?
		ts, err := template.New(name).Funcs(functions).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			//associa oe layouts encontrados com a variavel 'ts'
			//entender como funciona o ParseGlob
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/*.layout.tmpl", pathToTemplates))
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts
	}

	return myCache, nil
}
