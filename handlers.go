package main

import (
	"encoding/json"
	"net/http"
	"path"
	"regexp"
)

var digitValidator = regexp.MustCompile(`^[0-9]+$`)

func find(x string) int {
	for i, book := range books {
		if x == book.Id {
			return i
		}
	}
	return -1
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	
	if digitValidator.MatchString(id){
		i := find(id)
		if i == -1 {
			w.WriteHeader(404)
			return
		}else{
			dataJson, _ := json.Marshal(books[i])
			w.Header().Set("Content-Type", "application/json")
			w.Write(dataJson)
		}
	}else{
		dataJson, _ := json.Marshal(books)
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		
	}
	w.WriteHeader(200)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	book := Book{}
	json.Unmarshal(body, &book)
	books = append(books, book)
	w.WriteHeader(200)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 {
		w.WriteHeader(404)
	}else{
		len := r.ContentLength
		body := make([]byte, len)
		r.Body.Read(body)
		book := Book{}
		json.Unmarshal(body, &book)
		savedBook := books[i]

		if book.Title != ""{
			savedBook.Title = book.Title
		}
		if book.Edition != ""{
			savedBook.Edition = book.Edition
		}
		if book.Copyright != ""{
			savedBook.Copyright = book.Copyright
		}
		if book.Language != ""{
			savedBook.Language = book.Language
		}
		if book.Pages != ""{
			savedBook.Pages = book.Pages
		}
		if book.Author != ""{
			savedBook.Author = book.Author
		}
		if book.Publisher != ""{
			savedBook.Publisher = book.Publisher
		}

		books[i] = savedBook

		w.WriteHeader(200)
	}
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 {
		w.WriteHeader(404)
	}else{
		 books = append(books[:i], books[i+1:]...)
		 w.WriteHeader(200)
	}
	
	return
}
