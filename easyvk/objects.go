package easyvk

// An UserObject contains information about user.
// https://vk.com/dev/objects/user
type UserObject struct {
	ID         int `json:"id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Sex        int `json:"sex"`
	Nickname   string `json:"nickname"`
	MaidenName string `json:"maiden_name"`
	Domain     string `json:"domain"`
	ScreenName string `json:"screen_name"`
	Bdate      string `json:"bdate"`
	City struct {
		ID    int `json:"id"`
		Title string `json:"title"`
	} `json:"city"`
	Country struct {
		ID    int `json:"id"`
		Title string `json:"title"`
	} `json:"country"`
	Photo50                string `json:"photo_50"`
	Photo100               string `json:"photo_100"`
	Photo200               string `json:"photo_200"`
	PhotoMax               string `json:"photo_max"`
	Photo200Orig           string `json:"photo_200_orig"`
	Photo400Orig           string `json:"photo_400_orig"`
	PhotoMaxOrig           string `json:"photo_max_orig"`
	PhotoID                string `json:"photo_id"`
	HasPhoto               int `json:"has_photo"`
	HasMobile              int `json:"has_mobile"`
	IsFriend               int `json:"is_friend"`
	FriendStatus           int `json:"friend_status"`
	Online                 int `json:"online"`
	WallComments           int `json:"wall_comments"`
	CanPost                int `json:"can_post"`
	CanSeeAllPosts         int `json:"can_see_all_posts"`
	CanSeeAudio            int `json:"can_see_audio"`
	CanWritePrivateMessage int `json:"can_write_private_message"`
	CanSendFriendRequest   int `json:"can_send_friend_request"`
	MobilePhone            string `json:"mobile_phone"`
	HomePhone              string `json:"home_phone"`
	Site                   string `json:"site"`
	Status                 string `json:"status"`
	LastSeen struct {
		Time     int `json:"time"`
		Platform int `json:"platform"`
	} `json:"last_seen"`
	CropPhoto struct {
		Photo struct {
			ID       int `json:"id"`
			AlbumID  int `json:"album_id"`
			OwnerID  int `json:"owner_id"`
			Photo75  string `json:"photo_75"`
			Photo130 string `json:"photo_130"`
			Photo604 string `json:"photo_604"`
			Width    int `json:"width"`
			Height   int `json:"height"`
			Text     string `json:"text"`
			Date     int `json:"date"`
			PostID   int `json:"post_id"`
		} `json:"photo"`
		Crop struct {
			X  float64 `json:"x"`
			Y  float64 `json:"y"`
			X2 float64 `json:"x2"`
			Y2 float64 `json:"y2"`
		} `json:"crop"`
		Rect struct {
			X  int `json:"x"`
			Y  int `json:"y"`
			X2 int `json:"x2"`
			Y2 int `json:"y2"`
		} `json:"rect"`
	} `json:"crop_photo"`
	Verified         int `json:"verified"`
	FollowersCount   int `json:"followers_count"`
	Blacklisted      int `json:"blacklisted"`
	BlacklistedByMe  int `json:"blacklisted_by_me"`
	IsFavorite       int `json:"is_favorite"`
	IsHiddenFromFeed int `json:"is_hidden_from_feed"`
	CommonCount      int `json:"common_count"`
	Career           []interface{} `json:"career"`
	Military         []interface{} `json:"military"`
	University       int `json:"university"`
	UniversityName   string `json:"university_name"`
	Faculty          int `json:"faculty"`
	FacultyName      string `json:"faculty_name"`
	Graduation       int `json:"graduation"`
	HomeTown         string `json:"home_town"`
	Relation         int `json:"relation"`
	Personal struct {
		Religion   string `json:"religion"`
		InspiredBy string `json:"inspired_by"`
		PeopleMain int `json:"people_main"`
		LifeMain   int `json:"life_main"`
		Smoking    int `json:"smoking"`
		Alcohol    int `json:"alcohol"`
	} `json:"personal"`
	Interests    string `json:"interests"`
	Music        string `json:"music"`
	Activities   string `json:"activities"`
	Movies       string `json:"movies"`
	Tv           string `json:"tv"`
	Books        string `json:"books"`
	Games        string `json:"games"`
	Universities []interface{} `json:"universities"`
	Schools      []interface{} `json:"schools"`
	About        string `json:"about"`
	Relatives    []interface{} `json:"relatives"`
	Quotes       string `json:"quotes"`
	Deactivated  string `json:"deactivated"`
}

// A PhotoObject contains information about photo.
// https://vk.com/dev/objects/photo
type PhotoObject struct {
	ID      int `json:"id"`
	AlbumID int `json:"album_id"`
	OwnerID int `json:"owner_id"`
	UserID  int `json:"user_id"`
	Sizes []struct {
		Src    string `json:"src"`
		Width  int `json:"width"`
		Height int `json:"height"`
		Type   string `json:"type"`
	} `json:"sizes"`
	Text   string `json:"text"`
	Date   int `json:"date"`
	PostID int `json:"post_id"`
	Likes struct {
		UserLikes int `json:"user_likes"`
		Count     int `json:"count"`
	} `json:"likes"`
	Reposts struct {
		Count int `json:"count"`
	} `json:"reposts"`
	Comments struct {
		Count int `json:"count"`
	} `json:"comments"`
	CanComment int `json:"can_comment"`
	Tags struct {
		Count int `json:"count"`
	} `json:"tags"`
}

// A VideoObject contains information about video.
// https://vk.com/dev/objects/video
type VideoObject struct {
	ID          int `json:"id"`
	OwnerID     int `json:"owner_id"`
	Title       string `json:"title"`
	Duration    int `json:"duration"`
	Description string `json:"description"`
	Date        int `json:"date"`
	Comments    int `json:"comments"`
	Views       int `json:"views"`
	Width       int `json:"width"`
	Height      int `json:"height"`
	Photo130    string `json:"photo_130"`
	Photo320    string `json:"photo_320"`
	Photo800    string `json:"photo_800"`
	AddingDate  int `json:"adding_date"`
	Files struct {
		Mp4240 string `json:"mp4_240"`
		Mp4360 string `json:"mp4_360"`
		Mp4480 string `json:"mp4_480"`
		Mp4720 string `json:"mp4_720"`
	} `json:"files"`
	Player     string `json:"player"`
	CanAdd     int `json:"can_add"`
	CanComment int `json:"can_comment"`
	CanRepost  int `json:"can_repost"`
	Likes struct {
		UserLikes int `json:"user_likes"`
		Count     int `json:"count"`
	} `json:"likes"`
	Reposts struct {
		Count        int `json:"count"`
		UserReposted int `json:"user_reposted"`
	} `json:"reposts"`
	Repeat int `json:"repeat"`
}
