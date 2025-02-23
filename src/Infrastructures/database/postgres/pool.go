package postgres

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
	config "voucher-redeem-api/src/Commons"
)

type Database struct {
	DB *sqlx.DB
}

func NewDatabase() (*Database, error) {
	dbURL := getDatabaseURL()
	db, err := sqlx.Open("postgres", dbURL)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	err = db.Ping()
	if err != nil {
		err := db.Close()
		if err != nil {
			return nil, err
		}
		return nil, err
	}

	return &Database{DB: db}, nil
}

func getDatabaseURL() string {
	cnfg := config.LoadConfig()
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cnfg["user"], cnfg["password"], cnfg["host"], cnfg["port"], cnfg["database"])
}

func (d *Database) GetDB() *sqlx.DB {
	return d.DB
}
