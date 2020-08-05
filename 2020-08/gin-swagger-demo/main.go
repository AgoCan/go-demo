package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/agocan/go-demo/2020-08/gin-swagger-demo/docs"
)

// User ss
type User struct {
	Name string `json:"name"`
}

// @Summary 新增文章标签
// @Produce  json
// @Param name query string true "name"
// @Param state query int false "State"
// @Param created_by body User true "CreatedBy"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /api/v1/tags [post]
func handler(c *gin.Context) {

}

// @title das文档
// @version 1.0
// @description 这是测试文档

// @host localhost
// @BasePath /api/v1
func main() {
	r := gin.New()

	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json") // The url pointing to API definition
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	r.POST("/api/v1/tags", handler)
	r.Run()
}
