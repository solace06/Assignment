package internal

import "github.com/gin-gonic/gin"

func TaskRoutes(router *gin.RouterGroup) {
	r := router.Group("/tasks")

	r.POST("", CreateTask)
	r.GET("", ListTasks)

}

func CreateTask(c *gin.Context) {
	var task CreateTaskRequest

	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(400, "invalid request")
		return
	}

	err = ValidateTask(task)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	insertedTask, err := NewTask(task)
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	response := &Response{
		Data:    *insertedTask,
		Message: "task created successfully",
	}

	c.JSON(201, response)
}

func ListTasks(c *gin.Context) {

	tasks, err := GetAllTasks()
	if err != nil {
		c.JSON(500, err.Error())
		return
	}

	response := &Response{
		Data:    tasks,
		Message: "tasks fetched successfully",
	}

	c.JSON(200, response)
}
