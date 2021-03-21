package handlers

import (
	"encoding/json"
	"net/http"
	"github.com/manosriram/go-api/pkg/db"
	"log"
)

type user struct {
    name string
    age int
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var user User
	// json.NewDecoder(r.Body).Decode(&user)

	conn := db.GetConnection();
	rows, err := conn.Query("select * from user");
	var results []user;
	for rows.Next() {
	    var newUser user;
	    err = rows.Scan(&newUser.name, &newUser.age);
	    if err != nil {
		log.Fatal(err);
	    }

	    results = append(results, newUser);
	}

	json.NewEncoder(w).Encode(&results)
}
