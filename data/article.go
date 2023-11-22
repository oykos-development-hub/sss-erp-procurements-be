package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Article struct
type Article struct {
	ID             int       `db:"id,omitempty"`
	ItemID         int       `db:"item_id"`
	Title          string    `db:"title"`
	Description    string    `db:"description"`
	NetPrice       int       `db:"net_price"`
	VATPercentage  string    `db:"vat_percentage"`
	Manufacturer   *string   `db:"manufacturer"`
	Amount         *int      `db:"amount"`
	VisibilityType int       `db:"visibility_type"`
	CreatedAt      time.Time `db:"created_at,omitempty"`
	UpdatedAt      time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *Article) Table() string {
	return "articles"
}

// GetAll gets all records from the database, using upper
func (t *Article) GetAll(condition *up.Cond) ([]*Article, error) {
	collection := upper.Collection(t.Table())
	var all []*Article
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
func (t *Article) Get(id int) (*Article, error) {
	var one Article
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Article) Update(m Article) error {
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
func (t *Article) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Article) Insert(m Article) (int, error) {
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
