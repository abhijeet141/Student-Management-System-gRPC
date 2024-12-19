package main

import (
	"context"
	"encoding/json"
	pb "github/student-management-gRpc/proto"
	"log"
	"time"
)

func callRegisterServer(client pb.StudentManagementServiceClient, user *pb.User) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.Register(ctx, user)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
	response, err := json.MarshalIndent(map[string]string{"Server Response": res.Message}, "", " ")
	if err != nil {
		log.Fatalf("Error Marshalling Response %v", err)
	}
	log.Println(string(response))
}
