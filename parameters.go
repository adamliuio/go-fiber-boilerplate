package main

import (
	"log"
	"os"
	"strings"
)

var (
	IsTestMode bool
	Hostname   string
	IsLocal    bool
)

func init() {
	var err error
	if Hostname, err = os.Hostname(); err != nil {
		log.Fatalln(err)
	}
	IsTestMode = strings.Contains(os.Args[0], "test") // checking if it's in test mode
	IsLocal = Hostname == "Adams-MacBook-Pro.local"   // checking if the app in local
}
