package main

import (
	c "github.com/ahmed-saleh/playbook/cmd/cli/cmd"
	"github.com/ahmed-saleh/playbook/config"
	"github.com/ahmed-saleh/playbook/models"
)

func main() {
	config.Setup("ini/app.ini")
	models.Setup(config.MysqlSettings)
	c.Execute()
}
