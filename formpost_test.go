package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		panic(err)
	}

	firstName := req.PostForm.Get("first_name")
	lastName := req.PostForm.Get("last_name")

	fmt.Fprintf(writer, "Hello %s %s", firstName, lastName)
}

func Test_Name(t *testing.T) {
	reqBody := strings.NewReader("first_name=Raisa&last_name=Supriatna")
	request := httptest.NewRequest("POST", "localhost:8080", reqBody)
	request.Header.Add("content-type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	result := recorder.Result()
	body, _ := io.ReadAll(result.Body)

	fmt.Println(string(body))

}
