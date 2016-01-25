# Papercast

Podcast of your favourite articles

## Use case

Do you have **dozens of tabs open in your browser**, hoping to get to those articles some day? Save them to a free account on Instapaper and Papercast will create a personal podcast with those articles narrated. Now close those tabs and breathe freely!

## More Information

* High quality voice will narrate your articles to you
* Listen at any speed your podcast player allows
* Links are run through [GoOse](https://github.com/advancedlogic/GoOse) content extractor to only narrate the article itself
* No ads

## Running locally

Currently only supports **MacOS X**, since I use SpeechSynthesis.framework. Edit CFLAGS in both files under mactts/ to point at correct header path.

```bash
# go get
# go run *.go
```

## How To Use

1. Open https://www.instapaper.com/u
2. Click your email (top right corner), then **Download** and choose **RSS Feed**
3. Replace domain in the link from https://www.instapaper.com to http://papercast.xyz to end up with something like `http://papercast.xyz/rss/12345/s0m3h4sh`
4. Add it to your podcast player

## After Gopher Gala 2016

Port to Linux and run as a free service at https://papercast.xyz/

## Subscribe to updates

https://tinyletter.com/papercast
