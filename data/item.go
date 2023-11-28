package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Item struct
type Item struct {
	ID               int        `db:"id,omitempty"`
	Title            string     `db:"title"`
	BudgetIndentID   int        `db:"budget_indent_id"`
	PlanID           int        `db:"plan_id"`
	IsOpenProcurment bool       `db:"is_open_procurement"`
	ArticleType      string     `db:"article_type"`
	Status           *string    `db:"status"`
	SerialNumber     *string    `db:"serial_number"`
	DateOfPublishing *time.Time `db:"date_of_publishing"`
	DateOfAwarding   *time.Time `db:"date_of_awarding"`
	FileID           *int       `db:"file_id"`
	CreatedAt        time.Time  `db:"created_at,omitempty"`
	UpdatedAt        time.Time  `db:"updated_at"`
}

// Table returns the table name
func (t *Item) Table() string {
	return "items"
}

// GetAll gets all records from the database, using upper

func (t *Item) GetAll(page *int, size *int, condition *up.Cond, orders []interface{}) ([]*Item, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*Item
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

	err = res.OrderBy(orders...).All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using upper
func (t *Item) Get(id int) (*Item, error) {
	var one Item
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *Item) Update(m Item) error {
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
func (t *Item) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *Item) Insert(m Item) (int, error) {
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
