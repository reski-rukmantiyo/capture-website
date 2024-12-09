package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables from config.env
	err := godotenv.Load("config.env")
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// Take URL as an input to variable websiteUrl
	if len(os.Args) < 2 {
		fmt.Println("Please provide a URL as an argument.")
		return
	}
	websiteUrl := os.Args[1]

	dockerUrl := os.Getenv("DOCKER_URL")
	if dockerUrl == "" {
		fmt.Println("DOCKER_URL is not set in config.env")
		return
	}

	// create allocator context for use with creating a browser context later
	allocatorContext, cancel := chromedp.NewRemoteAllocator(context.Background(), dockerUrl)
	defer cancel()

	// create context
	ctx, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()

	var buf []byte
	err = chromedp.Run(ctx,
		chromedp.Navigate(websiteUrl),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.WaitReady(`img`, chromedp.ByQueryAll),
		chromedp.ActionFunc(func(ctx context.Context) error {
			return chromedp.Evaluate(`Promise.all(Array.from(document.images).map(img => img.complete ? Promise.resolve() : img.decode()))`, nil).Do(ctx)
		}),
		chromedp.FullScreenshot(&buf, 100), // Take a full screenshot
		chromedp.Sleep(2*time.Second),
	)

	if err != nil {
		fmt.Println(err)
	}

	err = os.WriteFile("output/output.png", buf, 0644)

	if err != nil {
		fmt.Println(err)
	}
}
