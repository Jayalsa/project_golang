package controllers

import (
	"fmt"
	"jayalsa/project_golang/entities"
	"jayalsa/project_golang/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	ProductService interfaces.IProduct
}

func InitProductController(productSvc interfaces.IProduct) *ProductController {
	return &ProductController{ProductService: productSvc}
}

func (p ProductController) InsertProduct(c *gin.Context) {
	fmt.Println("Invoked controller")
	var product entities.Product
	err := c.BindJSON(&product)
	if err != nil {
		return
	}
	result, err := p.ProductService.Insert(&product)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}
func (p ProductController) GetProducts(c *gin.Context) {
	result, err := p.ProductService.GetProducts()
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}
func (p ProductController) GetProductByID(c *gin.Context) {
	productID := c.Param("_id")
	product, err := p.ProductService.GetProductByID(productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch product"})
		return
	}

	if product == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Product not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, product)
}
