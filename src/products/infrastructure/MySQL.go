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

func (mysql *MySQL) GetAll() []domain.Product {
	query := "SELECT * FROM products"
	var employees []domain.Product
	rows := mysql.conn.FetchRows(query)

	if rows == nil {
		fmt.Println("No se obtuvieron los datos")
		return employees
	}
	defer rows.Close()

	for rows.Next() {
		var e domain.Product
		if err := rows.Scan(&e.Id, &e.Name, &e.Quantity, &e.BarCode); err != nil {
			fmt.Println("Error al escanear la fila: %w", err)
		} else {
			employees = append(employees, e)
		}
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error iterando sobre las filas: %w", err)
	}

	fmt.Println("Lista de empleados")
	return employees
}

func (mysql *MySQL) Update(id int, product domain.Product) (uint, error) {
	query := "UPDATE products SET name = ?, quantity = ?, barcode = ?,  WHERE id = ?"
	res, err := mysql.conn.ExecutePreparedQuery(query, product.Name, product.Quantity, product.BarCode, id)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta: %v", err)
		return 0, err
	}
	rows, _ := res.RowsAffected()
	return uint(rows), nil
}