package config

import (
	"embed"
	"html/template"
	"taskapi/controller"
	"taskapi/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(embeddedFiles embed.FS) {
	// 1.创建路由
	// 默认使用了2个中间件Logger(), Recovery()
	r := gin.Default()
	// 允许跨域请求
	r.Use(middleware.CORSHeaders())
	// 加载模板文件
	tmpl := template.Must(template.New("").ParseFS(embeddedFiles, "templates/**/*"))
	r.SetHTMLTemplate(tmpl)

	r.GET("/", controller.Home)
	r.GET("/tasks", controller.TaskListPage)
	r.GET("/tasks/:id", controller.TaskDetailPage)
	r.GET("/tasks/add", controller.TaskAddPage)
	r.GET("/tasks/edit/:id", controller.TaskEditPage)

	// api接口，使用RewriteResponse
	api := r.Group("/api")
	api.Use(middleware.RewriteResponse())
	api.GET("/hello", controller.Hello)
	api.POST("/task/add", controller.AddTask)
	api.POST("/task/delete", controller.DeleteTask)
	api.POST("/task/update", controller.UpdateTask)
	api.POST("/task/finish", controller.FinishTask)
	api.POST("/task/begin", controller.BeginTask)
	api.GET("/task/list", controller.ListTask)

	r.Run(":8000")
}
