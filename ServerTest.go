package main

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	mysqlTest()
}

func routeTest() {
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPerson).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePerson).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePerson).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func mysqlTest() {
	db, err := sql.Open("mysql",
		"root:@tcp(127.0.0.1:3306)/Test123")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var (
		ID int
		NAME string
		SEX string
	)
	rows, err := db.Query("select ID, NAME, SEX from John")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&ID, &NAME, &SEX)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(ID, NAME, SEX)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

type Person struct {
	ID        int    `json:"id,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
}

func GetPeople(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	var kid = Person{5, "Avery", "Durrant"}

	json.NewEncoder(w).Encode(kid)
}
func GetPerson(w http.ResponseWriter, r *http.Request)    {}
func CreatePerson(w http.ResponseWriter, r *http.Request) {}
func DeletePerson(w http.ResponseWriter, r *http.Request) {}
