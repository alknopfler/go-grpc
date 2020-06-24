package main

import (
  pb "gitlab.com/alknopfler/go-grpc/communications"
  "log"
  "golang.org/x/net/context"
  "google.golang.org/grpc"
  "errors"
  "net"
)

// Server is implementation proto interface
type Server struct{}

type ReqName struct{
  Name string
}

type Resp struct{
  Status string
}

func main() {
  grpcServer := grpc.NewServer()
  var server Server
  pb.RegisterCommunicatorServer(grpcServer, server)
  listen, err := net.Listen("tcp", "localhost:3000")
  if err != nil {
    log.Fatalf("could not listen to 0.0.0.0:3000 %v", err)
  }
  log.Println("Server starting...")
  log.Fatal(grpcServer.Serve(listen))
}

// Search function responsible to get the Country information
func (Server) TransferData(ctx context.Context, request *ReqName) (*Resp, error) {
  print(request)
  if request.Name == "alber" {
    log.Println("the name is: "+ request.Name)
    return &Resp{"OK"}, nil
  }
  return &Resp{""}, errors.New("error")
}