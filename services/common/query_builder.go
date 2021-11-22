package common

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"log"
	"strings"
	"time"
)

func Pagination(db *pg.DB, query string, params interface{}, limit int64, offset int64, total *int64, data interface{}) error {
	_, err := db.Query(pg.Scan(total), fmt.Sprintf("select count(*) as total from (%s) as sub", query), params)
	if err != nil {
		return err
	}

	_, err = db.Query(data, fmt.Sprintf("select * from (%s) as sub LIMIT  %d OFFSET %d", query, limit, offset), params)
	if err != nil {
		return err
	}

	return nil
}

func WhereText(field string, search string) string {
	return fmt.Sprintf("%s LIKE %s", field, "'%"+search+"%'")
}

func WhereTextArrayOr(field string, search []string) string {
	query_list := []string{}

	for _, text := range search {
		query_list = append(query_list, fmt.Sprintf("%s LIKE %s", field, "'%"+text+"%'"))
	}

	return strings.Join(query_list, " OR ")
}

func WhereTextArrayAnd(field string, search []string) string {
	query_list := []string{}

	for _, text := range search {
		query_list = append(query_list, fmt.Sprintf("%s LIKE %s", field, "'%"+text+"%'"))
	}

	return strings.Join(query_list, " AND ")
}

func WhereMatchText(field string, search string) string {
	return fmt.Sprintf("%s LIKE %s", field, "'"+search+"'")
}

func WhereMatch(field string, search string) string {
	return fmt.Sprintf("%s = %s", field, search)
}

func WhereNotMatch(field string, search string) string {
	return fmt.Sprintf("%s <> %s", field, search)
}

func ConvertTime(strTime string, layout string) *time.Time {
	if layout == "" {
		layout = "2006-01-02"
	}

	if strTime == "" {
		return nil
	}
	timeData, err := time.Parse(layout, strTime)
	if err != nil {
		log.Println(err)
		return nil
	}
	return &timeData
}
