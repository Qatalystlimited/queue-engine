package main

import (
	"context"
	"log"
	"net"
	"os"

	"github.com/Qatalystlimited/queue-engine/internal/db"
	qs "github.com/Qatalystlimited/queue-engine/internal/queue"
	pb "github.com/Qatalystlimited/queue-engine/proto"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedQueueServiceServer
	queue *qs.Service
}

func (s *server) JoinQueue(ctx context.Context, req *pb.JoinQueueRequest) (*pb.JoinQueueResponse, error) {
	ticketID, position, err := s.queue.JoinQueue(ctx, req.UserId, req.QueueId)
	if err != nil {
		return nil, err
	}
	return &pb.JoinQueueResponse{
		TicketId: ticketID,
		Position: position,
		Status:   "waiting",
	}, nil
}

func (s *server) GetPosition(ctx context.Context, req *pb.GetPositionRequest) (*pb.GetPositionResponse, error) {
	position, err := s.queue.GetPosition(ctx, req.UserId, req.QueueId)
	if err != nil {
		return nil, err
	}
	return &pb.GetPositionResponse{Position: position}, nil
}

func main() {
	godotenv.Load()

	database, err := db.Connect()
	if err != nil {
		log.Fatal("DB connection failed:", err)
	}
	defer database.Close()

	port := os.Getenv("PORT")
	if port == "" {
		port = os.Getenv("GRPC_PORT")
	}
	if port == "" {
		port = "50051"
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterQueueServiceServer(grpcServer, &server{
		queue: &qs.Service{DB: database},
	})

	log.Println("Server running on port", port)
	grpcServer.Serve(lis)
}
