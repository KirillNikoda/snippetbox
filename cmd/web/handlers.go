package main

import (
	"errors"
	"fmt"
	"knikoda/snippetbox/pkg/models"
	"net/http"
	"strconv"
)

func (a *app) home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		a.notFound(w)
		return
	}
	s, err := a.snippets.Latest()
	if err != nil {
		a.serverError(w, err)
		return
	}
	for _, snippet := range s {
		fmt.Fprintf(w, "%v\n", snippet)
	}

	// files := []string{
	// 	"./ui/html/home.page.gohtml",
	// 	"./ui/html/base.layout.gohtml",
	// 	"./ui/html/footer.partial.gohtml",
	// }

	// ts, err := template.ParseFiles(files...)
	// fmt.Println(ts.Name())
	// if err != nil {
	// 	a.serverError(w, err)
	// 	return
	// }

	// err = ts.Execute(w, nil)
	// if err != nil {
	// 	a.serverError(w, err)
	// }
}

func (a *app) showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		a.notFound(w)
		return
	}

	s, err := a.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			a.notFound(w)
			return
		}
		a.serverError(w, err)
	}

	fmt.Fprintf(w, "%v", s)
}

func (a *app) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		a.clientError(w, http.StatusMethodNotAllowed)
		return
	}

	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := "7"

	id, err := a.snippets.Insert(expires, title, content)
	if err != nil {
		a.serverError(w, err)
	}

	http.Redirect(w, r, fmt.Sprintf("snippet?id=id=%d", id), http.StatusSeeOther)

	w.Write([]byte("Create a new snippet..."))
}
