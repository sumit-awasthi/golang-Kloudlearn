package main

import (
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}

var DB *gorm.DB
var students []Student

func FindStudents(c *gin.Context) {
	var students []Student
	DB.Find(&students)

	c.JSON(http.StatusOK, gin.H{"data": students})
}

func FindStudent(c *gin.Context) {
	var student Student
	if err := DB.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": student})
}

func CreateStudent(c *gin.Context) {
	var input Student
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	student := Student{Id: input.Id, Name: input.Name}
	DB.Create(&student)

	c.JSON(http.StatusOK, gin.H{"data": "Record created."})
}

func UpdateStudent(c *gin.Context) {
	var student Student
	if err := DB.Where("id = ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input Student
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	DB.Model(&Student{}).Where("id = ?", c.Param("id")).Update(input)

	c.JSON(http.StatusOK, gin.H{"data": "Record updated."})
}

func DeleteStudent(c *gin.Context) {
	var student Student
	if err := DB.Where("id= ?", c.Param("id")).First(&student).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}
	DB.Where("id = ?", c.Param("id")).Delete(&student)

	c.JSON(http.StatusOK, gin.H{"data": "Record deleted."})
}

func main() {
	// Database connection
	db, err := gorm.Open("postgres", "user=postgres password=root dbname=books_store port=5432 sslmode=disable")

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Student{})

	DB = db

	r := gin.Default()

	// Routes
	r.GET("/students", FindStudents)
	r.GET("/students/:id", FindStudent)
	r.POST("/students", CreateStudent)
	r.PUT("/students/:id", UpdateStudent)
	r.DELETE("/students/:id", DeleteStudent)

	r.Run(":8000")
}
