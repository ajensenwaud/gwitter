package gwitter

import (
  "testing"
  "fmt"
  "github.com/mrjones/oauth"
  "log"
  "io/ioutil"
)

func Test33BasicAuth(t *testing.T) {
  at := &oauth.AccessToken{ 
    Token: "XXX",
    Secret: "XXX",
    AdditionalData: map[string]string{ "user_id": "14732690", "screen_name": "aojensen" },
  }
  cfg, err := ReadFromFile("/home/aj/.gwitterrc")
  
  c := oauth.NewConsumer(
    cfg.Main.ConsumerKey, 
    cfg.Main.ConsumerSecret,
    oauth.ServiceProvider{
          RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
          AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
          AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
        })

  response, err := c.Get("https://api.twitter.com/1.1/statuses/home_timeline.json", 
    map[string]string{"count": "1"},
    at)
  if err != nil { 
    log.Fatal(err)
  }
  defer response.Body.Close()
  bits, err := ioutil.ReadAll(response.Body)
  fmt.Println("The newest item in your home timeline is: " + string(bits))
}


