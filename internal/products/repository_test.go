package products

import (
	"errors"
	"github.com/go-playground/assert/v2"
	"github.com/ncostamagna/meli-bootcamp/pkg/store"
	"testing"
)

const errorGetAll = "error for GetAll"

func TestGetAll(t *testing.T) {

	dbMock := store.Mock{
		Data: nil,
		Err:  errors.New(errorGetAll),
	}
	storeMocked := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	myRepo:= NewRepository(&storeMocked)
	_,err := myRepo.GetAll()
	assert.Equal(t,err,errors.New(errorGetAll))
}