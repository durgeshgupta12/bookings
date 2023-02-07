package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/durgeshgupta12/bookings/internal/config"
	"github.com/durgeshgupta12/bookings/internal/models"
	"github.com/durgeshgupta12/bookings/internal/render"
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
	render.RenderTemplate(w, r, "home.hbs", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Again!"
	render.RenderTemplate(w, r, "about.hbs", &models.TemplateData{
		StringMap: stringMap,
	})
}

func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "make-reservation.hbs", &models.TemplateData{})
}
func (m *Repository) Generals(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "generals.hbs", &models.TemplateData{})
}
func (m *Repository) Majors(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "majors.hbs", &models.TemplateData{})
}
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "search-availability.hbs", &models.TemplateData{})
}
func (m *Repository) PostAvailability(w http.ResponseWriter, r *http.Request) {
	start := r.Form.Get("start")
	end := r.Form.Get("end")

	w.Write([]byte(fmt.Sprintf("start date is %s and end date is %s", start, end)))
}

type jsonResponse struct {
	OK      bool   `json:"ok"`
	Message string `json:"message"`
}

func (m *Repository) AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		OK:      true,
		Message: "Available",
	}

	out, err := json.MarshalIndent(resp, "", "     ")
	if err != nil {
		log.Println(err)
	}
	log.Println(string(out))
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, r, "contact.hbs", &models.TemplateData{})
}
