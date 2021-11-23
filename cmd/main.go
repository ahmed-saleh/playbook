package main

import (
	"github.com/ahmed-saleh/playbook/config"
	"github.com/ahmed-saleh/playbook/models"
	"github.com/ahmed-saleh/playbook/server"
)

func main() {
	config.Setup("ini/app.ini")
	models.Setup(config.MysqlSettings)
	s := server.Build(config.ServerSettings)
	s.ListenAndServe()
}
