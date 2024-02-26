package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/maickmachado/cestas-do-bem/internal/config"
	"github.com/maickmachado/cestas-do-bem/internal/handlers"
	"github.com/maickmachado/cestas-do-bem/internal/middlewares"
)

func Routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middlewares.NoSurf)
	mux.Use(middlewares.SessionLoad)

	mux.Get("/", handlers.Repo.Home)

	mux.Get("/about", handlers.Repo.About)
	mux.Get("/general", handlers.Repo.General)
	mux.Get("/major", handlers.Repo.Major)
	mux.Get("/contact", handlers.Repo.Contact)

	mux.Get("/beneficiary-register", handlers.Repo.BeneficiaryRegister)
	mux.Post("/beneficiary-register", handlers.Repo.PostBeneficiaryRegister)
	mux.Get("/beneficiary-sumary", handlers.Repo.BeneficiarySumary)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
