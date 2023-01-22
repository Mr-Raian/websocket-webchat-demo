package main

import (
	"context"
	"fmt"
	"websocket-webchat-demo/internal/ent"
	"websocket-webchat-demo/internal/handlers"
	"websocket-webchat-demo/internal/router"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	if err := godotenv.Load("./.env"); err != nil {
		log.Fatal("Can't load ./.env file")
	}

	connectToDB()

	defer handlers.DB.Close()
	e := router.New()
	go handlers.HandleMessages()

	e.Logger.Fatal(e.Start("127.0.0.1:8088"))
}

func connectToDB() {
	var err error
	handlers.DB, err = ent.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"),
			os.Getenv("DB_NAME"), os.Getenv("DB_PASS"), os.Getenv("DB_SSLMODE")))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool.
	if err := handlers.DB.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
