package s5_dbase

import "database/sql"

func CreateItemsTable(dbase *sql.DB) error {
	userStmt, err := dbase.Prepare(`
		CREATE TABLE items(
			"item_id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			"giver_id" integer NOT NULL,
			"receiver_id" integer,
			"name" text,
			"description" text,
			"image" blob
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