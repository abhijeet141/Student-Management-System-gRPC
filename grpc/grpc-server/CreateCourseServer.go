package main

import (
	"fmt"
	pb "github/student-management-gRpc/proto"
	"io"

	"github.com/beego/beego/v2/client/orm"
)

func (s *StudentManagementServer) CreateCourse(stream pb.StudentManagementService_CreateCourseServer) error {
	o := orm.NewOrm()
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.Message{Message: "Courses Created Successfully"})
		}
		if err != nil {
			return err
		}
		_, err = o.Insert(req)

		if err != nil {
			return fmt.Errorf("failed to insert course %w", err)
		}
	}
}
