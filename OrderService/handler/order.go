package handler

import (
    "context"
    "encoding/json"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/mmartin04/Docker/OrderService/db"
    "github.com/mmartin04/Docker/OrderService/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct{}

// Create - Erstellt eine neue Bestellung
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

// GetByID - Holt eine Bestellung anhand ihrer ID
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

// List - Listet alle Bestellungen auf
func (o *Order) List(w http.ResponseWriter, r *http.Request) {
    collection := db.GetCollection("orders")
    cursor, err := collection.Find(context.TODO(), bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(context.TODO())

    var orders []model.Order
    if err = cursor.All(context.TODO(), &orders); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(orders)
}

// UpdateByID - Aktualisiert eine Bestellung anhand ihrer ID
func (o *Order) UpdateByID(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var order model.Order
    err = json.NewDecoder(r.Body).Decode(&order)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    collection := db.GetCollection("orders")
    _, err = collection.UpdateByID(context.TODO(), objID, bson.M{"$set": order})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(order)
}

// DeleteByID - Löscht eine Bestellung anhand ihrer ID
func (o *Order) DeleteByID(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    collection := db.GetCollection("orders")
    _, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
