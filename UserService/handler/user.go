package handler

import (
    "context"
    "encoding/json"
    "net/http"
    "github.com/go-chi/chi/v5"
    "github.com/mmartin04/Docker/UserService/db"
    "github.com/mmartin04/Docker/UserService/model"
    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct{}

// Vorhandene Methoden: Create, GetByID

// List - Listet alle Benutzer auf
func (u *User) List(w http.ResponseWriter, r *http.Request) {
    collection := db.GetCollection("users")
    cursor, err := collection.Find(context.TODO(), bson.M{})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    defer cursor.Close(context.TODO())

    var users []model.User
    if err = cursor.All(context.TODO(), &users); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(users)
}

// UpdateByID - Aktualisiert einen Benutzer anhand seiner ID
func (u *User) UpdateByID(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    var user model.User
    err = json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    collection := db.GetCollection("users")
    _, err = collection.UpdateByID(context.TODO(), objID, bson.M{"$set": user})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)
}

// DeleteByID - Löscht einen Benutzer anhand seiner ID
func (u *User) DeleteByID(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
    objID, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    collection := db.GetCollection("users")
    _, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objID})
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}
