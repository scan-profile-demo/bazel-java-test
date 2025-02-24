package main

import (
	"errors"
	"fmt"
	"log"

	// "io"
	"net/http"
	"os"

	"github.com/Shulammite-Aso/bazel-demo-app/handlers"
	"github.com/antchfx/xmlquery"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello world")
	wadl, err := xmlquery.LoadURL("https://httpbin.org/get")
	if err != nil {
		panic(err)
	}

	attr := xmlquery.FindOne(wadl, "//application/@xmlns")
	fmt.Println(attr.InnerText())

	router := mux.NewRouter()

	router.HandleFunc("/greet", handlers.Greet).Methods("GET")
	router.HandleFunc("/greet-many", handlers.GreetMany).Methods("GET")

	address := ":5000"

	log.Printf("server started at port %v\n", address)

	err = http.ListenAndServe(address, router)

	if errors.Is(err, http.ErrServerClosed) {
		log.Printf("server closed\n")
	} else if err != nil {
		log.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
