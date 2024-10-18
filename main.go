package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/peacegauge/lenslocked/controllers"
	"github.com/peacegauge/lenslocked/templates"
	"github.com/peacegauge/lenslocked/views"
)

func main() {
	const port = ":5000"
	r := chi.NewRouter()

	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "layout-parts.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "layout-parts.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "layout-parts.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found, Opppps!", 404)
	})

	fmt.Println("Running server on port", port)
	http.ListenAndServe(port, r)
}
