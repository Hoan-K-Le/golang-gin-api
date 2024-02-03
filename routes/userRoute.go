package routes
import (
	"github.com/gin-gonic/gin"
	controller "github.com/Hoan-K-Le/golang-gin-api-ecom/controllers"
	
)

func UserRoute(router *gin.Engine) {
	router.POST("/signup", controller.SignUp())
	router.POST("/login", controller.Login())
	router.POST("/logout", controller.LogOut())
}