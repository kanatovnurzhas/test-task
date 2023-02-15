package main

import (
	"app2/config"
	"app2/internal/delivery"
	"app2/internal/repository"
	"app2/internal/server"
	"app2/internal/service"
	"fmt"
	"log"
)

func main() {
	configs := config.InitConfigs()
	coll, ctx, err := repository.ConnectDB(configs)
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err)
	}
	repo := repository.NewRepository(coll, ctx)
	services := service.NewService(*repo)
	handler := delivery.NewHandler(services)

	newServer := server.NewServerInit(configs)

	fmt.Printf("Starting server at port %s\nhttp://"+configs.Host+":%s/\n", configs.Port, configs.Port)
	if err := newServer.Run(handler.InitRoutes()); err != nil {
		log.Fatal(err)
	}
}
