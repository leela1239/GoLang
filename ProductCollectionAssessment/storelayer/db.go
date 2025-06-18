package storelayer

import (
	"database/sql"
	"product-details/models"
)

type Database struct {
	database *sql.DB
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{
		database: db,
	}
}

func (db *Database) GetProductsById(pid string) (*models.Product, error) {
	var product *models.Product
	query := `SELECT id,name FROM Product where id := $1`
	err := db.database.QueryRow(query, pid).Scan(&product.ProductId, &product.ProductName)
	if err != nil {
		return nil, err
	}
	variants, err := db.getVariantsById(product.ProductId)
	if err != nil {
		return nil, err
	}
	product.Variants = variants

	return nil, nil
}

func (db *Database) GetProductsByName(pname string) (*models.Product, error) {
	var product *models.Product
	query := `SELECT id,name FROM Product where name := $1`
	err := db.database.QueryRow(query, pname).Scan(&product.ProductId, &product.ProductName)
	if err != nil {
		return nil, err
	}
	varaints, err := db.getVariantsById(product.ProductId)
	if err != nil {
		return nil, err
	}
	product.Variants = varaints
	return nil, nil
}

func (db *Database) getVariantsById(pid string) ([]models.Variant, error) {
	var variants []models.Variant
	query := `SELECT id,pid,size,color,price,stock FROM Variant where pid = $1`
	rows, err := db.database.Query(query, pid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var variant models.Variant
		err := rows.Scan(&variant.Variantid, &variant.ProductId, &variant.Size, &variant.Color, &variant.Price, &variant.Stock)
		if err != nil {
			return nil, err
		}
		variants = append(variants, variant)
	}
	return variants, nil
}

func (db *Database) PostProducts(product *models.Product) error {
	query := `INSERT into Product(id,name) Valiues($1,$2) `
	_, err := db.database.Exec(query, product.ProductId, product.ProductName)
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) PutProducts(pid string, products *models.Product) {

}

func (db *Database) DeleteProducts(pid string) {

}

func (db *Database) DeleteVariant(pid string, vid string) {

}
