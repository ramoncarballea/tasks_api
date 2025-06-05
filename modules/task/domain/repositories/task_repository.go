package repositories

import (
	"tasks.com/modules/task/domain/models"
)

type TaskRepository interface {
	// GetAll returns a list of all tasks in the database,
	// along with the total count of tasks.
	//
	// The take parameter specifies the number of tasks to
	// return, while the offset parameter specifies the
	// offset from which to start returning tasks.
	GetAll(take uint, offset uint) ([]*models.Task, uint, error)
	// GetByID returns a task with the given ID.
	//
	// If no task with the given ID exists, GetByID returns nil and an error.
	GetByID(id uint) (*models.Task, error)
	// Add adds a new task to the database.
	//
	// The returned ID is the ID of the newly-created task.
	//
	// If an error occurs while adding the task, Add returns 0 and the error.
	Add(task models.Task) (uint, error)
	// Update updates a task in the database.
	//
	// If an error occurs while updating the task, Update returns the error.
	Update(task models.Task) error
	// Delete deletes a task with the given ID from the database.
	//
	// If the task does not exist, Delete returns an error.
	Delete(id uint) error
}
