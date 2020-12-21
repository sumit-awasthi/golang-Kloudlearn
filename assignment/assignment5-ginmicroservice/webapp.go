//assignment 5
package main

import (
	//"fmt"
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
	DOB      string `json:"dob"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

var err error

var dbmap = initDb()

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/form")
	checkErr(err, "sql.Open failed")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}
	err = dbmap.CreateTablesIfNotExists()
	checkErr(err, "Create tables failed")

	return dbmap
}
func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}

func main() {

	r := gin.Default()

	r.LoadHTMLGlob("template/*.html")
	r.GET("/displayForm", displayForm)
	r.GET("/display/", GetUsers)
	r.GET("/displaysingle/:email", GetUser)
	r.POST("/registration", CreateUser)
	r.PUT("/edituser/:email", EditUser)
	r.DELETE("/deleteuser/:email", DeleteUser)
	r.Run(":5000")
}

func displayForm(c *gin.Context) {
	//info := information{Name: "Sumit Awasthi", Age: 23}
	c.HTML(http.StatusOK, "registration.html", gin.H{
		"hl": "form", // info)
	})
}

func GetUsers(c *gin.Context) {

	var U []User
	_, err := dbmap.Select(&U, "select * from users")

	if err == nil {
		c.JSON(200, U)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}

}

func GetUser(c *gin.Context) {
	email := c.Param("email")
	//fmt.Println(email)
	var U []User
	_, err := dbmap.Select(&U, "select * from users where email = ? LIMIT 1", email)

	if err == nil {
		c.JSON(200, U)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}

}

func CreateUser(c *gin.Context) {

	name := c.PostForm("name")
	password := c.PostForm("password")
	gender := c.PostForm("gender")
	dob := c.PostForm("dob")
	address := c.PostForm("address")
	email := c.PostForm("email")
	phone := c.PostForm("phone")
	statement, err := dbmap.Prepare("INSERT INTO users VALUES(?,?,?,?,?,?,?)")

	_, err = statement.Exec(name, password, gender, dob, address, email, phone)
	if err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}
	c.JSON(200, gin.H{
		"error":   false,
		"message": "new user created",
	})
}

func EditUser(c *gin.Context) {
	email := c.Param("email")
	var Body User
	if err := c.ShouldBindJSON(&Body); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return

	}
	statement, err := dbmap.Prepare("Update users SET name = ?, password = ?, gender = ?, dob = ?, address = ?, phone = ? where email = ?")

	_, err = statement.Exec(Body.Name, Body.Password, Body.Gender, Body.DOB, Body.Address, Body.Phone, email)
	if err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return
	}
	c.JSON(404, gin.H{
		"error":   false,
		"message": "user details updated",
	})

}

func DeleteUser(c *gin.Context) {
	email := c.Param("email")

	statement, err := dbmap.Prepare("DELETE FROM users WHERE email = ?")

	_, err = statement.Exec(email)
	if err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid user id",
		})
		return
	}
	c.JSON(404, gin.H{
		"error":   false,
		"message": "User deleted",
	})

}
