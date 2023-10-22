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

	router.Run(":9000") // Corrected address
}

func getTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodo(context *gin.Context) {
	// Data from addTodo will be inside the request body and should come as JSON
	var newTodo = []todo{} // create an array of struct if you want to send multiple json data at once --> Take care about the append function

	err := context.BindJSON(&newTodo)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	todos = append(todos, newTodo...)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

// func getTodoIndexById(id string) (*todo, error) {
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
	context.IndentedJSON(http.StatusOK, todo)
}

func updateTodo(context *gin.Context) {
	// Get the ID parameter from the URL
	id := context.Param("id")

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

	// Respond with the updated todo
	context.IndentedJSON(http.StatusOK, todo)
}

func deleteTodo(context *gin.Context) {
	// Get the ID parameter from the URL
	id := context.Param("id")

	// Find the index of the todo item with the matching ID
	index, err := getTodoIndexById(id)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	// Remove the todo item from the todos slice
	todos = append(todos[:index], todos[index+1:]...)

	// Respond with a message to confirm the deletion
	context.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
