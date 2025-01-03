package models

import (
	"log"
	"time"

	"taskapi/util"
)

/*
CREATE TABLE `task` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(200) NOT NULL DEFAULT '',
  `tag` varchar(200) NOT NULL DEFAULT '',
  `status` tinyint NOT NULL DEFAULT 0,
  `detail` varchar(2000) NOT NULL DEFAULT '',
  `begin_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `finish_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB;
*/
// 结构体
type TaskTable struct {
	ID         uint   `gorm:"primarykey" json:"id"`
	Name       string `json:"name"`
	Tag        string `json:"tag"`
	Status     int    `json:"status"`
	Detail     string `json:"detail"`
	BeginTime  string `json:"begin_time"`
	FinishTime string `json:"finish_time"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}

// 状态枚举
const (
	TaskStatusInit   = 1
	TaskStatusDoing  = 2
	TaskStatusDone   = 3
	TaskStatusDelete = 4
)

func AddTask(data *TaskTable) int {
	data.Status = TaskStatusInit
	data.CreatedAt = GetCurrentTime()
	data.UpdatedAt = GetCurrentTime()
	err := util.Postgresdb.Create(&data).Error
	if err != nil {
		log.Fatal(err.Error())
		return 0
	}
	return int(data.ID)
}

func GetTask(id int) TaskTable {
	var Task TaskTable
	util.Postgresdb.Where("id = ?", id).First(&Task)
	return Task
}

func FinishTask(taskId int) {
	var Task TaskTable
	util.Postgresdb.Where("id = ?", taskId).First(&Task)
	if Task.Status == TaskStatusDone {
		return
	}
	Task.Status = TaskStatusDone
	Task.FinishTime = GetCurrentTime()
	Task.UpdatedAt = GetCurrentTime()
	util.Postgresdb.Save(&Task)
}

func BeginTask(taskId int) {
	var Task TaskTable
	util.Postgresdb.Where("id = ?", taskId).First(&Task)
	if Task.Status == TaskStatusDoing {
		return
	}
	Task.Status = TaskStatusDoing
	Task.BeginTime = GetCurrentTime()
	Task.UpdatedAt = GetCurrentTime()
	util.Postgresdb.Save(&Task)
}

func ListTask(keyword string) []TaskTable {
	var Tasks []TaskTable
	query := util.Postgresdb.Where("status != ?", TaskStatusDelete)
	if keyword != "" {
		query = query.Where("name like ?", "%"+keyword+"%")
	}
	query.Find(&Tasks).Limit(100)
	return Tasks
}

func DeleteTask(id uint) {
	var Task TaskTable
	util.Postgresdb.Where("id = ?", id).First(&Task)
	if Task.Status == TaskStatusDelete {
		return
	}
	Task.Status = TaskStatusDelete
	Task.UpdatedAt = GetCurrentTime()
	util.Postgresdb.Save(&Task)
}

func UpdateTask(data *TaskTable) {
	var Task TaskTable
	util.Postgresdb.Where("id = ?", data.ID).First(&Task)
	if Task.Status == TaskStatusDelete {
		return
	}
	if data.Name != "" {
		Task.Name = data.Name
	}
	if data.Tag != "" {
		Task.Tag = data.Tag
	}
	if data.Status != 0 {
		Task.Status = data.Status
		if data.Status == TaskStatusDone {
			Task.FinishTime = GetCurrentTime()
		} else if data.Status == TaskStatusDoing {
			Task.BeginTime = GetCurrentTime()
		} else if data.Status == TaskStatusInit {
			Task.BeginTime = ""
			Task.FinishTime = ""
		}
	}
	if data.Detail != "" {
		Task.Detail = data.Detail
	}
	Task.UpdatedAt = GetCurrentTime()
	util.Postgresdb.Save(&Task)
}

func GetCurrentTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
