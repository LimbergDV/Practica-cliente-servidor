package domain

type Iproduct interface {
	Save(product Product) (uint, error)
	GetAll() ([]Product) 
	// Delete(id int) (uint, error)
	Update(id int, product Product) (uint, error)
}