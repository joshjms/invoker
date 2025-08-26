package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/joshjms/invoker/api/workerpb"
	"google.golang.org/grpc"
)

type workerServer struct {
	workerpb.UnimplementedWorkerServiceServer
}

func (s *workerServer) Work(ctx context.Context, req *workerpb.WorkRequest) (*workerpb.WorkResponse, error) {
	startTime := time.Now()
	duration := time.Duration(req.GetDurationMs()) * time.Millisecond

	for {
		select {
		case <-time.After(time.Duration(duration) * time.Millisecond):
			return &workerpb.WorkResponse{
				StartAt: startTime.UnixMilli(),
				EndAt:   time.Now().UnixMilli(),
			}, nil
		default:
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen at port 50051: %v", err)
	}

	server := grpc.NewServer()
	workerpb.RegisterWorkerServiceServer(server, &workerServer{})

	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
