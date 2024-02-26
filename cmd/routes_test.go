package setup_test

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi"
	"github.com/maickmachado/cestas-do-bem/internal/config"
	"github.com/maickmachado/cestas-do-bem/internal/routes"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig

	mux := routes.Routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing
	default:
		t.Error(fmt.Printf("TestRoutes, Don't know type %T\n", v))
	}
}
