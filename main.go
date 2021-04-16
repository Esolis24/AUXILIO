package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/Esolis24/LabMoviles/controller"
	"github.com/Esolis24/LabMoviles/model"
	"github.com/gorilla/mux"
)

func createUser(w http.ResponseWriter, r *http.Request) { //Pal login
	var newUser model.User
	db := controller.Connect()
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Insert a valid user")
	}

	json.Unmarshal(reqBody, &newUser)
	rows := controller.Query("USUARIO", newUser.Email, newUser.Password, db)
	if rows == nil {
		fmt.Fprintf(w, "Insert a valid user")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// func getUsers(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(usuarios)
// }
// func getSpecificUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	//El nombre "id" es por lo que paso en el HandleFunc
// 	userId, err := strconv.Atoi(vars["id"]) //Lo convierte a un entero
// 	if err != nil {
// 		fmt.Fprintf(w, "Invalid ID")
// 		return
// 	}
// 	for _, user := range usuarios {
// 		if user.ID == userId {
// 			w.Header().Set("Content-Type", "application/json")
// 			json.NewEncoder(w).Encode(user)
// 		}
// 	}
// }
// func deleteUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	//El nombre "id" es por lo que paso en el HandleFunc
// 	userId, err := strconv.Atoi(vars["id"]) //Lo convierte a un entero
// 	if err != nil {
// 		fmt.Fprintf(w, "Invalid ID")
// 		return
// 	}
// 	for i, user := range usuarios {
// 		if user.ID == userId {
// 			usuarios = append(usuarios[:i], usuarios[i+1:]...)
// 			fmt.Fprintf(w, "La tarea con el ID: %v ha sido removida satisfactoriamente", userId)
// 		}
// 	}
// }
// func updateUser(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	userId, err := strconv.Atoi(vars["id"])
// 	var updatedUser usuario
// 	if err != nil {
// 		fmt.Fprintf(w, "Invalid ID")
// 	}
// 	reqBody, err := ioutil.ReadAll(r.Body) //Lee todos los datos que da el usuario
// 	if err != nil {
// 		fmt.Fprintf(w, "Please enter valid Data")
// 	}
// 	json.Unmarshal(reqBody, &updatedUser)

// 	for i, user := range usuarios {
// 		if user.ID == userId {
// 			usuarios = append(usuarios[:i], usuarios[i+1:]...)
// 			updatedUser.ID = userId
// 			usuarios = append(usuarios, updatedUser)

// 			fmt.Fprintf(w, "The user with ID:%v was updated successfully", userId)
// 		}
// 	}
// }
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido a mi API epa")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	//router.HandleFunc("/usuarios", getUsers).Methods("GET")    //Este es para consultar
	router.HandleFunc("/usuarios", createUser).Methods("POST") //Este para enviar
	//router.HandleFunc("/login/{id}/{password}", getSpecificUser).Methods("GET")
	// router.HandleFunc("/deleteUser/{id}", deleteUser).Methods("DELETE")
	// router.HandleFunc("/usuarios/{id}", updateUser).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3000", router))

}
