package learngoweb

import (
	"embed"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"text/template"
)

//go:embed templates/*.gohtml
var templates embed.FS

var myTemplates = template.Must(template.ParseFS(templates, "templates/*.gohtml"))

func TemplateCaching(req *http.Request, res http.ResponseWriter) {
	myTemplates.ExecuteTemplate(res, "simple.gohtml", "Hello HTML Template")
}

func TestTemplateCaching(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	res := httptest.NewRecorder()

	TemplateCaching(req, res)

	data, _ := io.ReadAll(res.Result().Body)
	fmt.Println(string(data))
}
