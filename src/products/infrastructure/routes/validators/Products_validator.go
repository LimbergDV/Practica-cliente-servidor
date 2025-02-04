package validators

import (
	"actividad-principal/src/products/domain"
	"errors"
)

func CheckProduct(product domain.Product) error{
	if product.Name == "" {
		return errors.New("No puedes poner un nombre vacio")
	}

	if product.Quantity < 0{
		return errors.New("El precio debe ser mayor a 0")
	}

	if product.Id < 0{
		return errors.New("El id debe de ser mayor a 0")
	}

	if product.BarCode == "" {
		return errors.New("No debes de poner el código de barras vacío")
	}
	
	return nil
}