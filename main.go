package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	University string `json:"university"`
	Course     string `json:"course"`
}

var Users []User = []User{
	{
		Id:         1,
		Name:       "Israel Morais",
		University: "UniJuazeiro",
		Course:     "Sistemas de Informação",
	},
	{
		Id:         2,
		Name:       "Bruno Ferreira",
		University: "UniJuazeiro",
		Course:     "Sistemas de Informação",
	},
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo")
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(Users)
}
func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
	}
	var newUser User
	json.Unmarshal(body, &newUser)
	Users = append(Users, newUser)

	encoder := json.NewEncoder(w)
	encoder.Encode(newUser)
}

func routerUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		listUsers(w, r)
	} else if r.Method == "POST" {
		create(w, r)
	}
}

func router() {
	http.HandleFunc("/", home)
	http.HandleFunc("/users", routerUsers)
	// http.HandleFunc("/users", createUsers)
}

func Server() {
	router()
	fmt.Println("Servidor rodando")
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
	Server()
}
