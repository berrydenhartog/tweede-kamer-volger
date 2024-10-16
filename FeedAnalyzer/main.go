package main

import (
	"encoding/xml"
	"flag"
	"io"
	"log"
	"net/http"
	"time"
)

var scheduleTime int
var url string
var running bool

// todo: split nicely in separate structs and add more content types
type Feed struct {
	XMLName xml.Name `xml:"feed"`
	Title   string   `xml:"title"`
	Link    []struct {
		Rel  string `xml:"rel,attr"`
		Href string `xml:"href,attr"`
	} `xml:"link"`
	Updated string `xml:"updated"`
	Author  struct {
		Name  string `xml:"name"`
		URI   string `xml:"uri"`
		Email string `xml:"email"`
	} `xml:"author"`
	Rights string  `xml:"rights"`
	ID     string  `xml:"id"`
	Entry  []Entry `xml:"entry"`
}

type Entry struct {
	Title  string `xml:"title"`
	ID     string `xml:"id"`
	Author struct {
		Name string `xml:"name"`
	} `xml:"author"`
	Updated  string `xml:"updated"`
	Category struct {
		Term string `xml:"term,attr"`
	} `xml:"category"`
	Link []struct {
		Rel  string `xml:"rel,attr"`
		Href string `xml:"href,attr"`
	} `xml:"link"`
	Content Content `xml:"content"`
}
type Content struct {
	Type    string `xml:"type,attr"`
	Verslag *struct {
		ID            string `xml:"id,attr"`
		Verwijderd    string `xml:"verwijderd,attr"`
		Bijgewerkt    string `xml:"bijgewerkt,attr"`
		ContentType   string `xml:"contentType,attr"`
		ContentLength string `xml:"contentLength,attr"`
		Vergadering   struct {
			XSIType string `xml:"xsi:type,attr"`
			Ref     string `xml:"ref,attr"`
		} `xml:"vergadering"`
		Soort  string `xml:"soort"`
		Status string `xml:"status"`
	} `xml:"verslag"`
	Unparsed []xml.Attr `xml:",any,attr"`
}

func parseXML(data []byte) (Feed, error) {
	var feed Feed
	err := xml.Unmarshal(data, &feed)
	return feed, err
}

func fetchXML(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}

func findNextLink(feed Feed) string {
	for _, link := range feed.Link {
		if link.Rel == "next" {
			return link.Href
		}
	}
	return ""
}

func parseEntries(entry []Entry) {
	// todo store entries in database

	// todo: convert entry to cloud event for new entry

	// todo: send cloud event to event bus (rabbitmq)

}

func scanSyncFeed() {
	if running {
		return
	}
	running = true
	defer func() {
		running = false
	}()
	nexturl := url

	for nexturl != "" {
		log.Printf("fetching %s", nexturl)
		data, err := fetchXML(nexturl)
		if err != nil {
			log.Printf("error fetching XML: %v", err)
			break
		}
		feed, err := parseXML(data)
		if err != nil {
			log.Printf("error parsing XML: %v", err)
			break
		}

		go parseEntries(feed.Entry)

		nexturl = findNextLink(feed)
		// todo: store last url with skiptoken in database
	}
}

func main() {

	flag.IntVar(&scheduleTime, "scheduleTime", 5, "Schedule time in seconds")
	flag.StringVar(&url, "url", "https://gegevensmagazijn.tweedekamer.nl/SyncFeed/2.0/Feed", "URL to scan")
	flag.Parse()
	log.Println("scheduleTime:", scheduleTime)
	log.Println("url:", url)

	ticker := time.NewTicker(time.Duration(scheduleTime) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		go scanSyncFeed()
	}
}
