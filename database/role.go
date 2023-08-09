package database

import (
	"database/sql"
)

type CreateRoleRes struct {
	Id   int64  `json:"id"`
	Role string `json:"role"`
}

func CreateRole(db *sql.DB, role string) (CreateRoleRes, error) {
	var c CreateRoleRes

	row := db.QueryRow(`
		INSERT INTO role (role)
		VALUES ($1) 
		RETURNING id, role
	`, role)

	err := row.Scan(
		&c.Id,
		&c.Role,
	)

	if err != nil {
		return c, err
	}

	return c, nil
}

func MountRole(db *sql.DB, idAccount int64, idRole int64) (int64, error) {
	var id int64

	row := db.QueryRow(`
		INSERT INTO account_role (account_id, role_id)
		VALUES ($1, $2)
		RETURNING id
	`, idAccount, idRole)

	err := row.Scan(&id)

	if err != nil {
		return id, err
	}

	return id, err
}
