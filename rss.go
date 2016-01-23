package main

import (
	"encoding/xml"
)

// RssFeedXML RssFeedXML
type RssFeedXML struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel *RssFeed
}

// RssImage RssImage
type RssImage struct {
	XMLName xml.Name `xml:"image"`
	URL     string   `xml:"url"`
	Title   string   `xml:"title"`
	Link    string   `xml:"link"`
	Width   int      `xml:"width,omitempty"`
	Height  int      `xml:"height,omitempty"`
}

// RssTextInput RssTextInput
type RssTextInput struct {
	XMLName     xml.Name `xml:"textInput"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Name        string   `xml:"name"`
	Link        string   `xml:"link"`
}

// RssFeed RssFeed
type RssFeed struct {
	XMLName       xml.Name `xml:"channel"`
	Title         string   `xml:"title"`       // required
	Link          string   `xml:"link"`        // required
	Description   string   `xml:"description"` // required
	Language      string   `xml:"language,omitempty"`
	Copyright     string   `xml:"copyright,omitempty"`
	PubDate       string   `xml:"pubDate,omitempty"`       // created or updated
	LastBuildDate string   `xml:"lastBuildDate,omitempty"` // updated used
	Category      string   `xml:"category,omitempty"`
	Generator     string   `xml:"generator,omitempty"`
	Docs          string   `xml:"docs,omitempty"`
	Cloud         string   `xml:"cloud,omitempty"`
	TTL           int      `xml:"ttl,omitempty"`
	Rating        string   `xml:"rating,omitempty"`
	SkipHours     string   `xml:"skipHours,omitempty"`
	SkipDays      string   `xml:"skipDays,omitempty"`
	Image         *RssImage
	TextInput     *RssTextInput
	Items         []*RssItem
}

// RssItem RssItem
type RssItem struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`       // required
	Link        string   `xml:"link"`        // required
	Description string   `xml:"description"` // required
	Author      string   `xml:"author,omitempty"`
	Category    string   `xml:"category,omitempty"`
	Comments    string   `xml:"comments,omitempty"`
	Enclosure   *RssEnclosure
	GUID        string `xml:"guid,omitempty"`    // Id used
	PubDate     string `xml:"pubDate,omitempty"` // created or updated
	Source      string `xml:"source,omitempty"`
}

// RssEnclosure RssEnclosure
type RssEnclosure struct {
	//RSS 2.0 <enclosure url="http://example.com/file.mp3" length="123456789" type="audio/mpeg" />
	XMLName xml.Name `xml:"enclosure"`
	URL     string   `xml:"url,attr"`
	Length  string   `xml:"length,attr"`
	Type    string   `xml:"type,attr"`
}

func feedFromInstapaper(insta *InstapaperFeed) *RssFeedXML {
	var feed = &RssFeedXML{Version: "2.0"}

	feed.Channel = &RssFeed{
		Title:       insta.Title,
		Link:        serverHost,
		Description: "Crafted with ♥ for you!",
	}

	for _, item := range insta.Items {
		feed.Channel.Items = append(feed.Channel.Items, item.toXML())
	}

	return feed
}
