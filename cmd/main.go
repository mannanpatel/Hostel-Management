package main

import (
	"hst_manag/internal/database"
	"hst_manag/internal/helper"
	"hst_manag/router"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

// Initialize the logger
var logger *log.Logger

func InitLogger() {
	// Open a log file in append mode
	file, err := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// Create a multi-writer to write to both the terminal and the log file
	multiWriter := io.MultiWriter(os.Stdout, file)

	// Create a new logger that uses the multi-writer
	logger = log.New(multiWriter, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	log.SetOutput(multiWriter)
}

func Init() {
	helper.LoadEnv()
	database.DBconnect()
	helper.CreateAdmin()
}

func main() {
	InitLogger()
	logger.Println("Starting application...")

	Init()
	r := gin.Default()
	r.Use(gin.LoggerWithWriter(logger.Writer()))
	router.Routes(r)
	router.AdminRoutes(r)
	logger.Println("Routes initialized, starting server on port 8082")
	r.Run(":8082")
}
