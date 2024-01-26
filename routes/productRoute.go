package routes
import (
	"github.com/gin-gonic/gin"
	controller "github.com/Hoan-K-Le/golang-gin-api-ecom/controllers"
	
)

func ProductRoute(router *gin.Engine) {
	router.GET("/all",controller.GetProducts)
	router.GET("/product/:id", controller.GetProductId)
	router.POST("/create", controller.AddProduct)
	router.DELETE("/product/:id", controller.DeleteProduct)
}