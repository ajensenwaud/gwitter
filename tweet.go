package gwitter

type Tweet struct { 
  CreatedAt string `json:"created_at"`
  Text string `json:"text"`
  Retweeted bool `json:"retweeted"`
}

type User struct { 
  Name string `json:"name"`
}

type TwitterAccessToken struct { 
  Token string
  Secret string
  ScreenName string
  UserId string
}
  
