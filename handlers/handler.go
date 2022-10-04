package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var DB = make(map[int]string)

func GenerateURL(c *gin.Context) {
	rawURL := c.PostForm("url")
	hash := hash(rawURL)
	DB[hash] = rawURL
	c.JSON(http.StatusFound, gin.H{
		"url": "http://127.0.0.1:8080/r/" + strconv.Itoa(hash),
	})

}

func hash(rawURL string) int {
	byteURL := []byte(rawURL)
	var hash int
	for index, number := range byteURL {
		hash += int(number) * (31 ^ (len(byteURL) - index - 1))
	}

	return hash
}

func Redirect(c *gin.Context) {
	url := c.Param("url")
	intURL, err := strconv.Atoi(url)
	if err != nil {
		c.JSON(http.StatusNoContent, gin.H{
			"error": err.Error(),
		})
	}
	redirectURL := DB[intURL]
	c.Redirect(http.StatusMovedPermanently, redirectURL)
}
