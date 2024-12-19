package main

import (
	"context"
	"fmt"
	pb "github/student-management-gRpc/proto"

	"github.com/beego/beego/v2/client/orm"
)

func (s *StudentManagementServer) Login(ctx context.Context, req *pb.User) (*pb.TokenId, error) {
	var user pb.User
	user.UserName = req.UserName
	user.Password = req.Password
	o := orm.NewOrm()
	userInfo := pb.User{UserName: user.UserName}
	err := o.Read(&userInfo, "UserName")
	if err != nil {
		return nil, fmt.Errorf("user not found %w", err)
	}
	if user.Password != userInfo.Password {
		return nil, fmt.Errorf("invalid Credentials")
	}
	token, err := GenerateJWT(userInfo.UserName)
	if err != nil {
		return nil, fmt.Errorf("error generating error %v", err)
	}
	return &pb.TokenId{
		Token: token,
	}, nil
}
