package repository

import (
	"errors"
	"message-queue/internals/models"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestAddProduct(t *testing.T) {
	tests := []struct {
		name        string
		input       models.Products
		expectedID  int
		buildStub   func(mock sqlmock.Sqlmock)
		expectedErr error
	}{
		{
			name: "successful product insertion",
			input: models.Products{
				ProductName:            "Test Product",
				ProductDescription:     "Description",
				ProductImage:           []string{"image1.jpg", "image2.jpg"},
				ProductPrice:           19.99,
				CompressedProductImage: []string{"compressed1.jpg", "compressed2.jpg"},
			},
			expectedID: 1,
			buildStub: func(mock sqlmock.Sqlmock) {
				// Use sqlmock.Arg matcher to check for the expected arguments.
				mock.ExpectQuery("^INSERT INTO products").WithArgs(
					"Test Product", "Description",
					pq.Array([]string{"image1.jpg", "image2.jpg"}), 19.99,
					pq.Array([]string{"compressed1.jpg", "compressed2.jpg"}),
				).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
			},
			expectedErr: nil,
		},

		{
			name: "failed product insertion",
			input: models.Products{
				ProductName:            "Test Product",
				ProductDescription:     "Description",
				ProductImage:           []string{"image1.jpg", "image2.jpg"},
				ProductPrice:           19.99,
				CompressedProductImage: []string{"compressed1.jpg", "compressed2.jpg"},
			},
			expectedID: 0,
			buildStub: func(mock sqlmock.Sqlmock) {
				mock.ExpectQuery("^INSERT INTO products (.+)$").
					WithArgs("Test Product", "Description", "{image1.jpg,image2.jpg}", 19.99, "{compressed1.jpg,compressed2.jpg}").
					WillReturnError(errors.New("insertion failed"))
			},
			expectedErr: errors.New("insertion failed"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// New() method from sqlmock package creates a sqlmock database connection and a mock to manage expectations.
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}
			defer db.Close()

			// Initialize the db instance with the mock db connection.
			gormDB, err := gorm.Open(postgres.New(postgres.Config{Conn: db}), &gorm.Config{})
			if err != nil {
				t.Fatalf("an error '%s' was not expected when initializing a mock db session", err)
			}

			// Create a product repository mock by passing a pointer to gorm.DB.
			productRepository := NewProductRepository(gormDB)

			// Before we actually execute our function, we need to expect the required DB actions.
			tt.buildStub(mock)

			// Call the actual method.
			actualID, actualErr := productRepository.AddProduct(tt.input)

			// Validate err is nil if we are not expecting to receive an error.
			if tt.expectedErr == nil {
				assert.NoError(t, actualErr)
			} else {
				// Validate whether expected and actual errors are the same.
				assert.Equal(t, tt.expectedErr, actualErr)
			}

			// Check the returned product ID.
			assert.Equal(t, tt.expectedID, actualID)

			// Check that all expectations were met.
			err = mock.ExpectationsWereMet()
			if err != nil {
				t.Errorf("Unfulfilled expectations: %s", err)
			}
		})
	}
}
