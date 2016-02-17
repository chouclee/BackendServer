package controllers

import (
	"backendAPI/models"
	"encoding/json"
	"github.com/astaxie/beego"
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
	json.Unmarshal(t.Ctx.Input.RequestBody, &task)
	taskid := models.AddTask(task)
	t.Data["json"] = map[string]string{"TaskId": taskid}
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

// @Title update
// @Description update the task
// @Param	taskId		path 	string	true		"The taskid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {task} models.Object
// @Failure 403 :taskId is empty
// @router /:taskId [put]
/*func (t *TaskController) Put() {
	taskId := t.Ctx.Input.Param(":taskId")
	var task models.Task
	json.Unmarshal(t.Ctx.Input.RequestBody, &task)

	err := models.UpdateT(taskId, t.Score)
	if err != nil {
		t.Data["json"] = err.Error()
	} else {
		t.Data["json"] = "update success!"
	}
	t.ServeJSON()
}*/

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
