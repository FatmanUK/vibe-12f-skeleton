package db

import (
	"database/sql"
	"errors"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
)

const failureCount = 2

var mock sqlmock.Sqlmock
var currentAttempt int

func setExpectations(mock *sqlmock.Sqlmock, testConnStr string) {
	errFail := errors.New("failed")
	currentAttempt++
	switch testConnStr {
		case "success": {
			(*mock).ExpectPing().WillReturnError(nil)
		}
		case "failure": {
			(*mock).ExpectPing().WillReturnError(errFail)
		}
		case "delayed": {
			o := (*mock).ExpectPing()
			if currentAttempt < (failureCount + 1) {
				o.WillReturnError(errFail)
			} else {
				o.WillReturnError(nil)
			}
		}
	}
}

func testInitDb(connStr string) (*sql.DB, error) {
	var db *sql.DB
	var err error
	db, mock, err = sqlmock.New(
		sqlmock.MonitorPingsOption(true),
	)
	if err == nil {
		setExpectations(&mock, connStr)
	}
	return db, err
}

func TestConnectDB_Success(t *testing.T) {
	currentAttempt = 0
	db, err := ConnectDB(testInitDb, "success", 5)
	defer db.Close()
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestConnectDb_CompleteFailure(t *testing.T) {
	currentAttempt = 0
	_, err := ConnectDB(testInitDb, "failure", 5)
	if err == nil {
		t.Error("Expected error, got none")
	}
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestConnectDb_DelayedSuccess(t *testing.T) {
	currentAttempt = 0
	db, err := ConnectDB(testInitDb, "delayed", 5)
	defer db.Close()
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}
	err = mock.ExpectationsWereMet()
	if err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
