package main

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type RssFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Description string    `xml:"description"`
		Link        string    `xml:"link"`
		Language    string    `xml:"language"`
		Item        []RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Title       string `xml:"title"`
	Description string `xml:"description"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string) (RssFeed, error) {
	// create a http client that timesout on 10 second
	httpClient := http.Client{Timeout: 10 * time.Second}

  // make a get request to url
	resp, err := httpClient.Get(url)
	if err != nil {
		return RssFeed{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RssFeed{}, err
	}

	rssFeed := RssFeed{}

	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return RssFeed{}, err
	}

	return rssFeed, nil
}
