package learngoweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateLayput(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/header.gohtml", "./templates/footer.gohtml", "./templates/layout.gohtml"))

	t.ExecuteTemplate(w, "layout", map[string]interface{}{
		"Title": "Templating Layout",
		"Name":  "Gozenx",
	})
}

func TestTemplateLayout(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080", nil)
	res := httptest.NewRecorder()

	TemplateLayput(res, req)
	result := res.Result()

	body, _ := io.ReadAll(result.Body)
	fmt.Println(string(body))
}
