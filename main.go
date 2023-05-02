package main

import (
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Course struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"desc"` //กำหนด key ใน json
}

var courses = []Course{
	{
		ID: "1", Name: "TDD", Description: "Lests ",
	},
	{
		ID: "2", Name: "AXD", Description: "Cusrghbsxyjkrxsj6r ",
	},
}

func main() {
	r := gin.Default()
	//courses
	r.GET("/courses", listCourses)
	r.GET("/courses/:id", getCourse)
	r.Run(":8080")
}
func listCourses(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, courses)
}
func getCourse(c *gin.Context) {
	id := c.Param("id")
	for _, course := range courses {
		if course.ID == id {
			c.IndentedJSON(http.StatusOK, course)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "course not found"})
	//
}
