package user

import (
	"fmt"
	"github.com/go-martini/martini"
	"io"
	"net/http"
)

func UserApiRoute(m martini.Router) {
	m.Get("/", UserList)
	m.Post("/register", UserRegister)
}

func UserList(res http.ResponseWriter, req *http.Request) {
	fmt.Println("UserList")
	io.WriteString(res, req.URL.String())
}

func UserRegister(res http.ResponseWriter, req *http.Request) {

}
