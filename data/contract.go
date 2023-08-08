package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Contract struct
type Contract struct {
	ID                  int        `db:"id,omitempty"`
	PublicProcurementID int        `db:"public_procurement_id,omitempty"`
	SupplierID          int        `db:"supplier_id,omitempty"`
	SerialNumber        string     `db:"serial_number"`
	DateOfSigning       time.Time  `db:"date_of_signing"`
	DateOfExpiry        *time.Time `db:"date_of_expiry"`
	NetValue            float32    `db:"net_value"`
	GrossValue          float32    `db:"gross_value"`
	FileID              *int       `db:"file_id"`
	CreatedAt           time.Time  `db:"created_at,omitempty"`
	UpdatedAt           time.Time  `db:"updated_at"`
}

// Table returns the table name
func (t *Contract) Table() string {
	return "contracts"
}

// GetAll gets all records from the database, using upper
func (t *Contract) GetAll(page *int, size *int, condition *up.Cond) ([]*Contract, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*Contract
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	total, err := res.Count()
	if err != nil {
		return nil, nil, err
	}

	if page != nil && size != nil {
		res = paginateResult(res, *page, *size)
	}

	err = res.All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using upper
func (t *Contract) Get(id int) (*Contract, error) {
	var one Contract
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Contract) Update(m Contract) error {
	m.UpdatedAt = time.Now()
	collection := upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using upper
func (t *Contract) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Contract) Insert(m Contract) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}
