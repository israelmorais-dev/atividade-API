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
		Name:       "Morais",
		University: "UniJuazeiro",
		Course:     "Sistemas de Informação",
	},
}

// @title           Swagger API-Users
// @version         1.0
// @description     Documentação da API de Users.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API-Users Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080

func main() {
	fmt.Println("Servidor rodando")
	// port := os.Getenv("PORT")
	// router := mux.NewRouter()
	http.HandleFunc("/", home)
	http.HandleFunc("/users", routerUsers)
	// http.HandleFunc("/users", createUsers)

	log.Fatal(http.ListenAndServe(":8080", nil))

}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bem vindo")
}

// ShowAllUsers godoc
// @Summary      Show all users
// @Description  List all users
// @Tags         users
// @Accept       json
// @Produce      json
// @Success      200  {object}  User
// @Router       /users [get]

func listUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(w)
	encoder.Encode(Users)
}

// ShowAllcreate godoc
// @Summary      Show all create
// @Description  Create all create
// @Tags         create
// @Accept       json
// @Produce      json
// @Success      201  {object}  User
// @Router       /users [post]

func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(201)
	body, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
	}
	var newUser User
	json.Unmarshal(body, &newUser)
	newUser.Id = len(Users) + 1
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
