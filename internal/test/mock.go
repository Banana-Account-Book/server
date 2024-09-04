package test

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMockDB() (*gorm.DB, sqlmock.Sqlmock) {
	// Create sqlmock database connection
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		panic("Failed to create mock database: " + err.Error())
	}

	// Use postgres driver with the sqlmock connection
	dialector := postgres.New(postgres.Config{
		Conn: mockDB,
	})

	// Open gorm DB connection
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic("Failed to open gorm database: " + err.Error())
	}

	return db, mock
}

// CloseMockDB closes the mock database connection
func CloseMockDB(mockDB *gorm.DB) {
	sqlDB, err := mockDB.DB()
	if err != nil {
		panic("Failed to get sql.DB from gorm.DB: " + err.Error())
	}
	sqlDB.Close()
}
