package storage

// WantedVehicle represents model.
type WantedVehicle struct {
	ID            string  `db:"id"`
	OVD           string  `db:"ovd"`
	Brand         *string `db:"brand"`
	Model         *string `db:"model"`
	Kind          *string `db:"kind"`
	Color         string  `db:"color"`
	Plates        *string `db:"plates"`
	BodyNumber    *string `db:"body_number"`
	ChassisNumber *string `db:"chassis_number"`
	EngineNumber  *string `db:"engine_number"`
	TheftDate     string  `db:"theft_date"`
	State         string  `db:"state"`
	InsertDate    string  `db:"insert_date"`
}

// Operation represents operation on a vehicle.
type Operation struct {
	ID      int64  `db:"id"`
	Date    string `db:"date"`
	Brand   string `db:"brand" `
	Model   string `db:"model"`
	Year    int32  `db:"year"`
	Color   string `db:"color"`
	Kind    string `db:"kind"`
	Body    string `db:"body"`
	Purpose string `db:"purpose"`
	Number  string `db:"number"`
}

type Transport []*WantedVehicle

// Len is the number of elements in the collection.
func (t Transport) Len() int {
	return len(t)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (t Transport) Less(i, j int) bool {
	return t[i].ID <= t[j].ID
}

// Swap swaps the elements with indexes i and j.
func (t Transport) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t Transport) Search(id string) int {
	pivot := len(t) / 2
	return t.search(pivot, id)
}

func (t Transport) search(i int, id string) int {
	if len(t) == 0 {
		return -1
	}

	if id == t[i].ID {
		return i
	}

	if id < t[i].ID {
		return Transport(t[:i]).search(len(t[:i])/2, id)
	} else {
		return Transport(t[i+1:]).search(len(t[i+1:])/2, id)
	}
}
