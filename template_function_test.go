package learngoweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
}

func (myPage MyPage) SayHello(name string) string {
	return "Hello " + name + ", My Name Is " + myPage.Name
}

func TemplateFunction(req *http.Request, res http.ResponseWriter) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ .SayHello "Goz" }}`))

	t.ExecuteTemplate(res, "FUNCTION", MyPage{
		Name: "Zenklot",
	})
}

func TestTemplateFunction(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	res := httptest.NewRecorder()

	TemplateFunction(req, res)

	data, _ := io.ReadAll(res.Result().Body)
	fmt.Println(string(data))
}

func TemplateFunctionGlobal(req *http.Request, res http.ResponseWriter) {
	t := template.Must(template.New("FUNCTION").Parse(`{{ len .Name }}`))

	t.ExecuteTemplate(res, "FUNCTION", MyPage{
		Name: "Zenklot",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	res := httptest.NewRecorder()

	TemplateFunctionGlobal(req, res)

	data, _ := io.ReadAll(res.Result().Body)
	fmt.Println(string(data))
}

func TemplateFunctionCreateGlobal(req *http.Request, res http.ResponseWriter) {

	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{ upper .Name }}`))

	t.ExecuteTemplate(res, "FUNCTION", MyPage{
		Name: "Zenklot",
	})
}

func TestTemplateFunctionCreateGlobal(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	res := httptest.NewRecorder()

	TemplateFunctionCreateGlobal(req, res)

	data, _ := io.ReadAll(res.Result().Body)
	fmt.Println(string(data))
}

func TemplateFunctionPipeline(req *http.Request, res http.ResponseWriter) {

	t := template.New("FUNCTION")
	t = t.Funcs(map[string]interface{}{
		"sayHello": func(name string) string {
			return "Hello " + name
		},
		"upper": func(value string) string {
			return strings.ToUpper(value)
		},
	})

	t = template.Must(t.Parse(`{{ sayHello .Name | upper }}`))

	t.ExecuteTemplate(res, "FUNCTION", MyPage{
		Name: "Zenklot",
	})
}

func TestTemplateFunctionPipeline(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "localhost:8080", nil)
	res := httptest.NewRecorder()

	TemplateFunctionPipeline(req, res)

	data, _ := io.ReadAll(res.Result().Body)
	fmt.Println(string(data))
}
