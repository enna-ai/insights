package core

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
)

func parseJSONFile(c echo.Context, fileName string, v interface{}) error {
	file, err := c.FormFile(fileName)
	if err != nil {
		return fmt.Errorf("error getting %s file: %w", fileName, err)
	}

	fileData, err := file.Open()
	if err != nil {
		return fmt.Errorf("error opening %s file: %w", fileName, err)
	}

	defer fileData.Close()

	err = json.NewDecoder(fileData).Decode(v)
	if err != nil {
		return fmt.Errorf("error decoding %s data: %w", fileName, err)
	}

	return nil
}
