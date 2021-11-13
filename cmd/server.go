package cmd

import (
	"github.com/gin-gonic/gin"
)

func Reg(g* gin.RouterGroup) {
	g.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 0,
		})
	})
}

func Run() error {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 0,
		})
	})
	Reg(r.Group("/config"))
	return r.Run()
}