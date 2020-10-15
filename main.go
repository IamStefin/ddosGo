package main

import (
	"flag"
	"github.com/iamstefin/ddosgo/ddosgo"
)

var Usage string = `
ddosgo
`

func main() {
	ddosgo.SetupCloseHandler()
	url := flag.String("url", "", "The Target URL")
	no := flag.Int("n", 200, "Maximum number of concurrent requests")
	flag.Parse()
	if *url != "" {
		ddosgo.Dodos(*url, *no)
	} else {
		flag.Usage()
	}
}
