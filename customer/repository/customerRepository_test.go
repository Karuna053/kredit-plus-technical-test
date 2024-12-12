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

func TestCreate(t *testing.T) {
	// Initialize variables.
	db, mock, _ := sqlmock.New()                         // Create *sql.DB and mock.
	gormdb, _ := gorm.Open(postgres.New(postgres.Config{ // Create gormDB
		Conn: db,
	}), &gorm.Config{})
	customerRepository := NewCustomerRepository(gormdb)

	// Mocks.
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO (.+) VALUES (.+)`).WithArgs(
		"",
		"",
		"",
		"",
		AnyTime{},
		float64(0),
		"",
		"",
		AnyTime{},
		AnyTime{},
	).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))
	mock.ExpectCommit()

	// Tries to run the Repository function.
	newCustomer := &domain.Customer{}
	customer, err := customerRepository.Create(context.TODO(), newCustomer)

	// Assert.
	assert.Nil(t, err)
	assert.NotNil(t, customer)
}

func TestUpdate(t *testing.T) {
	// Initialize variables.
	db, mock, _ := sqlmock.New()
	gormdb, _ := gorm.Open(postgres.New(postgres.Config{ // Create gormDB
		Conn: db,
	}), &gorm.Config{})
	customerRepository := NewCustomerRepository(gormdb)

	// Mocks.
	mock.ExpectBegin()
	mock.ExpectExec(`UPDATE (.+) SET .+`).WithArgs(
		"",
		"",
		"",
		"",
		AnyTime{},
		float64(0),
		"",
		"",
		AnyTime{},
		AnyTime{},
		1,
	).WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()

	// Tries to run the Repository function.
	customer, err := customerRepository.Update(context.TODO(), &domain.Customer{
		ID: 1,
	})

	// Asserts
	assert.Nil(t, err)
	assert.NotNil(t, customer)
}

func TestFetchByID(t *testing.T) {
	// Initialize variables.
	db, mock, _ := sqlmock.New()
	gormdb, _ := gorm.Open(postgres.New(postgres.Config{ // Create gormDB
		Conn: db,
	}), &gorm.Config{})
	customerRepository := NewCustomerRepository(gormdb)

	// Mocks.
	query := `SELECT(.*)`
	mock.ExpectQuery(query).WithArgs(
		uint(1),
		1,
	).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	// Tries to run the Repository function.
	customer, err := customerRepository.FetchByID(context.TODO(), uint(1))

	// Asserts
	assert.Nil(t, err)
	assert.NotNil(t, customer)
}

func TestFetchAll(t *testing.T) {
	// Initialize variables.
	db, mock, _ := sqlmock.New()
	gormdb, _ := gorm.Open(postgres.New(postgres.Config{ // Create gormDB
		Conn: db,
	}), &gorm.Config{})
	customerRepository := NewCustomerRepository(gormdb)

	// Mocks.
	query := `SELECT(.*)`
	mock.ExpectQuery(query).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("1"))

	// Tries to run the Repository function.
	customer, err := customerRepository.FetchAll(context.TODO())

	// Asserts
	assert.Nil(t, err)
	assert.NotNil(t, customer)

}

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
