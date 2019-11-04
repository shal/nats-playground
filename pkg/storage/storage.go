package storage

type Database interface {
	WantedVehicles(string) (Transport, error)
	InsertWantedVehicle(*WantedVehicle) error
	Operations(string) ([]Operation, error)
}

type Storage struct {
	db Database
}

// New returns new storage.
func New(db Database) *Storage {
	return &Storage{
		db: db,
	}
}

// WantedVehicles returns all wanted vehicles with stolen status.
func (s *Storage) WantedVehicles() (Transport, error) {
	return s.db.WantedVehicles("stolen")
}

// InsertWantedVehicle add new vehicle to the storage.
func (s *Storage) InsertWantedVehicle(v *WantedVehicle) error {
	return s.db.InsertWantedVehicle(v)
}

// Operations returns all operations, thats contains specified a number.
func (s *Storage) Operations(number string) ([]Operation, error) {
	return s.db.Operations(number)
}
