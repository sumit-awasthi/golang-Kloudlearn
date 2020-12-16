//Query in gin
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
) /*
func Post(c *gin.Context) {
    body :=c.Request.Body
    value,err :=ioutil.ReadAll(body)
    if err !=nil{
        fmt.Println(err,Error())
    }
	c.JSON(200, gin.H{
		"message": string(value),
	})
}*/
func QueryStrings(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}
func Pathparameters(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}
func main() {
	fmt.Println("check the working of API")
	r := gin.Default()
	r.GET("/query", QueryStrings)
	//r.POST("/", Post)
	r.GET("/path/:name/:age", Pathparameters)
	r.Run()
}
