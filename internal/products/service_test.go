package products

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/ncostamagna/meli-bootcamp/pkg/store"
	"github.com/stretchr/testify/assert"
)

func TestServiceGetAll(t *testing.T) {
	input := []Product{
		{
			ID:    1,
			Name:  "CellPhone",
			Type:  "Tech",
			Count: 3,
			Price: 250,
		}, {
			ID:    2,
			Name:  "Notebook",
			Type:  "Tech",
			Count: 10,
			Price: 1750.5,
		},
	}
	dataJson, _ := json.Marshal(input)
	dbMock := store.Mock{
		Data: dataJson,
		Err:  nil,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)
	result, err := myService.GetAll()
	assert.Equal(t, input, result)
	assert.Nil(t, err)
}

func TestServiceGetAllError(t *testing.T) {
	// Inicializacion (input/output)
	errorEsperado := errors.New("error for GetAll")
	dbMock := store.Mock{
		Data: nil,
		Err:  errorEsperado,
	}

	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)
	result, err := myService.GetAll()
	assert.Equal(t, errorEsperado, err)
	assert.Nil(t, result)
}

func TestStore(t *testing.T) {

	testProduct := Product{
		Name:  "CellPhone",
		Type:  "Tech",
		Count: 3,
		Price: 52.0,
	}
	dbMock := store.Mock{
		Data: nil,
		Err:  nil,
	}

	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)
	result, _ := myService.Store(testProduct.Name, testProduct.Type, testProduct.Count, testProduct.Price)
	assert.Equal(t, testProduct.Name, result.Name)
	assert.Equal(t, testProduct.Type, result.Type)
	assert.Equal(t, testProduct.Price, result.Price)
	assert.Equal(t, 1, result.ID)
}

func TestStoreError(t *testing.T) {

	testProduct := Product{
		Name:  "CellPhone",
		Type:  "Tech",
		Count: 3,
		Price: 52.0,
	}
	errorEsperado := errors.New("error for Storage")
	dbMock := store.Mock{
		Data: nil,
		Err:  errorEsperado,
	}

	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeStub)
	myService := NewService(myRepo)
	result, err := myService.Store(testProduct.Name, testProduct.Type, testProduct.Count, testProduct.Price)
	assert.Equal(t, errorEsperado, err)
	assert.Equal(t, Product{}, result)
}
