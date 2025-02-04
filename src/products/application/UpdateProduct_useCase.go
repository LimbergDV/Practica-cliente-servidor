package application

import "actividad-principal/src/products/domain"



type UpdateProduct struct {
	db domain.Iproduct
}

func NewUpdateEmployee( db domain.Iproduct) *UpdateProduct {
	return &UpdateProduct{db: db}
}

func (ue *UpdateProduct) Run (id int, employee domain.Product) (uint, error) {
	return ue.db.Update(id, employee)
}