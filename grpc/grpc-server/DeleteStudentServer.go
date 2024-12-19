package main

import (
	"context"
	"fmt"
	pb "github/student-management-gRpc/proto"

	"github.com/beego/beego/v2/client/orm"
)

func (s *StudentManagementServer) DeleteStudent(ctx context.Context, req *pb.StudentId) (*pb.Message, error) {
	var id int
	o := orm.NewOrm()
	id = int(req.Id)
	num, err := o.QueryTable("student").Filter("id", id).Delete()
	if err != nil {
		return nil, fmt.Errorf("failed to delete student %w", err)
	}
	if num == 0 {
		return nil, fmt.Errorf("student not found with id")
	}
	return &pb.Message{
		Message: "Student Deleted Successfully",
	}, nil
}
