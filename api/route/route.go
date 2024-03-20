package route

import (
	"image-processing-app/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter(env *bootstrap.Env, timeout time.Duration, gin *gin.Engine) {
	publicRouter := gin.Group("")

	NewConvertRouter(env, timeout, publicRouter)
	NewResizeRouter(env, timeout, publicRouter)
	NewCompressRouter(env, timeout, publicRouter)
	// protectedRouter := gin.Group("")
	// protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
}
