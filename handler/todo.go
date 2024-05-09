package handler

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/MassouAnas/ChiBackEnd/model"
	"github.com/MassouAnas/ChiBackEnd/repository/todo"
	"github.com/go-chi/chi/v5"
	"go.mongodb.org/mongo-driver/bson"
)

type Todo struct{
    Repo *todo.MongoRepo
}
func (to *Todo) Create(w http.ResponseWriter, r *http.Request) {
    var todo model.Todo

    if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }

	now := time.Now().Format("2 Jan 06 03:04PM")
	todo.CreatedAt = &now

    todo.TodoID = rand.Uint64()

    err := to.Repo.Insert(r.Context(), todo)
    if err != nil {
        fmt.Println("failed to insert todo:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    res, err := json.Marshal(todo)
    if err != nil {
        fmt.Println("failed to marshal todo:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write(res)
}

func (to *Todo) List(w http.ResponseWriter, r *http.Request){
	res, err := to.Repo.FindAll(r.Context())
	if err != nil {
		fmt.Println("failed to get all todos", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	OurResponse, err := json.Marshal(res)
	if err != nil{
		fmt.Println("failed to marshal todos:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
	}
	w.WriteHeader(http.StatusOK)
    w.Write(OurResponse)
	 
}

func (to *Todo) ListByID(w http.ResponseWriter, r *http.Request){
	idParam := chi.URLParam(r, "id")

	v, err := strconv.ParseUint(idParam, 10,64)
	if err != nil {
		fmt.Println("failed to parse the id string into an int")
	}


	res, err := to.Repo.FindByID(r.Context(), v)
	if err != nil {
		fmt.Println("failed to get the todo", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	OurResponse, err := json.Marshal(res)
	if err != nil{
		fmt.Println("failed to marshal the todo:", err)
        w.WriteHeader(http.StatusInternalServerError)
        return
	}
	w.WriteHeader(http.StatusOK)
    w.Write(OurResponse)
}
func (to *Todo) DeleteById(w http.ResponseWriter, r *http.Request){
	idParam := chi.URLParam(r, "id")

	v, err := strconv.ParseUint(idParam, 10,64)
	if err != nil {
		fmt.Println("failed to parse the id string into an int")
	}
	err = to.Repo.Delete(r.Context(), v)
		if err != nil {
			fmt.Println("failed to delete the todo",err)
			return
		}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todo deleted"))

}
func (to *Todo) Update(w http.ResponseWriter, r *http.Request){
	idParam := chi.URLParam(r, "id")

	v, err := strconv.ParseUint(idParam, 10,64)
	if err != nil {
		fmt.Println("failed to parse the id string into an int")
		return
	}

	var updatedTodo map[string]interface{}

	err = json.NewDecoder(r.Body).Decode(&updatedTodo)
	if err != nil {
		fmt.Println("failed to decode json body", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	mappingJSONtoBSON := bson.M{}
	for key, value := range updatedTodo{
		mappingJSONtoBSON[key]=value
	}

	err = to.Repo.Update(r.Context(), v, mappingJSONtoBSON)
		if err != nil {
			fmt.Println("failed to delete the todo",err)
			return
		}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Todo updated successfully"))
}