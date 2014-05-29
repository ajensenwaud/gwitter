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

package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/ajensenwaud/gwitter"
	"github.com/ajensenwaud/gwitter/term"
)

func main() {

	configFlag := flag.Bool("configure", false, "Configure .gwitterrc")
	consumerKeyFlag := flag.String("key", "", "Twitter API consumer key")
	consumerSecretFlag := flag.String("secret", "", "Twitter API consumer secret")
	postFlag := flag.String("post", "", "Post a new entry to Twitter")
	listFlag := flag.Int("list", 0, "List the most recent tweets in your main feed")
	allRecentFlag := flag.Bool("all", false, "List all unread tweets")
	flag.Parse()

	//fmt.Println("configFlag: ", *configFlag)
	// fmt.Println("postFlag: ", *postFlag)
	// fmt.Println("listFlag: ", *listFlag)
	//fmt.Println("allRecentFlag: ", *allRecentFlag)

	// If -configure:
	if *configFlag {
		if len(*consumerKeyFlag) == 0 || len(*consumerSecretFlag) == 0 {
			log.Fatal("You must specify the API consumer secret and consumer key in order to configure Gwitter")
			os.Exit(-1)
		}
		cons := gwitter.ConfigureConsumer(*consumerKeyFlag, *consumerSecretFlag)
		at, err := gwitter.AuthenticateFirstTime(cons)
		if err != nil {
			log.Fatal("Error in authentication: ", err)
			os.Exit(-1)
		}
		fmt.Println("Write the following to ~/.gwitterrc:")
		fmt.Println("[Main]")
		fmt.Println("ConsumerKey = ", *consumerKeyFlag)
		fmt.Println("ConsumerSecret = ", *consumerSecretFlag)
		fmt.Println("")
		fmt.Println("[AccessToken]")
		fmt.Println("Token = ", at.Token)
		fmt.Println("Secret = ", at.Secret)
		fmt.Println("ScreenName = ", at.ScreenName)
		fmt.Println("UserId = ", at.UserId)
		os.Exit(0)
	}

	// If we are not configuring for the first time, it means we are either posting or listing
	// This means we need to configure the consumer API and authenticate in order to do anything:
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	rcFilePath := usr.HomeDir + "/.gwitterrc"
	consumer, err := gwitter.ConfigureApi(rcFilePath)
	if err != nil {
		log.Fatal("Error in loading .gwitterrc file", err)
		os.Exit(-1)
	}
	at, err := gwitter.ConfigureAccessToken(rcFilePath)
	if err != nil {
		log.Fatal("Error in loading access token: ", err)
		os.Exit(-1)
	}

	// If -post <msg> to Twitter
	if len(*postFlag) > 0 {
		gwitter.PostTweet(*postFlag, at, consumer)
		fmt.Println("Tweet posted!")
	}

	// TODO: Implement
	if (*listFlag) > 0 {
		tweets, err := gwitter.GetTimeline(at, consumer, *listFlag)
		if err != nil {
			log.Fatal(err)
			os.Exit(-1)
		}
		printTweets(*tweets)
	}

	// TODO: Implement
	if *allRecentFlag {
	}
}

func printTweets(tweets []gwitter.Tweet) {
	for v := range tweets {
		t := tweets[v]
		// terminal.Stdout.Colorf("@{w}%v\n", t.Text)
		// terminal.Stdout.Colorf("@{y}%s @{w}at @{c}%v@{w}:", t.User.Name, t.CreatedAt).Nl()
		//	terminal.Stdout.Colorf("@{|}@{/}%s", t.Text).Nl()
		fmt.Printf(term.FgYellow+"%s "+term.FgWhite+"(@"+term.FgBlue+"%s"+term.FgWhite+") at "+term.FgCyan+"%v"+term.Reset+":\n", t.User.Name, t.User.ScreenName, t.CreatedAt)
		fmt.Printf(term.FgWhite+"%s\n\n", t.Text)
	}
}
