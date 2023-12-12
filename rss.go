package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
)

type RSS struct {
	Channel   Channel `xml:"channel"`
	XmlnsAtom string  `xml:"_xmlns:atom"`
	Version   string  `xml:"_version"`
}

type Channel struct {
	Title         string        `xml:"title"`
	Link          []LinkElement `xml:"link"`
	Description   string        `xml:"description"`
	Generator     string        `xml:"generator"`
	Language      string        `xml:"language"`
	LastBuildDate string        `xml:"lastBuildDate"`
	Item          []Item        `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	PubDate     string `xml:"pubDate"`
	GUID        string `xml:"guid"`
	Description string `xml:"description"`
}

type LinkClass struct {
	Href   string `xml:"_href"`
	Rel    string `xml:"_rel"`
	Type   string `xml:"_type"`
	Prefix string `xml:"__prefix"`
}

type LinkElement struct {
	LinkClass *LinkClass
	String    *string
}

func fetchFromFeed(url string) (RSS, error) {
	client := http.Client{}
	resp, err := client.Get(url)
	if err != nil {
		return RSS{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return RSS{}, fmt.Errorf("HTTP request failed with status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSS{}, err
	}

	rss := RSS{}
	xml.Unmarshal(body, &rss)
	return rss, nil
}
