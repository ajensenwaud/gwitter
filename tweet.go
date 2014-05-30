package gwitter

type Tweet struct {
	CreatedAt            string `json:"created_at"`
	Text                 string `json:"text"`
	Retweeted            bool   `json:"retweeted"`
	RetweetedStatus      *Tweet `json:"retweeted_status`
	User                 User   `json:"user"`
	Source               string `json:"source"`
	Id                   int64  `json:"id"`
	IdStr                string `json:"id_str"`
	InReplyToScreenName  string `json:"in_reply_to_screen_name"`
	InReplyToStatusId    int64  `json:"in_reply_to_status_id"`
	InReplyToStatusIdStr string `json:"in_reply_to_status_id_str"`
	InReplyToUserId      int64  `json:"in_reply_to_user_id"`
	InReplyToUserIdStr   string `json:"in_reply_to_user_id_str"`
	RetweetCount         int    `json:"retweet_count"`
	Truncated            bool   `json:"truncated"`
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
