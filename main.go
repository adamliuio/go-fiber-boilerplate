package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	if err := godotenv.Load(); err != nil {
		log.Fatalln("Error loading .env file: ", err)
	}
}

func main() {
	log.Println("OK:", os.Getenv("OK"))
	log.Println("IsTestMode:", IsTestMode)
	log.Println("Hostname:", Hostname)
	server()
}
