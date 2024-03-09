package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine){
	v1 := router.Group("/api/v1")
	{
		v1.GET("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Get Openings endpoint",
			})
		})

		v1.POST("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Post Openings endpoint",
			})
		})

		v1.DELETE("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Delete Openings endpoint",
			})
		})

		v1.PUT("/opening", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Put Openings endpoint",
			})
		})

		v1.GET("/list-openings", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "Get All Openings endpoint",
			})
		})
	}
}