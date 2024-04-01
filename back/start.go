package main

import (
	"net/http"
	"syj/hope/entity"
	"syj/hope/service"
	"syj/hope/task"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	task.StartAutoGenerateSay(60000)
	r := gin.Default()
	r.Use(CORSMiddleware())
	r.GET("/says", func(c *gin.Context) {
		rg := service.GetSays()
		c.JSON(http.StatusOK, rg)
	})
	r.GET("/comments", func(c *gin.Context) {
		rg := service.GetComments(c.Query("sayId"))
		c.JSON(http.StatusOK, rg)
	})
	r.GET("/likes", func(c *gin.Context) {
		rg := service.GetLike(c.Query("sayId"))
		c.JSON(http.StatusOK, rg)
	})
	r.POST("/comment", func(c *gin.Context) {
		con := c.PostForm("content")
		nickName := c.PostForm("nickName")
		sayId := c.PostForm("sayId")
		comment := entity.Comment{
			SayId:      sayId,
			NickName:   nickName,
			Content:    con,
			RecordTime: time.Now().Format("2006-01-02 15:04:05"),
		}
		service.InsertComment(comment)
		c.JSON(http.StatusOK, nil)
	})
	r.POST("/like", func(c *gin.Context) {
		sayId := c.PostForm("sayId")
		service.Like(sayId)
		c.JSON(http.StatusOK, nil)
	})
	r.Run(":1111")
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, DELETE, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}
