package main

import (
	"context"
	"encoding/json"
	pb "github/student-management-gRpc/proto"
	"log"
	"time"
)

func callDeleteStudentByIdServer(client pb.StudentManagementServiceClient, studentId *pb.StudentId) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.DeleteStudent(ctx, studentId)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
	response, err := json.MarshalIndent(map[string]string{"Server Response": res.Message}, "", " ")
	if err != nil {
		log.Fatalf("Error Marshalling Response %v", err)
	}
	log.Println(string(response))
}
