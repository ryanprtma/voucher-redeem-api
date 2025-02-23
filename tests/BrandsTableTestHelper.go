package tests

import (
	"github.com/jmoiron/sqlx"
)

type BrandsTableTestHelper struct {
	db *sqlx.DB
}

func NewBrandsTableTestHelper(db *sqlx.DB) *BrandsTableTestHelper {
	return &BrandsTableTestHelper{db: db}
}

func (b *BrandsTableTestHelper) AddBrand(id, name string, createdAt, updatedAt string) error {
	query := `INSERT INTO brands (id, name, created_at, updated_at) VALUES (:id, :name, :created_at, :updated_at)`
	_, err := b.db.NamedExec(query, map[string]interface{}{
		"id":         id,
		"name":       name,
		"created_at": createdAt,
		"updated_at": updatedAt,
	})
	return err
}

func (b *BrandsTableTestHelper) FindBrandById(id string) ([]map[string]interface{}, error) {
	query := `SELECT * FROM brands WHERE id = :id`
	rows, err := b.db.NamedQuery(query, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	defer func(rows *sqlx.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var results []map[string]interface{}
	for rows.Next() {
		row := make(map[string]interface{})
		if err := rows.MapScan(row); err != nil {
			return nil, err
		}
		results = append(results, row)
	}

	return results, nil
}

func (b *BrandsTableTestHelper) CleanTable() error {
	_, err := b.db.Exec(`TRUNCATE TABLE brands`)
	return err
}
