package cmd

import (
	"log"

	"github.com/seiixasgustavo/ToDo-fullstack-backend/pkg/config"
	"github.com/seiixasgustavo/ToDo-fullstack-backend/pkg/controllers"
)

func RunServer() {
	db, err := config.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	router := controllers.Router(db)
	router.Logger.Fatal(router.Start(":8000"))
}
