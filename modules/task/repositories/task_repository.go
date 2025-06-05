package repositories

import (
	"tasks.com/config/database"
	"tasks.com/modules/task/domain/models"
	abstractRepositories "tasks.com/modules/task/domain/repositories"
)

type taskRepository struct {
	abstractRepositories.TaskRepository
	conn database.Connection
}

func NewTaskRepository(conn database.Connection) abstractRepositories.TaskRepository {
	return &taskRepository{
		conn: conn,
	}
}

func (r *taskRepository) GetAll(take, offset uint) ([]*models.Task, uint, error) {
	db, err := r.conn.Open()
	if err != nil {
		return nil, 0, err
	}
	defer db.Close()

	rows, err := db.Query(`SELECT id, name, description, expires_at FROM "tasks" ORDER BY "id" LIMIT $1 OFFSET $2`, take, offset)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	tasks := make([]*models.Task, 0)
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Name, &task.Description, &task.ExpiresAt)
		if err != nil {
			return nil, 0, err
		}

		tasks = append(tasks, &task)
	}

	var total uint
	err = db.QueryRow(`SELECT COUNT(*) FROM tasks`).Scan(&total)
	if err != nil {
		return nil, 0, err
	}

	return tasks, total, nil
}

func (r *taskRepository) GetByID(id uint) (*models.Task, error) {
	db, err := r.conn.Open()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var task models.Task
	query := `SELECT id, name, description, expires_at, created_at, updated_at WHERE id = $1`
	err = db.QueryRow(query, id).Scan(&task.ID,
		&task.Name,
		&task.Description,
		&task.ExpiresAt,
		&task.CreatedAt,
		&task.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *taskRepository) Add(task models.Task) (uint, error) {
	db, err := r.conn.Open()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	query := `INSERT INTO "tasks" (name,
								   description,
		                           expires_at,
								   created_at,
								   updated_at)
              VALUES ($1, $2, $3, $4, $5) returning "id"`

	var id uint
	err = db.QueryRow(query,
		task.Name,
		task.Description,
		task.ExpiresAt,
		task.CreatedAt,
		task.UpdatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *taskRepository) Update(task models.Task) error {
	db, err := r.conn.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	command := `UPDATE "tasks" SET name = $1, description = $2, expires_at = $3 WHERE id = $4`
	_, err = db.Exec(command, task.Name, task.Description, task.ExpiresAt, task.ID)
	if err != nil {
		return err
	}

	return nil
}

func (r *taskRepository) Delete(id uint) error {
	db, err := r.conn.Open()
	if err != nil {
		return err
	}
	defer db.Close()

	command := `DELETE FROM "tasks" WHERE id = $1`
	_, err = db.Exec(command, id)
	if err != nil {
		return err
	}

	return nil
}
