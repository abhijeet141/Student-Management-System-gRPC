package main

import (
	"context"
	"fmt"
	pb "github/student-management-gRpc/proto"

	"github.com/beego/beego/v2/client/orm"
)

func (s *StudentManagementServer) GetStudentById(ctx context.Context, req *pb.StudentId) (*pb.Student, error) {
	var id int
	var student pb.Student
	o := orm.NewOrm()
	id = int(req.Id)
	err := o.QueryTable("student").Filter("id", id).RelatedSel("course").One(&student)
	if err != nil {
		return nil, fmt.Errorf("failed to insert user %w", err)
	}
	if err == orm.ErrNoRows {
		return nil, fmt.Errorf("no student found with given ID")
	}
	return &pb.Student{
		Id: student.Id, Name: student.Name, Age: student.Age, User: student.User, Course: student.Course,
	}, nil
}
