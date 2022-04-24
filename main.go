package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	University string `json:"university"`
	Course     string `json:"course"`
}

type allUsers []User

var users = allUsers{
	{
		Id:         1,
		Name:       "Israel Morais",
		University: "UniJuazeiro",
		Course:     "Sistemas de Informação",
	},
}

func main() {
	log.Println("Starting API")
	router := mux.NewRouter()
	router.HandleFunc("/", Home)
	router.HandleFunc("/users", GetAllUsers).Methods("GET")

	http.ListenAndServe(":8080", router)
}

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Aplicação em execução!\n")
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	log.Println("Acessando o endpoint get all user")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
