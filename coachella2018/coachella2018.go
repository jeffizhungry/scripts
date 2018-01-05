package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/lestrrat/go-libxml2/clib"
	"github.com/lestrrat/go-libxml2/dom"
	"github.com/lestrrat/go-libxml2/parser"
	"github.com/lestrrat/go-libxml2/types"
	"github.com/pkg/errors"
)

var flagLinksOnly bool
var flagTries int

var argURL string

func init() {
	argURL = os.Args[1]

	flag.BoolVar(&flagLinksOnly, "linksonly", false, "Don't open chrome tabs, only print links")
	flag.IntVar(&flagTries, "tries", 20, "Specify number of links to create")
	flag.Parse()
}

func main() {
	var links []string
	for i := 0; i < flagTries; i++ {

		// Get new HTML
		fmt.Println("----- FETCHING HTML -----")
		html, err := getHTML(argURL)
		if err != nil {
			fmt.Println("Failed to get HTML:", err.Error())
			return
		}

		// Parse for link with queue id
		fmt.Println("----- PARSE QUEUE LINK -----")
		link, err := parseQueueLink(html)
		if err != nil {
			fmt.Println("Failed to get HTML:", err.Error())
			return
		}
		links = append(links, link)
	}

	// Open links in new chrome tab
	for _, link := range links {
		fmt.Println("----- OPENING QUEUE LINK:", link)
		go openNewChromeWindow(link)
	}
	time.Sleep(30 * time.Second)
}

// getHTML fetches an HTML by URL.
func getHTML(url string) (string, error) {
	c := http.Client{}
	resp, err := c.Get(url)
	if err != nil {
		return "", errors.Wrap(err, "geturl")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "readbody")
	}
	return string(body), nil
}

var reLink = regexp.MustCompile(`^(http[s]?://go.festivalticketing.com/)`)

// parseQueueLink
func parseQueueLink(s string) (string, error) {
	doc, err := parseHTMLWithCustomDefaults(s)
	if err != nil {
		return "", err
	}
	result, err := doc.Find("//a")
	if err != nil {
		return "", err
	}
	var foundLink string
	for _, node := range result.NodeList() {
		el, ok := node.(types.Element)
		if !ok {
			continue
		}
		hrefAttr, err := el.GetAttribute("href")
		if err != nil {
			return "", err
		}
		link := hrefAttr.Value()

		fmt.Println("-- parseQueueLink: a tag: ", link)
		if reLink.MatchString(link) && !strings.Contains(link, "exit") && !strings.Contains(link, "javascriptdisabled") {
			foundLink = link
			fmt.Println("-- parseQueueLink: FOUND LINK: ", foundLink)
		}
	}
	return foundLink, nil
}

func openNewChromeWindow(url string) {
	c := exec.Command("/Applications/Google Chrome.app/Contents/MacOS/Google Chrome", "--new-window", "--kiosk", url)
	if err := c.Run(); err != nil {
		fmt.Println(err.Error())
	}
}

// parseHTMLWithCustomDefaults parses HTML into libxml doc.
func parseHTMLWithCustomDefaults(html string) (types.Document, error) {
	doc, err := parseHTMLString(html,
		parser.HTMLParseRecover|
			parser.HTMLParseNoDefDTD|
			parser.HTMLParseNoNet|
			parser.HTMLParseIgnoreEnc|
			parser.HTMLParseNoError|
			parser.HTMLParseNoWarning|
			parser.HTMLParseCompact,
	)
	return doc, err
}

// parse html into libxml2 dom doc
// we use custom parse function to specify utf8 encoding, since the lib does not
func parseHTMLString(content string, options ...parser.HTMLOption) (types.Document, error) {
	var option parser.HTMLOption
	if len(options) > 0 {
		option = options[0]
	} else {
		option = parser.DefaultHTMLOptions
	}
	docptr, err := clib.HTMLReadDoc(content, "", "utf8", int(option))
	if err != nil {
		return nil, errors.Wrap(err, "parse: readdoc")
	}
	if docptr == 0 {
		return nil, errors.Wrap(clib.ErrInvalidDocument, "parse: invalid doc")
	}
	return dom.WrapDocument(docptr), nil
}
