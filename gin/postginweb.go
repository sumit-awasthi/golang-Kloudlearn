//basic post operation in web using gin
package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("template/*.html")
	//r.GET("/", displayString)
	//r.GET("/user/:name", userInfo)
	r.GET("/html", displayHtml)
	r.POST("/greetings", greet)
	r.Run(":5050")
}
func greet(c *gin.Context) {
	name := c.PostForm("name")
	out := "hello " + name + "how old are u"
	c.String(http.StatusOK, out)
}

/*
type information struct {
	Name string
	Age  int
}
*/
func displayHtml(c *gin.Context) {
	//info := information{Name: "Sumit Awasthi", Age: 23}
	c.HTML(http.StatusOK, "form.html", gin.H{
		"hl": "form", // info)
	})
}

/*
func userInfo(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"User Name is": name,
	})
}

func displayString(c *gin.Context) {
	c.String(http.StatusOK, "HELLO")
}
*/
