package main

import (
	"encoding/json"
	"fmt"
	r "github.com/dancannon/gorethink"
)

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

// Metodo para convertir a JSON
func printObj(v interface{}) []byte {
	vBytes, _ := json.Marshal(v)
	return vBytes
}

// Metodo para suscribirse a una tabla y escuchar los nuevos eventos
func Suscribe(table string) {
	result, err := r.Table(table).Changes().Run(session)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("*** Escuchando: ***")
	var rs interface{}
	for result.Next(&rs) {
		fmt.Println("*** Nuevo ingreso: ***")
		data_JSON := printObj(rs)
		fmt.Println(string(data_JSON))
	}
}

func main() {
	Suscribe("users")
}
