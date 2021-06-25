package server

import (

	"github.com/Aziz145/testaziz/tree/master/controller"


	"github.com/gin-gonic/gin"
)



func Routes(router *gin.Engine, inDB *controller.InDB) {


	router.POST("/user/", inDB.CreateUser)
	// router.GET("/user/:limit/:offset", inDB.GetUserLimit)
	router.GET("/user/:id", inDB.GetUser)
	router.PUT("/user/:id", inDB.UpdateUser)
	router.DELETE("/user/:id",inDB.DeleteUser)


}
