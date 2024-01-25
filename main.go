package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"text/template"

	"github.com/labstack/echo/v4"
)

type PageVariables struct {
	Message               string
	UsersNotFollowingBack []string
	NotFollowingUsers     []string
}

type User struct {
	Href      string `json:"href"`
	Value     string `json:"value"`
	Timestamp int    `json:"timestamp"`
}

type FollowersFile struct {
	Title          string        `json:"title"`
	MediaListData  []interface{} `json:"media_list_data"`
	StringListData []User        `json:"string_list_data"`
}

type FollowingFile struct {
	RelationshipsFollowing []FollowersFile `json:"relationships_following"`
}

func handler(c echo.Context) error {
	pageVariables := PageVariables{
		Message: "GO",
	}

	return c.Render(http.StatusOK, "main.html", pageVariables)
}

func notFound(c echo.Context) error {
	return c.Render(http.StatusNotFound, "404.html", nil)
}

func main() {
	e := echo.New()

	e.Renderer = NewTemplateRenderer("public/layouts/base.html", "public/views/", "public/components/")

	e.GET("/", handler)
	e.GET("/*", notFound)

	e.POST("/upload", func(c echo.Context) error {
		followersFile, err := c.FormFile("followers")
		if err != nil {
			fmt.Printf("Error getting followers file: %s\n", err)
			return c.String(http.StatusBadRequest, "Error getting followers file")
		}

		followersData, err := followersFile.Open()
		if err != nil {
			fmt.Printf("Error opening followers file: %s\n", err)
			return c.String(http.StatusInternalServerError, "Error opening followers file")
		}
		defer followersData.Close()

		followingFile, err := c.FormFile("following")
		if err != nil {
			fmt.Printf("Error getting following file: %s\n", err)
			return c.String(http.StatusBadRequest, "Error getting following file")
		}

		followingData, err := followingFile.Open()
		if err != nil {
			fmt.Printf("Error opening following file: %s\n", err)
			return c.String(http.StatusInternalServerError, "Error opening following file")
		}
		defer followingData.Close()

		var followers []FollowersFile
		err = json.NewDecoder(followersData).Decode(&followers)
		if err != nil {
			fmt.Printf("Error decoding followers data: %s\n", err)
			return c.String(http.StatusInternalServerError, "Error decoding followers data")
		}

		var following FollowingFile
		err = json.NewDecoder(followingData).Decode(&following)
		if err != nil {
			fmt.Printf("Error decoding following data: %s\n", err)
			return c.String(http.StatusInternalServerError, "Error decoding following data")
		}

		notFollowingBack := findUsersImNotFollowing(following.RelationshipsFollowing, followers)
		usersNotFollowingBack := findNotFollowingBack(followers, following.RelationshipsFollowing)

		pageVariables := PageVariables{
			Message:               "Files Uploaded",
			UsersNotFollowingBack: usersNotFollowingBack,
			NotFollowingUsers:     notFollowingBack,
		}

		return c.Render(http.StatusOK, "main.html", pageVariables)
	})

	e.Logger.Fatal(e.Start(":8080"))
}

func findNotFollowingBack(followers []FollowersFile, following []FollowersFile) []string {
	followersMap := make(map[string]bool)

	for _, follower := range followers {
		followersMap[follower.StringListData[0].Value] = true
	}

	var usersNotFollowingMe []string
	for _, follow := range following {
		if _, ok := followersMap[follow.StringListData[0].Value]; !ok {
			usersNotFollowingMe = append(usersNotFollowingMe, follow.StringListData[0].Value)
		}
	}

	return usersNotFollowingMe
}

func findUsersImNotFollowing(followings []FollowersFile, followers []FollowersFile) []string {
	followingsMap := make(map[string]bool)

	for _, following := range followings {
		followingsMap[following.StringListData[0].Value] = true
	}

	var notFollowingBack []string
	for _, follower := range followers {
		if _, ok := followingsMap[follower.StringListData[0].Value]; !ok {
			notFollowingBack = append(notFollowingBack, follower.StringListData[0].Value)
		}
	}

	return notFollowingBack
}

func NewTemplateRenderer(baseTemplate, templateDir, componentDir string) *TemplateRenderer {
	return &TemplateRenderer{
		baseTemplate: template.Must(template.ParseFiles(baseTemplate)),
		templates:    template.Must(template.ParseGlob(templateDir + "*.html")),
		components: template.Must(template.ParseGlob(componentDir + "*.html")),
	}
}

type TemplateRenderer struct {
	baseTemplate *template.Template
	templates    *template.Template
	components   *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	if err := t.baseTemplate.ExecuteTemplate(w, "base.html", data); err != nil {
		return err
	}

	if err := t.templates.ExecuteTemplate(w, name, data); err != nil {
		return err
	}

	return t.components.ExecuteTemplate(w, "footer", data)
}
