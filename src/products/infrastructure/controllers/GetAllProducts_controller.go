package controllers

import (
	"actividad-principal/src/products/application"
	"actividad-principal/src/products/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
)

type GetAllEmployeesController struct {
	app *application.GetAllEmployees
}

func NewGetAllEmployeesController() *GetAllEmployeesController {
	mysql := infrastructure.GetMySQL()
	app := application.NewGetAllProducts(mysql)
	return &GetAllEmployeesController{app: app}
}

func (ctrl *GetAllEmployeesController) Run(c *gin.Context) {
	res := ctrl.app.Run()

	if len(res) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"status": false, "error": "No se encontró ningún empleado"})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"employees": res})
	}
}