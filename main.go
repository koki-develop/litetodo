package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	if err := run(); err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}
}

type Task struct {
	ID        uint      `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type TaskCreateInput struct {
	Title string `json:"title"`
}

type TaskUpdateInput struct {
	Title     *string `json:"title"`
	Completed *bool   `json:"completed"`
}

func run() error {
	/*
	 * Open database
	 */

	db, err := gorm.Open(sqlite.Open("./todo.db"), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	/*
	 * Migrate database
	 */

	if err := db.AutoMigrate(&Task{}); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	/*
	 * Setup routes
	 */

	e := echo.New()

	e.GET("/tasks", func(c echo.Context) error {
		var ts []Task
		if err := db.Find(&ts).Error; err != nil {
			return fmt.Errorf("failed to find tasks: %w", err)
		}

		return c.JSON(http.StatusOK, ts)
	})

	e.GET("/tasks/:id", func(c echo.Context) error {
		var t Task
		if err := db.First(&t, c.Param("id")).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(http.StatusNotFound, map[string]string{"message": "Not Found"})
			}
			return fmt.Errorf("failed to find task: %w", err)
		}

		return c.JSON(http.StatusOK, t)
	})

	e.POST("/tasks", func(c echo.Context) error {
		var ipt TaskCreateInput
		if err := c.Bind(&ipt); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad Request"})
		}

		var t = Task{Title: ipt.Title}
		if err := db.Create(&t).Error; err != nil {
			return fmt.Errorf("failed to create task: %w", err)
		}

		return c.JSON(http.StatusOK, t)
	})

	e.PATCH("/tasks/:id", func(c echo.Context) error {
		var ipt TaskUpdateInput
		if err := c.Bind(&ipt); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"message": "Bad Request"})
		}

		var t Task
		if err := db.First(&t, c.Param("id")).Error; err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(http.StatusNotFound, map[string]string{"message": "Not Found"})
			}
			return fmt.Errorf("failed to find task: %w", err)
		}

		if ipt.Title != nil {
			t.Title = *ipt.Title
		}
		if ipt.Completed != nil {
			t.Completed = *ipt.Completed
		}
		if err := db.Save(&t).Error; err != nil {
			return fmt.Errorf("failed to save task: %w", err)
		}

		return c.JSON(http.StatusOK, t)
	})

	e.DELETE("/tasks/:id", func(c echo.Context) error {
		var t Task
		if err := db.First(&t, c.Param("id")).Error; err != nil {
			return fmt.Errorf("failed to find task: %w", err)
		}

		if err := db.Delete(&t).Error; err != nil {
			return fmt.Errorf("failed to delete task: %w", err)
		}

		return c.NoContent(http.StatusNoContent)
	})

	/*
	 * Start server
	 */

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := e.Start(fmt.Sprintf(":%s", port)); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
