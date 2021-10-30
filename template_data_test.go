package learngoweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(w, "name.gohtml", map[string]interface{}{
		"Title": "Template Data Map",
		"Name":  "Raisa",
		"Address": map[string]interface{}{
			"Street": "Jati, Sindangrasa Ciamis",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080", nil)
	res := httptest.NewRecorder()

	TemplateDataMap(res, req)
	result := res.Result()

	body, _ := io.ReadAll(result.Body)
	fmt.Println(string(body))
}

type Address struct {
	Street string
}
type Page struct {
	Title   string
	Name    string
	Address Address
}

func TemplateDataStruct(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))

	t.ExecuteTemplate(w, "name.gohtml", Page{
		Title: "Template Data Map",
		Name:  "Raisa",
		Address: Address{
			Street: "Jati, Sindangrasa, Ciamis",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080", nil)
	res := httptest.NewRecorder()

	TemplateDataStruct(res, req)
	result := res.Result()

	body, _ := io.ReadAll(result.Body)
	fmt.Println(string(body))
}
