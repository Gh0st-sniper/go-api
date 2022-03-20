package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"io/ioutil"
)

func HomePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"home": "Hello World",
	})
}

func PostPage(c *gin.Context) {
	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(200, gin.H{
		"home": string(value),
	})
}

func Query(c *gin.Context) {
	name := c.Query("name")
	age := c.Query("age")

	c.JSON(200, gin.H{
		"name": name,
		"age":  age,
	})
}

func Path(c *gin.Context) {
	name := c.Param("name")
	age := c.Param("age")

	c.JSON(200, gin.H{

		"name": name,
		"age":  age,
	})
}

func main() {

	server := gin.Default()

	server.GET("/", HomePage)
	server.POST("/", PostPage)
	server.GET("/query", Query)
	server.GET("/query/:name/:age", Path)
	server.Run()
}
