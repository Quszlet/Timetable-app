package main

import (
	"github.com/Quszlet/timetable-app/bootstrap"
	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	env := app.Env
	gin := gin.Default()
	gin.Run(env.ServerAddress)
}
