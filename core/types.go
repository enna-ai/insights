package core

type PageVariables struct {
	UsersNotFollowingMeBack []string
	UsersImNotFollowingBack []string
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
