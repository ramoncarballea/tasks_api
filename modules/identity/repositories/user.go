package repositories

import (
	"github.com/google/uuid"
	"tasks.com/config/database"
	"tasks.com/modules/identity/domain/models"
	"tasks.com/modules/identity/domain/repositories"
	"time"
)

type userRepository struct {
	repositories.UserRepository
	db database.Connection
}

func NewUserRepository(db database.Connection) repositories.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) GetAll(take, offset uint) ([]*models.User, uint, error) {
	con, err := r.db.Open()
	if err != nil {
		return nil, 0, err
	}
	defer con.Close()

	rows, err := con.Query(`SELECT id, first_name, last_name, email FROM "users" ORDER BY "id" LIMIT $1 OFFSET $2`, take, offset)
	if err != nil {
		return nil, 0, err
	}

	var users []*models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
		if err != nil {
			return nil, 0, err
		}

		users = append(users, &user)
	}

	var count uint
	err = con.QueryRow(`SELECT COUNT(1) FROM "users"`).Scan(&count)
	if err != nil {
		return nil, 0, err
	}

	return users, count, nil
}

func (r *userRepository) GetByID(id uuid.UUID) (*models.User, error) {
	con, err := r.db.Open()
	if err != nil {
		return nil, err
	}
	defer con.Close()

	query := `SELECT 
    			u.id, 
    			u.first_name, 
    			u.last_name, 
    			u.email, 
    			u.created_at, 
    			u.updated_at,
    			r.id,
				r.name,
				p.id,
				p.name
			  FROM "users" u 
			  LEFT JOIN "user_roles" ur ON ur.user_id = u.id
			  LEFT JOIN "roles" r ON r.id = ur.role_id
			  LEFT JOIN "role_permissions" rp ON rp.role_id = r.id
			  LEFT JOIN "permissions" p ON p.id = rp.permission_id
			  WHERE u.id = $1`

	rows, err := con.Query(query, id)
	if err != nil {
		return nil, err
	}

	var user models.User
	var rolesMap = make(map[uint]*models.Role)

	for rows.Next() {
		var role models.Role
		var permission models.Permission
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email, &user.CreatedAt, &user.UpdatedAt, &role.ID, &role.Name, &permission.ID, &permission.Name)
		if err != nil {
			return nil, err
		}

		if sRole, ok := rolesMap[role.ID]; !ok {
			role.Permissions = append(role.Permissions, &permission)
			rolesMap[role.ID] = &role
			user.Roles = append(user.Roles, &role)
		} else {
			sRole.Permissions = append(sRole.Permissions, &permission)
		}
	}

	return &user, nil
}

func (r *userRepository) Add(user models.User) (uuid.UUID, error) {
	con, err := r.db.Open()
	if err != nil {
		return uuid.Nil, err
	}
	defer con.Close()

	query := `INSERT INTO "users" (first_name, password, last_name, email, created_at) VALUES ($1, $2, $3, $4, $5) returning "id"`
	var id uuid.UUID
	err = con.QueryRow(query, user.FirstName, user.Password, user.LastName, user.Email, time.Now()).Scan(&id)
	if err != nil {
		return uuid.Nil, err
	}

	tx, err := con.Begin()
	if err != nil {
		return uuid.Nil, err
	}

	for _, role := range user.Roles {
		query = `INSERT INTO "user_roles" (user_id, role_id) VALUES ($1, $2)`
		_, err = tx.Exec(query, id, role.ID)
		if err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				return uuid.Nil, txErr
			}
			return uuid.Nil, err
		}
	}

	if err = tx.Commit(); err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (r *userRepository) Update(user models.User) error {
	con, err := r.db.Open()
	if err != nil {
		return err
	}
	defer con.Close()

	tx, err := con.Begin()
	if err != nil {
		return err
	}

	query := `UPDATE "users" SET first_name = $1, last_name = $2, email = $3, updated_at = $4 WHERE id = $5`
	_, err = tx.Exec(query, user.FirstName, user.LastName, user.Email, time.Now(), user.ID)
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return txErr
		}
		return err
	}

	_, err = tx.Exec(`DELETE FROM "user_roles" WHERE user_id = $1`, user.ID)
	if err != nil {
		if txErr := tx.Rollback(); txErr != nil {
			return txErr
		}
		return err
	}

	for _, role := range user.Roles {
		query = `INSERT INTO "user_roles" (user_id, role_id) VALUES ($1, $2)`
		_, err = tx.Exec(query, user.ID, role.ID)
		if err != nil {
			if txErr := tx.Rollback(); txErr != nil {
				return txErr
			}
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(id uuid.UUID) error {
	con, err := r.db.Open()
	if err != nil {
		return err
	}
	defer con.Close()

	query := `DELETE FROM "users" WHERE id = $1`
	_, err = con.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}
