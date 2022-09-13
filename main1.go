package main

// import (
// 	"bufio"
// 	"context"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"os"

// 	"github.com/gin-gonic/gin"
// 	"github.com/mafredri/cdp"
// 	"github.com/mafredri/cdp/devtool"
// 	"github.com/mafredri/cdp/protocol/dom"
// 	"github.com/mafredri/cdp/protocol/page"
// 	"github.com/mafredri/cdp/rpcc"
// )

// func main() {

// 	router := gin.Default()
//     router.GET("/getPDF", run )
//     router.Run(":8080")

// 	// err := run(5 * time.Second)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// }

// func run(c *gin.Context) {
// 	// ctx, cancel := context.WithTimeout(context.Background(), timeout)
// 	// defer cancel()

// 	// ctx, cancel := devtoo.NewContext(context.Background())
//     // defer cancel()

// 	// Use the DevTools HTTP/JSON API to manage targets (e.g. pages, webworkers).
// 	ctx:= context.Background()
// 	devt := devtool.New("http://13.233.204.151:9222")
// 	pt, _ := devt.Create(ctx)
// 	defer devt.Close(ctx, pt)

// 	// Initiate a new RPC connection to the Chrome DevTools Protocol target.
// 	conn, err := rpcc.DialContext(ctx, pt.WebSocketDebuggerURL)
// 	if err != nil {
// 		log.Fatal(err)
// 		// return err
// 	}
// 	defer conn.Close() // Leaving connections open will leak memory.

// 	baseBrowser := cdp.NewClient(conn)

// 	// Open a DOMContentEventFired client to buffer this event.
// 	domContent, err := baseBrowser.Page.DOMContentEventFired(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 		// return err
// 	}
// 	defer domContent.Close()

// 	// Enable events on the Page domain, it's often preferrable to create
// 	// event clients before enabling events so that we don't miss any.
// 	if err = baseBrowser.Page.Enable(ctx); err != nil {
// 		log.Fatal(err)
// 		// return err
// 	}

// 	// Create the Navigate arguments with the optional Referrer field set.
// 	navArgs := page.NewNavigateArgs("https://www.google.com").
// 		SetReferrer("https://duckduckgo.com")
// 	nav, err := baseBrowser.Page.Navigate(ctx, navArgs)
// 	if err != nil {
// 		log.Fatal(err)
// 		// return err
// 	}

// 	// Wait until we have a DOMContentEventFired event.
// 	if _, err = domContent.Recv(); err != nil {
// 		log.Fatal(err)
// 		// return err
// 	}

// 	fmt.Printf("Page loaded with frame ID: %s\n", nav.FrameID)

// 	// Fetch the document root node. We can pass nil here
// 	// since this method only takes optional arguments.
// 	doc, err := baseBrowser.DOM.GetDocument(ctx, nil)
// 	if err != nil {
// 		log.Fatal(err)
// 		// return err
// 	}

// 	// Get the outer HTML for the page.
// 	result, err := baseBrowser.DOM.GetOuterHTML(ctx, &dom.GetOuterHTMLArgs{
// 		NodeID: &doc.Root.NodeID,
// 	})
// 	if err != nil {
// 		log.Fatal(err)
// 		// return err
// 	}

// 	fmt.Printf("HTML: %s\n", result.OuterHTML)

// 	// Capture a screenshot of the current page.
// 	screenshotName := "screenshot.jpg"
// 	screenshotArgs := page.NewCaptureScreenshotArgs().
// 		SetFormat("jpeg").
// 		SetQuality(80)
// 	screenshot, err := baseBrowser.Page.CaptureScreenshot(ctx, screenshotArgs)
// 	if err != nil {
// 		log.Fatal(err)
// 		// return err
// 	}
// 	if err = ioutil.WriteFile(screenshotName, screenshot.Data, 0o644); err != nil {
// 		log.Fatal(err)
// 		// return err
// 	}

// 	fmt.Printf("Saved screenshot: %s\n", screenshotName)

// 	pdfName := "page.pdf"
// 	f, err := os.Create(pdfName)
// 	if err != nil {
// 		log.Fatal(err)
// 		// return err
// 	}

// 	pdfArgs := page.NewPrintToPDFArgs().
// 		SetTransferMode("ReturnAsStream") // Request stream.
// 	pdfData, err := baseBrowser.Page.PrintToPDF(ctx, pdfArgs)
// 	if err != nil {
// 		log.Fatal(err)
// 		// return err
// 	}

// 	sr := baseBrowser.NewIOStreamReader(ctx, *pdfData.Stream)
// 	r := bufio.NewReader(sr)

// 	// Write to file in ~r.Size() chunks.
// 	_, err = r.WriteTo(f)
// 	if err != nil {
// 		log.Fatal(err)
// 		// return err
// 	}

// 	err = f.Close()
// 	if err != nil {

// 	}

// 	fmt.Printf("Saved PDF: %s\n", pdfName)

// 	//Set header
// 	c.Writer.Header().Set("Content-type", "application/octet-stream")
// 	c.File("./page.pdf")

// 	// return nil
// }