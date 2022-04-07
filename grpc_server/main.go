package main

import (
	"context"
	"fmt"
	proto "github.com/Icorp/docker-lesson/calc_proto"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	proto.UnimplementedCalcServiceServer
}

func (s *server) Add(ctx context.Context, request *proto.CalcRequest) (*proto.CalcResponse, error) {
	return &proto.CalcResponse{Result: request.A + request.B}, nil
}

func (s *server) Subtract(ctx context.Context, request *proto.CalcRequest) (*proto.CalcResponse, error) {
	return &proto.CalcResponse{Result: request.A - request.B}, nil
}

func (s *server) Multiply(ctx context.Context, request *proto.CalcRequest) (*proto.CalcResponse, error) {
	return &proto.CalcResponse{Result: request.A * request.B}, nil
}

func (s *server) Divide(ctx context.Context, request *proto.CalcRequest) (*proto.CalcResponse, error) {
	return &proto.CalcResponse{Result: request.A / request.B}, nil
}

func (s *server) mustEmbedUnimplementedAddServiceServer() {
	//TODO implement me
	panic("implement me")
}

func main() {
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	fmt.Println("GRPC Server is running on port :4040")
	proto.RegisterCalcServiceServer(srv, &server{})
	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
