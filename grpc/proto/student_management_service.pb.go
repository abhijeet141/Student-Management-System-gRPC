// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v5.29.1
// source: proto/student_management_service.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_proto_student_management_service_proto protoreflect.FileDescriptor

var file_proto_student_management_service_proto_rawDesc = []byte{
	0x0a, 0x26, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a,
	0x1c, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x5f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1a, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74,
	0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x6e, 0x6f, 0x5f, 0x70, 0x61, 0x72, 0x61, 0x6d, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x49, 0x64, 0x5f, 0x73, 0x74, 0x72, 0x75, 0x63, 0x74, 0x75,
	0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x22, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x72,
	0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1d, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x72, 0x75,
	0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x20, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x73, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x5f, 0x49, 0x64, 0x5f, 0x73, 0x74,
	0x72, 0x75, 0x63, 0x74, 0x75, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x32, 0x99, 0x03,
	0x0a, 0x18, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x6d,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x27, 0x0a, 0x08, 0x52, 0x65,
	0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55,
	0x73, 0x65, 0x72, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x12, 0x24, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0b, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x64, 0x12, 0x32, 0x0a, 0x0e, 0x47, 0x65, 0x74,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x42, 0x79, 0x49, 0x64, 0x12, 0x10, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x0e, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x2d, 0x0a,
	0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x12, 0x0d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x6f, 0x75, 0x72, 0x73, 0x65, 0x1a, 0x0e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x36, 0x0a, 0x0e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4e, 0x6f, 0x50, 0x61, 0x72, 0x61, 0x6d, 0x1a, 0x12,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x30, 0x01, 0x12, 0x2f, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53, 0x74,
	0x75, 0x64, 0x65, 0x6e, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x2f, 0x0a, 0x0d, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x53,
	0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x31, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x53, 0x74, 0x75, 0x64, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_proto_student_management_service_proto_goTypes = []any{
	(*User)(nil),        // 0: proto.User
	(*StudentId)(nil),   // 1: proto.StudentId
	(*Course)(nil),      // 2: proto.Course
	(*NoParam)(nil),     // 3: proto.NoParam
	(*Student)(nil),     // 4: proto.Student
	(*Message)(nil),     // 5: proto.Message
	(*TokenId)(nil),     // 6: proto.TokenId
	(*StudentList)(nil), // 7: proto.StudentList
}
var file_proto_student_management_service_proto_depIdxs = []int32{
	0, // 0: proto.StudentManagementService.Register:input_type -> proto.User
	0, // 1: proto.StudentManagementService.Login:input_type -> proto.User
	1, // 2: proto.StudentManagementService.GetStudentById:input_type -> proto.StudentId
	2, // 3: proto.StudentManagementService.CreateCourse:input_type -> proto.Course
	3, // 4: proto.StudentManagementService.GetAllStudents:input_type -> proto.NoParam
	4, // 5: proto.StudentManagementService.CreateStudent:input_type -> proto.Student
	4, // 6: proto.StudentManagementService.UpdateStudent:input_type -> proto.Student
	1, // 7: proto.StudentManagementService.DeleteStudent:input_type -> proto.StudentId
	5, // 8: proto.StudentManagementService.Register:output_type -> proto.Message
	6, // 9: proto.StudentManagementService.Login:output_type -> proto.TokenId
	4, // 10: proto.StudentManagementService.GetStudentById:output_type -> proto.Student
	5, // 11: proto.StudentManagementService.CreateCourse:output_type -> proto.Message
	7, // 12: proto.StudentManagementService.GetAllStudents:output_type -> proto.StudentList
	5, // 13: proto.StudentManagementService.CreateStudent:output_type -> proto.Message
	4, // 14: proto.StudentManagementService.UpdateStudent:output_type -> proto.Student
	5, // 15: proto.StudentManagementService.DeleteStudent:output_type -> proto.Message
	8, // [8:16] is the sub-list for method output_type
	0, // [0:8] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_student_management_service_proto_init() }
func file_proto_student_management_service_proto_init() {
	if File_proto_student_management_service_proto != nil {
		return
	}
	file_proto_course_structure_proto_init()
	file_proto_user_structure_proto_init()
	file_proto_message_structure_proto_init()
	file_proto_no_param_structure_proto_init()
	file_proto_token_Id_structure_proto_init()
	file_proto_student_list_structure_proto_init()
	file_proto_student_structure_proto_init()
	file_proto_student_Id_structure_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_student_management_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_student_management_service_proto_goTypes,
		DependencyIndexes: file_proto_student_management_service_proto_depIdxs,
	}.Build()
	File_proto_student_management_service_proto = out.File
	file_proto_student_management_service_proto_rawDesc = nil
	file_proto_student_management_service_proto_goTypes = nil
	file_proto_student_management_service_proto_depIdxs = nil
}
