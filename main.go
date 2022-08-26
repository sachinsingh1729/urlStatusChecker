package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// map for storing websites url and there status
var URLs = make(map[string]string)

func main() {

	fmt.Println("Starting Server")
	router := gin.Default()
	router.GET("/Websites", getWebsites)
	router.POST("/Websites", postWebsites)

	router.Run("localhost:8080")
}

func getWebsites(c *gin.Context) {

	if url, ok := c.GetQuery("name"); ok {
		fmt.Println("status of", url, URLs[url])
	} else {
		for url, status := range URLs {
			fmt.Println("status of", url, status)
		}
	}
	return
}

func postWebsites(c *gin.Context) {
	var newWebsites []string

	// Call BindJSON to bind the received JSON to
	// newWebsites.
	if err := c.BindJSON(&newWebsites); err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan string)
	// add new Websites to URLs
	for _, url := range newWebsites {
		go addStatus(url, ch)
	}

	for url := range ch {
		go func(link string) {
			// continuously checking status of website after 1 minute
			time.Sleep(60 * time.Second)
			addStatus(link, ch)
		}(url)
	}
	return

}

// for adding the status of website in map
func addStatus(url string, c chan string) {
	if _, err := http.Get("http://" + url); err != nil {
		fmt.Println(url, "is down")
		URLs[url] = "down"
	} else {
		fmt.Println(url, "is up")
		URLs[url] = "up"
	}
	c <- url
	return
}
