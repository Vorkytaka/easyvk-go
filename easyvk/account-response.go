package easyvk

// An Info describes a set of user's info.
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

// A ProfileInfo describes a set of user's info.
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

// A Counters describes a set of user's counters.
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

// A Permissions describes a set of app's permissions.
type Permissions struct {
	Notify        bool
	Friends       bool
	Photos        bool
	Audio         bool
	Video         bool
	Pages         bool
	Status        bool
	Notes         bool
	Messages      bool
	Wall          bool
	Ads           bool
	Offline       bool
	Docs          bool
	Groups        bool
	Notifications bool
	Stats         bool
	Email         bool
	Market        bool
}

// A Blacklist describes a user's blacklist.
type Blacklist struct {
	Response struct {
		Count int `json:"count"`
		Items []struct {
			ID int `json:"id"`
			FirstName string `json:"first_name"`
			LastName string `json:"last_name"`
			Deactivated string `json:"deactivated"`
		} `json:"items"`
	} `json:"response"`
}