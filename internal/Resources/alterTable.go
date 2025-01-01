package Resources

import (
	"fmt"

	"github.com/eediallo/real_time_forum/internal/db"
)

func AlterSessionTable() error {

	var err error
	err = db.InitDB()
	if err != nil {
		return fmt.Errorf("ERROR: fail to initialize db: %w", err)
	}

	tx, errB := db.DB.Begin()
	if errB != nil {
		return fmt.Errorf("ERROR: Fail to begin db %w", err)
	}

	// Copy data from old table to new table
	copyDataStmt := `
	 INSERT INTO Session_New (SessionID, UserID, CreatedAt)
	 SELECT SessionID, UserID, CreatedAt FROM Session;
	 `
	_, err = tx.Exec(copyDataStmt)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("ERROR: Fail to execute copiedData %w", err)
	}

	// Drop the old table
	dropOldTableStmt := `
	 DROP TABLE IF EXISTS Session;
	 `
	_, err = db.DB.Exec(dropOldTableStmt)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("ERROR: Fail to drop table Session %w", err)
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("ERROR: Fail to to commit transaction %w", err)
	}

	fmt.Println("Username column removed successfully.")

	return nil

}
