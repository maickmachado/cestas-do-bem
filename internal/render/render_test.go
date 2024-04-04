package render

import (
	"net/http"
	"testing"

	"github.com/maickmachado/cestas-do-bem/internal/models"
)

func TestAddDefaultData(t *testing.T) {

	var td models.TemplateData

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	sessionManager.Put(r.Context(), "flash", "123")

	result := AddDefaultData(&td, r)

	if result.Flash != "123" {
		t.Error("Flash value of 123 not found")
	}

}

func TestRenderTemplate(t *testing.T) {
	// ./../../ is beacause you put the path were you are running it - it's running where the render package is
	pathToTemplates = "./../../templates"
	tc, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}

	app.TemplateCache = tc

	var w myWriter

	r, err := getSession()
	if err != nil {
		t.Error(err)
	}

	err = RenderTemplate(&w, r, "home.page.tmpl", &models.TemplateData{})
	if err != nil {
		t.Error("error rendering template to browser")
	}

	err = RenderTemplate(&w, r, "non-existent.page.tmpl", &models.TemplateData{})
	if err == nil {
		t.Error("rendered template that does not exist")
	}

}

func getSession() (*http.Request, error) {

	r, err := http.NewRequest("GET", "/some-url", nil)

	if err != nil {
		return nil, err
	}

	ctx := r.Context()
	ctx, _ = sessionManager.Load(ctx, r.Header.Get("X-Session"))
	r = r.WithContext(ctx)

	return r, nil
}

func TestNewTemplates(t *testing.T) {
	NewTemplates(app)
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplates = "./../../templates"
	_, err := CreateTemplateCache()
	if err != nil {
		t.Error(err)
	}
}
