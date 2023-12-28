package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	jwt "test/util"
)

func main() {
	r := gin.Default()
	jwtPayLoad := jwt.JwtPayLoad{
		Appid:  "wenli",
		Appkey: "a3b50661fa14c198813711b29acef97e",
	}
	token, _ := jwt.GenToken(jwtPayLoad)
	fmt.Println("2222=====================" + token)
	user, err := jwt.ParseToken(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("========================")
	fmt.Println(user)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": 1,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
