package core

import (
	"encoding/json"
	"fmt"

	"github.com/labstack/echo/v4"
)

func ParseJSONFiles(c echo.Context) ([]string, []string, error) {
	var followers []FollowersFile
	var followings FollowingFile

	err := parseJSONFile(c, "followers", &followers)
	if err != nil {
		fmt.Printf("Error parsing followers: %s\n", err)
		return nil, nil, err
	}

	err = parseJSONFile(c, "following", &followings)
	if err != nil {
		fmt.Printf("Error parsing followings: %s\n", err)
		return nil, nil, err
	}

	usersImNotFollowingBack := FindUsersImNotFollowingBack(followings.RelationshipsFollowing, followers)
	usersNotFollowingMeBack := FindUsersNotFollowingMeBack(followers, followings.RelationshipsFollowing)

	return usersImNotFollowingBack, usersNotFollowingMeBack, nil
}

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
