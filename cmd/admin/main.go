package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/hellodoge/courses-tg-bot/internal/admin/handler"
	"github.com/hellodoge/courses-tg-bot/internal/admin/server"
	"github.com/hellodoge/courses-tg-bot/internal/repository"
	"github.com/hellodoge/courses-tg-bot/internal/service"
	"github.com/hellodoge/courses-tg-bot/pkg/mdb"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {

	if err := initConfig(); err != nil {
		logrus.Fatal("error initializing configs: ", err)
	}

	if err := godotenv.Load(); err != nil {
		logrus.Error("error while loading .env file: ", err)
	}

	db, err := sqlx.Connect("mysql", os.Getenv("DB_URI_MYSQL"))
	if err != nil {
		logrus.Fatal("error while connecting to database:", err)
	}

	client, err := mdb.Connect(os.Getenv("MONGODB_URI"))
	if err != nil {
		logrus.Fatal("error while connecting to mongo db:", err)
	}

	repo := repository.NewRepository(db, client)
	services := service.NewService(repo)
	handlers := handler.NewHandler(services)

	httpServer := server.InitServer(server.Config{
		Port:    uint16(viper.GetUint("port")),
		Timeout: viper.GetDuration("timeout"),
	}, handlers.InitRoutes())

	if err := httpServer.Run(); err != nil {
		logrus.Fatalf("error occurred while running http server: %s", err)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("admin")
	viper.SetDefault("port", server.DefaultPort)
	viper.SetDefault("timeout", server.DefaultTimeout)
	return viper.ReadInConfig()
}
