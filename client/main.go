package main

import (
	"context"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
)

var (
	Root     *template.Template
	NotFound *template.Template
)

func init() {
	var err error

	Root, err = template.ParseFiles("./web/templates/root.html")
	if err != nil {
		panic(err)
	}

	NotFound, err = template.ParseFiles("./web/templates/notfound.html")
	if err != nil {
		panic(err)
	}
}

func main() {
	var err error

	router := http.NewServeMux()

	webRouter := mux.NewRouter()

	webRouter.HandleFunc("/", handleRoot())
	webRouter.NotFoundHandler = handleNotFound()

	router.Handle("/", webRouter)

	staticHandler := http.StripPrefix(
		"/static/",
		http.FileServer(http.Dir("./web/static")),
	)

	router.Handle("/static/", staticHandler)

	srv := http.Server{
		Addr:    ":8089",
		Handler: router,
	}

	log.Println("Start apiserver, port :8089")

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

// Handel
func handleRoot() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		Root.Execute(w, nil)
	}
}

func handleNotFound() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		NotFound.Execute(w, nil)
	}
}
