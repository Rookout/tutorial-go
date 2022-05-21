package todos

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

var todoRepository = NewInMemoryTodoRepository()

func RegisterEndPoints(e *echo.Echo) {
	e.File("/", "web/index.html")
	e.Static("/web", "web")
	e.GET("/todos", getAllTodosHandler)
	e.POST("/todos", createTodoHandler)
	e.DELETE("/todos", deleteAllTodosHandler)

	e.GET("/todos/:id", getTodoHandler)
	e.DELETE("/todos/:id", deleteTodoHandler)
	e.PATCH("/todos/:id", updateTodoHandler)
}

func getAllTodosHandler(c echo.Context) error {
	todos := todoRepository.GetAll()
	todosJson, _ := json.Marshal(&todos)
	return c.String(http.StatusOK, string(todosJson))
}

func createTodoHandler(c echo.Context) (err error) {
	todo := new(Todo)
	if err = c.Bind(todo); err != nil {
		return err
	}

	todoRepository.Create(todo)
	todoJson, _ := json.Marshal(&todo)
	return c.String(http.StatusCreated, string(todoJson))
}

func deleteAllTodosHandler(c echo.Context) (err error) {
	todoRepository.DeleteAll()
	return c.NoContent(http.StatusNoContent)
}

func getTodoHandler(c echo.Context) (err error) {
	id := c.Param("id")
	if todo, err := todoRepository.Get(id); err != nil {
		return c.String(http.StatusNotFound, "Todo note was not found")
	} else {
		todoJson, _ := json.Marshal(&todo)
		return c.String(http.StatusOK, string(todoJson))
	}
}

func deleteTodoHandler(c echo.Context) (err error) {
	id := c.Param("id")
	if err := todoRepository.Delete(id); err != nil {
		return c.String(http.StatusNotFound, "Todo note was not found")
	} else {
		return c.NoContent(http.StatusNoContent)
	}
}

func updateTodoHandler(c echo.Context) (err error) {
	id := c.Param("id")

	todo, err := todoRepository.Get(id)
	if err != nil {
		return c.String(http.StatusNotFound, "Todo note was not found")
	}
	if err = c.Bind(todo); err != nil {
		return err
	}

	if err := todoRepository.Update(todo); err != nil {
		return c.String(http.StatusNotFound, "Todo note was not found")
	} else {
		todoJson, _ := json.Marshal(&todo)
		return c.String(http.StatusOK, string(todoJson))
	}
}
