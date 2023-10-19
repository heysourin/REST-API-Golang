package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "Getting a space suit", Completed: false},
	{ID: "2", Item: "Getting a spaceship", Completed: false},
	{ID: "3", Item: "Going to Mars", Completed: false},
}

func main() {
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.POST("/todos", addTodo)
	router.Run(":9000") // Corrected address
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}


func addTodo(context *gin.Context) {
	// Data from addTodo will be inside the request body and should come as JSON
	var newTodo = []todo{}// create an array of struct if you want to send multiple json data at once --> Take care about the append function

	err := context.BindJSON(&newTodo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todos = append(todos, newTodo...)
	context.IndentedJSON(http.StatusCreated, newTodo)
}
