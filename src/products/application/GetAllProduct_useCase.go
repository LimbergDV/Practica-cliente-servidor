package application

import "actividad-principal/src/products/domain"



type GetAllEmployees struct {
	db domain.Iproduct
}

func NewGetAllProducts(db domain.Iproduct) *GetAllEmployees {
	return &GetAllEmployees{db: db}
}

func (le *GetAllEmployees) Run () []domain.Product {
	return le.db.GetAll()
}