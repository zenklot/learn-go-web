package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestHeader(writer http.ResponseWriter, req *http.Request) {
	cont_tipe := req.Header.Get("content-type")
	fmt.Fprint(writer, cont_tipe)
}

func ResponseHeader(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Add("X-Powered-By", "Gozenx")
	fmt.Fprint(writer, "OK")
}

func TestReqHeader(t *testing.T) {
	request := httptest.NewRequest("GET", "localhost:8080", nil)
	request.Header.Add("content-type", "application/json")
	recorder := httptest.NewRecorder()

	RequestHeader(recorder, request)
	ResponseHeader(recorder, request)

	poweredBy := recorder.Header().Get("X-Powered-By")
	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))
	fmt.Println(poweredBy)

}
