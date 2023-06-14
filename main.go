package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sulis96/quinzia-golang-instrumentations/config"
	"github.com/sulis96/quinzia-golang-instrumentations/internal/controller"
	"github.com/sulis96/quinzia-golang-instrumentations/internal/repository"
	"github.com/sulis96/quinzia-golang-instrumentations/internal/service"
	"github.com/sulis96/quinzia-golang-instrumentations/pkg/db"
	"github.com/sulis96/quinzia-golang-instrumentations/router"
)

var (
	appConfig          *config.AppConfig
	dbConfig           *config.DbConfig
	tracerConfig       *config.InstrumentationConfig
	postgresClientRepo *db.Database
)

func main() {
	//load environment variable
	if err := godotenv.Load(); err != nil {
		log.Fatal(err)
	}
	// init config
	appConfig = config.InitAppConfig()
	dbConfig = config.InitDatabaseConfig()
	tracerConfig = config.InitInstrumentationConfig()
	// init database
	postgresClientRepo, err := db.NewDatabase(dbConfig)
	if err != nil {
		os.Exit(1)
	}
	defer postgresClientRepo.Close()

	iRepository := repository.NewRepository(postgresClientRepo)
	iService := service.NewService(iRepository)
	iController := controller.NewController(iService)
	iRouter := router.NewRouter(iController)

	iRouter.RunHTTP(appConfig)
}
