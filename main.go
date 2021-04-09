package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type user struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

var DB []*user

func main() {
	fmt.Println("users api")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("welcome to users api"))

	})
	http.HandleFunc("/api/v1/users", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method)
		if request.Method == http.MethodPost {
			var userObject user
			err := json.NewDecoder(request.Body).Decode(&userObject)
			if err != nil {
				fmt.Println("coudnt  read requestBody", err)
				writer.Write([]byte("couldnt read requestBody"))
				return

			}
			DB = append(DB, &userObject)

			writer.Write([]byte(" user created succesfully"))

		} else if request.Method == http.MethodGet {
			err := json.NewEncoder(writer).Encode(DB)
			if err != nil {
				fmt.Println("coudnt  read requestBody", err)
				writer.Write([]byte("couldnt read requestBody"))
				return

			}

		}

	})
	err := http.ListenAndServe("0.0.0.0:8200", nil)
	if err != nil {
		fmt.Println("could'nt start server", err)
	}
}
