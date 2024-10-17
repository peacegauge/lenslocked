package main

import (
	"html/template"
	"os"
)

func main() {
	t, err := template.ParseFiles("index.gohtml")
	if err != nil {
		panic(err)
	}

	user := struct {
		Name string
	}{
		Name: "Susan Tsichler",
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
