package learngoweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Ifstet struct {
	Title string
	Name  string
}

func TemplateAction(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))

	t.ExecuteTemplate(w, "if.gohtml", Ifstet{
		Title: "Template Action",
	})
}

func TestTemplateAction(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080", nil)
	res := httptest.NewRecorder()

	TemplateAction(res, req)
	result := res.Result()

	body, _ := io.ReadAll(result.Body)
	fmt.Println(string(body))
}

func TemplateActionPerbandingan(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))

	t.ExecuteTemplate(w, "comparator.gohtml", map[string]int32{
		"FinalVal": 50,
	})
}

func TestTemplateActionPerbandingan(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080", nil)
	res := httptest.NewRecorder()

	TemplateActionPerbandingan(res, req)
	result := res.Result()

	body, _ := io.ReadAll(result.Body)
	fmt.Println(string(body))
}

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))

	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Title": "Template Action Range",
		"Hobi": []string{
			"Memancing",
			"Bernyanyi",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080", nil)
	res := httptest.NewRecorder()

	TemplateActionRange(res, req)
	result := res.Result()

	body, _ := io.ReadAll(result.Body)
	fmt.Println(string(body))
}

func TemplateActionWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/address.gohtml"))

	t.ExecuteTemplate(w, "address.gohtml", map[string]interface{}{
		"Title": "Template Action With",
		"Name":  "Gozenx",
		"Address": map[string]interface{}{
			"Street": "Jati, Sindangrasa",
			"City":   "Ciamis",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080", nil)
	res := httptest.NewRecorder()

	TemplateActionWith(res, req)
	result := res.Result()

	body, _ := io.ReadAll(result.Body)
	fmt.Println(string(body))
}
