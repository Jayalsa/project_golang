package routes

import (
	"jayalsa/project_golang/controllers"
	"jayalsa/project_golang/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine, p controllers.ProductController) {
	product := r.Group("/api/product")
	product.POST("/insert", p.InsertProduct)
	product.GET("/getProducts", p.GetProducts)
	product.GET("/getProducts/:id", p.GetProductByID)

}
func SecuredRoutes(g *gin.Engine, a *controllers.AuthController) {
	g.Use(middleware.Authenticate())
	g.POST("/api/users/logout", a.Logout)
}
func AppRoutes(r *gin.Engine, a controllers.AuthController) {
	//user routes
	user := r.Group("/api/user")

	user.POST("/register", a.Register)
	user.POST("/login", a.Login)
}
