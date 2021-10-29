package learngoweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, req *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = req.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprintf(writer, "Success Create Cookie")
}

func GetCookie(writer http.ResponseWriter, req *http.Request) {
	cookie, err := req.Cookie("token")
	if err != nil {
		fmt.Fprintf(writer, "No Cookie")
	} else {
		fmt.Fprintf(writer, "Hello %s", cookie.Value)
	}

}

func TestCookies(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080?name=zenklot", nil)
	res := httptest.NewRecorder()

	SetCookie(res, req)

	cookies := res.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("%s : %s\n", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	req := httptest.NewRequest("GET", "localhost:8080", nil)
	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = "Ini Isi Cookie"
	req.AddCookie(cookie)

	res := httptest.NewRecorder()

	GetCookie(res, req)

	response := res.Result()
	body, _ := io.ReadAll(response.Body)
	fmt.Println(string(body))

}
