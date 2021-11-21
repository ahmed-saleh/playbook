package main

import (
	"github.com/ahmed-saleh/playbook/cmd/server"
	"github.com/ahmed-saleh/playbook/config"
	"github.com/ahmed-saleh/playbook/models"
)

func main() {
	config.Setup("ini/app.ini")
	models.Setup(config.MysqlSettings)
	server.Start(config.ServerSettings)
}
