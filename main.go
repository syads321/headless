// Command eval is a chromedp example demonstrating how to evaluate javascript
// and retrieve the result.
package main

import (
	"context"
	"io/ioutil"
	"log"

	"time"

	"github.com/chromedp/chromedp"
)

func main() {
	// create context
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		// set up a proxy (such as Fiddler) and uncomment the next two lines to see the network requests if it still does not work.
		//chromedp.ProxyServer("localhost:8866"),
		//chromedp.Flag("ignore-certificate-errors", true),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.159 Safari/537.36"),
	)
	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()
	ctx, cancel = chromedp.NewContext(
		ctx,
		// uncomment the next line to see the CDP messages
		//chromedp.WithDebugf(log.Printf),
	)
	defer cancel()
	var buf []byte

	// run task list
	var res string
	//var path string
	//path = "#react-root > section > main > div > header > div > div > span > img"
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.instagram.com/pktabah/?__a=1`),
		RunWithTimeOut(&ctx, 50, chromedp.Tasks{
			chromedp.FullScreenshot(&buf, 60),
		}),
		//chromedp.Evaluate("document.querySelector('"+path+"').src", &res),
	)
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("fullScreenshot2.png", buf, 0o644); err != nil {
		log.Fatal(err)
	}

	log.Printf("window object keys: %v", res)
}

func RunWithTimeOut(ctx *context.Context, timeout time.Duration, tasks chromedp.Tasks) chromedp.ActionFunc {
	return func(ctx context.Context) error {
		timeoutContext, cancel := context.WithTimeout(ctx, timeout*time.Second)
		time.Sleep(15 * time.Second)
		defer cancel()
		return tasks.Do(timeoutContext)
	}
}
