package main

import (
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type webData struct {
	name       string
	linkCount  int
	imageCount int
	links      []string
}

//var m = map[string]*webData{"foo.com": {"foo.com", 100, 100, []string{"a.txt", "b.txt"}}}

var websiteMap map[string]webData

func initLinkParser() {
	//var m = map[string]*webData{"foo.com": {"foo.com", 100, 100, []string{"a.txt", "b.txt"}}}
	websiteMap = make(map[string]webData)
}

/*
https://www.devdungeon.com/content/web-scraping-go#download_a_url
https://stackoverflow.com/questions/39292672/how-to-build-a-map-of-struct-and-append-values-to-it-using-go
*/
func parseURL(name string, URL string) {

	resp, err := http.Get(URL)
	if err != nil {
		println("error reading url")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		println("Error. Incorrect response")
	}

	/*bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			println("error reading url")
		}
		pageContent := string(bodyBytes)
	    ///println(bodyString)

	    linkStartIndex := strings.Index(pageContent, "<a href>")*/

	document, err := goquery.NewDocumentFromReader(resp.Body)

	var count int = 0
	var links []string
	//links []

	document.Find("a[href]").Each(func(index int, item *goquery.Selection) {
		href, _ := item.Attr("href")
		//fmt.Printf("link: %s - anchor text: %s\n", href, item.Text())
		count++
		links = append(links, href)
	})

	//println("links parsed. count:", count)
	wData := webData{name, count, count, links}

	websiteMap[name] = wData

	//return websiteMap

}

func main() {
	//println("html link parser")

	var wg sync.WaitGroup
	wg.Add(3)
	initLinkParser()

	go func() {
		defer wg.Done()
		parseURL("flipkart", "http://www.flipkart.com")
	}()

	go func() {
		defer wg.Done()
		parseURL("amazon", "http://www.amazon.in")
	}()

	go func() {
		defer wg.Done()
		parseURL("snapdeal", "http://www.snapdeal.com")
	}()

	wg.Wait()

	for _, element := range websiteMap {
		println("name:", element.name, "linkcount:", element.linkCount)

		/*for _, link := range element.links {
			println("link:", link)
		}*/
	}

}
