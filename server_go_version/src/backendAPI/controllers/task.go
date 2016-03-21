package controllers

import (
	"backendAPI/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
)

// Operations about spark tasks
type TaskController struct {
	beego.Controller
}

// @Title create
// @Description create task
// @Param	body		body 	models.Object	true		"The task content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (t *TaskController) Post() {
	var task models.Task
	//var err error
	json.Unmarshal(t.Ctx.Input.RequestBody, &task)
	task.CreateAt = time.Now()
	taskid := models.AddTask(task)
	ta, _ := models.GetTask(taskid)
	ta.Start()
	t.Data["json"] = map[string]string{"TaskId": taskid, "Status": "Started"}
	t.ServeJSON()
}

// @Title Get
// @Description find task by taskid
// @Param	taskId		path 	string	true		"the taskid you want to get"
// @Success 200 {task} models.Object
// @Failure 403 :taskId is empty
// @router /:taskId [get]
func (t *TaskController) Get() {
	taskId := t.Ctx.Input.Param(":taskId")
	if taskId != "" {
		task, err := models.GetTask(taskId)
		if err != nil {
			t.Data["json"] = err.Error()
		} else {
			t.Data["json"] = task
		}
	}
	t.ServeJSON()
}

// @Title Check Status
// @Description check job status by taskid
// @Param	taskId		path 	string	true		"the taskid you want to get"
// @Success 200 {task} models.Object
// @Failure 403 :taskId is empty
// @router /check/:taskId [get]
func (t *TaskController) Check() {
	taskId := t.Ctx.Input.Param(":taskId")
	if taskId != "" {
		task, err := models.GetTask(taskId)
		if err != nil {
			t.Data["json"] = err.Error()
		} else {
			result, err := models.DBGet(taskId)
			//var duration string // TODO
			if err != nil {
				t.Data["json"] = map[string]string{
					"duration":  "6.341 secs",
					"startTime": task.CreateAt.String(),
					"result":    result,
					"status":    "",
					"jobId":     taskId,
				}
			} else {
				t.Data["json"] = err.Error()
			}
		}
	}
	t.ServeJSON()
}

// @Title GetAll
// @Description get all tasks
// @Success 200 {task} models.Task
// @Failure 403 :taskId is empty
// @router / [get]
func (t *TaskController) GetAll() {
	task := models.GetAllTasks()
	t.Data["json"] = task
	t.ServeJSON()
}

// @Title delete
// @Description delete the task
// @Param	taskId		path 	string	true "The taskId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 taskId is empty
// @router /:taskId [delete]
func (t *TaskController) Delete() {
	taskId := t.Ctx.Input.Param(":taskId")
	models.Delete(taskId)
	t.Data["json"] = "delete success!"
	t.ServeJSON()
}
