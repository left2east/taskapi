package controller

import (
	"net/http"
	"strconv"
	"taskapi/models"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	title := "首页"
	c.HTML(http.StatusOK, "home/index", gin.H{"title": title})
}

func TaskListPage(c *gin.Context) {
	title := "任务列表"
	tasklist := models.ListTask("")
	c.HTML(http.StatusOK, "task/list", gin.H{"title": title, "tasks": tasklist})
}

func TaskDetailPage(c *gin.Context) {
	title := "任务详情"
	c.HTML(http.StatusOK, "task/detail", gin.H{"title": title})
}

func TaskAddPage(c *gin.Context) {
	title := "新增任务"
	c.HTML(http.StatusOK, "task/add", gin.H{"title": title})
}

func TaskEditPage(c *gin.Context) {
	taskId := c.Param("id")
	id, err := strconv.Atoi(taskId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task id"})
		return
	}
	task := models.GetTask(id)
	title := "编辑任务"
	c.HTML(http.StatusOK, "task/edit", gin.H{"title": title, "task": task})
}
