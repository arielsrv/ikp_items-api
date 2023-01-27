package app

import (
	"fmt"
	"ikp_items-api/src/main/app/infrastructure/database"
	"log"
	"net/http"

	"ikp_items-api/src/main/app/config"
	"ikp_items-api/src/main/app/config/env"
	"ikp_items-api/src/main/app/handlers"
	"ikp_items-api/src/main/app/server"
	"ikp_items-api/src/main/app/services"
)

var dbClient = ProvideDBClient()

func Run() error {
	app := server.New(server.Config{
		Recovery:  true,
		Swagger:   true,
		RequestID: true,
		Logger:    true,
	})

	pingService := services.NewPingService()
	pingHandler := handlers.NewPingHandler(pingService)

	itemService := services.NewItemService(dbClient)
	itemHandler := handlers.NewItemHandler(itemService)

	server.RegisterHandler(pingHandler)
	server.RegisterHandler(itemHandler)

	server.Register(http.MethodGet, "/ping", server.Resolve[handlers.PingHandler]().Ping)
	server.Register(http.MethodPost, "/items", server.Resolve[handlers.ItemHandler]().Create)

	host := config.String("HOST")
	if env.IsEmpty(host) && !env.IsDev() {
		host = "0.0.0.0"
	} else {
		host = "127.0.0.1"
	}

	port := config.String("PORT")
	if env.IsEmpty(port) {
		port = "8080"
	}

	address := fmt.Sprintf("%s:%s", host, port)

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://%s:%s/ping in the browser", host, port)

	return app.Start(address)
}

func ProvideDBClient() database.IDbClient {
	connectionString := config.String("PROD_CONNECTION_STRING")
	mySQLClient := database.NewMySQLClient(connectionString)

	return database.NewDBClient(mySQLClient)
}
