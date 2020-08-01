package main

import (
  "flag"
  "github.com/iamstefin/ddosgo/ddosgo"
)

var Usage string = `
ddosgo
`

func main()  {
  ddosgo.SetupCloseHandler()
  url := flag.String("url", "", "The target URL")
  flag.Parse()
  ddosgo.Dodos(*url)
}
