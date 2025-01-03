package main

import (
	"embed"
	"taskapi/config"
	"taskapi/router"
	"taskapi/util"
)

//go:embed templates/*
var embeddedFiles embed.FS

//go:embed config/*.yml
var embeddedConfigFiles embed.FS

func main() {
	config.InitConfig(embeddedConfigFiles)
	router.InitRouter(embeddedFiles)
	util.InitPostgresDb()
}
