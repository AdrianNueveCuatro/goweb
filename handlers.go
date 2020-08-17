package main

import (
    "encoding/json"
    "net/http"
    "path"
)

func find(x string) int {
    for i, book := range books {
        if x == book.Id {
            return i
        }
    }
    return -1
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	// Lee el id y lo busca en el diccionario
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)

	// Obtiene toda la informacion en caso de no ingresar id
	dataJson, err := json.Marshal(books)
	if i == -1 {
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		return
	}

	// Imprime el libro especificado en el url
	dataJson2, err := json.Marshal(books[i])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson2)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	//leer json y guardar en variable
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	book := Book{}
	json.Unmarshal(body, &book)

	//guarda la info en el diccionario
	books = append(books, book)
	w.WriteHeader(200)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	//leer json y guardar en variable
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	book := Book{}
	json.Unmarshal(body, &book)

	// Lee el id y lo busca en el diccionario
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 {
		return
	}

	//modifica el libro especifico en el diccionario
	if book.Id != "" {
		books[i].Id = book.Id
	}
	if book.Title != "" {
		books[i].Title = book.Title
	}
	if book.Edition != "" {
		books[i].Edition = book.Edition
	}
	if book.Copyright != "" {
		books[i].Copyright = book.Copyright
	}
	if book.Language != "" {
		books[i].Language = book.Language
	}
	if book.Pages != "" {
		books[i].Pages = book.Pages
	}
	if book.Author != "" {
		books[i].Author = book.Author
	}
	if book.Publisher != "" {
		books[i].Publisher = book.Publisher
	}
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	// Lee el id y lo busca en el diccionario
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 {
		return
	}
	//guarda la info en el diccionario
	books = append(books[:i], books[i+1:]...)
	w.WriteHeader(200)
	return
}
