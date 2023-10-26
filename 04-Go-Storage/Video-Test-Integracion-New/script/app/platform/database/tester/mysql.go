package tester

import (
	"database/sql"
)

// NewMySQLTester returns a new instance of MySQLTester.
func NewMySQLTester(db *sql.DB, dbTestName string) *MySQLTester {
	return &MySQLTester{
		db: db,
		dbTestName: dbTestName,
	}
}

// NewMySQLTester returns a new instance of MySQLTester.
type MySQLTester struct {
	// db is the database connection.
	db *sql.DB
	// dbTestName is the name of the testable database.
	dbTestName string
}

// SetUp sets up the tester.
// - connected to the testable database.
func (m *MySQLTester) SetUp(query ...string) (err error) {
	// create database test
	_, err = m.db.Exec("CREATE DATABASE " + m.dbTestName)
	if err != nil {
		return
	}

	// use database test
	_, err = m.db.Exec("USE " + m.dbTestName)
	if err != nil {
		return
	}

	// run queries
	for _, q := range query {
		_, err = m.db.Exec(q)
		if err != nil {
			return
		}
	}
		
	return
}

// TearDown tears down the tester.
// - drop the testable database.
func (m *MySQLTester) TearDown() (err error) {
	// drop database
	_, err = m.db.Exec("DROP DATABASE IF EXISTS " + m.dbTestName)
	if err != nil {
		return
	}

	return
}