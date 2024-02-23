package dto

import "github.com/gin-gonic/gin"

//StoreParam 入参
type StoreParam []StoreParamItem

//StoreParamItem 文件信息
type StoreParamItem struct {
	Path    string `json:"path"`
	Content string `json:"content"`
}

//NewStoreParam 实例化参数
func NewStoreParam(c *gin.Context) (StoreParam, error) {
	result := StoreParam{}
	err := c.Bind(&result)
	if err == nil {
		return result, nil
	}
	return result, err
}
