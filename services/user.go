package services

import (
	"apm/model"
	"apm/pkg/entities"
	"apm/pkg/errors"
	"apm/pkg/hash"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"time"
)

type UserService struct {
	db *pg.DB
}

func NewUserService(db *pg.DB) *UserService {
	return &UserService{db: db}
}

func (u UserService) List(search *model.SearchBase) (*model.ListUserResult, error) {

	var rs = &model.ListUserResult{
		Skip:    search.Skip,
		Total:   0,
		Limit:   search.Limit,
		Records: make([]entities.User, 0),
	}
	query := u.db.Model(&rs.Records)

	if search.Search != "" {
		query = query.WhereGroup(func(q *orm.Query) (*orm.Query, error) {
			q = q.WhereOr("username LIKE ?", search.Search).
				WhereOr("first_name LIKE ?", search.Search).
				WhereOr("last_name LIKE ?", search.Search).
				WhereOr("email LIKE ?", search.Search)
			return q, nil
		})
	}

	queryCount := query.Clone()
	var err error
	rs.Total, err = queryCount.Count()
	if err != nil {
		_, err = errors.PGError(err)
		return nil, err
	}

	err = query.Relation("Tenant").
		Relation("Department").
		Limit(search.Limit).
		Offset(search.Skip).
		Column("user.id", "username", "first_name", "last_name", "email").
		Select()
	if err == pg.ErrNoRows {
		return rs, nil
	}
	if err != nil {
		_, err = errors.PGError(err)
		return nil, err
	}
	return rs, nil
}

func (u UserService) Get(id int64) (*entities.User, error) {
	data := new(entities.User)
	data.Id = id
	err := u.db.Model(data).
		Relation("Tenant").
		Relation("Department").
		WherePK().
		Column("user.id", "username", "first_name", "last_name", "permissions", "email", "address", "phone", "user.tenant_id", "user.department_id", "user.created_at", "user.updated_at").
		Select()
	if err == pg.ErrNoRows {
		return nil, errors.NotFound("not found user")
	}
	if err != nil {
		_, err = errors.PGError(err)
		return nil, err
	}
	return data, nil
}

func (u UserService) Store(data *entities.User) (pg.Result, error) {
	pass, err := hash.HashPassword(data.Password)
	if err != nil {
		return nil, err
	}
	data.Password = pass
	rs, err := u.db.Model(data).Insert()
	if err != nil {
		_, err = errors.PGError(err)
		return nil, err
	}
	return rs, nil
}

func (u UserService) Update(id int64, data *entities.User) (pg.Result, error) {
	data.Id = id
	timeNow := time.Now()
	data.UpdatedAt = &timeNow
	rs, err := u.db.Model(data).
		Column("username", "first_name", "last_name", "email", "address", "phone", "permissions", "department_id", "updated_at").
		WherePK().Update()
	if err != nil {
		_, err = errors.PGError(err)
		return nil, err
	}
	return rs, nil
}

func (u UserService) Delete(id int64) (pg.Result, error) {
	data := new(entities.User)
	data.Id = id
	rs, err := u.db.Model(data).
		WherePK().
		Delete()
	if err != nil {
		_, err = errors.PGError(err)
		return nil, err
	}
	return rs, nil
}
