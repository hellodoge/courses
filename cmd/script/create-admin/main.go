package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hellodoge/courses-tg-bot/internal/repository"
	"github.com/hellodoge/courses-tg-bot/internal/service"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"strings"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("error while loading .env file:", err)
	}

	db, err := sqlx.Connect("mysql", os.Getenv("DB_DATA_SOURCE_NAME_MYSQL"))
	if err != nil {
		log.Fatalln("error while connecting to database:", err)
	}

	repo := repository.NewRepository(db, &mongo.Client{})
	services := service.NewService(repo)

	description := strings.Join(os.Args[1:], " ")

	token, err := services.NewAdmin(description)
	if err != nil {
		log.Fatalln("error while creating admin:", err)
	}

	fmt.Println(token)
}
