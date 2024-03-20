package route

import (
	"image-processing-backend/api/controller"
	"image-processing-backend/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func NewConvertRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	cc := &controller.ConvertController{
		Env: env,
	}

	group.POST("/convert", cc.Convert)
}
