package main

import (
	"log"
	"os"
	"path/filepath"

	"server1/cmd/router"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	pathDir := "cd ../../"
	err := godotenv.Load(filepath.Join(pathDir, ".env"))
	if err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err)
	}
}

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	db, err := gorm.Open("postgres", `host=`+dbHost+` port=`+dbPort+` user=`+dbUser+` dbname=`+dbName+` password=`+dbPassword+` sslmode=disable`)
	serverPORT := os.Getenv("SERVER_PORT")
	if err != nil {
		log.Fatal("failed to connect to database", err)
	}
	defer db.Close()
	db.LogMode(true)
	handler := initHandler(db)
	router.API(*handler, serverPORT)
}
