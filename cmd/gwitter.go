package main

import ( 
  "github.com/ajensenwaud/gwitter"
  "fmt"
  "log"
)

func main() { 
  consumer, err := gwitter.ConfigureApi("/home/aj/.gwitterrc")
  if err != nil { 
    log.Fatal("ConfigureApi error:", err)
    return
  }
  
  at, err := gwitter.ConfigureAccessToken("/home/aj/.gwitterrc")
  if err != nil { 
    log.Fatal("ConfigureAccessToken error:", err)
  }
  /*
  if err != nil { 
    fmt.Println(err)
  }
  at, err := gwitter.AuthenticateFirstTime(consumer)
  if err != nil { 
    fmt.Println(err)
  } else { 
    fmt.Println("OK: ", at)
  }*/
  t := "horse! neigh!"
  fmt.Println("Sending the following tweet: ", t)
  gwitter.PostTweet(t, at, consumer) 
}
