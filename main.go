package main

import (
	"embed"
	"taskapi/config"
	"taskapi/models"

	)

//go:embed templates/*
var embeddedFiles embed.FS

func main() {
	models.InitPostgresDb()
	config.InitRouter(embeddedFiles)
}
