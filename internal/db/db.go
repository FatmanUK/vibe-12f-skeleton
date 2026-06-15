package db

import (
	"database/sql"
	"fmt"
	"time"
	_ "github.com/lib/pq"
)

type initFnType = func(string) (*sql.DB, error)

func InitDb(connStr string) (*sql.DB, error) {
	return sql.Open("postgres", connStr)
}

func ConnectDB(
		initFn initFnType,
		connStr string,
		maxAttempts int,
		) (*sql.DB, error) {
	var db *sql.DB
	var err error

	msgSuccess := "Successfully connected to the database"
	msgFmtErrConn := "Connection attempt %d failed: %v\n"
	msgFmtErrPing := "Ping attempt %d failed: %v\n"
	msgFmtErrFail := "Could not connect after %d attempts: %w"
	for attempt := 1; attempt <= maxAttempts; attempt++ {
		db, err = initFn(connStr)
		delayTime := time.Duration(attempt*attempt) * time.Second
		if err != nil {
			fmt.Printf(msgFmtErrConn, attempt, err)
			time.Sleep(delayTime)
			continue
		}

		// Check the connection
		err = db.Ping()
		if err == nil {
			fmt.Println(msgSuccess)
			return db, nil
		} else {
			fmt.Printf(msgFmtErrPing, attempt, err)
			db.Close()
			time.Sleep(delayTime)
		}
	}
	return nil, fmt.Errorf(msgFmtErrFail, maxAttempts, err)
}
