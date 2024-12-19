package main

import (
	"context"
	"encoding/json"
	pb "github/student-management-gRpc/proto"
	"log"
	"time"
)

func callGetStudentByIdServer(client pb.StudentManagementServiceClient, studentId *pb.StudentId) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.GetStudentById(ctx, studentId)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
	response, err := json.MarshalIndent(map[string]map[string]interface{}{
		"Server Response": {
			"Id":       res.Id,
			"Name":     res.Name,
			"Age":      res.Age,
			"CourseId": res.Course,
		}}, "", "\t")
	if err != nil {
		log.Fatalf("Error Marshalling Response %v", err)
	}
	log.Println(string(response))
}
