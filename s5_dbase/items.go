package s5_dbase

import (
	"database/sql"
	"fmt"
	"io"
)

func CreateItemsTable(dbase *sql.DB) error {
	userStmt, err := dbase.Prepare(`
		CREATE TABLE items(
			"item_id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
			"giver_id" integer NOT NULL,
			"receiver_id" text,
			"name" text,
			"description" text,
			"image" blob,
			FOREIGN KEY (giver_id) REFERENCES users (user_name),
			FOREIGN KEY (receiver_id) REFERENCES users (user_name)
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

func AddItem(dbase *sql.DB, giver_id, item_name, description string, image []byte) error {
	if len(image) > 100_000 {
		return fmt.Errorf("image too big for %s", item_name)
	}
	stmt, err := dbase.Prepare(`INSERT INTO items (giver_id, name, description, image) VALUES (?, ?, ?, ?);`)
	if err != nil {
		return fmt.Errorf("could not Prepare Item creator %s: %w", item_name, err)
	}
	_, err = stmt.Exec(giver_id, item_name, description, image)
	if err != nil {
		return fmt.Errorf("could not add %s to Database : %w", item_name, err)
	}
	return nil
}

func GetImage(db *sql.DB, item_id int, w io.Writer) error {
	row := db.QueryRow(`SELECT image FROM items WHERE item_id=?;`, item_id)
	var pic *sql.RawBytes
	if err := row.Scan(&pic); err != nil {
		return fmt.Errorf("could not get pic %d : %w", item_id, err)
	}
	w.Write(*pic)
	return nil
}
