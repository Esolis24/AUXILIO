package route

import (
	"net/http"

	"github.com/Esolis24/LabMoviles/controller"
	"github.com/Esolis24/LabMoviles/model"
	"github.com/gin-gonic/gin"
)

func UsersLogin(c *gin.Context) {
	user := model.User{}
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db, _ := c.Get("db")
	conn := controller.DBInit()
	err = controller.GetUser(c)

}
