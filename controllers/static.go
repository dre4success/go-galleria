package controllers

import "galleria.com/views"

//NewStatic function to render static pages
func NewStatic() *Static {
	return &Static{
		HomeView:    views.NewView("bootstrap", "static/home"),
		ContactView: views.NewView("bootstrap", "static/contact"),
		FaqView:     views.NewView("bootstrap", "static/faq"),
	}
}

// Static types for View pages
type Static struct {
	HomeView    *views.View
	ContactView *views.View
	FaqView     *views.View
}
