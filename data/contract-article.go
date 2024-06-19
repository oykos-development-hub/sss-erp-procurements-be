package data

import (
	"time"

	up "github.com/upper/db/v4"
	newErrors "gitlab.sudovi.me/erp/procurements-api/pkg/errors"
)

// Contract article struct
type ContractArticle struct {
	ID         int       `db:"id,omitempty"`
	ArticleID  int       `db:"article_id"`
	ContractID int       `db:"contract_id"`
	NetValue   *int      `db:"net_value"`
	GrossValue *int      `db:"gross_value"`
	CreatedAt  time.Time `db:"created_at,omitempty"`
	UpdatedAt  time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *ContractArticle) Table() string {
	return "contract_articles"
}

// GetAll gets all records from the database, using Upper
func (t *ContractArticle) GetAll(condition *up.Cond, orders []interface{}) ([]*ContractArticle, *uint64, error) {
	collection := Upper.Collection(t.Table())
	var all []*ContractArticle
	var res up.Result

	if condition != nil {
		res = collection.Find(*condition)
	} else {
		res = collection.Find()
	}

	total, err := res.Count()
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "upper count")
	}

	err = res.OrderBy(orders...).All(&all)
	if err != nil {
		return nil, nil, newErrors.Wrap(err, "upper order by")
	}

	return all, &total, err
}

// Get gets one record from the database, by id, using Upper
func (t *ContractArticle) Get(id int) (*ContractArticle, error) {
	var one ContractArticle
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, newErrors.Wrap(err, "upper get")
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *ContractArticle) Update(m ContractArticle) error {
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return newErrors.Wrap(err, "upper update")
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *ContractArticle) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return newErrors.Wrap(err, "upper delete")
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *ContractArticle) Insert(m ContractArticle) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, newErrors.Wrap(err, "upper insert")
	}

	id := getInsertId(res.ID())

	return id, nil
}
