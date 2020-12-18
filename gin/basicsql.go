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

var Users []User
var Users2 = []User{}

// db is an object of sql connection

var err error

var dbmap = initDb()

func initDb() *gorp.DbMap {
	db, err := sql.Open("mysql", "root:akshit#1@tcp(localhost:3306)/user")
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

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Next()
	}
}

func main() {

	r := gin.Default()

	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", GetUsers)
		userRoutes.POST("/", CreateUser)
		userRoutes.PUT("/:id", EditUser)
		userRoutes.DELETE("/:id", DeleteUser)
	}

	if err := r.Run(":5000"); err != nil {
		log.Fatal(err.Error())
	}
}

func GetUsers(c *gin.Context) {

	_, err := dbmap.Select(&Users, "select * from users")

	if err == nil {
		c.JSON(200, Users)
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
		"error": false,
	})
}

func EditUser(c *gin.Context) {
	id := c.Param("id")
	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "invalid request body",
		})
		return

	}

	for i, u := range Users {
		if u.ID == id {
			Users[i].Name = reqBody.Name
			Users[i].Age = reqBody.Age
		}
		c.JSON(200, gin.H{
			"error": false,
		})
		return
	}

	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})

}

func DeleteUser(C *gin.Context) {
	id := C.Param("id")

	for i, u := range Users {
		if u.ID == id {
			Users = append(Users[:i], Users[i+1:]...)

			C.JSON(200, gin.H{
				"error": false,
			})
			return
		}
	}
	C.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})
}
