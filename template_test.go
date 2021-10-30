package learngoweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SimpleHTML(w http.ResponseWriter, r *http.Request) {
	templateText := `<html><body>{{.}}</body></html>`
	// t, err := template.New("SIMPLE").Parse(templateText)
	// if err != nil {
	// 	panic(err)
	// }
	t := template.Must(template.New("SIMPLE").Parse(templateText))
	er := t.ExecuteTemplate(w, "SIMPLE", "Hallo HTML Template")
	if er != nil {
		panic(er)
	}

}

func TestSimpleHTML(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080", nil)
	res := httptest.NewRecorder()

	SimpleHTML(res, req)
	result := res.Result()

	body, _ := io.ReadAll(result.Body)
	fmt.Println(string(body))
}

func SimpleHTMLFile(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/simple.gohtml"))

	er := t.ExecuteTemplate(w, "simple.gohtml", "Hallo HTML Template")
	if er != nil {
		panic(er)
	}

}

func TestSimpleHTMLFile(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080", nil)
	res := httptest.NewRecorder()

	SimpleHTMLFile(res, req)
	result := res.Result()

	body, _ := io.ReadAll(result.Body)
	fmt.Println(string(body))
}

func TemplateDirectory(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseGlob("./templates/*.gohtml"))

	er := t.ExecuteTemplate(w, "simple.gohtml", "Hallo HTML Template")
	if er != nil {
		panic(er)
	}

}

func TestTemplateDIR(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080", nil)
	res := httptest.NewRecorder()

	TemplateDirectory(res, req)
	result := res.Result()

	body, _ := io.ReadAll(result.Body)
	fmt.Println(string(body))
}

//go:embed templates/*.gohtml
var templateEmbed embed.FS

func TemplateEmbed(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFS(templateEmbed, "templates/*.gohtml"))

	er := t.ExecuteTemplate(w, "simple.gohtml", "Hallo HTML Template")
	if er != nil {
		panic(er)
	}

}

func TestTemplateEmbed(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080", nil)
	res := httptest.NewRecorder()

	TemplateEmbed(res, req)
	result := res.Result()

	body, _ := io.ReadAll(result.Body)
	fmt.Println(string(body))
}
