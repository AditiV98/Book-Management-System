package main

import (
	handlerAdmin "Book_Management_System/handlers/admin"
	handlerBook "Book_Management_System/handlers/book"
	handlerIssue "Book_Management_System/handlers/bookIssue"
	handlerUser "Book_Management_System/handlers/users"
	adminService "Book_Management_System/services/admin"
	bookService "Book_Management_System/services/book"
	issueService "Book_Management_System/services/bookIssue"
	"Book_Management_System/services/token"
	userService "Book_Management_System/services/users"
	"Book_Management_System/stores/admin"
	"Book_Management_System/stores/book"
	"Book_Management_System/stores/bookIssue"
	"Book_Management_System/stores/users"
	"database/sql"
	"github.com/gin-contrib/cors"
	"github.com/joho/godotenv"
	"log"
	"os"

	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // Import the MySQL driver
)

const appName = "book_management_system"

func main() {
	connectionString := "root:password@tcp(localhost:2001)/library"
	var err error
	mysql, err := sql.Open("mysql", connectionString)
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}
	defer mysql.Close()

	err = godotenv.Load("./configs/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	webClientID := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")

	tokenService := token.New(webClientID)

	adminStore := admin.New(mysql)
	adminService := adminService.New(adminStore, tokenService)
	adminHandler := handlerAdmin.New(adminService)

	bookStore := book.New(mysql)
	bookSvc := bookService.New(bookStore)
	bookHandler := handlerBook.New(bookSvc)

	userStore := users.New(mysql)
	userSvc := userService.New(userStore)
	userHandler := handlerUser.New(userSvc)

	issueStore := bookIssue.New(mysql)
	issueSvc := issueService.New(issueStore)
	issueHandler := handlerIssue.New(issueSvc)

	app := gin.New()

	app.Use(func(c *gin.Context) {
		c.Header("Cross-Origin-Opener-Policy", "same-origin")
		c.Next()
	})

	// Configure CORS middleware
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Replace with your React frontend's URL
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	app.Use(cors.New(config))

	app.POST("/admin/login", adminHandler.Login)
	app.POST("/admin/logout", adminHandler.Logout)

	app.POST("/book", bookHandler.Create)
	app.PUT("/book/:id", bookHandler.Update)
	app.GET("/book", bookHandler.GetAll)
	app.GET("/book/:id", bookHandler.GetByID)
	app.DELETE("/book/:id", bookHandler.Delete)

	app.POST("/user", userHandler.Create)
	app.PUT("/user/:id", userHandler.Update)
	app.GET("/user", userHandler.GetAll)
	app.GET("/user/:id", userHandler.GetByID)
	app.DELETE("/user/:id", userHandler.Delete)

	app.POST("/issue", issueHandler.Create)
	app.PUT("/issue/:id", issueHandler.Update)
	app.GET("/issue", issueHandler.GetAll)
	app.GET("/issue/:id", issueHandler.GetByID)
	app.DELETE("/issue/:id", issueHandler.Delete)

	app.Run(":8080")
}
