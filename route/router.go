package route

import (
	"github.com/Esolis24/LabMoviles/controller"	"github.com/gin-gonic/gin"
)


func Router() {
	db := controller.DBInit()
	inDB:= &controller.InDB{DB: db}
	route:=gin.Default();

	route.GET("/", controller.Index)
	route.GET("/user/:email:password", inDB.GetUser)
	route.GET("/users", inDB.GetUsers)
	route.POST("/user", inDB.CreateUser)
	route.PUT("/user", inDB.UpdateUser)
	route.DELETE("/user/:id", inDB.DeleteUser)
	route.Run();
}
