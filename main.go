package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/assignment/resultview/config"
	"github.com/assignment/resultview/controller"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	db, err := config.Connect()
	defer db.Close()
	if err != nil {
		fmt.Println("Connection established.... hoy nay")
		//return nil, err
		//panic(err)
	}
	router.HandleFunc("/", controller.Index).Methods("GET")
	router.HandleFunc("/showresult", controller.ShowResult).Methods("GET")
	router.HandleFunc("/viewresult", controller.ViewResult).Methods("POST")

	//http.HandleFunc("/", controller.Index)

	log.Fatal(http.ListenAndServe(":8000", router))
}
