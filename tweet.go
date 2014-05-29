package gwitter

type Tweet struct {
	CreatedAt string `json:"created_at"`
	Text      string `json:"text"`
	Retweeted bool   `json:"retweeted"`
	User      User   `json:"user"`
	Source    string `json:"source"`
}

type User struct {
	Name       string `json:"name"`
	ScreenName string `json:"screen_name"`
}

type TwitterAccessToken struct {
	Token      string
	Secret     string
	ScreenName string
	UserId     string
}
