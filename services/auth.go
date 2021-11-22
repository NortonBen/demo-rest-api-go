package services

import (
	"apm/model"
	"apm/pkg/entities"
	"apm/pkg/errors"
	"apm/pkg/hash"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-pg/pg/v10"
	"time"
)

type AuthService struct {
	db           *pg.DB
	accessSecret string
}

func NewAuthService(db *pg.DB, AccessSecret string) *AuthService {
	return &AuthService{
		db:           db,
		accessSecret: AccessSecret,
	}
}

func (a AuthService) Login(data *model.Login) (*model.LoginResult, error) {
	user := new(entities.User)
	err := a.db.Model(user).Where("username = ?", data.Username).Select()
	if err == pg.ErrNoRows {
		return nil, errors.BadRequest("not-auth", "not-auth")
	}
	if err != nil {
		_, err = errors.PGError(err)
		return nil, err
	}
	if !hash.CheckPasswordHash(data.Password, user.Password) {
		return nil, errors.BadRequest("not-auth", "not-auth")
	}

	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = user.Id
	atClaims["username"] = user.Username
	atClaims["first_name"] = user.FirstName
	atClaims["last_name"] = user.LastName
	atClaims["tenant_id"] = user.TenantId
	atClaims["permissions"] = user.Permissions
	atClaims["department_id"] = user.DepartmentId
	atClaims["exp"] = time.Now().Add(time.Hour * 10).Unix()

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(a.accessSecret))
	if err != nil {
		return nil, err
	}

	return &model.LoginResult{
		AccessToken: token,
	}, nil
}
