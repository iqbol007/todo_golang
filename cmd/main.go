package main

import (
	"github.com/spf13/viper"
	"log"
	todo "todoAppGo"
	"todoAppGo/pakackge/handler"
	"todoAppGo/pakackge/reposithory"
	"todoAppGo/pakackge/service"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initialization configs: %s", err.Error())

	}
	repos := reposithory.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("error occured when run http server: %s", err.Error())
	}
}
func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
