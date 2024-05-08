package handler

import (
	"fmt"
	"net/http"
)

type Todo struct{

}
func (To *Todo) Create(w http.ResponseWriter, r *http.Request){
	fmt.Println("Create a todo")
}

func (To *Todo) List(w http.ResponseWriter, r *http.Request){
	fmt.Println("List all todo")
}

func (To *Todo) ListByID(w http.ResponseWriter, r *http.Request){
	fmt.Println("List a todo by ID")
}
func (To *Todo) Update(w http.ResponseWriter, r *http.Request){
	fmt.Println("Update a todo")
}
func (To *Todo) DeleteById(w http.ResponseWriter, r *http.Request){
	fmt.Println("Delete / archive a todo by ID")
}