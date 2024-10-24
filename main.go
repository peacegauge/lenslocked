package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/peacegauge/lenslocked/controllers"
	"github.com/peacegauge/lenslocked/models"
	"github.com/peacegauge/lenslocked/templates"
	"github.com/peacegauge/lenslocked/views"
)

func main() {
	const port = ":5000"
	r := chi.NewRouter()

	//connecting to db
	cfg := models.DefaultPostgresConfig()

	db, err := models.Open(cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected")

	userService := models.UserService{
		DB: db,
	}

	//Home handler
	tpl := views.Must(views.ParseFS(templates.FS, "home.gohtml", "tailwind.gohtml"))
	r.Get("/", controllers.StaticHandler(tpl))

	//Contact handler
	tpl = views.Must(views.ParseFS(templates.FS, "contact.gohtml", "tailwind.gohtml"))
	r.Get("/contact", controllers.StaticHandler(tpl))

	//Faq handler
	tpl = views.Must(views.ParseFS(templates.FS, "faq.gohtml", "tailwind.gohtml"))
	r.Get("/faq", controllers.FAQ(tpl))

	//Signup handler
	usersC := controllers.Users{
		UserService: &userService,
	}
	usersC.Templates.New = views.Must(views.ParseFS(templates.FS, "signup.gohtml", "tailwind.gohtml"))
	r.Get("/signup", usersC.New)

	//signup new
	r.Get("/signup", usersC.New)
	r.Post("/users", usersC.Create)

	//NotFound handler
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found, Opppps!", 404)
	})

	fmt.Println("Running server on port", port)
	http.ListenAndServe(port, r)
}
