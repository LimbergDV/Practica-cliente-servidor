package routes

import (
	"actividad-principal/src/products/infrastructure/controllers"

	"github.com/gin-gonic/gin"
)

func Routes (r *gin.Engine){

	productsRoutes := r.Group("/products") 
	{

	productsRoutes.POST("/", controllers.NewCreateProductController().Run)
	productsRoutes.GET("/", controllers.NewGetAllEmployeesController().Run)
	productsRoutes.PUT("/:id", controllers.NewUpdateProductController().Run)
	// productsRoutes.DELETE("/:id", controllers.NewDeleteProductByIdController().Run)

	}
}