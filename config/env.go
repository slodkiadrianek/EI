package config

import (
	"log"
	"os"
)

type Env struct {
	Port   string
	DbLink string
}

func SetConfig() *Env {
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}

	dbLink := os.Getenv("DbLink")
	if dbLink == "" {
		log.Fatal("DbLink environment variable is not set")
	}

	return &Env{
		Port:   port,
		DbLink: dbLink,
	}
}
