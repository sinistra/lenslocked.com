package controllers

import "sinistra/lenslocked.com/views"

type Static struct {
	Home    *views.View
	Contact *views.View
	Faq     *views.View
}

func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
		Faq:     views.NewView("bootstrap", "static/faq"),
	}
}
