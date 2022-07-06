package main

import (
	"encoding/json"
	entity "ex6/Entity"
	"ex6/db"
	"fmt"
	"io"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/go-chi/chi"
)

func main() {
	db.InitDatabase()
	router := chi.NewRouter()
	router.Route("/api", func(r chi.Router) {
		r.Post("/location", PostLocation)
		r.Get("/location", GetLocation)
	})
	http.ListenAndServe(":8080", router)
}

func PostLocation(w http.ResponseWriter, r *http.Request) {
	location := entity.Location{}
	body := r.Body
	decodedBody, _ := io.ReadAll(body)
	json.Unmarshal(decodedBody, &location)
	value := fmt.Sprintf("My location is %v, %v", location.Latitude, location.Longitude)
	db.GetDB().Create(&location)
	fmt.Fprintf(w, value)

}

func GetLocation(w http.ResponseWriter, r *http.Request) {
	city := r.URL.Query().Get("city")
	fmt.Fprint(w, city)
	var location = entity.Location{}
	db.GetDB().Where("city=?", city).Find(&location)
	fmt.Fprintln(w, location)
}
