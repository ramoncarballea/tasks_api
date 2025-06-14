package repositories

import (
	"tasks.com/config/database"
	"tasks.com/modules/identity/domain/models"
	"tasks.com/modules/identity/domain/repositories"
	"time"
)

type permissionRepository struct {
	repositories.PermissionRepository
	db database.Connection
}

func NewPermissionRepository(db database.Connection) repositories.PermissionRepository {
	return &permissionRepository{
		db: db,
	}
}

func (r *permissionRepository) GetAll() ([]*models.Permission, error) {
	con, err := r.db.Open()
	if err != nil {
		return nil, err
	}
	defer con.Close()

	rows, err := con.Query(`SELECT id, name FROM "permissions"`)
	if err != nil {
		return nil, err
	}

	var permissions []*models.Permission
	for rows.Next() {
		var permission models.Permission
		err := rows.Scan(&permission.ID, &permission.Name)
		if err != nil {
			return nil, err
		}

		permissions = append(permissions, &permission)
	}

	return permissions, nil
}

func (r *permissionRepository) GetByName(name string) (*models.Permission, error) {
	con, err := r.db.Open()
	if err != nil {
		return nil, err
	}
	defer con.Close()

	var permission models.Permission
	query := `SELECT id, name FROM "permissions" WHERE name = $1`
	err = con.QueryRow(query, name).Scan(&permission.ID, &permission.Name)
	if err != nil {
		return nil, err
	}

	return &permission, nil
}

func (r *permissionRepository) Add(name string) (uint, error) {
	con, err := r.db.Open()
	if err != nil {
		return 0, err
	}
	defer con.Close()

	query := `INSERT INTO "permissions" (name, created_at) VALUES ($1, $2) returning "id"`
	var id uint
	err = con.QueryRow(query, name, time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *permissionRepository) AddRange(permissions []models.Permission) error {
	con, err := r.db.Open()
	if err != nil {
		return err
	}
	defer con.Close()

	tx, err := con.Begin()
	if err != nil {
		return err
	}

	for _, permission := range permissions {
		query := `INSERT INTO "permissions" (name, created_at) VALUES ($1, $2) returning "id"`
		var id uint
		err = tx.QueryRow(query, permission.Name, time.Now()).Scan(&id)
		if err != nil {
			txErr := tx.Rollback()
			if txErr != nil {
				return txErr
			}
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
