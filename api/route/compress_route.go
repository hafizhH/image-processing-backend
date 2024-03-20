package route

import (
	"image-processing-backend/api/controller"
	"image-processing-backend/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func NewCompressRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	cc := &controller.CompressController{
		Env: env,
	}

	group.POST("/compress", cc.Compress)
}
