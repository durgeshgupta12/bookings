package handler

import (
	"net/http"

	"github.com/durgeshgupta12/bookings/pkg/config"
	"github.com/durgeshgupta12/bookings/pkg/models"
	"github.com/durgeshgupta12/bookings/pkg/render"
)

// Repo the repository used by handlers
var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// new repo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// newhandlers sets repository for handlers
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.hbs", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again!"
	render.RenderTemplate(w, "about.hbs", &models.TemplateData{
		StringMap: stringMap,
	})
}
