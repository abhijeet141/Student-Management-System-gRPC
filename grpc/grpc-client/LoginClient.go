package main

import (
	"context"
	"encoding/json"
	pb "github/student-management-gRpc/proto"
	"log"
	"time"
)

func callLoginServer(client pb.StudentManagementServiceClient, user *pb.User) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := client.Login(ctx, user)
	if err != nil {
		log.Fatalf("Server error: %v", err)
	}
	response, err := json.MarshalIndent(map[string]map[string]string{
		"Server Response": {
			"TokenId": "Bearer " + res.Token,
		}}, "", "\t")
	if err != nil {
		log.Fatalf("Error Marshalling Response %v", err)
	}
	log.Println(string(response))
}
