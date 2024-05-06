package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"xyz-books/register"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	booksController := register.BooksInit()

	// Handle missing environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("PORT environment variable not set, using default port 8080")
	}

	router.LoadHTMLGlob("frontend/index.html")

	book := router.Group("v1/api/books")
	book.GET("/:id", booksController.GetBook)

	srv := &http.Server{
		Addr:         "0.0.0.0:" + port,
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalf("There was an error with the http server: %v", err)
	}

	log.Printf("Server listening on http://localhost:%s", port)
}
