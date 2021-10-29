package learngoweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, req *http.Request) {
	if req.URL.Query().Get("name") != "" {
		http.ServeFile(writer, req, "./resources/index.html")
	} else {
		http.ServeFile(writer, req, "./resources/notfound.html")
	}
}

func TestServeFile(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

//go:embed resources/index.html
var resourceOk string

//go:embed resources/notfound.html
var resourcesNotfound string

func ServeFileEmbed(writer http.ResponseWriter, req *http.Request) {
	if req.URL.Query().Get("name") != "" {
		fmt.Fprint(writer, resourceOk)
	} else {
		fmt.Fprint(writer, resourcesNotfound)
	}
}

func TestServeFileEmbed(t *testing.T) {
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
