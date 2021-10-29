package learngoweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServer_handler(t *testing.T) {
	var handler http.HandlerFunc = func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "Hello World")
	}
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServer_serverMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "Hello World")
	})
	mux.HandleFunc("/about/", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "Halaman About")
	})
	mux.HandleFunc("/about/profile/", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "Halaman Profile")
	})
	mux.HandleFunc("/about/contact/", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "Halaman Contact")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServer_request(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, request.Method)
		fmt.Fprint(writter, request.RequestURI)
	})
	mux.HandleFunc("/about/", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "Halaman About")
	})
	mux.HandleFunc("/about/profile/", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "Halaman Profile")
	})
	mux.HandleFunc("/about/contact/", func(writter http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writter, "Halaman Contact")
	})

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
