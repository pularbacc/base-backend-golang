package database

import "database/sql"

type ProfileReq struct {
	Name       string `json:"name"`
	Account_id int64  `json:"account_id"`
}

type ProfileRes struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Account_id string `json:"account_id"`
}

func CreateProfile(db *sql.DB, p ProfileReq) (ProfileRes, error) {
	var profile ProfileRes

	row := db.QueryRow(`
		INSERT INTO profile (name, account_id)
		VALUES ($1, $2)
		RETURNING id, name, account_id
	`, p.Name, p.Account_id)

	err := row.Scan(
		&profile.Id,
		&profile.Name,
		&profile.Account_id,
	)

	if err != nil {
		return profile, err
	}

	return profile, nil
}
