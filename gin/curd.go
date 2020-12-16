//curd operation in gin
package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	ID   string `json:"id"`
	NAME string `json:"name"`
	Age  int    `json:"age"`
}

var Users []User //nil
//var Users2 = []User{} //[]empty slice
func main() {
	r := gin.Default()
	userRoutes := r.Group("/users")

	userRoutes.GET("/", getUsers)
	userRoutes.POST("/", CreateUser)
	userRoutes.PUT("/:id", Updateuser)
	userRoutes.DELETE("/:id", Deleteuser)
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err.Error())
		//r.Run()
	}
}
func getUsers(c *gin.Context) {
	c.JSON(200, Users)
}
func CreateUser(c *gin.Context) {
	var Body User
	if err := c.ShouldBindJSON(&Body); err != nil {
		c.JSON(422, gin.H{
			"message": "invalid request body",
		})
		return
	}
	Body.ID = uuid.New().String()
	Users = append(Users, Body)
	c.JSON(200, gin.H{
		"message": "NO error",
	})
}
func Updateuser(c *gin.Context) {
	id := c.Param("id")
	var Body User
	if err := c.ShouldBindJSON(&Body); err != nil {
		c.JSON(422, gin.H{
			"message": "invalid request body",
		})
		return
	}
	for i, u := range Users {
		if u.ID == id {
			Users[i].NAME = Body.NAME
			Users[i].Age = Body.Age

			c.JSON(200, gin.H{
				"message": "No Error found",
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"message": "invalid user id ",
	})
}
func Deleteuser(c *gin.Context) {
	id := c.Param("id")

	for i, u := range Users {
		if u.ID == id {
			Users = append(Users[:i], Users[i+1:]...)
			c.JSON(200, gin.H{
				"message": "No error",
			})
			return
		}
	}
	c.JSON(404, gin.H{
		"nessage": "invalid user id",
	})
}
