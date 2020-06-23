// https://blog.logrocket.com/how-to-build-a-rest-api-with-golang-using-gin-and-gorm/

package main

import (
	"net/http"
	"userapi/controllers"

	"userapi/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"Welcome": "User Management API"})
	})

	r.GET("/users", controllers.Findusers)
	r.POST("/createuser", controllers.Createuser)
	r.GET("/users/:id", controllers.Finduser)
	r.PATCH("/updateusers/:id", controllers.Updateuser)
	r.DELETE("/deleteusers/:id", controllers.DeleteUser)

	r.Run(":80")
}
