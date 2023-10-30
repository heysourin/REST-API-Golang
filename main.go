package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
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

	router.GET("/todos", getTodos)          //Read
	router.GET("/todos/:id", getTodo)       //Read
	router.POST("/todos", addTodo)          //Create
	router.PUT("/todos/:id", updateTodo)    //Update
	router.DELETE("/todos/:id", deleteTodo) //Delete

	router.Run(":9000") // Port on local machine
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	var newTodo = []todo{} // Creating a variable of todo struct to send payload with it

	err := context.BindJSON(&newTodo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todos = append(todos, newTodo...) // adding the payload data to the existing data
	context.IndentedJSON(http.StatusCreated, newTodo)
}

// Checkpoint: requested id exists or not
func getTodoIndexById(id string) (int, error) {
	for i, t := range todos {
		if t.ID == id {
			return i, nil
		}
	}
	return -1, errors.New("Todo not found")
}

func getTodo(context *gin.Context) {
	id := context.Param("id")
	todo, err := getTodoIndexById(id)
	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return

	}
	context.IndentedJSON(http.StatusOK, todos[todo])
}

func updateTodo(context *gin.Context) {
	id := context.Param("id") // Get the ID parameter from the URL

	// Find the todo item with the matching ID
	todo, err := getTodoIndexById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	// Bind the JSON data from the request body into the existing todo struct
	if err := context.BindJSON(todo); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.IndentedJSON(http.StatusOK, todo) // Respond with the updated todo
}

func deleteTodo(context *gin.Context) {
	id := context.Param("id") // Get the ID parameter from the URL

	// Find the index of the todo item with the matching ID
	index, err := getTodoIndexById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todos = append(todos[:index], todos[index+1:]...) // Remove the todo item from the todos slice

	context.JSON(http.StatusOK, gin.H{"message": "Todo deleted"}) // Respond with a message to confirm the deletion
}
