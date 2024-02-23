package controller

import (
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"pipeline-common-k8s-sidecar/app/dto"

	"github.com/gin-gonic/gin"
)

//Text 文本工具
type Text struct{}

//Store 存储文本
func (t *Text) Store(c *gin.Context) {
	param, err := dto.NewStoreParam(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	for _, item := range param {
		err := os.MkdirAll(filepath.Dir(item.Path), 0777)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		err = ioutil.WriteFile(item.Path, []byte(item.Content), 0777)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.JSON(http.StatusOK, "")
}
