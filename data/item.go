package data

import (
	"context"
	"errors"
	"fmt"
	"time"

	up "github.com/upper/db/v4"

	"gitlab.sudovi.me/erp/procurements-api/contextutil"
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

// GetAll gets all records from the database, using Upper

func (t *Item) GetAll(page *int, size *int, condition *up.Cond, orders []interface{}) ([]*Item, *uint64, error) {
	collection := Upper.Collection(t.Table())
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

// Get gets one record from the database, by id, using Upper
func (t *Item) Get(id int) (*Item, error) {
	var one Item
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *Item) Update(ctx context.Context, m Item) error {
	m.UpdatedAt = time.Now()
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return errors.New("user ID not found in context")
	}

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(m.ID)
		if err := res.Update(&m); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *Item) Delete(ctx context.Context, id int) error {
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return errors.New("user ID not found in context")
	}

	err := Upper.Tx(func(sess up.Session) error {
		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())
		res := collection.Find(id)
		if err := res.Delete(); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *Item) Insert(ctx context.Context, m Item) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	userID, ok := contextutil.GetUserIDFromContext(ctx)
	if !ok {
		return 0, errors.New("user ID not found in context")
	}

	var id int

	err := Upper.Tx(func(sess up.Session) error {

		query := fmt.Sprintf("SET myapp.user_id = %d", userID)
		if _, err := sess.SQL().Exec(query); err != nil {
			return err
		}

		collection := sess.Collection(t.Table())

		var res up.InsertResult
		var err error

		if res, err = collection.Insert(m); err != nil {
			return err
		}

		id = getInsertId(res.ID())

		return nil
	})

	if err != nil {
		return 0, err
	}

	return id, nil
}
