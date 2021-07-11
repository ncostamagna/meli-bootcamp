package products

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/ncostamagna/meli-bootcamp/pkg/store"
	"github.com/stretchr/testify/assert"
)

const errorGetAll = "error for GetAll"

func TestGetAllError(t *testing.T) {
	// Inicializacion (input/output)
	errorEsperado := errors.New(errorGetAll)
	dbMock := store.Mock{
		Data: nil,
		Err:  errorEsperado,
	}
	storeMocked := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeMocked)
	// Ejecución
	_, err := myRepo.GetAll()
	// validacion
	assert.Equal(t, err, errorEsperado)
}

func TestGetAll(t *testing.T) {
	// Inicializacion (input/output)
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
	storeMocked := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}
	myRepo := NewRepository(&storeMocked)
	// Ejecución
	resp, _ := myRepo.GetAll()
	// validacion
	assert.Equal(t, resp, input)
}
