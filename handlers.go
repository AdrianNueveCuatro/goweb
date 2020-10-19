package main

import (
	"encoding/json"
	"net/http"
	"path"
)

func find(x string) int {
	if x == "" {
		return -1
	}
	for i, volcano := range volcanoList {
		if x == volcano.Number {
			return i
		}
	}
	return -2
}

func handleGet(w http.ResponseWriter, r *http.Request) (err error) {
	// Lee el id y lo busca en el diccionario
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 {
		return
	}
	// Obtiene toda la informacion en caso de no ingresar id
	dataJson, err := json.Marshal(append(volcanoList[1:]))
	if i == -2 {
		w.Header().Set("Content-Type", "application/json")
		w.Write(dataJson)
		return
	}

	// Imprime el libro especificado en el url
	dataJson2, err := json.Marshal(volcanoList[i])
	w.Header().Set("Content-Type", "application/json")
	w.Write(dataJson2)
	return
}

func handlePut(w http.ResponseWriter, r *http.Request) (err error) {
	//leer json y guardar en variable
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	book := Volcano{}
	json.Unmarshal(body, &book)

	//guarda la info en el diccionario
	volcanoList = append(volcanoList, book)
	w.WriteHeader(200)
	return
}

func handlePost(w http.ResponseWriter, r *http.Request) (err error) {
	//leer json y guardar en variable
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	volcano := Volcano{}
	json.Unmarshal(body, &volcano)

	// Lee el id y lo busca en el diccionario
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 || i == -2 {
		return
	}

	//modifica el libro especifico en el diccionario
	if volcano.Number != "" {
		volcanoList[i].Number = volcano.Number
	}
	if volcano.Name != "" {
		volcanoList[i].Name = volcano.Name
	}
	if volcano.Country != "" {
		volcanoList[i].Country = volcano.Country
	}
	if volcano.Region != "" {
		volcanoList[i].Region = volcano.Region
	}
	if volcano.Type != "" {
		volcanoList[i].Type = volcano.Type
	}
	if volcano.ActivityEvidence != "" {
		volcanoList[i].ActivityEvidence = volcano.ActivityEvidence
	}
	if volcano.LastKnownEruption != "" {
		volcanoList[i].LastKnownEruption = volcano.LastKnownEruption
	}
	if volcano.Latitude != "" {
		volcanoList[i].Latitude = volcano.Latitude
	}
	if volcano.Longitude != "" {
		volcanoList[i].Longitude = volcano.Longitude
	}
	if volcano.Elevation != "" {
		volcanoList[i].Elevation = volcano.Elevation
	}
	if volcano.DominantRockType != "" {
		volcanoList[i].DominantRockType = volcano.DominantRockType
	}
	if volcano.TectonicSetting != "" {
		volcanoList[i].TectonicSetting = volcano.TectonicSetting
	}
	w.WriteHeader(200)
	return
}

func handleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	// Lee el id y lo busca en el diccionario
	id := path.Base(r.URL.Path)
	checkError("Parse error", err)
	i := find(id)
	if i == -1 || i == -2 {
		return
	}
	//guarda la info en el diccionario
	volcanoList = append(volcanoList[:i], volcanoList[i+1:]...)
	w.WriteHeader(200)
	return
}
