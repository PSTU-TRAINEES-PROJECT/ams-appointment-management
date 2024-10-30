package utils

import (
	"ams-appointment-management/app/common/logger"
	"strconv"

	"github.com/labstack/echo/v4"
)

func ParseQueryParamAsInt(context echo.Context, paramName string) (int, error) {
	value, err := strconv.Atoi(context.Param(paramName))
	if err != nil {
		logger.Error(err)
		return 0, err
	}
	return value, nil
}
