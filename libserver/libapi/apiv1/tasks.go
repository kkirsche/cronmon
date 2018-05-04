package apiv1

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorhill/cronexpr"
	"github.com/jmoiron/sqlx"
	"github.com/kkirsche/cronmon/libdb"
	"github.com/lib/pq"

	"github.com/labstack/echo"
)

type task struct {
	ID             int         `json:"id" db:"id"`
	Name           string      `json:"name" db:"name"`
	Description    string      `json:"description" db:"description"`
	CronExpression string      `json:"cron_expression" db:"cron_expression"`
	CreatedAt      pq.NullTime `json:"created_at" db:"created_at"`
	CreatedBy      string      `json:"created_by" db:"created_by"`
	UpdatedAt      pq.NullTime `json:"updated_at" db:"updated_at"`
	UpdatedBy      string      `json:"updated_by" db:"updated_by"`
	LastStarted    pq.NullTime `json:"last_started" db:"last_started"`
	LastCompleted  pq.NullTime `json:"last_completed" db:"last_completed"`
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

	t.CreatedAt.Time = time.Now().UTC()
	t.CreatedBy = "system"
	t.UpdatedAt.Time = time.Now().UTC()
	t.UpdatedBy = "system"

	if t.CronExpression == "" || t.Name == "" {
		return c.JSON(
			http.StatusBadRequest,
			message{Message: "missing required task field"})
	}

	_, err := cronexpr.Parse(t.CronExpression)
	if err != nil {
		return c.JSON(
			http.StatusBadRequest,
			message{Message: fmt.Sprintf(
				"cron expression error: %s", err.Error()),
			})
	}

	db, err := sqlx.Open(libdb.Type, libdb.ConnectionURL)
	if err != nil {
		c.Logger().Errorf("failed to connect to database with error: %s", err.Error())
		return c.JSON(http.StatusInternalServerError, message{Message: "internal server error"})
	}
	defer db.Close()

	err = db.QueryRow(`INSERT INTO tasks
		(name, description, cron_expression, created_at, created_by, updated_at, updated_by)
		VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`,
		t.Name, t.Description, t.CronExpression, t.CreatedAt.Time, t.CreatedBy, t.UpdatedAt.Time, t.UpdatedBy,
	).Scan(&t.ID)
	if err != nil {
		c.Logger().Errorf("failed to insert task into database with error: %s", err.Error())
		return c.JSON(http.StatusBadRequest, message{Message: "invalid request"})
	}

	return c.JSON(http.StatusCreated, t)
}

// GetTask is used to retrieve a specific task object based on it's ID
func GetTask(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound, message{Message: "Not Found"})
	}

	db, err := sqlx.Open(libdb.Type, libdb.ConnectionURL)
	if err != nil {
		c.Logger().Errorf("failed to connect to database with error: %s", err.Error())
		return err
	}
	defer db.Close()

	t := task{}
	err = db.Get(&t, "SELECT * FROM tasks WHERE id=$1", id)
	if err != nil {
		c.Logger().Errorf("failed to retrieve task with error: %s", err.Error())
		return c.JSON(http.StatusBadRequest, message{Message: "invalid request"})
	}

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
