package s5_dbase

import (
	"database/sql"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func CreateUsersTable(dbase *sql.DB) error {
	userStmt, err := dbase.Prepare(`
		CREATE TABLE users(			
			"user_name" text NOT NULL PRIMARY KEY,
			"pass_hash" text
 		);
	`)
	if err != nil {
		return err
	}

	_, err = userStmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

func AddUser(db *sql.DB, username, pass string) error {
	phash, err := bcrypt.GenerateFromPassword([]byte(pass), 14)
	if err != nil {
		return fmt.Errorf("could not make password for %s, %w", pass, err)
	}

	stmt, err := db.Prepare(`INSERT INTO users (user_name, pass_hash) VALUES (?, ?)`)
	if err != nil {
		return fmt.Errorf("could not prepare to add %s to database : %w", username, err)
	}
	_, err = stmt.Exec(username, phash)
	if err != nil {
		return fmt.Errorf("could not add %s to database : %w", username, err)
	}
	return nil
}

func CheckLogin(db *sql.DB, username, password string) (bool, error) {
	row := db.QueryRow(`SELECT pass_hash FROM users WHERE user_name=?`, username)
	var phash []byte
	err := row.Scan(&phash)
	if err != nil {
		return false, fmt.Errorf("User does not exist : %w", err)
	}
	// error is password fail
	return bcrypt.CompareHashAndPassword(phash, []byte(password)) == nil, nil
}
