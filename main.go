package main

import (
  "os"

  "github.com/iamstefin/ddosgo/cmd"
)

func main()  {
  ddosgo.SetupCloseHandler()
  url := os.Args[1]
  ddosgo.Dodos(url)
}
