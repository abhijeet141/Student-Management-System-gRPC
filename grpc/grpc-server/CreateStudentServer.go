package main

import (
	"context"
	"fmt"
	pb "github/student-management-gRpc/proto"

	"github.com/beego/beego/v2/client/orm"
)

func (s *StudentManagementServer) CreateStudent(ctx context.Context, req *pb.Student) (*pb.Message, error) {
	var student pb.Student
	o := orm.NewOrm()
	student.Name = req.Name
	student.Age = req.Age
	student.Course = req.Course
	student.User = req.User
	_, err := o.Insert(&student)
	if err != nil {
		return nil, fmt.Errorf("failed to insert student %w", err)
	}
	return &pb.Message{
		Message: "Student Created Successfully",
	}, nil
}
