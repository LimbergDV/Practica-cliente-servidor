package controllers

import (
	"actividad-principal/src/products/application"
	"actividad-principal/src/products/domain"
	"actividad-principal/src/products/infrastructure"
	"fmt"

	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UpdateProductController struct {
	app *application.UpdateProduct
}

func NewUpdateProductController() *UpdateProductController {
	mysql := infrastructure.GetMySQL()
	app := application.NewUpdateEmployee(mysql)
	return &UpdateProductController{app: app}
}

func (ctrl *UpdateProductController) Run(c *gin.Context) {
	id := c.Param("id")
	var employee domain.Product

	idEmployee, _ := strconv.ParseUint(id, 10, 64)

	if err := c.ShouldBindJSON(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	rowsAffected, _ := ctrl.app.Run(int(idEmployee), employee)

	if rowsAffected == 0 {
		fmt.Println("No se pudo actualizar el empleado")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el empleado"})
		return
	}

	// Enviar una respuesta exitosa
	c.JSON(http.StatusOK, gin.H{
		"message": "Employee updated successfully",
	})
}
