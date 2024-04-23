package main

import (
    "net/http"

    "github.com/gorilla/mux"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "crud/handlers"
    "crud/models"
)

func main() {
    db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&models.Book{})

    handler := handlers.Handlers{DB: db}

    router := mux.NewRouter()

    router.HandleFunc("/books", handler.GetBooks).Methods("GET")
    router.HandleFunc("/books/{id}", handler.GetBook).Methods("GET")
    router.HandleFunc("/books", handler.CreateBook).Methods("POST")
    router.HandleFunc("/books/{id}", handler.UpdateBook).Methods("PUT")
    router.HandleFunc("/books/{id}", handler.DeleteBook).Methods("DELETE")

    http.ListenAndServe(":8080", router)
}
