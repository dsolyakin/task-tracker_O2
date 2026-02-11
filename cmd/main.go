package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dsolyakin/task-tracker/internal/database"
	"github.com/dsolyakin/task-tracker/internal/http"
	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", host, user, password, dbname, port, sslmode)

	db := database.InitDB(dsn)
	r := gin.Default()
	http.InitRoutes(r, db)
	r.Run()
}
