package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "api-go-learning/Microservices-gRPC/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("🚀 Starting gRPC Client...")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")
	fmt.Println("🔗 Connecting to: localhost:50051")
	fmt.Println("━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━")

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("❌ Failed to connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Call SayHello
	name := "World"
	fmt.Printf("\n📤 Sending request: SayHello(\"%s\")\n", name)
	
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("❌ Could not greet: %v", err)
	}
	
	fmt.Printf("📥 Response: %s\n", r.GetMessage())
	fmt.Println("\n✅ gRPC call successful!")
}
