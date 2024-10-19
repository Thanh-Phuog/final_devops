package main

import (
	"book_mana/database"
	"book_mana/models"
	"book_mana/routes"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateBook(t *testing.T) {
	// Kết nối database test
	database.Connect()
	//database.DB.AutoMigrate(&models.Book{})

	gin.SetMode(gin.TestMode)
	router := routes.SetupRouter()

	book := models.Book{Title: "Test Book", Author: "Test Author", Year: 2023}
	body, _ := json.Marshal(book)

	req, _ := http.NewRequest("POST", "/books", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var createdBook models.Book
	err := json.Unmarshal(rr.Body.Bytes(), &createdBook)
	assert.NoError(t, err)
	assert.Equal(t, book.Title, createdBook.Title)
}

func TestGetBooks(t *testing.T) {
	// Kết nối database test
	database.Connect()

	gin.SetMode(gin.TestMode)
	router := routes.SetupRouter()

	req, _ := http.NewRequest("GET", "/books", nil)
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
}
