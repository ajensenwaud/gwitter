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
  "gopkg.in/gcfg.v1"
)

type Config struct { 
  Main struct { 
    ConsumerKey string
    ConsumerSecret string
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

