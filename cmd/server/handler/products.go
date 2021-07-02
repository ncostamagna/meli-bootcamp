package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ncostamagna/meli-bootcamp/internal/products"
	"github.com/ncostamagna/meli-bootcamp/pkg/web"
)

type request struct {
	Name string `json:"nombre"`
	Type string `json:"tipo"`
	Count int   `json:"cantidad"`
	Price float64 `json:"precio"`
}

type Product struct {
	service products.Service
}

func NewProduct(p products.Service) *Product {
	return &Product{
		service: p,
	}
}


func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")

		if token != "123456" {
			ctx.JSON(401, web.NewResponse(401, nil, "Token inválido"))
			return
		}

		p, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(400, web.NewResponse(400,nil, err.Error()))
			return
		}

		if len(p) == 0 {
			ctx.JSON(404, web.NewResponse(404,nil, "No hay productos almacenados"))
			return
		}

		ctx.JSON(200, web.NewResponse(200,p, ""))
	}
}

func (c *Product) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")

		if token != "123456" {
			ctx.JSON(401,  web.NewResponse(401, nil, "Token inválido"))
			return
		}

		var req request

		if err := ctx.Bind(&req); err != nil {
			ctx.JSON(400,  web.NewResponse(400, nil, err.Error()))
			return
		}

		if req.Name == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El nombre del producto es requerido"))
			return
		}

		if req.Type == "" {
			ctx.JSON(400, web.NewResponse(400, nil, "El tipo del producto es requerido"))
			return
		}

		if req.Count == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "La cantidad es requerida"))
			return
		}

		if req.Price == 0 {
			ctx.JSON(400, web.NewResponse(400, nil, "El precio es requerido"))
			return
		}

		p, err := c.service.Store(req.Name, req.Type, req.Count, req.Price)
		if err != nil {
			ctx.JSON(400,  web.NewResponse(400, nil, err.Error()))
			return
		}
		ctx.JSON(200, web.NewResponse(200,p, ""))
	}
}