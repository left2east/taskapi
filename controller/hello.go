package controller

import (
	"taskapi/models"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	user := models.User{
		Name:  "zhangsan",
		Email: "abc@123.com",
	}
	
	c.JSON(0, user)
}
