package route

import (
	"image-processing-app/api/controller"
	"image-processing-app/bootstrap"
	"time"

	"github.com/gin-gonic/gin"
)

func NewConvertRouter(env *bootstrap.Env, timeout time.Duration, group *gin.RouterGroup) {
	cc := &controller.ConvertController{
		Env: env,
	}

	group.POST("/convert", cc.Convert)
}
