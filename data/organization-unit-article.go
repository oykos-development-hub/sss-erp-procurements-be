package data

import (
	"time"

	up "github.com/upper/db/v4"
)

// OrganizationUnitArticle struct
type OrganizationUnitArticle struct {
	ID                  int       `db:"id,omitempty"`
	ArticleID           int       `db:"article_id"`
	OrganizationUnitID  int       `db:"organization_unit_id"`
	Amount              int       `db:"amount"`
	Status              string    `db:"status"`
	IsRejected          bool      `db:"is_rejected"`
	UsedArticles        int       `db:"used_articles"`
	RejectedDescription *string   `db:"rejected_description"`
	CreatedAt           time.Time `db:"created_at,omitempty"`
	UpdatedAt           time.Time `db:"updated_at"`
}

// Table returns the table name
func (t *OrganizationUnitArticle) Table() string {
	return "organization_unit_articles"
}

// GetAll gets all records from the database, using Upper
func (t *OrganizationUnitArticle) GetAll(condition *up.Cond) ([]*OrganizationUnitArticle, error) {
	collection := Upper.Collection(t.Table())
	var all []*OrganizationUnitArticle
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

// Get gets one record from the database, by id, using Upper
func (t *OrganizationUnitArticle) Get(id int) (*OrganizationUnitArticle, error) {
	var one OrganizationUnitArticle
	collection := Upper.Collection(t.Table())

	res := collection.Find(up.Cond{"id": id})
	err := res.One(&one)
	if err != nil {
		return nil, err
	}
	return &one, nil
}

// Update updates a record in the database, using Upper
func (t *OrganizationUnitArticle) Update(m OrganizationUnitArticle) error {
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res := collection.Find(m.ID)
	err := res.Update(&m)
	if err != nil {
		return err
	}
	return nil
}

// Delete deletes a record from the database by id, using Upper
func (t *OrganizationUnitArticle) Delete(id int) error {
	collection := Upper.Collection(t.Table())
	res := collection.Find(id)
	err := res.Delete()
	if err != nil {
		return err
	}
	return nil
}

// Insert inserts a model into the database, using Upper
func (t *OrganizationUnitArticle) Insert(m OrganizationUnitArticle) (int, error) {
	m.CreatedAt = time.Now()
	m.UpdatedAt = time.Now()
	collection := Upper.Collection(t.Table())
	res, err := collection.Insert(m)
	if err != nil {
		return 0, err
	}

	id := getInsertId(res.ID())

	return id, nil
}
