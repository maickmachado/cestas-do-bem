package handlers

import (
	"log"
	"net/http"

	"github.com/maickmachado/cestas-do-bem/internal/config"
	"github.com/maickmachado/cestas-do-bem/internal/forms"
	"github.com/maickmachado/cestas-do-bem/internal/models"
	"github.com/maickmachado/cestas-do-bem/internal/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepository creates a new repository
func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{
		App: app,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(repo *Repository) {
	Repo = repo
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)
	stringMap := make(map[string]string)
	stringMap["test"] = "1 2 3 testando 1 2 3"

	// send data to the template
	render.RenderTemplate(w, r, "home.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) BeneficiaryRegister(w http.ResponseWriter, r *http.Request) {
	var emptyRegister models.Register
	data := make(map[string]interface{})
	data["register"] = emptyRegister

	stringMap := make(map[string]string)
	stringMap["teste"] = "1 2 3 testando 1 2 3"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "beneficiary-register.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
		Form:      forms.New(nil),
		Data:      data,
	})
}

func (m *Repository) PostBeneficiaryRegister(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	register := models.Register{
		InputName:         r.Form.Get("inputName"),
		InputStreet:       r.Form.Get("inputStreet"),
		InputStreetNumber: r.Form.Get("inputStreetNumber"),
		InputComp:         r.Form.Get("inputComp"),
		InputPhone:        r.Form.Get("inputPhone"),
	}

	form := forms.New(r.PostForm)

	form.Required("inputName", "inputStreet", "inputStreetNumber", "inputComp", "inputPhone")
	form.MinLength("inputName", 3, r)
	//TODO: adicionar um IsEmail aqui, um validador - aula 69 ou anterior

	if !form.Valid() {
		data := make(map[string]interface{})
		data["register"] = register

		render.RenderTemplate(w, r, "beneficiary-register.page.tmpl", &models.TemplateData{
			Form: form,
			Data: data,
		})
		return
	}

	m.App.Session.Put(r.Context(), "register", register)

	http.Redirect(w, r, "/beneficiary-sumary", http.StatusSeeOther)
}

func (m *Repository) BeneficiarySumary(w http.ResponseWriter, r *http.Request) {

	register, ok := m.App.Session.Get(r.Context(), "register").(models.Register)
	if !ok {
		log.Println("cannot get item from session")
		m.App.Session.Put(r.Context(), "error", "cant get register from session")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	m.App.Session.Remove(r.Context(), "register")

	data := make(map[string]interface{})
	data["register"] = register

	render.RenderTemplate(w, r, "beneficiary-sumary.page.tmpl", &models.TemplateData{
		Data: data,
	})

}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["teste"] = "1 2 3 testando 1 2 3"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) General(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["teste"] = "1 2 3 testando 1 2 3"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Major(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["teste"] = "1 2 3 testando 1 2 3"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {

	stringMap := make(map[string]string)
	stringMap["teste"] = "1 2 3 testando 1 2 3"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.RenderTemplate(w, r, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
