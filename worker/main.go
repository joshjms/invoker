package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/joshjms/invoker/api/workerpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

type workerServer struct {
	workerpb.UnimplementedWorkerServiceServer
}

func (s *workerServer) Work(ctx context.Context, req *workerpb.WorkRequest) (*workerpb.WorkResponse, error) {
	startTime := time.Now()
	duration := time.Duration(req.GetDurationMs()) * time.Millisecond

	timer := time.NewTimer(duration)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			return &workerpb.WorkResponse{
				StartAt: startTime.UnixMilli(),
				EndAt:   time.Now().UnixMilli(),
			}, nil
		default:
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("failed to listen at port 8080: %v", err)
	}

	server := grpc.NewServer()
	workerpb.RegisterWorkerServiceServer(server, &workerServer{})

	healthSrv := health.NewServer()
	healthpb.RegisterHealthServer(server, healthSrv)

	log.Println("worker running at port 8080")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
