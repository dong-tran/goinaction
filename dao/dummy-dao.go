package dao

import (
	"github.com/dong-tran/goinaction/model"
	"github.com/jmoiron/sqlx"
	"golang.org/x/xerrors"
)

type DummyRepository interface {
	FindAll() ([]model.Dummy, error)
	Insert(tx *sqlx.Tx, d *model.Dummy) error
}

type dummy struct {
	stmt map[int]*sqlx.NamedStmt
}

const (
	findAllQueryIndex = iota
	insertQueryIndex
)

const (
	findAllQuery = "select * from dummy order by id"
	inserQuery   = `insert into dummy
	(name) values (:name)`
)

func NewDummyRepository(db *sqlx.DB) (DummyRepository, error) {
	var statements = map[int]*sqlx.NamedStmt{}
	var err error
	statements[findAllQueryIndex], err = db.PrepareNamed(findAllQuery)
	if err != nil {
		return nil, xerrors.Errorf("%+v", err)
	}
	statements[insertQueryIndex], err = db.PrepareNamed(inserQuery)
	if err != nil {
		return nil, xerrors.Errorf("%+v", err)
	}
	return &dummy{
		stmt: statements,
	}, nil
}

func (r dummy) FindAll() ([]model.Dummy, error) {
	var result []model.Dummy
	if err := r.stmt[findAllQueryIndex].Select(&result, map[string]interface{}{}); err != nil {
		return nil, err
	}
	return result, nil
}

func (r dummy) Insert(tx *sqlx.Tx, d *model.Dummy) error {
	var err error
	if tx != nil {
		_, err = tx.NamedStmt(r.stmt[insertQueryIndex]).Exec(map[string]interface{}{
			"name": d.Name,
		})
	} else {
		_, err = r.stmt[insertQueryIndex].Exec(map[string]interface{}{
			"name": d.Name,
		})
	}
	return err
}
