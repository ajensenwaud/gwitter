package gwitter 

import (
  "code.google.com/p/gcfg"
)

type Config struct { 
  Main struct { 
    ConsumerKey string
    ConsumerSecret string
    // VerificationCode string
  }
  AccessToken struct {
    Token string
    Secret string
    UserId string
    ScreenName string
  }
}

func ReadFromFile(f string) (Config, error) { 
  var cfg Config
  err := gcfg.ReadFileInto(&cfg, f)
  return cfg, err
}

