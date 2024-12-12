package repository

import (
	"context"
	"database/sql/driver"
	"kredit-plus/domain"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateSuccess(t *testing.T) {
	// Initialize variables.
	db, mock, _ := sqlmock.New()                         // Create *sql.DB and mock.
	gormdb, _ := gorm.Open(postgres.New(postgres.Config{ // Create gormDB
		Conn: db,
	}), &gorm.Config{})
	customerTenorRepository := NewCustomerTenorRepository(gormdb)

	// Mocks.
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO (.+) VALUES (.+)`).WithArgs(
		0,
		0,
		float64(0),
		AnyTime{},
		AnyTime{},
	).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
	mock.ExpectCommit()

	// Tries to run the Repository function.
	newCustomerTenor := &domain.CustomerTenor{}
	customerTenor, err := customerTenorRepository.Create(context.TODO(), newCustomerTenor)

	// Assert.
	assert.Nil(t, err)
	assert.NotNil(t, customerTenor)
}

func TestUpdateSuccess(t *testing.T) {
	// Initialize variables.
	db, mock, _ := sqlmock.New()                         // Create *sql.DB and mock.
	gormdb, _ := gorm.Open(postgres.New(postgres.Config{ // Create gormDB
		Conn: db,
	}), &gorm.Config{})
	customerTenorRepository := NewCustomerTenorRepository(gormdb)

	// Mocks.
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE (.+) SET .+`).WithArgs(
		0,
		0,
		float64(0),
		AnyTime{},
		AnyTime{},
		1,
	).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// Tries to run the Repository function.
	customerTenor, err := customerTenorRepository.Update(context.TODO(), &domain.CustomerTenor{
		ID: 1, // Assign ID to the struct to force gorm to use UPDATE and not INSERT
	})

	// Assert.
	assert.Nil(t, err)
	assert.NotNil(t, customerTenor)
}

func TestFetchByIDSuccess(t *testing.T) {
	// Initialize variables.
	db, mock, _ := sqlmock.New()                         // Create *sql.DB and mock.
	gormdb, _ := gorm.Open(postgres.New(postgres.Config{ // Create gormDB
		Conn: db,
	}), &gorm.Config{})
	customerTenorRepository := NewCustomerTenorRepository(gormdb)

	// Mocks.
	query := `SELECT(.*)`
	mock.ExpectQuery(query).WithArgs(
		uint(1),
		1,
	).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	// Tries to run the Repository function.
	customerTenor, err := customerTenorRepository.FetchByID(context.TODO(), uint(1))

	// Assert.
	assert.Nil(t, err)
	assert.NotNil(t, customerTenor)
}

func TestFetchAllSuccess(t *testing.T) {
	// Initialize variables.
	db, mock, _ := sqlmock.New()                         // Create *sql.DB and mock.
	gormdb, _ := gorm.Open(postgres.New(postgres.Config{ // Create gormDB
		Conn: db,
	}), &gorm.Config{})
	customerTenorRepository := NewCustomerTenorRepository(gormdb)

	// Mocks.
	query := `SELECT(.*)`
	mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	// Tries to run the Repository function.
	customerTenor, err := customerTenorRepository.FetchAll(context.TODO())

	// Assert.
	assert.Nil(t, err)
	assert.NotNil(t, customerTenor)
}

// func TestDeleteSuccess(t *testing.T) {
// 	// Initialize variables.
// 	db, mock, _ := sqlmock.New()                         // Create *sql.DB and mock.
// 	gormdb, _ := gorm.Open(postgres.New(postgres.Config{ // Create gormDB
// 		Conn: db,
// 	}), &gorm.Config{})
// 	customerTenorRepository := NewCustomerTenorRepository(gormdb)

// 	// Mocks.
// 	mock.ExpectBegin()
// 	mock.ExpectExec(`DELETE FROM (.+) WHERE id = .+`).WithArgs(
// 		1,
// 	).WillReturnResult(sqlmock.NewResult(0, 1))
// 	// mock.ExpectCommit()

// 	// Tries to run the Repository function.
// 	err := customerTenorRepository.Delete(context.TODO(), &domain.CustomerTenor{})

// 	// Assert.
// 	assert.Nil(t, err)
// }

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
