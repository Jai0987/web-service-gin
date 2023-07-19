package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var (
	oauthConfig *oauth2.Config
	db          *sql.DB
)

func main() {
	var err error

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		// Use localhost as the default database URL if DATABASE_URL environment variable is not set
		databaseURL = "postgres://jaikash12:jaikash12@localhost/ginauth?sslmode=disable"
	}

	// Database connection
	db, err = sql.Open("postgres", databaseURL)
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}
	defer db.Close()

	r := gin.Default()

	// User details API endpoints
	r.POST("/users", createUser)
	r.GET("/users/:id", getUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	// Account operations API endpoints
	r.GET("/accounts/:id", getAccount)
	r.POST("/accounts/:id/pay", payBill)
	r.GET("/accounts/:id/due", getDueDate)
	r.GET("/accounts/:id/score", getCreditScore)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}
	fmt.Println("Server is running on port:", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatal("Failed to start server: ", err)
	}
}

func init() {
	// Load environment variables from the cred.env file
	err := godotenv.Load("cred.env")
	if err != nil {
		log.Fatal("Error loading cred.env file: ", err)
	}

	// Read the environment variables for CLIENT_ID and CLIENT_SECRET
	clientID := os.Getenv("CLIENT_ID")
	if clientID == "" {
		log.Fatal("CLIENT_ID environment variable is not set.")
	}

	clientSecret := os.Getenv("CLIENT_SECRET")
	if clientSecret == "" {
		log.Fatal("CLIENT_SECRET environment variable is not set.")
	}

	// Create the OAuth2 config using the environment variables
	oauthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:9999/auth/google/callback",
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
	}
}
