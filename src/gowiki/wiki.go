package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Page struct {
	Title string
	Body  []byte
}

//method that save a content of Page struct into a text file
func (p *Page) save() error {
	filename := "data/" + p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//get the content of a Page struct from a text file
func getPageData(title string) (*Page, error) {
	filename := "data/" + title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func deletePageData(title string) error {
	filename := "data/" + title + ".txt"
	err := os.Remove(filename)
	if err != nil {
		return err
	}
	return nil
}

func homePageView(w http.ResponseWriter, r *http.Request) {
	files, _ := ioutil.ReadDir("data")
	var pages []*Page
	for _, f := range files {
		p, _ := getPageData(f.Name()[:len(f.Name())-len(filepath.Ext(f.Name()))])
		pages = append(pages, p)
	}
	t, _ := template.ParseFiles("template/home.html")
	t.Execute(w, pages)
}

func createPageView(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	if r.Method == "POST" {
		body, _ := io.ReadAll(r.Body)
		fmt.Println(body)
	}
	t, _ := template.ParseFiles("template/create.html")
	t.Execute(w, r)
}

func detailPageView(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/detail/"):]
	p, err := getPageData(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	t, _ := template.ParseFiles("template/detail.html")
	t.Execute(w, p)
}

func editPageView(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	if title != "" {
		p, err := getPageData(title)
		if err != nil {
			p = &Page{Title: title}
		}
		t, _ := template.ParseFiles("template/edit.html")
		t.Execute(w, p)
	} else {
		http.Redirect(w, r, "/home", http.StatusFound)
	}
}

//Handler for Saving Page
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	if title != "" && body != "" {
		p := &Page{Title: title, Body: []byte(body)}
		err := p.save()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

//Handler for Deleting Page
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/delete/"):]
	if title != "" {
		err := deletePageData(title)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	http.Redirect(w, r, "/home/", http.StatusFound)
}

func main() {
	_, err_data := ioutil.ReadDir("data")
	if err_data != nil {
		os.Mkdir("data", 0755)
	}

	http.HandleFunc("/detail/", detailPageView)
	http.HandleFunc("/edit/", editPageView)
	http.HandleFunc("/home/", homePageView)
	http.HandleFunc("/create/", createPageView)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/delete/", deleteHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
