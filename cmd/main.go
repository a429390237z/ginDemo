package main

import (
	"encoding/json"
	"fmt"
	"ginDemo/model"
	"ginDemo/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {

		t := model.NewPerson("张三", 22, "男")
		p := service.NewPeople()
		p.AddPerson(t)
		fmt.Println(p)

		data, err := json.Marshal(p)
		if err != nil {
			fmt.Println("json failed:", err)
			return
		}
		c.JSON(200, gin.H{
			"message": string(data),
		})
	})

	if err := r.Run(); err != nil {
		fmt.Println("start err:", err)
		return
	}
}
