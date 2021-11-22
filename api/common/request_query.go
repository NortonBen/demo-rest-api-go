package common

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func GetSkipLimit(cxt *fiber.Ctx) (int, int) {
	limit := GetNumberToQuery(cxt, "limit", 50)
	skip := GetNumberToQuery(cxt, "skip", 0)
	return int(skip), int(limit)
}

func GetNumberToQuery(cxt *fiber.Ctx, name string, value int64) int64 {
	numberN := value
	number := cxt.Query(name)

	if number != "" {
		rs, err := strconv.ParseInt(number, 10, 64)
		if err == nil {
			numberN = rs
		}
	}

	return numberN
}
