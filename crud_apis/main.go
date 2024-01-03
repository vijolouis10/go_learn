package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Todo struct {
	ID   int
	TASK string
}

var todos []Todo

func getTodos(c *gin.Context) {
	c.JSON(http.StatusOK, todos)
}

func getTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for index, todo := range todos {
		if todos[index].ID == id {
			c.JSON(http.StatusOK, todo)
			return
		}
	}
}

func createTodo(c *gin.Context) {
	var todo Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	todos = append(todos, todo)

	c.JSON(http.StatusOK, gin.H{"message": "Person created successfully"})
}

func updateTodo(c *gin.Context) {
	var todo Todo

	id, _ := strconv.Atoi(c.Param("id"))

	var found bool
	for index, t := range todos {
		if t.ID == id {
			found = true
			if !found {
				c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
				return
			} else {
				todos[index].TASK = todo.TASK
			}
		}
		c.JSON(http.StatusOK, gin.H{"message": "Todo updated successfully"})
	}
}

func deleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	for index, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:index], todos[index+1:]...)
		}
		c.JSON(http.StatusOK, gin.H{"message": "Person deleted successfully"})
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/getTodos", getTodos)
	r.GET("/getTodo/:id", getTodo)
	r.POST("/createTodo", createTodo)
	r.PUT("/updateTodo/:id", updateTodo)
	r.DELETE("/deleteTodo/:id", deleteTodo)

	return r
}

func main() {
	router := setupRouter()
	if err := router.Run(":8000"); err != nil {
		log.Fatalf("Failed to run server %s", err)
	}
}
