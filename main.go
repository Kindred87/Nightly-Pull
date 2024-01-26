package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
	"github.com/joho/godotenv"
)

func main() {

	if _, err := os.Stat("creds.env"); err == nil {
		err = godotenv.Load("creds.env")
		if err != nil {
			panic(fmt.Sprintf("Error loading creds.env file: %s", err))
		}
	}

	l := launcher.New().Leakless(false).Headless(true)

	browser := rod.New().ControlURL(l.MustLaunch()).MustConnect()

	defer browser.MustClose()

	page := stealth.MustPage(browser)
	defer page.Close()

	log.Println("Navigating to site...")
	page.MustNavigate(os.Getenv("SITE_URL"))

	page.MustWaitLoad()

	log.Println("Clicking login button...")
	page.MustElement("#navbarNav > ul > li.order-1.order-sm-2.nav-item.d-sm-flex.flex-sm-column.justify-content-sm-center.ms-sm-auto > a").MustClick()

	page.MustWaitLoad()

	siteName := os.Getenv("SITE_NAME")
	username := os.Getenv("SITE_USERNAME")
	password := os.Getenv("SITE_PASSWORD")

	log.Println("Entering username and password...")
	page.MustElement("#email").MustInput(username)
	page.MustElement("#password").MustInput(password)

	log.Println("Logging in...")
	page.MustElement("#btnLogin").MustClick()

	page.MustWaitLoad()

	log.Println("Navigating to site...")
	ele := page.MustElementR("#app > div.wrapper > div > main > div > div.row.justify-content-center > div:nth-child(1) > div > div.card-body > div > a", siteName)
	ele.MustClick()

	page.MustWaitLoad()

	log.Println("Checking if login was successful...")
	ele = page.MustElementR("#company-name", siteName)
	if ele != nil {
		fmt.Println("Element found")
	} else {
		fmt.Println("Element not found")
	}

	time.Sleep(5 * time.Second)
}
