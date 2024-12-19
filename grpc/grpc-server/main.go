package main

import (
	pb "github/student-management-gRpc/proto"
	"log"
	"net"
	"os"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

const (
	port = ":8080"
)

type StudentManagementServer struct {
	pb.StudentManagementServiceServer
}

func init() {
	orm.RegisterModel(new(pb.Course), new(pb.User), new(pb.Student))
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	connectionstring := os.Getenv("CONNECTION_STRING")
	if connectionstring == "" {
		log.Fatal("Connection String is not set in the environment")
	}
	orm.RegisterDataBase("default", "mysql", connectionstring)
	err = orm.RunSyncdb("default", false, true)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	log.Println("Database connection is established")
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to start the server %v", err)
	}
	log.Println("APP started and running on PORT 8080")
	grpcserver := grpc.NewServer()
	pb.RegisterStudentManagementServiceServer(grpcserver, &StudentManagementServer{})
	err = grpcserver.Serve(lis)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
