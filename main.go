// main.go
package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

type Animal struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Especie   string `json:"especie"`
	Edad      string `json:"edad"`
	Dueno     string `json:"dueno"`
	Telefono  string `json:"telefono"`
	Direccion string `json:"direccion"`
	Barrio    string `json:"barrio"`
}

var Animales []Animal
var idCounter = 1

func main() {
	http.HandleFunc("/", serveTemplate)
	http.HandleFunc("/api/animales", handleAnimales)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Servidor corriendo en http://localhost:8084")
	http.ListenAndServe(":8084", nil)
}

func serveTemplate(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("template/index.html"))
	tmpl.Execute(w, nil)
}

func handleAnimales(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(Animales)
	case http.MethodPost:
		var a Animal
		json.NewDecoder(r.Body).Decode(&a)
		a.ID = idCounter
		idCounter++
		Animales = append(Animales, a)
		json.NewEncoder(w).Encode(a)
	case http.MethodPut:
		var updated Animal
		json.NewDecoder(r.Body).Decode(&updated)
		for i, a := range Animales {
			if a.ID == updated.ID {
				Animales[i] = updated
				break
			}
		}
		json.NewEncoder(w).Encode(updated)
	case http.MethodDelete:
		idStr := r.URL.Query().Get("id")
		id, _ := strconv.Atoi(idStr)
		for i, a := range Animales {
			if a.ID == id {
				Animales = append(Animales[:i], Animales[i+1:]...)
				break
			}
		}
		json.NewEncoder(w).Encode(map[string]string{"status": "deleted"})
	}
}
