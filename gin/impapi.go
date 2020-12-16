//basic api implementation using gin
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

//GET
func Homepage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is the GET page",
	})
}

//POST
/*
func Post(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is the POST page",
	})
}*/

func Post(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}
	c.JSON(200, gin.H{
		"message": string(value),
	})
}

//Put
func Put(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is the PuT page",
	})
}

//DELETE
func Delete(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "this is the DELETE page",
	})
}
func main() {
	fmt.Println("check the working of API")

	r := gin.Default()
	r.GET("/", Homepage)
	r.POST("/", Post)
	r.PUT("/", Put)
	r.DELETE("/", Delete)
	r.Run()
}
