package repositories

import (
	"tasks.com/config/database"
	"tasks.com/modules/identity/domain/models"
	"tasks.com/modules/identity/domain/repositories"
	"time"
)

type roleRepository struct {
	repositories.RoleRepository
	db database.Connection
}

func NewRoleRepository(db database.Connection) repositories.RoleRepository {
	return &roleRepository{
		db: db,
	}
}

func (r *roleRepository) GetAll() ([]*models.Role, error) {
	con, err := r.db.Open()
	if err != nil {
		return nil, err
	}
	defer con.Close()

	rows, err := con.Query(`SELECT id, name FROM "roles"`)
	if err != nil {
		return nil, err
	}

	var roles []*models.Role
	for rows.Next() {
		var role models.Role
		err := rows.Scan(&role.ID, &role.Name)
		if err != nil {
			return nil, err
		}

		roles = append(roles, &role)
	}

	return roles, nil
}

func (r *roleRepository) GetByName(name string) (*models.Role, error) {
	con, err := r.db.Open()
	if err != nil {
		return nil, err
	}
	defer con.Close()

	var role models.Role

	query := `SELECT r.id, r.name, p.id, p.name FROM "roles" r 
                    LEFT JOIN "role_permissions" rp ON r.id = rp.role_id
                    LEFT JOIN "permissions" p ON p.id = rp.permission_id
    		  WHERE r.name = $1`

	rows, err := con.Query(query, name)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var permission models.Permission
		err = rows.Scan(&role.ID, &role.Name, &permission.ID, &permission.Name)
		if err != nil {
			return nil, err
		}

		role.Permissions = append(role.Permissions, &permission)
	}

	return &role, nil
}

func (r *roleRepository) Add(role models.Role) (uint, error) {
	con, err := r.db.Open()
	if err != nil {
		return 0, err
	}
	defer con.Close()

	query := `INSERT INTO "roles" (name, created_at) VALUES ($1, $2) returning "id"`

	var id uint
	err = con.QueryRow(query, role.Name, time.Now()).Scan(&id)
	if err != nil {
		return 0, err
	}

	tx, err := con.Begin()
	if err != nil {
		return 0, err
	}

	for _, permission := range role.Permissions {
		_, err := tx.Exec(`INSERT INTO "role_permissions" (role_id, permission_id) VALUES ($1, $2)`, id, permission.ID)
		if err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				return 0, txErr
			}
			return 0, err
		}
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *roleRepository) Delete(id uint) error {
	con, err := r.db.Open()
	if err != nil {
		return err
	}
	defer con.Close()

	query := `DELETE FROM "roles" WHERE id = $1`
	_, err = con.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *roleRepository) Update(role models.Role) error {
	con, err := r.db.Open()
	if err != nil {
		return err
	}
	defer con.Close()

	query := `UPDATE "roles" SET name = $1 WHERE id = $2`
	_, err = con.Exec(query, role.Name, role.ID)
	if err != nil {
		return err
	}

	tx, err := con.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec(`DELETE FROM "role_permissions" WHERE role_id = $1`, role.ID)
	if err != nil {
		txErr := tx.Rollback()
		if txErr != nil {
			return txErr
		}
		return err
	}

	for _, permission := range role.Permissions {
		_, err := tx.Exec(`INSERT INTO "role_permissions" (role_id, permission_id) VALUES ($1, $2)`, role.ID, permission)
		if err != nil {
			if txErr := tx.Rollback(); txErr != nil {
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

func (r *roleRepository) AddRange(roles []models.Role) error {
	con, err := r.db.Open()
	if err != nil {
		return err
	}
	defer con.Close()

	tx, err := con.Begin()
	if err != nil {
		return err
	}

	for _, role := range roles {
		query := `INSERT INTO "roles" (name, created_at) VALUES ($1, $2) RETURNING "id"`
		err = tx.QueryRow(query, role.Name, time.Now()).Scan(&role.ID)
		if err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				return txErr
			}
			return err
		}

		for _, permission := range role.Permissions {
			_, err := tx.Exec(`INSERT INTO "role_permissions" (role_id, permission_id) VALUES ($1, $2)`, role.ID, permission.ID)
			if err != nil {
				if txErr := tx.Rollback(); txErr != nil {
					return txErr
				}
				return err
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
