package handlers

import (
	"golang-hotel-booking-system/cmd/pkg/config"
	"golang-hotel-booking-system/cmd/pkg/models"
	"golang-hotel-booking-system/cmd/pkg/render"
	"net/http"
)

var Repo *Repository
type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "Hey this is a context data."
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: stringMap})
}