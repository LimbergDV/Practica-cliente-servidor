package main

import (
	products "actividad-principal/src/products/infrastructure"
	productsRoutes "actividad-principal/src/products/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

func main (){
	products.GoMySQL()
	r := gin.Default()
	productsRoutes.Routes(r)
	r.Run(":8081")  
}