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
	"fmt"
	"log"
	"net/http"

	"github.com/mrjones/oauth"
)

type GwitterError struct {
	Message string
	Code    int
}

func (err GwitterError) Throw(code int, msg string) *GwitterError {
	return &GwitterError{Code: code, Message: msg}
}

func (err GwitterError) Error() string {
	return err.Message
}

func ConfigureConsumer(consumerKey string, consumerSecret string) *oauth.Consumer {
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
		Token:      cfg.AccessToken.Token,
		Secret:     cfg.AccessToken.Secret,
		UserId:     cfg.AccessToken.UserId,
		ScreenName: cfg.AccessToken.ScreenName}
	return at, nil

}
func AuthenticateFirstTime(c *oauth.Consumer) (*TwitterAccessToken, error) {
	requestToken, url, err := c.GetRequestTokenAndUrl("oob")

	if err != nil {
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
		Token:      accessToken.Token,
		Secret:     accessToken.Secret,
		ScreenName: accessToken.AdditionalData["screen_name"],
		UserId:     accessToken.AdditionalData["user_id"],
	}
	return tat, nil
}

func twitterAccessTokenToOauthAccessToken(at *TwitterAccessToken) *oauth.AccessToken {
	return &oauth.AccessToken{Token: at.Token, Secret: at.Secret, AdditionalData: map[string]string{"user_id": at.UserId, "screen_name": at.ScreenName}}
}

func PostTweet(t string, at *TwitterAccessToken, consumer *oauth.Consumer) error {
	if len(t) > 140 {
		errstr := fmt.Sprintf("Your tweet exceeds 140 characters (it is exactly %d character(s) long)", len(t))
		return GwitterError{Code: 1, Message: errstr}
	}
	oauthAt := twitterAccessTokenToOauthAccessToken(at)
	resp, err := consumer.Post("https://api.twitter.com/1.1/statuses/update.json", map[string]string{"status": t}, oauthAt)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	return nil
}

func GetTimeline(at *TwitterAccessToken, consumer *oauth.Consumer, count int) (*[]Tweet, error) {
	var tweets []Tweet
	resp, err := get("https://api.twitter.com/1.1/statuses/home_timeline.json",
		map[string]string{"count": fmt.Sprintf("%v", count)},
		consumer,
		at)

	if err != nil {
		return nil, err
	}

	// DEBUG:
	// defer resp.Body.Close()
	// contents, err := ioutil.ReadAll(resp.Body)
	//fmt.Printf("%s\n", string(contents))
	// END DEBUG

	err = decodeResponse(resp, &tweets)
	if err != nil {
		return nil, err
	}

	return &tweets, nil
}

func get(url string, qs map[string]string, consumer *oauth.Consumer, at *TwitterAccessToken) (*http.Response, error) {
	resp, err := query("get", url, qs, consumer, at)
	return resp, err
}

func post(url string, qs map[string]string, consumer *oauth.Consumer, at *TwitterAccessToken) (*http.Response, error) {
	resp, err := query("post", url, qs, consumer, at)
	return resp, err
}

func query(t string, url string, qs map[string]string, consumer *oauth.Consumer, at *TwitterAccessToken) (*http.Response, error) {
	oauthAt := twitterAccessTokenToOauthAccessToken(at)
	var resp *http.Response
	var err error
	if t == "get" {
		resp, err = consumer.Get(url, qs, oauthAt)
	} else if t == "post" {
		resp, err = consumer.Post(url, qs, oauthAt)
	}

	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return resp, nil
}
