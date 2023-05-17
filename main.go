// GithubActionsTest
package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("Start")

	log.Printf("2 + 2 = %d\n", AddTwoNumbers(2, 2))

	r := SetupRouter()
	r.Run()

	log.Println("End")
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/truth", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("2 + 2 = %d. (if it is not 4 then we are in trouble)", AddTwoNumbers(2, 2)),
		})
	})
	r.POST("/addTwoNumbers", AddTwoNumbersHandler)
	return r
}

func AddTwoNumbersHandler(c *gin.Context) {
	n1Str := c.PostForm("n1")
	n2Str := c.PostForm("n2")

	n1, err1 := strconv.Atoi(n1Str)
	n2, err2 := strconv.Atoi(n2Str)

	if err1 != nil || err2 != nil {
		c.JSON(422, gin.H{
			"message": "Invalid arguments",
		})
	} else {
		ans := AddTwoNumbers(n1, n2)
		c.JSON(200, gin.H{
			"ans": ans,
			"message": fmt.Sprintf("%d + %d = %d", n1, n2, ans),
		})
	}
}

func AddTwoNumbers(n1, n2 int) int {
	return n1 + n2
}
