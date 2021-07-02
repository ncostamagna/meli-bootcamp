package products

import (
	"fmt"
	"github.com/ncostamagna/meli-bootcamp/pkg/store"
)

type Product struct {
	ID       int `json:"id"`
	Name   string `json:"nombre"`
	Type     string `json:"tipo"`
	Count int `json:"cantidad"`
	Price   float64 `json:"precio"`
}

type Repository interface{
	GetAll() ([]Product, error)
	Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error)
	LastID() (int, error)
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (r *repository) GetAll() ([]Product, error) {
	var ps []Product
	fmt.Println("entra")
	r.db.Read(&ps)
	return ps, nil
}

func (r *repository) Store(id int, nombre, tipo string, cantidad int, precio float64) (Product, error) {
	var ps []Product
	r.db.Read(&ps)
	p := Product{id, nombre, tipo, cantidad, precio}
	ps = append(ps, p)
	if err := r.db.Write(ps); err != nil {
		return Product{}, err
	}
	return p, nil
}

func (r *repository) LastID() (int, error) {
	var ps []Product
	if err := r.db.Read(&ps); err != nil{
		return 0, err
	}

	if len(ps) == 0 {
		return 0, nil
	}

	return ps[len(ps) - 1].ID, nil
}