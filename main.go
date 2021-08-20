// Command eval is a chromedp example demonstrating how to evaluate javascript
// and retrieve the result.
package main

import (
	"context"
	"log"

	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		// set up a proxy (such as Fiddler) and uncomment the next two lines to see the network requests if it still does not work.
		//chromedp.ProxyServer("localhost:8866"),
		//chromedp.Flag("ignore-certificate-errors", true),
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.106 Safari/537.36"),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel = chromedp.NewContext(
		ctx,
		// uncomment the next line to see the CDP messages
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// run task list
	var res string
	var path string
	path = "#react-root > section > main > div > header > div > div > span > img"
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.instagram.com/malangjualrumah/`),
		chromedp.WaitVisible(path),
		chromedp.Evaluate("document.querySelector('"+path+"').src", &res),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("window object keys: %v", res)
}
