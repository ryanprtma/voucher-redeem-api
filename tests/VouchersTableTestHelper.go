package tests

import "github.com/jmoiron/sqlx"

type VouchersTableTestHelper struct {
	db *sqlx.DB
}

func NewVouchersTableTestHelper(db *sqlx.DB) *VouchersTableTestHelper {
	return &VouchersTableTestHelper{db: db}
}

func (v *VouchersTableTestHelper) AddVoucher(id, brandID, name string, point, price, stock int, description, createdAt, updatedAt string) error {
	query := `INSERT INTO vouchers (id, brand_id, name, point, price, stock, description, created_at, updated_at) 
			  VALUES (:id, :brand_id, :name, :point, :price, :stock, :description, :created_at, :updated_at)`
	_, err := v.db.NamedExec(query, map[string]interface{}{
		"id":          id,
		"brand_id":    brandID,
		"name":        name,
		"point":       point,
		"price":       price,
		"stock":       stock,
		"description": description,
		"created_at":  createdAt,
		"updated_at":  updatedAt,
	})
	return err
}

func (v *VouchersTableTestHelper) FindVoucherById(id string) ([]map[string]interface{}, error) {
	query := `SELECT * FROM vouchers WHERE id = :id`
	rows, err := v.db.NamedQuery(query, map[string]interface{}{"id": id})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

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

func (v *VouchersTableTestHelper) CleanTable() error {
	_, err := v.db.Exec(`TRUNCATE TABLE vouchers`)
	return err
}
