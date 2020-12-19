package main

import (
	"database/sql"

	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type User struct {
	ID   string
	Name string
	Age  int
}

var err error

var dbmap = initDb()

func initDb() *gorp.DbMap {
	//db, err = sql.Open("mysql", "root:password123@(127.0.0.1:3306)/user)
	db, err := sql.Open("mysql", "root:password123@tcp(localhost:3306)/user")
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

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", GetUsers)
		userRoutes.GET("/:id", GetUser)
		userRoutes.POST("/", CreateUser)
		userRoutes.PUT("/:id", EditUser)
		userRoutes.DELETE("/:id", DeleteUser)
	}

	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
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
	id := c.Param("id")
	var U []User
	_, err := dbmap.Select(&U, "select * from users where id = ? LIMIT 1", id)

	if err == nil {
		c.JSON(200, U)
	} else {
		c.JSON(404, gin.H{"error": "user not found"})
	}

}

func CreateUser(c *gin.Context) {
	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return

	}
	statement, err := dbmap.Prepare("INSERT INTO users(id,name,age)VALUES(?,?,?)")

	reqBody.ID = uuid.New().String()

	_, err = statement.Exec(reqBody.ID, reqBody.Name, reqBody.Age)
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
	id := c.Param("id")
	var Body User
	if err := c.ShouldBindJSON(&Body); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return

	}
	statement, err := dbmap.Prepare("Update users SET name = ?, age = ? where id = ?")

	_, err = statement.Exec(Body.Name, Body.Age, id)
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
	id := c.Param("id")

	statement, err := dbmap.Prepare("DELETE FROM users WHERE id = ?")

	_, err = statement.Exec(id)
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
