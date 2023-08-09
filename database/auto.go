package database

import "database/sql"

type Cmd struct {
	Id     int64  `json:"id"`
	Method string `json:"method"`
	Url    string `json:"url"`
	Body   string `json:"body"`
}

type Auto struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	Cmds []Cmd  `json:"cmds"`
}

func CreateAuto(db *sql.DB, a Auto) (Auto, error) {
	// Insert the "auto" record
	query := "INSERT INTO auto (name) VALUES ($1) RETURNING id"
	var autoID int64
	err := db.QueryRow(query, a.Name).Scan(&autoID)
	if err != nil {
		return Auto{}, err
	}

	// Insert the "cmd" records associated with the "auto" record
	cmdQuery := "INSERT INTO auto_api (method, url, body, auto_id) VALUES ($1, $2, $3, $4)"
	for _, cmd := range a.Cmds {
		_, err := db.Exec(cmdQuery, cmd.Method, cmd.Url, cmd.Body, autoID)
		if err != nil {
			return Auto{}, err
		}
	}

	// Return the created "auto" record with the updated ID
	a.Id = autoID
	return a, nil
}

func UpdateAuto(db *sql.DB, a Auto) error {
	// Update the "auto" record
	query := "UPDATE auto SET name = $1 WHERE id = $2"
	_, err := db.Exec(query, a.Name, a.Id)
	if err != nil {
		return err
	}

	// Delete existing "cmd" records associated with the "auto" record
	deleteQuery := "DELETE FROM auto_api WHERE auto_id = $1"
	_, err = db.Exec(deleteQuery, a.Id)
	if err != nil {
		return err
	}

	// Insert the updated "cmd" records associated with the "auto" record
	cmdQuery := "INSERT INTO auto_api (method, url, body, auto_id) VALUES ($1, $2, $3, $4)"
	for _, cmd := range a.Cmds {
		_, err := db.Exec(cmdQuery, cmd.Method, cmd.Url, cmd.Body, a.Id)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetAutoList(db *sql.DB) ([]Auto, error) {
	// Get the list of all "auto" records
	query := "SELECT id, name FROM auto"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var autos []Auto
	for rows.Next() {
		var auto Auto
		err := rows.Scan(&auto.Id, &auto.Name)
		if err != nil {
			return nil, err
		}
		autos = append(autos, auto)
	}

	return autos, nil
}

func GetAutoByID(db *sql.DB, id int64) (Auto, error) {
	// Get the "auto" record by ID
	query := "SELECT id, name FROM auto WHERE id = $1"
	row := db.QueryRow(query, id)

	var auto Auto
	err := row.Scan(&auto.Id, &auto.Name)
	if err != nil {
		return Auto{}, err
	}

	// Get the associated "cmd" records
	cmdQuery := "SELECT id, method, url, body FROM auto_api WHERE auto_id = $1"
	rows, err := db.Query(cmdQuery, id)
	if err != nil {
		return Auto{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var cmd Cmd
		err := rows.Scan(&cmd.Id, &cmd.Method, &cmd.Url, &cmd.Body)
		if err != nil {
			return Auto{}, err
		}
		auto.Cmds = append(auto.Cmds, cmd)
	}

	return auto, nil
}

func DeleteAuto(db *sql.DB, id int64) error {
	// Delete the "auto" record
	query := "DELETE FROM auto WHERE id = $1"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	// Delete the associated "cmd" records
	deleteQuery := "DELETE FROM auto_api WHERE auto_id = $1"
	_, err = db.Exec(deleteQuery, id)
	if err != nil {
		return err
	}

	return nil
}
