package service

import (
	"errors"
	"testing"

	"github.com/l4rma/go-rest-api/server/db/entity"
	"github.com/l4rma/go-rest-api/server/db/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	bookRepository repository.BookRepository = repository.NewPostgresRepository()
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) Open() error {
	return errors.New("Error")
}

func (mock *MockRepository) Close() error {
	return errors.New("Error")
}

func (mock *MockRepository) Save(book *entity.Book) (int64, error) {
	args := mock.Called()
	res := args.Get(0)
	return res.(int64), args.Error(1)
}

func (mock *MockRepository) FindAll() ([]*entity.Book, error) {
	args := mock.Called()
	res := args.Get(0)
	return res.([]*entity.Book), args.Error(1)
}

func (mock *MockRepository) InsertDummyData(repo repository.BookRepository) {
	return
}

func TestValidateNilBook(t *testing.T) {
	testService := NewBookService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "The book is nil", err.Error())
}

func TestValidateEmptyTitle(t *testing.T) {
	book := &entity.Book{ID: 1, Title: "", Author: "Lars", Year: 2022}

	testService := NewBookService(nil)

	err := testService.Validate(book)

	assert.NotNil(t, err)
	assert.Equal(t, "The book has no title", err.Error())
}

func TestCreate(t *testing.T) {
	mockRepo := new(MockRepository)

	var id int64 = 1
	var year int16 = 2022
	book := entity.Book{ID: id, Title: "BokTittel", Author: "Lars", Year: year}

	// Mock repo method return value
	mockRepo.On("Save").Return(id, nil)

	testService := NewBookService(mockRepo)

	result, _ := testService.Create(&book)

	mockRepo.AssertExpectations(t)

	assert.Equal(t, id, result)
}

func TestFindAll(t *testing.T) {
	mockRepo := new(MockRepository)

	var id int64 = 1
	var year int16 = 2022
	book := entity.Book{ID: id, Title: "BokTittel", Author: "Lars", Year: year}

	// Mock repo method return value
	mockRepo.On("FindAll").Return([]*entity.Book{&book}, nil)

	testService := NewBookService(mockRepo)

	result, _ := testService.FindAll()

	// Mock Assertion: Behavioral
	mockRepo.AssertExpectations(t)

	// Mock Assertion
	assert.Equal(t, id, result[0].ID)
	assert.Equal(t, "BokTittel", result[0].Title)
	assert.Equal(t, "Lars", result[0].Author)
	assert.Equal(t, year, result[0].Year)
}
