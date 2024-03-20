package main

import (
	"image-processing-backend/api/middleware"
	"image-processing-backend/api/route"
	"image-processing-backend/bootstrap"
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
