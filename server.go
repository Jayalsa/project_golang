package main

import (
	"context"
	"fmt"
	"jayalsa/project_golang/config"
	"jayalsa/project_golang/controllers"
	"jayalsa/project_golang/routes"
	"jayalsa/project_golang/services"
	"log"

	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoClient *mongo.Client
	err         error
	ctx         context.Context
	server      *gin.Engine
)

func main() {
	// ctx = context.Background()
	server = gin.Default()
	InitializeDatabase()
	InitializeProducts()
	InitializeAuthentication()
	ctx1, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer mongoClient.Disconnect(ctx1)
	server.Run(":4000")
}
func home(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "WELCOME TO GIN",
	})
}
func DbOperation() {
	mongoClient, err = config.ConnectDataBase()
	if err != nil {
		fmt.Println("Error in Connection to the Database")
	}
}
func InitializeDatabase() {
	mongoClient, err = config.ConnectDataBase()
	if err != nil {
		log.Fatalf("Unable to connect to Database", err)
	} else {
		fmt.Println("Connected to Database")
	}

}
func InitializeProducts() {
	productCollection := config.GetCollection(mongoClient, "ecommerce", "products")
	productSvc := services.InitProductService(productCollection)
	productCtrl := controllers.InitProductController(productSvc)
	routes.ProductRoutes(server, *productCtrl)
}
func InitializeAuthentication() {
	collection := config.GetCollection(mongoClient, "ecommerce", "users")
	authSvc := services.InitUserService(collection)
	authCtrl := controllers.InitAuthController(authSvc)
	routes.AppRoutes(server, *authCtrl)
	routes.SecuredRoutes(server, *&authCtrl)
}
