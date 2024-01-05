package main

import (
	"log"

	dagestan "github.com/Way-Flare/dagestan-backend"
	"github.com/Way-Flare/dagestan-backend/internal/database/psql"
	"github.com/Way-Flare/dagestan-backend/internal/service"
	"github.com/Way-Flare/dagestan-backend/internal/transport/rest"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := psql.NewRepository()
	services := service.NewService(repos)
	handlers := rest.NewHandler(*services)

	srv := new(dagestan.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
