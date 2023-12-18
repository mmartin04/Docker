package handler

import (
    "context"
    "encoding/json"
    "net/http"
    "github.com/SimonPietrucha/Docker/OrderService/db"
    "github.com/SimonPietrucha/Docker/OrderService/model"
    "github.com/go-chi/chi/v5"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct{}

func (o *Order) Create(w http.ResponseWriter, r *http.Request) {
    var order model.Order
    err := json.NewDecoder(r.Body).Decode(&order)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    collection := db.GetCollection("orders")
    result, err := collection.InsertOne(context.TODO(), order)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(result)
}

func (o *Order) GetByID(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var order model.Order
    collection := db.GetCollection("orders")
    err = collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&order)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(order)
}

// Weitere CRUD-Methoden hier implementieren...
