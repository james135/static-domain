package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file - Exiting")
	}

	// Store IP to check for changes
	cachedIP := ""

	// Construct a new Cloudflare API object
	cf, err := new(cf).New()
	if err != nil {
		log.Fatal("Error setting up Cloudflare - Exiting")
	}

	// Log errors to a file
	f, err := os.OpenFile("static-domain.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v - Exiting", err)
	}
	defer f.Close()
	log.SetOutput(f)

	// Loop every 5 seconds and find this machines public IP
	for {
		liveIP, err := findIPAddress()
		if err != nil {
			log.Println("*** Error fetching IP ***")
			log.Println(err)
			continue
		}
		// Public IP different to locally cached IP - Update Cloudflare DNS records
		if cachedIP != liveIP {
			err = cf.updateAllRecords(liveIP)
			if err != nil {
				log.Println("*** Error updating Cloudflare ***")
				log.Println(err)
				continue
			}
			log.Printf("*** IP updated to %s ***", liveIP)
			cachedIP = liveIP
		}
		time.Sleep(time.Second * 5)
	}
}
