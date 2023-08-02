package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Plan struct
type Plan struct {
	ID               int        `db:"id,omitempty"`
	Year             string     `db:"year"`
	Title            string     `db:"title"`
	Active           bool       `db:"active"`
	SerialNumber     *string    `db:"serial_number"`
	DateOfPublishing *time.Time `db:"date_of_publishing"`
	DateOfClosing    *time.Time `db:"date_of_closing"`
	PreBudgetID      *int       `db:"pre_budget_id"`
	FileID           *int       `db:"file_id"`
	CreatedAt        time.Time  `db:"created_at,omitempty"`
	UpdatedAt        time.Time  `db:"updated_at"`
}

// Table returns the table name
func (t *Plan) Table() string {
	return "plans"
}

// GetAll gets all records from the database, using upper
func (t *Plan) GetAll(page *int, size *int, condition *up.Cond) ([]*Plan, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*Plan
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
func (t *Plan) Get(id int) (*Plan, error) {
	var one Plan
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Plan) Update(m Plan) error {
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
func (t *Plan) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Plan) Insert(m Plan) (int, error) {
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
