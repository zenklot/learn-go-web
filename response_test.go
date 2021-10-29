package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func ResponseCode(writer http.ResponseWriter, req *http.Request) {

	name := req.URL.Query().Get("name")

	if name == "" {
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Name is empty")
	} else {
		fmt.Fprintf(writer, "Haloo %s", name)
	}
}
func TestResponseCode(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080/hallo?name=goz", nil)
	res := httptest.NewRecorder()

	ResponseCode(res, req)

	result := res.Result()

	body, _ := io.ReadAll(result.Body)
	fmt.Println(result.StatusCode)
	fmt.Println(result.Status)
	fmt.Println(string(body))

}
