package apiv1

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/kkirsche/cronmon/libdb"
	// enables the use of database/sql and sqlx
	_ "github.com/lib/pq"

	"github.com/labstack/echo"
)

type task struct {
	ID             string    `json:"id" db:"id"`
	Name           string    `json:"name" db:"name"`
	Description    string    `json:"description" db:"description"`
	CronExpression string    `json:"cron_expression" db:"cron_expression"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	CreatedBy      string    `json:"created_by" db:"created_by"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
	UpdatedBy      string    `json:"updated_by" db:"updated_by"`
	LastStarted    time.Time `json:"last_started" db:"last_started"`
	LastCompleted  time.Time `json:"last_completed" db:"last_completed"`
}

type message struct {
	Message string `json:"message"`
}

func validateTask(t task) bool {
	if t.Name == "" || t.UpdatedBy == "" || t.CreatedBy == "" {
		return false
	}

	return true
}

// CreateTask is used to create a new task object
func CreateTask(c echo.Context) error {
	t := task{}
	if err := c.Bind(&t); err != nil {
		return err
	}

	t.CreatedAt = time.Now().UTC()
	t.CreatedBy = "system"
	t.UpdatedAt = time.Now().UTC()
	t.UpdatedBy = "system"

	db, err := sql.Open(libdb.Type, libdb.ConnectionURL)
	if err != nil {
		c.Logger().Errorf("failed to connect to database with error: %s", err.Error())
		return err
	}
	defer db.Close()

	err = db.QueryRow(`INSERT INTO tasks
		(name, description, cron_expression, created_at, created_by, updated_at, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		t.Name, t.Description, t.CronExpression, t.CreatedAt, t.CreatedBy, t.UpdatedAt, t.UpdatedBy).Scan(&t.ID)
	if err != nil {
		c.Logger().Errorf("failed to insert task into database with error: %s", err.Error())
		return err
	}

	return c.JSON(http.StatusCreated, t)
}

// GetTask is used to retrieve a specific task object based on it's ID
func GetTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, message{Message: "Not Found"})
	}
	t := task{}
	fmt.Println(id)

	return c.JSON(http.StatusOK, t)
}

// GetTasks is used to retrieve all tasks (or at least a page of tasks)
func GetTasks(c echo.Context) error {
	tasks := []task{task{}}
	return c.JSON(http.StatusOK, tasks)
}

// UpdateTask is used to update an existing task object
func UpdateTask(c echo.Context) error {
	t := new(task)
	if err := c.Bind(t); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, t)
}

// DeleteTask is used to delete a task object
func DeleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, message{Message: "Not Found"})
	}
	fmt.Println(id)

	return c.NoContent(http.StatusNoContent)
}

func StartTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, message{Message: "Not Found"})
	}
	fmt.Println(id)

	return c.NoContent(http.StatusNoContent)
}

func CompleteTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, message{Message: "Not Found"})
	}
	fmt.Println(id)
	return c.NoContent(http.StatusNoContent)
}
