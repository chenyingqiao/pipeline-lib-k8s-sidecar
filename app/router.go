package app

import (
	"pipeline-common-k8s-sidecar/app/controller"

	"github.com/gin-gonic/gin"
)

//RegisterRouter 注册路由
func RegisterRouter() *gin.Engine {
	r := gin.Default()

	text := controller.Text{}

	r.POST("/text-store", text.Store)
	return r
}
