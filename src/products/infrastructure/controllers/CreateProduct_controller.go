package controllers

import (
	"actividad-principal/src/products/application"
	"actividad-principal/src/products/domain"
	"actividad-principal/src/products/infrastructure"
	"actividad-principal/src/products/infrastructure/routes/validators"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateProductController struct {
	app *application.CreateProduct
}

//Este constructor permite inyectar la dependencia del caso del uso (application.CreateProduct) al controlador.
func NewCreateProductController()  *CreateProductController {
	mysql := infrastructure.NewMySQL()
	app := application.NewCreateProduct(mysql)
	return &CreateProductController{app: app}
}

func (ctrl *CreateProductController) Run(c *gin.Context){
	var products domain.Product

	if err := c.ShouldBindJSON(&products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status":false, "error": "Datos invalidos " + err.Error()})
		return 
	}

	if err := validators.CheckProduct(products); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {"status": false, "error": "Datos invalidos " + err.Error()})
	} 

	rowsAffected, err := ctrl.app.Run(products)

	if err != nil {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	return

	if rowsAffected == 0 {
		fmt.Print("hola")
	}
	} else {
		c.JSON(http.StatusCreated, gin.H {"mensaje": "Producto creado"})
		c.JSON(http.StatusOK, products)
	}
}


