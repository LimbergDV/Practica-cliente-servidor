package application

import "actividad-principal/src/products/domain"





type CreateProduct struct{
	db domain.Iproduct
}

func NewCreateProduct (db domain.Iproduct) *CreateProduct{
	return &CreateProduct{db:db}
}

func (cp *CreateProduct) Run (product domain.Product)  (uint, error){ 
	return cp.db.Save(product)
}