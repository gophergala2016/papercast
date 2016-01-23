package main

import (
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
)

const (
	serverHost = "http://papercast.io/"
)

func proxyLink(link string) string {
	u, _ := url.Parse(serverHost)
	u.Path = "/proxy"
	q := u.Query()
	q.Set("url", link)
	u.RawQuery = q.Encode()
	return u.String()
}

func main() {
	web := gin.Default()

	web.Static("/static", "./assets")
	web.StaticFile("/", "./assets/index.html")

	web.LoadHTMLGlob("templates/*")
	web.GET("/rss/:user/:hash", func(c *gin.Context) {
		userID := c.Param("user")
		hash := c.Param("hash")

		ifNoneMatch := c.Request.Header.Get("If-None-Match")

		i := newInstapaper(userID, hash)
		err := i.fetchInstapaperFeed(ifNoneMatch)
		if err != nil {
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

	web.Run(":3000")
}
