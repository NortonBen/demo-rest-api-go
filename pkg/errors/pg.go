package errors

import (
	"github.com/go-pg/pg/v10"
	"log"
)

type ErrorCM struct {
	message string
}

func (e ErrorCM) Error() string {
	return e.message
}

func PGError(err error) (pg.Error, error) {
	switch err.(type) {
	case pg.Error:
		{
			pgError := err.(pg.Error)
			log.Println(pgError.Field('H'), ";", pgError.Field('M'), ";", pgError.Field('C'), ";", pgError.Field('W'), ";", pgError.Field('D'))
			return pgError, &Error{
				Code: 500,
				Message: pgError.Field('M'),
				ShortCode: pgError.Field('M'),
			}
		}
	default:
		return nil, err
	}
}


