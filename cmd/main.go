package main

import (
	"image-processing-app/api/middleware"
	"image-processing-app/api/route"
	"image-processing-app/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	gin.Use(middleware.ErrorHandler)

	route.SetupRouter(env, timeout, gin)

	gin.Run(env.ServerAddress)
}
