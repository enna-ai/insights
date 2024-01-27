package core

func FindUsersNotFollowingMeBack(followers []FollowersFile, following []FollowersFile) []string {
	followersMap := make(map[string]bool)

	for _, follower := range followers {
		followersMap[follower.StringListData[0].Value] = true
	}

	var listUsersNotFollowingBack []string

	for _, follow := range following {
		if _, ok := followersMap[follow.StringListData[0].Value]; !ok {
			listUsersNotFollowingBack = append(listUsersNotFollowingBack, follow.StringListData[0].Value)
		}
	}

	return listUsersNotFollowingBack
}

func FindUsersImNotFollowingBack(followings []FollowersFile, followers []FollowersFile) []string {
	followingsMap := make(map[string]bool)

	for _, following := range followings {
		followingsMap[following.StringListData[0].Value] = true
	}

	var listUsersImNotFollowingBack []string

	for _, follower := range followers {
		if _, ok := followingsMap[follower.StringListData[0].Value]; !ok {
			listUsersImNotFollowingBack = append(listUsersImNotFollowingBack, follower.StringListData[0].Value)
		}
	}

	return listUsersImNotFollowingBack
}
