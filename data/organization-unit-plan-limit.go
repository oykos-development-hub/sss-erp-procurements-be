package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// OrganizationUnitPlanLimit struct
type OrganizationUnitPlanLimit struct {
	ID                 int       `db:"id,omitempty"`
	ItemID             int       `db:"item_id"`
	OrganizationUnitID int       `db:"organization_unit_id"`
	Limit              int       `db:"limit_value"`
	CreatedAt          time.Time `db:"created_at"`
	UpdatedAt          time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *OrganizationUnitPlanLimit) Table() string {
	return "organization_unit_plan_limits"
}

// GetAll gets all records from the database, using upper
func (t *OrganizationUnitPlanLimit) GetAll(condition *up.Cond) ([]*OrganizationUnitPlanLimit, error) {
	collection := upper.Collection(t.Table())
	var all []*OrganizationUnitPlanLimit
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	err := res.OrderBy("created_at desc").All(&all)
	if err != nil {
		return nil, err
	}

	return all, err
}

// Get gets one record from the database, by id, using upper
func (t *OrganizationUnitPlanLimit) Get(id int) (*OrganizationUnitPlanLimit, error) {
	var one OrganizationUnitPlanLimit
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *OrganizationUnitPlanLimit) Update(m OrganizationUnitPlanLimit) error {
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
func (t *OrganizationUnitPlanLimit) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *OrganizationUnitPlanLimit) Insert(m OrganizationUnitPlanLimit) (int, error) {
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
