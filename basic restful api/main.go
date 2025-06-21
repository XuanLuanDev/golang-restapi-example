package main

import(
	"myapi/apis"
	"net/http"
    "github.com/gorilla/mux"
    "myapi/midleware"

)

func main() {
    router := mux.NewRouter()
    router.Use(midleware.RefererHandler)
    router.HandleFunc("/api/v1/product/getall", apis.GetProducts).Methods("GET")
    router.HandleFunc("/api/v1/product/create", apis.CreateProduct).Methods("POST")
    router.HandleFunc("/api/v1/product/update", apis.UpdateProduct).Methods("PUT")
    router.HandleFunc("/api/v1/product/delete", apis.Delete).Methods("DELETE")
    apiRouter := router.PathPrefix("/api/product").Subrouter()
    apiRouter.Use(midleware.AuthHandler)
    err := http.ListenAndServe(":5000", router)
    if err != nil {
        panic(err)
    }
}