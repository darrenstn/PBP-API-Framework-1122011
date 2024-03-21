package main

import (
	"eksplorasi2/controllers"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gobuffalo/buffalo"
)

func main() {
	// Create a new Buffalo application
	app := buffalo.New(buffalo.Options{})

	// Define routes
	app.GET("/users", controllers.GetAllUsers)
	app.POST("/user", controllers.InsertUser)
	app.PUT("/user", controllers.UpdateUser)
	app.DELETE("/user/{user_id}", controllers.DeleteUser)

	// Serve the application
	fmt.Println("Connected to port 8888")
	log.Println("Connected to port 8888")

	// Create a new http.Server
	server := &http.Server{
		Addr:    ":8888", // Specify the desired port here
		Handler: app,
	}

	// Serve the application
	fmt.Println("Starting server on :8888")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
