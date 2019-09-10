package controllers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"sinistra/lenslocked.com/context"
	"sinistra/lenslocked.com/models"
	"sinistra/lenslocked.com/views"
	"strconv"
)

func NewGalleries(gs models.GalleryService) *Galleries {
	return &Galleries{
		New:      views.NewView("bootstrap", "galleries/new"),
		ShowView: views.NewView("bootstrap", "galleries/show"),
		gs:       gs,
	}
}

type Galleries struct {
	New      *views.View
	ShowView *views.View
	gs       models.GalleryService
}

type GalleryForm struct {
	Title string `schema:"title"`
}

// POST /galleries
func (g *Galleries) Create(w http.ResponseWriter, r *http.Request) {
	var vd views.Data
	var form GalleryForm
	if err := parseForm(r, &form); err != nil {
		vd.SetAlert(err)
		g.New.Render(w, vd)
		return
	}

	user := context.User(r.Context())
	gallery := models.Gallery{
		Title:  form.Title,
		UserID: user.ID,
	}
	if err := g.gs.Create(&gallery); err != nil {
		vd.SetAlert(err)
		g.New.Render(w, vd)
		return
	}
	fmt.Fprintln(w, gallery)
}

// GET /galleries/:id
func (g *Galleries) Show(w http.ResponseWriter, r *http.Request) {
	// First we get the vars like we saw earlier. We do this
	// so we can get variables from our path, like the "id"
	// variable.
	vars := mux.Vars(r)
	// Next we need to get the "id" variable from our vars.
	idStr := vars["id"]
	// Our idStr is a string, so we use the Atoi function
	// provided by the strconv package to convert it into an
	// integer. This function can also return an error, so we
	// need to check for errors and render an error.
	id, err := strconv.Atoi(idStr)
	if err != nil {
		// If there is an error we will return the StatusNotFound
		// status code, as the page requested is for an invalid
		// gallery ID, which means the page doesn't exist.
		http.Error(w, "Invalid gallery ID", http.StatusNotFound)
		return
	}
	// This line is only to prevent the compiler from complaining
	// about us not using the id parameter. We will use it later.
	_ = id
	// Finally we need to lookup the gallery with the ID we
	// have, but we haven't written that code yet! For now we
	// will create a temporary gallery to test our view.
	gallery := models.Gallery{
		Title: "A temporary fake gallery with ID: " + idStr,
	}
	// We will build the views.Data object and set our gallery
	// as the Yield field, but technically we do not need
	// to do this and could just pass the gallery into the
	// Render method because of the type switch we coded into
	// the Render method.
	var vd views.Data
	vd.Yield = gallery
	g.ShowView.Render(w, vd)
}
