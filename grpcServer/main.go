package main

import (
	"context"
	"fmt"
	"jayalsa/project_golang/product"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	product.UnimplementedProductServiceServer
}

func (*server) CreateProduct(ctx context.Context, request *product.Product) (*product.ProductResponse, error) {
	return &product.ProductResponse{Id: "1",
		ErrorMessage: "", Success: request.Name + " Successfully Created"}, nil
}
func main() {
	lis, err := net.Listen("tcp", ":6000")
	defer lis.Close()
	if err != nil {
		log.Printf(err.Error())
		return
	}
	s := grpc.NewServer()
	product.RegisterProductServiceServer(s, &server{})
	fmt.Println("Server listening on port :6000")
	if err := s.Serve(lis); err != nil {
		fmt.Print(err.Error())
	}
}
