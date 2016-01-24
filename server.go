package main

import (
	"log"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

const (
	serverHost = "http://papercast.io/"
)

func generateProxyLink(link string) (string, error) {
	s, _ := url.Parse(serverHost)
	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}

	s.Path = "/proxy/" + u.Scheme + "/" + u.Host + u.RequestURI()
	s.RawQuery = u.Query().Encode()
	return s.String(), nil
}

func main() {
	web := gin.Default()

	web.Static("/static", "./assets")
	web.StaticFile("/", "./assets/index.html")

	web.GET("/rss/:user/:hash", func(c *gin.Context) {
		userID := c.Param("user")
		hash := c.Param("hash")

		ifNoneMatch := c.Request.Header.Get("If-None-Match")

		i := newInstapaper(userID, hash)
		err := i.fetchInstapaperFeed(ifNoneMatch)
		if err != nil {
			log.Println(err)
			c.XML(http.StatusBadRequest, nil)
			return
		}

		if i.notModified {
			c.Header("Etag", i.etag)
			c.XML(http.StatusNotModified, nil)
			return
		}

		feed := feedFromInstapaper(i.feed)
		c.XML(http.StatusOK, feed)
	})

	web.GET("/proxy/:scheme/:host/*path", func(c *gin.Context) {
		scheme := c.Param("scheme")
		host := c.Param("host")
		path := c.Param("path")
		query := c.Request.URL.Query().Encode()

		u := url.URL{Scheme: scheme, Host: host, Path: path, RawQuery: query}

		article, err := extractFromURL(u)
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "")
			return
		}

		audioBuffer, err := speak(article)
		if err != nil {
			log.Println(err)
			c.String(http.StatusInternalServerError, "")
			return
		}

		err = audioBuffer.WriteTo(c.Writer)
		if err != nil {
			log.Println(err)
		}
	})

	web.Run(":3000")
}
