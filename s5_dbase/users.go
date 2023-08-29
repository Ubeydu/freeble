package s5_dbase

import "database/sql"

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
