package infrastructure

import (
	"actividad-principal/src/core"
	"actividad-principal/src/products/domain"
	"fmt"
	"log"
)


type MySQL struct{
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}

	return &MySQL{conn: conn}
}

func (mysql *MySQL) Save(product domain.Product) (uint, error) {
	query := "INSERT INTO products (name, quantity, barcode) VALUES (?, ?, ?)"
	res, err := mysql.conn.ExecutePreparedQuery(query, product.Name, product.Quantity, product.BarCode)

	if err != nil {
		fmt.Println("Error al preparar la consulta: %v", err)
		return 0, err
	}

	id, _ := res.LastInsertId()

	fmt.Println("Producto creado")
	return uint(id), nil
}