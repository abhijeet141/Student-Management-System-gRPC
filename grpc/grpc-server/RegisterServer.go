package main

import (
	"context"
	"fmt"
	pb "github/student-management-gRpc/proto"

	"github.com/beego/beego/v2/client/orm"
)

func (s *StudentManagementServer) Register(ctx context.Context, req *pb.User) (*pb.Message, error) {
	var user pb.User
	user.UserName = req.UserName
	user.Password = req.Password
	user.Role = req.Role
	o := orm.NewOrm()
	_, err := o.Insert(&user)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user %w", err)
	}
	return &pb.Message{
		Message: "User Registered Successfully",
	}, nil
}
