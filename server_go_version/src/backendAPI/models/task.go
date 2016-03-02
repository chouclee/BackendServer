package models

import (
	"errors"
	"log"
	"strconv"
	"time"
)

var (
	Tasks map[string]*Task
)

type Task struct {
	TaskId    string
	UserId    string
	ModelTpye string
	CreatedAt Time
	Algorithm string
}

func init() {
	Tasks = make(map[string]*Task)
}

func (t *Task) Start() {
	id := t.TaskId
	// test database insert
	// client.Insert(id, "/home/a/b/trainingdata;/home/testingdata")
	// test database lookup
	storedPaths, err := dblib.client.Get(id)
	if err != nil {
		//w.Write([]byte("Job id does not exist!\n"))
		return
	}
	paths := strings.Split(storedPaths, ";") // split by ";"
	if len(paths) != 2 {                     // check paths
		log.Println("Paths in database is worng: " + storedPaths)
		//w.Write([]byte("Paths in database is worng: " + storedPaths))
		return
	}
	//w.Write([]byte("Job ID:" + jobId + "\n"))
	//w.Write([]byte("Training Data: " + paths[0] + "\n"))
	//w.Write([]byte("Testing Data:" + paths[1] + "\n"))

	// call external python script
	go func(id int) {
		//outFile := "/home/honeycomb/HoneyBuzzard/output/result_" + jobId + ".json"
		outDir := "/home/honeycomb/HoneyBuzzard/output"
		//output,err := exec.Command("/bin/spark-submit", "/home/honeycomb/SparkTeam/PySpark.py",
		//	paths[0], paths[1], outDir).Output()
		log.Println("testing data: " + paths[0])
		log.Println("training data:" + paths[1])
		log.Println("out dir: " + outDir)
		err := exec.Command("/bin/spark-submit", "/home/honeycomb/SparkTeam/PySpark.py",
			paths[0], paths[1], outDir).Run()
		if err != nil {
			log.Println(err)
		}
		//if err != nil {
		//	log.Fatal(err) // caution: log.Fatal may terminate the program
		//} else {
		//log.Println(string(output))
		filePath := outDir + "/result_" + jobId
		log.Println(filePath)
		err = os.Rename(outDir+"/part-00000", filePath)
		if err != nil {
			log.Println("file does not exist")
			return
		}
		dblib.client.Insert(id, filePath)
		//}
	}(id)

}

func AddTask(task Task) (TaskId string) {
	task.TaskId = "honeycomb" + strconv.FormatInt(time.Now().UnixNano(), 10)
	task.CreatedAt = time.Now()
	Tasks[task.TaskId] = &task
	return task.TaskId
}

func GetTask(TaskId string) (task *Task, err error) {
	if v, ok := Tasks[TaskId]; ok {
		return v, nil
	}
	return nil, errors.New("TaskId Not Exist")
}

func GetAllTasks() map[string]*Task {
	return Tasks
}

func DeleteTask(TaskId string) {
	delete(Tasks, TaskId)
}
