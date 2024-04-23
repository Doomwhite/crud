package handlers

import (
    "encoding/json"
    "crud/models"
    "net/http"

    "github.com/gorilla/mux"
    "gorm.io/gorm"
)

type Handlers struct {
    DB *gorm.DB
}

func (handler *Handlers) GetBooks(writer http.ResponseWriter, request *http.Request) {
    var books []models.Book
    handler.DB.Find(&books)
    json.NewEncoder(writer).Encode(books)
}

func (handler *Handlers) GetBook(writer http.ResponseWriter, request *http.Request) {
    params := mux.Vars(request)
    var book models.Book
    handler.DB.First(&book, params["id"])
    if book.ID == 0 {
        http.Error(writer, "Book not found", http.StatusNotFound)
    } else {
        json.NewEncoder(writer).Encode(book)
    }
}

func (h *Handlers) CreateBook(w http.ResponseWriter, r *http.Request) {
    var book models.Book
    json.NewDecoder(r.Body).Decode(&book)
    h.DB.Create(&book)
    json.NewEncoder(w).Encode(book)
}

func (handler *Handlers) UpdateBook(writer http.ResponseWriter, request *http.Request) {
    params := mux.Vars(request)
    var book models.Book
    handler.DB.First(&book, params["id"])
    json.NewDecoder(request.Body).Decode(&book)
    handler.DB.Save(&book)
    json.NewEncoder(writer).Encode(book)
}

func (handler *Handlers) DeleteBook(writer http.ResponseWriter, request *http.Request) {
    params := mux.Vars(request)
    var book models.Book
    handler.DB.Delete(&book, params["id"])
    if book.ID == 0 {
        http.Error(writer, "Book not found", http.StatusNotFound)
    } else {
        json.NewEncoder(writer).Encode(book)
    }
}
