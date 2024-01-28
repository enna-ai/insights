package core

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func HandleFileError(c echo.Context, err error, status int, message string) bool {
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		c.String(status, message)
		return true
	}

	return false
}
