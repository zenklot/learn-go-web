package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(writer, "Hello World")
}

func TestHttp(t *testing.T) {
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	recorder := httptest.NewRecorder()

	HelloHandler(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(response.StatusCode)
	fmt.Println(response.Status)
	fmt.Println(string(body))
}

func SayHello(writer http.ResponseWriter, request *http.Request) {
	names := request.URL.Query()["name"]
	fmt.Fprint(writer, strings.Join(names, " "))
}

func TestQuery(t *testing.T) {
	request := httptest.NewRequest("GET", "localhost:8080/hello?name=gozenx&name=raisa&name=supriatna", nil)

	recorder := httptest.NewRecorder()

	SayHello(recorder, request)

	response := recorder.Result()

	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}
