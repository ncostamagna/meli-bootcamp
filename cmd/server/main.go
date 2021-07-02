package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ncostamagna/meli-bootcamp/cmd/server/handler"
	"github.com/ncostamagna/meli-bootcamp/internal/products"
	"github.com/ncostamagna/meli-bootcamp/pkg/store"
)

func main() {
	db := store.New(store.FileType, "./products.json")
	repo := products.NewRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)

	r := gin.Default()
	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}