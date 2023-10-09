package main

import (
	"example/go-rest-api/middleware"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	type Book struct {
		Id     string `json:"id"`
		Title  string `json:"title"`
		Author string `json:"author"`
	}
	var books = []Book{
		{Id: "1", Title: "The Go Programming Language", Author: "<NAME>"},
		{Id: "2", Title: "The Go Programming Language", Author: "<NAME>"},
		{Id: "3", Title: "The Go Programming Language", Author: "<NAME>"},
		{Id: "4", Title: "The Go Programming Language", Author: "<NAME>"},
		{Id: "5", Title: "The Go Programming Language", Author: "<NAME>"},
		{Id: "6", Title: "The Go Programming Language", Author: "<NAME>"},
		{Id: "7", Title: "The Go Programming Language", Author: "<NAME>"},
		{Id: "8", Title: "The Go Programming Language", Author: "<NAME>"},
		{Id: "9", Title: "The Go Programming Language", Author: "<NAME>"},
	}
	router := gin.Default()
	// router.GET("/", getHello)
	// auth := gin.BasicAuth(gin.Accounts{
	// 	"user":  "pass",
	// 	"user2": "pass2",
	// 	"user3": "pass3",
	// })
	router.GET("/getParam/:name/:age", getParam)
	router.GET("/books", func(c *gin.Context) {
		c.JSON(200, books)
	})
	router.POST("/", getDataPost)
	admin := router.Group("/admin", middleware.Authenticate()) // appying miidleware to the group
	{
		admin.GET("/", getHello)
	}
	// router.Use(middleware.Authenticate)
	router.GET("/getData", middleware.Authenticate(),middleware.AddHeader, getData)
	router.GET("/getData2", getData2)
	router.GET("/getdata3", getData3)
	server := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	server.ListenAndServe()
	// router.Run(":8080")
}
func getHello(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hell World",
	})
}
func getParam(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")
	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})

}
func getDataPost(c *gin.Context) {
	body := c.Request.Body
	value, _ := ioutil.ReadAll(body)
	name := c.Query("name")
	age := c.Query("age")
	c.JSON(200, gin.H{
		"message": "Hell World from post",
		"value":   string(value),
		"name":    name,
		"age":     age,
	})
}
func getData(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "HI I am getData method",
	})
}
func getData2(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "HI I am getData2 method",
	})
}
func getData3(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "HI I am getData3 method",
	})
}
