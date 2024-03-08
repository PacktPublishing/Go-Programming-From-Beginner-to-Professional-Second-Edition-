package main

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type Record struct {
	ID    int
	Name  string
	Value string
}

type Database struct {
	conn *sql.DB
}

func NewDatabase(conn *sql.DB) *Database {
	return &Database{conn: conn}
}

func (d *Database) InsertRecord(ctx context.Context, record Record) error {
	_, err := d.conn.ExecContext(ctx, "INSERT INTO records (id, name, value) VALUES ($1, $2, $3)", record.ID, record.Name, record.Value)
	return err
}

func (d *Database) GetRecordByID(ctx context.Context, id int) (Record, error) {
	var record Record
	row := d.conn.QueryRowContext(ctx, "SELECT id, name, value FROM records WHERE id = $1", id)
	err := row.Scan(&record.ID, &record.Name, &record.Value)
	return record, err
}

func TestDatabaseIntegration(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	testRecord := Record{
		ID:    1,
		Name:  "TestRecord",
		Value: "TestValue",
	}

	mock.ExpectExec("INSERT INTO records").WithArgs(testRecord.ID, testRecord.Name, testRecord.Value).WillReturnResult(sqlmock.NewResult(1, 1))
	rows := sqlmock.NewRows([]string{"id", "name", "value"}).AddRow(testRecord.ID, testRecord.Name, testRecord.Value)
	mock.ExpectQuery("SELECT id, name, value FROM records").WillReturnRows(rows)

	dbInstance := NewDatabase(db)
	err = dbInstance.InsertRecord(context.Background(), testRecord)
	assert.NoError(t, err, "Error inserting record into the database")

	retrievedRecord, err := dbInstance.GetRecordByID(context.Background(), 1)
	assert.NoError(t, err, "Error retrieving record from the database")
	assert.Equal(t, testRecord, retrievedRecord, "Retrieved record does not match the inserted record")

	// Ensure all expectations were met
	assert.NoError(t, mock.ExpectationsWereMet())
}
