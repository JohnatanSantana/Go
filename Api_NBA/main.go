package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/johnatansantana/Go/Api_NBA/config"
	. "github.com/johnatansantana/Go/Api_NBA/config/dao"
	nbarouter "github.com/johnatansantana/Go/Api_NBA/router"
)

var dao = NbaDAO{}
var config = Config{}

func init() {
	config.Read()

	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/nbas", nbarouter.GetAll).Methods("GET")
	r.HandleFunc("/api/v1/nbas/{id}", nbarouter.GetByID).Methods("GET")
	r.HandleFunc("/api/v1/nbas", nbarouter.Create).Methods("POST")
	r.HandleFunc("/api/v1/nbas/{id}", nbarouter.Update).Methods("PUT")
	r.HandleFunc("/api/v1/nbas/{id}", nbarouter.Delete).Methods("DELETE")

	var port = ":3000"
	fmt.Println("Server running in port:", port)
	log.Fatal(http.ListenAndServe(port, r))
}
