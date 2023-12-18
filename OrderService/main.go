package main

import (
    "context"
    "fmt"
    "github.com/mmartin04/Docker/OrderService/Anwendung"
    "github.com/mmartin04/Docker/OrderService/db"
)

func main() {
    db.Connect() // Initialisiert die MongoDB-Verbindung

    app := Anwendung.New()
    err := app.Start(context.TODO())
    if err != nil {
        fmt.Println("failed to start app:", err)
    }
}
