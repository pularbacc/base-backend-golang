package database

import (
	"database/sql"
	"fmt"
)

type AccountInfo struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	Name  string `json:"name"`
}

type AccountCreateReq struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Account struct {
	Id    int64  `json:"id"`
	Email string `json:"email"`
}

func CreateAccount(db *sql.DB, accountReq AccountCreateReq) (Account, error) {
	var account Account

	row := db.QueryRow(`
		INSERT INTO account (email, password)
		VALUES ($1, $2) 
		RETURNING id, email
	`, accountReq.Email, accountReq.Password)

	err := row.Scan(
		&account.Id,
		&account.Email,
	)

	if err != nil {
		return account, err
	}

	return account, nil
}

func GetListAccount(db *sql.DB) ([]Account, error) {
	var accounts []Account

	rows, err := db.Query(`
		SELECT id, email
		FROM account`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var acc Account
		if err := rows.Scan(&acc.Id, &acc.Email); err != nil {
			return nil, err
		}
		accounts = append(accounts, acc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return accounts, nil
}

func GetAccountInfoById(db *sql.DB, id int64) (AccountInfo, error) {
	var account AccountInfo

	row := db.QueryRow(`
		SELECT
			u.id AS account_id,
			u.email,
			r.role,
			p.name
		FROM
			account u
			INNER JOIN account_role ur ON u.id = ur.account_id
			INNER JOIN role r ON ur.role_id = r.id
			INNER JOIN profile p ON u.id = p.account_id
		WHERE
			u.id = $1`, id)

	err := row.Scan(
		&account.Id,
		&account.Email,
		&account.Role,
		&account.Name,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return account, fmt.Errorf("not found account")
		}

		return account, err
	}

	return account, nil
}

func GetAccountInfoByEmail(db *sql.DB, email string) (AccountInfo, error) {
	var account AccountInfo

	row := db.QueryRow(`
		SELECT
			u.id AS account_id,
			u.email,
			r.role,
			p.name
		FROM
			account u
			INNER JOIN account_role ur ON u.id = ur.account_id
			INNER JOIN role r ON ur.role_id = r.id
			INNER JOIN profile p ON u.id = p.account_id
		WHERE
			u.email = $1`, email)

	err := row.Scan(
		&account.Id,
		&account.Email,
		&account.Role,
		&account.Name,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return account, fmt.Errorf("not found account")
		}

		return account, err
	}

	return account, nil
}

func GetPassByEmail(db *sql.DB, email string) (string, error) {
	var password string

	row := db.QueryRow(`
		SELECT password 
		FROM account 
		WHERE email = $1`, email)

	err := row.Scan(
		&password,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return password, fmt.Errorf("not found account")
		}

		return password, err
	}

	return password, nil
}
