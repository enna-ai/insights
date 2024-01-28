package core

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func AnalyzeInstagramFollowersJSON(c echo.Context) ([]string, []string, error) {
	var followers []FollowersFile
	var followings FollowingFile

	err := parseJSONFile(c, "followers", &followers)
	if err != nil {
		return nil, nil, fmt.Errorf("error (2): Invalid JSON")
		// return nil, nil, err
	}

	err = parseJSONFile(c, "following", &followings)
	if err != nil {
		return nil, nil, fmt.Errorf("error (2): Invalid JSON")
		// return nil, nil, err
	}

	return CheckFollowStatusJSON(followings.RelationshipsFollowing, followers), CheckFollowStatusJSON(followers, followings.RelationshipsFollowing), nil
}

func CheckFollowStatusJSON(followersFile, followingsFile []FollowersFile) []string {	
	m := make(map[string]bool)

	for _, a := range followersFile {
		m[a.StringListData[0].Value] = true
	}

	var result []string

	for _, b := range followingsFile {
		if _, ok := m[b.StringListData[0].Value]; !ok {
			result = append(result, b.StringListData[0].Value)
		}
	}

	return result
}

func AnalyzeInstagramFollowersHTML(c echo.Context) ([]string, []string, error) {
	followers, err := parseHTMLFile(c, "followers")
	if err != nil {
		return nil, nil, fmt.Errorf("error (2): Invalid HTML")
		// return nil, nil, err
	}

	followings, err := parseHTMLFile(c, "following")
	if err != nil {
		return nil, nil, fmt.Errorf("error (2): Invalid HTML")
		// return nil, nil, err
	}

	return followers, followings, nil
}

func CheckFollowStatusHTML(followersFile, followingsFile []string) []string {
	m := make(map[string]struct{})
	var result []string

	for _, a := range followersFile {
		m[a] = struct{}{}
	}

	for _, b := range followingsFile {
		if _, ok := m[b]; !ok {
			result = append(result, b)
		}
	}

	return result
}
