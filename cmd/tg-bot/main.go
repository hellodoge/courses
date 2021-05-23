package main

import (
	_ "github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/hellodoge/courses-tg-bot/internal/repository"
	"github.com/hellodoge/courses-tg-bot/internal/service"
	"github.com/hellodoge/courses-tg-bot/internal/telegram"
	"github.com/hellodoge/courses-tg-bot/pkg/database"
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

	db, err := sqlx.Connect("mysql", database.Config{
		Host:     viper.GetString("db.host"),
		Port:     uint16(viper.GetInt("db.port")),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
	}.GetMySQLDataSourceName())

	if err != nil {
		logrus.Fatal("error while connecting to db: ", err)
	}

	client, err := mdb.Connect(os.Getenv("MONGODB_URI"))
	if err != nil {
		logrus.Fatalln("error while connecting to mongo db:", err)
	}

	repo := repository.NewRepository(db, client)
	services := service.NewService(repo)

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TG_BOT_TOKEN"))
	if err != nil {
		logrus.Fatal("error initializing bot: ", err)
	}

	tg := telegram.NewBot(bot, services)

	channel, err := tg.InitUpdateChannel()
	if err != nil {
		logrus.Fatal("error initializing update channel: ", err)
	}

	tg.HandleUpdates(channel)
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
