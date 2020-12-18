//basic post operation in web using gin
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("template/*")
	r.GET("/html", displayHtml)
	r.POST("/greetings", greet)
	r.Run(":5050")
}
func greet(c *gin.Context) {
	name := c.PostForm("name")
	out := "hello " + name + "how old are u"
	c.String(http.StatusOK, out)
}

func displayHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", gin.H{
		"hl": "form",
	})
}
