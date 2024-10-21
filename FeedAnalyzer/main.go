package main

import (
	"encoding/xml"
	"flag"
	"io"
	"log"
	"net/http"
	"time"
)

var running bool

type Element struct {
	XMLName xml.Name
	Content string `xml:",innerxml"` // Capture the inner XML content
}

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
    UnknownParts []Element `xml:",any"`
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
    UnknownParts []Element `xml:",any"`
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
	UnknownParts []Element `xml:",any"`
}

func parseMainXML(data []byte) (Feed, error) {
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

func parseEntry(entry Entry) {

    //insert into postgresdb


    if entry.UnknownParts != nil {
        log.Printf("entry.UnknownParts: %v\n", entry.UnknownParts)
    }
    if entry.Content.UnknownParts != nil {
        log.Printf("entry.Category: %v\n", entry.Category)
        log.Printf("entry.Content.type: %s\n", entry.Content.Type)
        log.Printf("entry.Content.UnknownParts: %v\n", entry.Content.UnknownParts)
    }

}

func scanSyncFeed(url string) {
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
		feed, err := parseMainXML(data)
		if err != nil {
			log.Printf("error parsing XML: %v", err)
			break
		}

        if feed.UnknownParts != nil {
            log.Printf("feed.UnknownParts: %v\n", feed.UnknownParts)
        }

        for _, entry := range feed.Entry {
            go parseEntry(entry)
        }

		nexturl = findNextLink(feed)
		// todo: store last url with skiptoken in database
	}
}

func main() {
    var scheduleTime int
    var url string

    var database_user string
    var database_password string
    var database_database string
    var database_server string
    var database_port int

    var kafka_server string
    var kafka_topic string


	flag.IntVar(&scheduleTime, "scheduleTime", 5, "Schedule time in seconds")

	flag.StringVar(&url, "url", "https://gegevensmagazijn.tweedekamer.nl/SyncFeed/2.0/Feed", "URL to scan")

    flag.StringVar(&database_user, "db_user", "tweedekamer", "User to connect to db")
    flag.StringVar(&database_password, "db_password", "changethis", "password to connect to db")
    flag.StringVar(&database_database, "db_database", "tweedekamer", "Database name to connect to")
    flag.StringVar(&database_server, "db_server", "postgres", "database server to connect to")
    flag.IntVar(&database_port, "db_port", 5432, "database port to connect to")

    flag.StringVar(&kafka_server, "kafka_server", "kafka:9092", "kafka server to connect to")
    flag.StringVar(&kafka_topic, "kafka_topic", "tweedekamer", "kafka topic to connect to")

	flag.Parse()

	log.Println("scheduleTime:", scheduleTime)
	log.Println("url:", url)

	ticker := time.NewTicker(time.Duration(scheduleTime) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		go scanSyncFeed(url)
	}
}
