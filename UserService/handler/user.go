package handler

import (
    "context"
    "encoding/json"
    "net/http"
    "github.com/SimonPietrucha/Docker/UserService/db"
    "github.com/SimonPietrucha/Docker/UserService/model"
    "github.com/go-chi/chi/v5"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
    var user model.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    collection := db.GetCollection("users")
    result, err := collection.InsertOne(context.TODO(), user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(result)
}

func (u *User) GetByID(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var user model.User
    collection := db.GetCollection("users")
    err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}

