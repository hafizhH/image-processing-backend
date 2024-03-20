package route

import (
	"image-processing-backend/api/controller"
	"image-processing-backend/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func NewResizeRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	rc := &controller.ResizeController{
		Env: env,
	}

	group.POST("/resize", rc.Resize)
}
