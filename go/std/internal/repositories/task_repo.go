package repositories

import (
	"database/sql"
	"go-std/internal/models"

	"github.com/google/uuid"
)

type taskRepositoryImpl struct {
	db *sql.DB
}

type TaskRepository interface {
	CreateTask(task *models.Task) error
	GetTask(id uuid.UUID) (*models.Task, error)
	UpdateTask(id uuid.UUID, task *models.Task) error
	DeleteTask(id uuid.UUID) error
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepositoryImpl{db: db}
}

func (r *taskRepositoryImpl) CreateTask(task *models.Task) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if task.ID == uuid.Nil {
		task.ID, err = uuid.NewV7()
		if err != nil {
			return err
		}
	}

	_, err = tx.Exec("INSERT INTO tasks (title, description) VALUES (?, ?)", task.Title, task.Description)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *taskRepositoryImpl) GetTask(id uuid.UUID) (*models.Task, error) {
	row := r.db.QueryRow("SELECT id, title, description FROM tasks WHERE id = ?", id)
	var task models.Task
	err := row.Scan(&task.ID, &task.Title, &task.Description)
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *taskRepositoryImpl) GetAllTasks(limit int, offset int) (tasks []*models.Task, err error) {
	rows, err := r.db.Query("SELECT id, title, description FROM tasks LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks = make([]*models.Task, 0, limit)
	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Title, &task.Description)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, &task)
	}

	return tasks, nil
}

func (r *taskRepositoryImpl) UpdateTask(id uuid.UUID, task *models.Task) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE tasks SET title = ?, description = ? WHERE id = ?", task.Title, task.Description, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *taskRepositoryImpl) DeleteTask(id uuid.UUID) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec("UPDATE tasks SET deleted_at = NOW() WHERE id = ?", id)
	if err != nil {
		return err
	}

	return tx.Commit()
}
