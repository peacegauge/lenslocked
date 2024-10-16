package main

import (
	"fmt"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	//Setting header
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>Welcome to the awesome vam</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Contact Page</h1><p>To get in touch, email me at <a href=\"mailto:teddy@example.com\">teddy@example.com</a></p>")
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	//Setting header
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprintf(w, "<h1>FAQ Section</h1><p>Q: Is there a free version?</p><p>A: Yes! We offer a  free trial for 30 days</p><br><p>Q: What are your support hours?</p><p>A: We have support staff answering emails 24/7, though times may be a bit slower on weekends.</p><br><p>Q: How do I contact support?</p><p>A: Email us - support@lenslocked.com</p><br>")
}

// func pathHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.URL.Path {
// 	case "/":
// 		homeHandler(w, r)
// 	case "/contact":
// 		contactHandler(w, r)
// 	default:
// 		http.Error(w, "Page not found, Opppps!", 404)
// 	}
// }

type Router struct{}

func (router Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		homeHandler(w, r)
	case "/contact":
		contactHandler(w, r)
	case "/faq":
		faqHandler(w, r)
	default:
		http.Error(w, "Page not found, Opppps!", 404)
	}
}

func main() {
	const port = ":5000"
	var router Router

	fmt.Println("Running server on port", port)
	http.ListenAndServe(port, router)
}
