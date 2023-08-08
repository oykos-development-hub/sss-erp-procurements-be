package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// Contract article struct
type ContractArticle struct {
	ID         int       `db:"id,omitempty"`
	ArticleID  int       `db:"article_id"`
	ContractID int       `db:"contract_id"`
	Amount     int       `db:"amount"`
	NetValue   *float32  `db:"net_value"`
	GrossValue *float32  `db:"gross_value"`
	CreatedAt  time.Time `db:"created_at,omitempty"`
	UpdatedAt  time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *ContractArticle) Table() string {
	return "contract_articles"
}

// GetAll gets all records from the database, using upper
func (t *ContractArticle) GetAll(condition *up.Cond) ([]*ContractArticle, *uint64, error) {
	collection := upper.Collection(t.Table())
	var all []*ContractArticle
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

	err = res.All(&all)
	if err != nil {
		return nil, nil, err
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using upper
func (t *ContractArticle) Get(id int) (*ContractArticle, error) {
	var one ContractArticle
	collection := upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using upper
func (t *ContractArticle) Update(m ContractArticle) error {
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
func (t *ContractArticle) Delete(id int) error {
	collection := upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using upper
func (t *ContractArticle) Insert(m ContractArticle) (int, error) {
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
