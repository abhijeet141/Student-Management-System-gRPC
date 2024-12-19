package main

import (
	pb "github/student-management-gRpc/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {
	conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect %v", err)

	}
	defer conn.Close()
	client := pb.NewStudentManagementServiceClient(conn)

	user := &pb.User{
		UserName: "john@example.com",
		Password: "adminPass153",
		Role:     "admin",
	}
	callRegisterServer(client, user)

	// user := &pb.User{
	// 	UserName: "max.john@example.com",
	// 	Password: "adminPass123",
	// }

	callLoginServer(client, user)

	course := &pb.CourseList{
		Course: []*pb.Course{
			{
				Title: "Introduction to Go Programming",
			},
			{
				Title: "Web Development with React",
			},
			{
				Title: "Machine Learning and AI Basics",
			},
		},
	}
	callCreateCourseServer(client, course)

	student := &pb.Student{
		Name: "Aditya",
		Age:  26,
		Course: &pb.Course{
			Id: 100,
		},
		User: &pb.User{
			Id: 1,
		},
	}
	callCreateStudentServer(client, student)
	studentId := &pb.StudentId{
		Id: 6,
	}
	callGetStudentByIdServer(client, studentId)
	callDeleteStudentByIdServer(client, studentId)
}
