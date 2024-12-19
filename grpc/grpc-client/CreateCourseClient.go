package main

import (
	"context"
	"encoding/json"
	pb "github/student-management-gRpc/proto"
	"log"
	"time"
)

func callCreateCourseServer(client pb.StudentManagementServiceClient, course *pb.CourseList) {
	log.Println("Client Side Streaming Started")
	stream, err := client.CreateCourse(context.Background())
	if err != nil {
		log.Fatalf("Error while creating stream: %v", err)
	}
	for _, course := range course.Course {
		log.Println(course.Title)
		err := stream.Send(&pb.Course{
			Title: course.Title,
		})
		if err != nil {
			log.Fatalf("Error while sending course details: %v", err)
		}
		log.Printf("Sent Course: %v", course)
		time.Sleep(1 * time.Second)
	}
	log.Println("Client Side Streaming Finished")

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Printf("Error while receiving response: %v", err)
		return
	}
	response, err := json.MarshalIndent(map[string]string{"Server Response": res.Message}, "", "\t")
	if err != nil {
		log.Fatalf("Error Marshalling Response %v", err)
	}
	log.Println(string(response))
}
