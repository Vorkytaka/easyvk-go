package easyvk

type Info struct {
	Response struct {
		Country         string `json:"country"`
		HTTPSRequired   int `json:"https_required"`
		TwoFaRequired   int `json:"2fa_required"`
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
		Bdate           string `json:"bdate"`
		BdateVisibility int `json:"bdate_visibility"`
		HomeTown        string `json:"home_town"`
		Country struct {
			ID    int `json:"id"`
			Title string `json:"title"`
		} `json:"country"`
		City struct {
			ID    int `json:"id"`
			Title string `json:"title"`
		} `json:"city"`
		Status string `json:"status"`
		Phone  string `json:"phone"`
	} `json:"response"`
}
