package main

import (
	"log"      //creating logger manually
	"net/http" //provide http clint and server

	"github.com/bala3663/go-project-book-management/routes"
	"github.com/gorilla/mux"                  //http request for multiplexer & used for request routing and dispatching
	_ "github.com/jinzhu/gorm/dialects/mysql" //ORM library for dealing with relational databases
)

func main() {
	r := mux.NewRouter() //intialize router and return new router instance
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)                                 //handle the string of given pattern
	log.Fatal(http.ListenAndServe("localhost:9010", r)) //handle request on incoming server of http
}
