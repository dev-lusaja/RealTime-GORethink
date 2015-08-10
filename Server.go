package main

import (
	"fmt"
	r "github.com/dancannon/gorethink"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type User struct {
	Name string `gorethink:"name"`
}

var session *r.Session

// Inicializamos la conexion a la DB
func init() {
	var err error
	session, err = r.Connect(r.ConnectOpts{
		Address:  "127.0.0.1:28015",
		Database: "GolangDB",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("conexion establecida a la DB")
	}
}

func UserUtils(res http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		name := req.FormValue("name")
		user := &User{name}
		_, err := r.Table("users").Insert(user).RunWrite(session)
		if err != nil {
			fmt.Print("Error: %s", err)
		} else {
			io.WriteString(res, "Insertado con exito")
		}
	}
	if req.Method == "DELETE" {
		id := req.FormValue("id")
		_, err := r.Table("users").Get(id).Delete().RunWrite(session)
		if err != nil {
			fmt.Print("Error: %s", err)
		} else {
			io.WriteString(res, "Eliminado con exito")
		}
	}
	if req.Method == "PUT" {
		id := req.FormValue("id")
		name := req.FormValue("name")
		user := &User{name}
		_, err := r.Table("users").Get(id).Update(user).RunWrite(session)
		if err != nil {
			fmt.Print("Error: %s", err)
		} else {
			io.WriteString(res, "Actualizado con exito")
		}
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/Users", UserUtils)
	http.ListenAndServe(":5000", r)
}
