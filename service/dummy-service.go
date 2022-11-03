package service

import (
	"github.com/dong-tran/goinaction/dao"
	"github.com/dong-tran/goinaction/dto"
	"github.com/dong-tran/goinaction/model"
	"github.com/dong-tran/goinaction/utils"
	"golang.org/x/xerrors"
)

type DummyService struct{}

func (s DummyService) GetAll() ([]*dto.Dummy, error) {
	db, err := utils.NewConnection()
	if err != nil {
		return nil, xerrors.Errorf("%+v", err)
	}
	defer db.Close()
	repo, err := dao.NewDummyRepository(db)
	if err != nil {
		return nil, xerrors.Errorf("%+v", err)
	}
	models, err := repo.FindAll()
	if err != nil {
		return nil, xerrors.Errorf("%+v", err)
	}
	var result []*dto.Dummy
	for _, m := range models {
		id := m.ID
		result = append(result, &dto.Dummy{
			ID:   &id,
			Name: m.Name,
		})
	}
	return result, nil
}

func (s DummyService) Insert(d *dto.Dummy) error {
	db, err := utils.NewConnection()
	if err != nil {
		return xerrors.Errorf("%+v", err)
	}
	defer db.Close()
	repo, err := dao.NewDummyRepository(db)
	if err != nil {
		return xerrors.Errorf("%+v", err)
	}

	err = repo.Insert(nil, &model.Dummy{
		ID:   0,
		Name: d.Name,
	})
	if err != nil {
		return xerrors.Errorf("%+v", err)
	}
	return nil
}
