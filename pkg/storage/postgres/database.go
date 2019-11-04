package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/opencars/wanted/pkg/storage"
)

type database struct {
	db *sqlx.DB
}

func (db *database) WantedVehicles(state string) (storage.Transport, error) {
	var vehicles []*storage.WantedVehicle
	err := db.db.Select(&vehicles, `SELECT * FROM wanted_vehicles WHERE state = $1`, state)
	if err != nil {
		return nil, fmt.Errorf("failed to select: %w", err)
	}
	return vehicles, nil
}

func (db *database) InsertWantedVehicle(v *storage.WantedVehicle) error {
	_, err := db.db.NamedExec(`INSERT INTO wanted_vehicles (id, ovd, brand, model, kind, color, plates, body_number, chassis_number, engine_number, theft_date, insert_date, state) VALUES (:id, :ovd, :brand, :model, :kind, :color, :plates, :body_number, :chassis_number, :engine_number, :theft_date, :insert_date, :state)`, v)
	if err != nil {
		return fmt.Errorf("failed to insert: %w", err)
	}

	return nil
}

func (db *database) Operations(number string) ([]storage.Operation, error) {
	var operations []storage.Operation
	err := db.db.Select(&operations, `SELECT id,date,brand,model,year,color,kind,body,purpose,number FROM operations WHERE number = $1 ORDER BY id DESC`, number)
	if err != nil {
		return nil, fmt.Errorf("failed to select: %w", err)
	}
	return operations, nil
}

func New(host string, port int, user, password, dbname string) (storage.Database, error) {
	info := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", info)
	if err != nil {
		return nil, fmt.Errorf("connection failed: %w", err)
	}

	return &database{
		db: db,
	}, nil
}
