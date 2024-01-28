package core

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/net/html"
)

func parseHTMLFile(c echo.Context, fileName string) ([]string, error) {
	file, err := c.FormFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("error getting %s file: %w", fileName, err)
	}

	fileData, err := file.Open()
	if err != nil {
		return nil, fmt.Errorf("error opening %s file: %w", fileName, err)
	}

	defer fileData.Close()

	doc, err := html.Parse(fileData)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s HTML: %w", fileName, err)
	}

	users := extractUsersFromHTML(doc)
	return users, nil
}

func extractUsersFromHTML(n *html.Node) []string {
	var users []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" && strings.Contains(attr.Val, "instagram.com") {
					username := strings.TrimPrefix(attr.Val, "https://www.instagram.com/")
					username = strings.TrimSuffix(username, "/")
					users = append(users, username)
				}
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}

	f(n)

	return users
}
