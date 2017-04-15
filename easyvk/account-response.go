package easyvk

type Info struct {
	Response struct {
		Country         string `json:"country"`
		HTTPS           int `json:"https_required"`
		TwoFactor       int `json:"2fa_required"`
		OwnPostsDefault int `json:"own_posts_default"`
		NoWallReplies   int `json:"no_wall_replies"`
		Intro           int `json:"intro"`
		Lang            int `json:"lang"`
	} `json:"response"`
}

type ProfileInfo struct {
	Response struct {
		FirstName       string `json:"first_name"`
		LastName        string `json:"last_name"`
		ScreenName      string `json:"screen_name"`
		Sex             int `json:"sex"`
		Relation        int `json:"relation"`
		Birthday        string `json:"bdate"`
		BirthVisibility int `json:"bdate_visibility"`
		Hometown        string `json:"home_town"`
		Status          string `json:"status"`
		Phone           string `json:"phone"`
		Country struct {
			ID    int `json:"id"`
			Title string `json:"title"`
		} `json:"country"`
		City struct {
			ID    int `json:"id"`
			Title string `json:"title"`
		} `json:"city"`
	} `json:"response"`
}

type Counters struct {
	Response struct {
		Friends            int `json:"friends"`
		FriendsSuggestions int `json:"friends_suggestions"`
		Messages           int `json:"messages"`
		Photos             int `json:"photos"`
		Videos             int `json:"videos"`
		Gifts              int `json:"gifts"`
		Events             int `json:"events"`
		Groups             int `json:"groups"`
		Notifications      int `json:"notifications"`
		SDK                int `json:"sdk"`
		AppRequests        int `json:"app_requests"`
	} `json:"response"`
}
