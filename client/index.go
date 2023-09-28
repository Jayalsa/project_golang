package main

import (
	"context"
	"fmt"
	"jayalsa/project_golang/product"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":6000", grpc.WithInsecure())
	if err != nil {
		fmt.Println(err.Error())
	}
	defer conn.Close()
	c := product.NewProductServiceClient(conn)
	consumer := product.Product{ID: "101", Name: "Iphone15", Description: "Iphone", Price: 85000}
	result, err := c.CreateProduct(context.Background(), &consumer)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}
}
