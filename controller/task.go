package controller

import (
	"log"
	"strconv"
	"taskapi/models"

	"github.com/gin-gonic/gin"
)

func AddTask(c *gin.Context) {
	var data models.TaskTable
	err := c.BindJSON(&data)
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(200, gin.H{"success": false})
		return
	}
	id := models.AddTask(&data)
	if id == 0 {
		log.Fatal("数据库插入失败")
		c.JSON(200, gin.H{"success": false})
	} else {
		c.JSON(200, gin.H{"id": id, "success": true})
	}
}

func FinishTask(c *gin.Context) {
	taskId := c.PostForm("id")
	if taskId == "" {
		c.JSON(200, gin.H{"success": false})
		return
	}
	// 转化为int类型
	id, err := strconv.Atoi(taskId)
	if err != nil {
		c.JSON(200, gin.H{"success": false})
		return
	}

	models.FinishTask(id)
	c.JSON(200, gin.H{"success": true})
}

func BeginTask(c *gin.Context) {
	taskId := c.PostForm("id")
	if taskId == "" {
		c.JSON(200, gin.H{"success": false})
		return
	}
	// 转化为int类型
	id, err := strconv.Atoi(taskId)
	if err != nil {
		c.JSON(200, gin.H{"success": false})
		return
	}

	models.BeginTask(id)
	c.JSON(200, gin.H{"success": true})
}

func ListTask(c *gin.Context) {
	keyword := c.Query("keyword")
	tasks := models.ListTask(keyword)
	c.JSON(200, gin.H{"tasks": tasks})
}

func DeleteTask(c *gin.Context) {
	taskId := c.PostForm("id")
	if taskId == "" {
		c.JSON(200, gin.H{"success": false})
		return
	}
	// 转化为int类型
	id, err := strconv.Atoi(taskId)
	if err != nil {
		c.JSON(200, gin.H{"success": false})
		return
	}

	models.DeleteTask(id)
	c.JSON(200, gin.H{"success": true})
}

func UpdateTask(c *gin.Context) {
	var data models.TaskTable
	err := c.BindJSON(&data)
	if err != nil {
		log.Fatal(err.Error())
		c.JSON(200, gin.H{"success": false})
		return
	}
	models.UpdateTask(&data)
	c.JSON(200, gin.H{"success": true})
}
