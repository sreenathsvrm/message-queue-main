package service

// import (
// 	"errors"
// 	"message-queue/internals/models"
// 	"message-queue/repository/mock"
// 	"testing"

// 	"github.com/go-playground/assert/v2"
// 	"github.com/golang/mock/gomock"
// )

// func TestAddProductService(t *testing.T) {
// 	tests := []struct {
// 		name          string
// 		input         models.Products
// 		expectedErr   error
// 		buildRepoStub func(mockRepo *mock.MockProductRepository)
// 		buildProducer func(mockProducer *mockproducer.MockProducer)
// 	}{
// 		{
// 			name: "successful product addition",
// 			input: models.Products{
// 				ProductName:            "Test Product",
// 				ProductDescription:     "Description",
// 				ProductImage:           []string{"image1.jpg", "image2.jpg"},
// 				ProductPrice:           19.99,
// 				CompressedProductImage: []string{"compressed1.jpg", "compressed2.jpg"},
// 			},
// 			buildRepoStub: func(mockRepo *mock.MockProductRepository) {
// 				mockRepo.EXPECT().AddProduct(gomock.Any()).Return(1, nil)
// 			},
// 			buildProducer: func(mockProducer *mockproducer.MockProducer) {
// 				mockProducer.EXPECT().Produce(gomock.Any()).Return(nil)
// 			},
// 			expectedErr: nil,
// 		},
// 		{
// 			name: "repository failure",
// 			input: models.Products{
// 				ProductName:            "Test Product",
// 				ProductDescription:     "Description",
// 				ProductImage:           []string{"image1.jpg", "image2.jpg"},
// 				ProductPrice:           19.99,
// 				CompressedProductImage: []string{"compressed1.jpg", "compressed2.jpg"},
// 			},
// 			buildRepoStub: func(mockRepo *mock.MockProductRepository) {
// 				mockRepo.EXPECT().AddProduct(gomock.Any()).Return(0, errors.New("insertion failed"))
// 			},
// 			buildProducer: func(mockProducer *mockproducer.MockProducer) {
// 				// No expectations for the producer as the repo is expected to fail.
// 			},
// 			expectedErr: errors.New("insertion failed"),
// 		},
// 		{
// 			name: "producer failure",
// 			input: models.Products{
// 				ProductName:            "Test Product",
// 				ProductDescription:     "Description",
// 				ProductImage:           []string{"image1.jpg", "image2.jpg"},
// 				ProductPrice:           19.99,
// 				CompressedProductImage: []string{"compressed1.jpg", "compressed2.jpg"},
// 			},
// 			buildRepoStub: func(mockRepo *mock.MockProductRepository) {
// 				mockRepo.EXPECT().AddProduct(gomock.Any()).Return(1, nil)
// 			},
// 			buildProducer: func(mockProducer *mockproducer.MockProducer) {
// 				mockProducer.EXPECT().Produce(gomock.Any()).Return(errors.New("producer error"))
// 			},
// 			expectedErr: errors.New("producer error"),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			ctrl := gomock.NewController(t)
// 			defer ctrl.Finish()

// 			mockRepo := mock.NewMockProductRepository(ctrl)
// 			mockProducer := mockproducer.NewMockProducer(ctrl)

// 			// Set up the repository and producer expectations
// 			tt.buildRepoStub(mockRepo)
// 			tt.buildProducer(mockProducer)

// 			// Create the product service with the mocked repository and producer
// 			productSvc := NewProductService(mockRepo, mockProducer)

// 			// Call the service function
// 			actualErr := productSvc.AddProduct(tt.input)

// 			// Validate the error
// 			if tt.expectedErr == nil {
// 				assert.NoError(t, actualErr)
// 			} else {
// 				assert.Equal(t, tt.expectedErr, actualErr)
// 			}
// 		})
// 	}
// }
