package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

type Person struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

type people []*Person

var People people

func init() {
	People = append(People,
		&Person{
			ID:      0,
			Name:    "Михаил",
			Surname: "Никишкин",
		},
		&Person{
			ID:      1,
			Name:    "",
			Surname: "Никишкин",
		},
		&Person{
			ID:      2,
			Name:    "Михаил",
			Surname: "Никишкин",
		},
		&Person{
			ID:      3,
			Name:    "Михаил",
			Surname: "Никишкин"},
		&Person{
			ID:      4,
			Name:    "Михаил",
			Surname: "Никишкин",
		},
	)
}

func main() {

	var err error

	router := mux.NewRouter()

	router.HandleFunc("/", handlePeopleGET()).Methods("GET")
	router.HandleFunc("/person", handlePersonGET()).Methods("GET")
	router.HandleFunc("/person", handlePersonPOST()).Methods("POST")
	router.HandleFunc("/person", handlePersonPUT()).Methods("PUT")
	router.HandleFunc("/person", handlePersonDELETE()).Methods("DELETE")

	srv := http.Server{
		Addr:    ":9543",
		Handler: router,
	}

	log.Println("Start apiserver, port :9543")

	idleConnsClosed := make(chan struct{})

	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		// setting a timeout to shut down the server
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(ctx); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err = srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
	<-idleConnsClosed

}

func handlePeopleGET() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("handlePeopleGET")
		responce(w, r, http.StatusOK, People)
	}
}

func handlePersonGET() http.HandlerFunc {

	type req struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("handlePersonGET")

		pid := &req{}

		if err := json.NewDecoder(r.Body).Decode(pid); err != nil {
			responceError(w, r, http.StatusBadRequest, err)
			return
		}

		peop := People.FindByID(pid.ID)

		if peop != nil {
			responce(w, r, http.StatusAccepted, peop)
		} else {
			responce(w, r, http.StatusNotFound, nil)
		}
	}
}

func handlePersonPOST() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("handlePersonPOST")

		per := &Person{}

		if err := json.NewDecoder(r.Body).Decode(per); err != nil {
			responceError(w, r, http.StatusBadRequest, err)
			return
		}

		if per.Name == "" || per.Surname == "" {
			responce(w, r, http.StatusUnauthorized, People)
			return
		}

		per.ID = len(People)
		People = append(People, per)

		responce(w, r, http.StatusAccepted, People)

	}
}

func handlePersonPUT() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("handlePersonPUT")

		per := &Person{}

		if err := json.NewDecoder(r.Body).Decode(per); err != nil {
			responceError(w, r, http.StatusBadRequest, err)
			return
		}

		if per.Name == "" || per.Surname == "" {
			responce(w, r, http.StatusUnauthorized, People)
			return
		}

		People.UpdatePerson(per)

		responce(w, r, http.StatusAccepted, People)

	}
}

func handlePersonDELETE() http.HandlerFunc {

	type req struct {
		ID int `json:"id"`
	}

	return func(w http.ResponseWriter, r *http.Request) {

		log.Println("handlePersonDELETE")

		pid := &req{}

		if err := json.NewDecoder(r.Body).Decode(pid); err != nil {
			responceError(w, r, http.StatusBadRequest, err)
			return
		}

		People.DeletePerson(pid.ID)
		responce(w, r, http.StatusOK, People)
	}
}

// responces functions
func responceError(w http.ResponseWriter, r *http.Request, code int, err error) {
	responce(w, r, code, map[string]string{"error :": err.Error()})
}

func responce(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

// People mathod type

func (p people) FindByID(id int) *Person {
	for _, val := range p {
		if val.ID == id {
			return val
		}
	}

	return nil
}

func (p people) UpdatePerson(per *Person) {
	for _, val := range p {
		if val.ID == per.ID {
			val.Name = per.Name
			val.Surname = per.Surname
			return
		}
	}
}

func (p people) DeletePerson(id int) {
	for i := 0; i < len(p); i++ {
		if p[i].ID == id {
			p[i] = nil
			return
		}
	}
}
