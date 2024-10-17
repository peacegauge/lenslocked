package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
)

func executeTemplate(w http.ResponseWriter, filepath string) {
	//Setting header
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	tpl, err := template.ParseFiles(filepath)
	if err != nil {
		log.Printf("parsing template: %v", err)
		http.Error(w, "There was an error parsing the template", http.StatusInternalServerError)
		return
	}

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Printf("executing template: %v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "home.gohtml")
	executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	tplPath := filepath.Join("templates", "contact.gohtml")
	executeTemplate(w, tplPath)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	//Setting header
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>FAQ Section</h1><p>Q: Is there a free version?</p><p>A: Yes! We offer a  free trial for 30 days</p><br><p>Q: What are your support hours?</p><p>A: We have support staff answering emails 24/7, though times may be a bit slower on weekends.</p><br><p>Q: How do I contact support?</p><p>A: Email us - support@lenslocked.com</p><br>")
}

func main() {
	const port = ":5000"

	r := chi.NewRouter()
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found, Opppps!", 404)
	})

	fmt.Println("Running server on port", port)
	http.ListenAndServe(port, r)
}
