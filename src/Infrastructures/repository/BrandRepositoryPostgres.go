package repository

import (
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
	"voucher-redeem-api/src/Commons/exceptions"
	"voucher-redeem-api/src/Domains/brands/entities"
)

type postgresDB struct {
	db *sqlx.DB
}

func NewPostgresDB(db *sqlx.DB) postgresDB {
	return postgresDB{db: db}
}

func (p postgresDB) verifyAvailableBrandName(name string) error {
	var count int

	query := "SELECT COUNT(*) FROM brands WHERE name = $1"

	err := p.db.Get(&count, query, name)

	if err != nil {
		return err
	}

	if count > 0 {
		return exceptions.NewInvariantError("nama brand tidak tersedia")
	}

	return nil
}

func (p postgresDB) InsertBrand(createBrand entities.CreateBrand) (*entities.Brand, error) {
	err := p.verifyAvailableBrandName(createBrand.Name)
	if err != nil {
		return nil, err
	}

	id := uuid.NewString()
	createdAt := time.Now()

	query := "INSERT INTO brands (id, name, created_at, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, name, created_at, updated_at"

	row := p.db.QueryRowx(query, id, createBrand.Name, createdAt, createdAt)

	var returnedID, returnedName string
	var returnedCreatedAt time.Time
	var returnedUpdatedAt time.Time

	err = row.Scan(&returnedID, &returnedName, &returnedCreatedAt, &returnedUpdatedAt)
	if err != nil {
		return nil, err
	}

	domainPayload := map[string]string{
		"id":         returnedID,
		"name":       returnedName,
		"created_at": returnedCreatedAt.Format(time.RFC3339),
		"updated_at": returnedUpdatedAt.Format(time.RFC3339),
	}

	createdBand, err := entities.NewBrand(domainPayload)

	if err != nil {
		return nil, err
	}

	return createdBand, nil
}
