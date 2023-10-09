package main

import (
	"github.com/sirupsen/logrus"
	"log"
	"main/internal/configs"
	"main/internal/db"
	"main/internal/handlers"
	"main/internal/repository"
	"main/internal/router"
	"main/internal/service"
	"main/pkg/logger"
	"net/http"
)

func main() {
	myLog, err := logger.GetLogger()
	if err != nil {
		log.Fatal(err)
	}
	err = start(myLog)
	if err != nil {
		myLog.Error(err)
	}
}

func start(log *logrus.Logger) error {
	conf, err := configs.GetConfigs()
	if err != nil {
		log.Error(err)
	}
	connectToDb, err := db.GetConnection(conf)
	if err != nil {
		log.Error(err)
	}
	newRepository := repository.GetRepository(connectToDb, log)
	newService := service.GetService(newRepository, log)
	newHandlers := handlers.GetHandler(newService, log)
	newRouter := router.GetRouter(newHandlers)
	server := http.Server{
		Addr:    conf.Host + conf.Port,
		Handler: newRouter,
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Error(err)
	}
	return nil
}
