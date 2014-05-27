/*
 * -----------------------------------------------------------------------------------
 * "THE BEER-WARE LICENSE" (Revision 42):
 * @aojensen wrote this file. As long as you retain this notice you
 * can do whatever you want with this stuff. If we meet some day, and you think
 * this stuff is worth it, you can buy me a beer in return.
 * Anders Jensen-Waud
 * 
 * For more information about the BEER-WARE license please see phk's web site: 
 * http://people.freebsd.org/~phk/
 * -----------------------------------------------------------------------------------
 */

package gwitter 

import ( 
  "github.com/mrjones/oauth"
  "log"
  "fmt"
  // "io/ioutil"
)

func ConfigureConsumer(consumerKey string, consumerSecret string) (*oauth.Consumer) { 
  c := oauth.NewConsumer(
    consumerKey, 
    consumerSecret, 
    oauth.ServiceProvider{
          RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
          AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
          AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
        })
  return c
}

func ConfigureApi(cfgFile string) (*oauth.Consumer, error) { 
 cfg, err := ReadFromFile(cfgFile)
  if err != nil {
    return nil, err
  }
  c := oauth.NewConsumer(
    cfg.Main.ConsumerKey, 
    cfg.Main.ConsumerSecret,
    oauth.ServiceProvider{
          RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
          AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
          AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
        })
  return c, nil 
}

func ConfigureAccessToken(cfgFile string) (*TwitterAccessToken, error) { 
  cfg, err := ReadFromFile(cfgFile)
  if err != nil { 
    return nil, err
  }
  at := &TwitterAccessToken{
    Token: cfg.AccessToken.Token, 
    Secret: cfg.AccessToken.Secret, 
    UserId: cfg.AccessToken.UserId, 
    ScreenName: cfg.AccessToken.ScreenName}
  return at, nil

}
func AuthenticateFirstTime(c *oauth.Consumer) (*TwitterAccessToken, error) { 
   requestToken, url, err := c.GetRequestTokenAndUrl("oob")

    if  err != nil { 
    log.Fatal(err)
    log.Fatal("GetRequestTokenAndUrl failed")
    return nil, err
  }
  fmt.Println("Go to this url: " + url)
  fmt.Print("Enter the verification code here:")

  verificationCode := ""
  fmt.Scanln(&verificationCode)

  accessToken, err := c.AuthorizeToken(requestToken, verificationCode) 
  if err != nil { 
    log.Fatal(err)
    log.Fatal("AuthorizeToken failed")
    return nil, err
  }
  fmt.Println("Add the following information to .gwitterrc:")
  fmt.Println("[AccessToken]")
  fmt.Println("Token: " + accessToken.Token)
  fmt.Println("Secret: " + accessToken.Secret)
  fmt.Println("ScreenName: " + accessToken.AdditionalData["screen_name"])
  fmt.Println("UserId: " + accessToken.AdditionalData["user_id"])

  tat := &TwitterAccessToken{
    Token: accessToken.Token, 
    Secret: accessToken.Secret, 
    ScreenName: accessToken.AdditionalData["screen_name"], 
    UserId: accessToken.AdditionalData["user_id"],
  }
  return tat, nil
}
func PostTweet(t string, at *TwitterAccessToken, consumer *oauth.Consumer) {
  oauthAt := &oauth.AccessToken{ Token: at.Token, Secret: at.Secret, AdditionalData: map[string]string { "user_id": at.UserId, "screen_name": at.ScreenName }}
  resp, err := consumer.Post("https://api.twitter.com/1.1/statuses/update.json", map[string]string{ "status": t}, oauthAt)
  if err != nil { 
    log.Fatal(err) 
  }
  fmt.Println(resp) 
    
}


