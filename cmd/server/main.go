package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"xyz-books/register"

	"github.com/gin-contrib/cors"
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

	config := cors.Config{
		AllowOrigins:  []string{"http://localhost:5173"},
		AllowMethods:  []string{"PUT", "PATCH", "GET"},
		AllowHeaders:  []string{"Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
		MaxAge:        12 * time.Hour,
	}
	router.Use(cors.New(config))

	book := router.Group("v1/api/book")
	book.GET("/list", booksController.ListBooks)
	book.GET("/:id", booksController.GetBook)
	book.PATCH("/:id", booksController.EditBook)

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
