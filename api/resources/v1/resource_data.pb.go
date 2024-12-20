// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        (unknown)
// source: resources/v1/resource_data.proto

package resources

import (
	_ "buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go/buf/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type K8SClusterResourceData_ClusterStatus int32

const (
	K8SClusterResourceData_CLUSTER_STATUS_UNSPECIFIED K8SClusterResourceData_ClusterStatus = 0
	K8SClusterResourceData_CLUSTER_STATUS_OTHER       K8SClusterResourceData_ClusterStatus = 1
	K8SClusterResourceData_READY                      K8SClusterResourceData_ClusterStatus = 2
	K8SClusterResourceData_FAILED                     K8SClusterResourceData_ClusterStatus = 3
	K8SClusterResourceData_OFFLINE                    K8SClusterResourceData_ClusterStatus = 4
)

// Enum value maps for K8SClusterResourceData_ClusterStatus.
var (
	K8SClusterResourceData_ClusterStatus_name = map[int32]string{
		0: "CLUSTER_STATUS_UNSPECIFIED",
		1: "CLUSTER_STATUS_OTHER",
		2: "READY",
		3: "FAILED",
		4: "OFFLINE",
	}
	K8SClusterResourceData_ClusterStatus_value = map[string]int32{
		"CLUSTER_STATUS_UNSPECIFIED": 0,
		"CLUSTER_STATUS_OTHER":       1,
		"READY":                      2,
		"FAILED":                     3,
		"OFFLINE":                    4,
	}
)

func (x K8SClusterResourceData_ClusterStatus) Enum() *K8SClusterResourceData_ClusterStatus {
	p := new(K8SClusterResourceData_ClusterStatus)
	*p = x
	return p
}

func (x K8SClusterResourceData_ClusterStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (K8SClusterResourceData_ClusterStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_v1_resource_data_proto_enumTypes[0].Descriptor()
}

func (K8SClusterResourceData_ClusterStatus) Type() protoreflect.EnumType {
	return &file_resources_v1_resource_data_proto_enumTypes[0]
}

func (x K8SClusterResourceData_ClusterStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use K8SClusterResourceData_ClusterStatus.Descriptor instead.
func (K8SClusterResourceData_ClusterStatus) EnumDescriptor() ([]byte, []int) {
	return file_resources_v1_resource_data_proto_rawDescGZIP(), []int{0, 0}
}

type K8SClusterResourceData_KubeVendor int32

const (
	K8SClusterResourceData_KUBE_VENDOR_UNSPECIFIED K8SClusterResourceData_KubeVendor = 0
	K8SClusterResourceData_KUBE_VENDOR_OTHER       K8SClusterResourceData_KubeVendor = 1
	K8SClusterResourceData_AKS                     K8SClusterResourceData_KubeVendor = 2
	K8SClusterResourceData_EKS                     K8SClusterResourceData_KubeVendor = 3
	K8SClusterResourceData_IKS                     K8SClusterResourceData_KubeVendor = 4
	K8SClusterResourceData_OPENSHIFT               K8SClusterResourceData_KubeVendor = 5
	K8SClusterResourceData_GKE                     K8SClusterResourceData_KubeVendor = 6
)

// Enum value maps for K8SClusterResourceData_KubeVendor.
var (
	K8SClusterResourceData_KubeVendor_name = map[int32]string{
		0: "KUBE_VENDOR_UNSPECIFIED",
		1: "KUBE_VENDOR_OTHER",
		2: "AKS",
		3: "EKS",
		4: "IKS",
		5: "OPENSHIFT",
		6: "GKE",
	}
	K8SClusterResourceData_KubeVendor_value = map[string]int32{
		"KUBE_VENDOR_UNSPECIFIED": 0,
		"KUBE_VENDOR_OTHER":       1,
		"AKS":                     2,
		"EKS":                     3,
		"IKS":                     4,
		"OPENSHIFT":               5,
		"GKE":                     6,
	}
)

func (x K8SClusterResourceData_KubeVendor) Enum() *K8SClusterResourceData_KubeVendor {
	p := new(K8SClusterResourceData_KubeVendor)
	*p = x
	return p
}

func (x K8SClusterResourceData_KubeVendor) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (K8SClusterResourceData_KubeVendor) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_v1_resource_data_proto_enumTypes[1].Descriptor()
}

func (K8SClusterResourceData_KubeVendor) Type() protoreflect.EnumType {
	return &file_resources_v1_resource_data_proto_enumTypes[1]
}

func (x K8SClusterResourceData_KubeVendor) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use K8SClusterResourceData_KubeVendor.Descriptor instead.
func (K8SClusterResourceData_KubeVendor) EnumDescriptor() ([]byte, []int) {
	return file_resources_v1_resource_data_proto_rawDescGZIP(), []int{0, 1}
}

type RhelHostResourceData_HostStatus int32

const (
	RhelHostResourceData_UNSPECIFIED RhelHostResourceData_HostStatus = 0
	RhelHostResourceData_DISABLED    RhelHostResourceData_HostStatus = 1
	RhelHostResourceData_ENABLED     RhelHostResourceData_HostStatus = 2
)

// Enum value maps for RhelHostResourceData_HostStatus.
var (
	RhelHostResourceData_HostStatus_name = map[int32]string{
		0: "UNSPECIFIED",
		1: "DISABLED",
		2: "ENABLED",
	}
	RhelHostResourceData_HostStatus_value = map[string]int32{
		"UNSPECIFIED": 0,
		"DISABLED":    1,
		"ENABLED":     2,
	}
)

func (x RhelHostResourceData_HostStatus) Enum() *RhelHostResourceData_HostStatus {
	p := new(RhelHostResourceData_HostStatus)
	*p = x
	return p
}

func (x RhelHostResourceData_HostStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (RhelHostResourceData_HostStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_resources_v1_resource_data_proto_enumTypes[2].Descriptor()
}

func (RhelHostResourceData_HostStatus) Type() protoreflect.EnumType {
	return &file_resources_v1_resource_data_proto_enumTypes[2]
}

func (x RhelHostResourceData_HostStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use RhelHostResourceData_HostStatus.Descriptor instead.
func (RhelHostResourceData_HostStatus) EnumDescriptor() ([]byte, []int) {
	return file_resources_v1_resource_data_proto_rawDescGZIP(), []int{1, 0}
}

type K8SClusterResourceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExternalClusterId string                               `protobuf:"bytes,219571597,opt,name=external_cluster_id,proto3" json:"external_cluster_id,omitempty"`
	ClusterStatus     K8SClusterResourceData_ClusterStatus `protobuf:"varint,499346904,opt,name=cluster_status,proto3,enum=resources.v1.K8SClusterResourceData_ClusterStatus" json:"cluster_status,omitempty"`
	KubeVendor        K8SClusterResourceData_KubeVendor    `protobuf:"varint,264191642,opt,name=kube_vendor,proto3,enum=resources.v1.K8SClusterResourceData_KubeVendor" json:"kube_vendor,omitempty"`
	KubeVersion       string                               `protobuf:"bytes,395858490,opt,name=kube_version,proto3" json:"kube_version,omitempty"`
}

func (x *K8SClusterResourceData) Reset() {
	*x = K8SClusterResourceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_v1_resource_data_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *K8SClusterResourceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*K8SClusterResourceData) ProtoMessage() {}

func (x *K8SClusterResourceData) ProtoReflect() protoreflect.Message {
	mi := &file_resources_v1_resource_data_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use K8SClusterResourceData.ProtoReflect.Descriptor instead.
func (*K8SClusterResourceData) Descriptor() ([]byte, []int) {
	return file_resources_v1_resource_data_proto_rawDescGZIP(), []int{0}
}

func (x *K8SClusterResourceData) GetExternalClusterId() string {
	if x != nil {
		return x.ExternalClusterId
	}
	return ""
}

func (x *K8SClusterResourceData) GetClusterStatus() K8SClusterResourceData_ClusterStatus {
	if x != nil {
		return x.ClusterStatus
	}
	return K8SClusterResourceData_CLUSTER_STATUS_UNSPECIFIED
}

func (x *K8SClusterResourceData) GetKubeVendor() K8SClusterResourceData_KubeVendor {
	if x != nil {
		return x.KubeVendor
	}
	return K8SClusterResourceData_KUBE_VENDOR_UNSPECIFIED
}

func (x *K8SClusterResourceData) GetKubeVersion() string {
	if x != nil {
		return x.KubeVersion
	}
	return ""
}

type RhelHostResourceData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HostId     string                          `protobuf:"bytes,219571597,opt,name=host_id,proto3" json:"host_id,omitempty"`
	HostStatus RhelHostResourceData_HostStatus `protobuf:"varint,75440785,opt,name=host_status,proto3,enum=resources.v1.RhelHostResourceData_HostStatus" json:"host_status,omitempty"`
	OsVersion  string                          `protobuf:"bytes,395858490,opt,name=os_version,proto3" json:"os_version,omitempty"`
}

func (x *RhelHostResourceData) Reset() {
	*x = RhelHostResourceData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_resources_v1_resource_data_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RhelHostResourceData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RhelHostResourceData) ProtoMessage() {}

func (x *RhelHostResourceData) ProtoReflect() protoreflect.Message {
	mi := &file_resources_v1_resource_data_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RhelHostResourceData.ProtoReflect.Descriptor instead.
func (*RhelHostResourceData) Descriptor() ([]byte, []int) {
	return file_resources_v1_resource_data_proto_rawDescGZIP(), []int{1}
}

func (x *RhelHostResourceData) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

func (x *RhelHostResourceData) GetHostStatus() RhelHostResourceData_HostStatus {
	if x != nil {
		return x.HostStatus
	}
	return RhelHostResourceData_UNSPECIFIED
}

func (x *RhelHostResourceData) GetOsVersion() string {
	if x != nil {
		return x.OsVersion
	}
	return ""
}

var File_resources_v1_resource_data_proto protoreflect.FileDescriptor

var file_resources_v1_resource_data_proto_rawDesc = []byte{
	0x0a, 0x20, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x62, 0x75, 0x66, 0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbd, 0x04,
	0x0a, 0x16, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x3b, 0x0a, 0x13, 0x65, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18,
	0x8d, 0xcb, 0xd9, 0x68, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01,
	0x52, 0x13, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x63, 0x6c, 0x75, 0x73, 0x74,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x6d, 0x0a, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72,
	0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0xd8, 0xdb, 0x8d, 0xee, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x32, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x2e, 0x4b, 0x38, 0x73, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x43, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x42, 0x0d, 0xba, 0x48, 0x0a, 0xc8, 0x01, 0x01, 0x82, 0x01, 0x04,
	0x10, 0x01, 0x20, 0x00, 0x52, 0x0e, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x63, 0x0a, 0x0b, 0x6b, 0x75, 0x62, 0x65, 0x5f, 0x76, 0x65, 0x6e,
	0x64, 0x6f, 0x72, 0x18, 0x9a, 0xfd, 0xfc, 0x7d, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2f, 0x2e, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x38, 0x73, 0x43,
	0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x61,
	0x74, 0x61, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x42, 0x0d, 0xba,
	0x48, 0x0a, 0xc8, 0x01, 0x01, 0x82, 0x01, 0x04, 0x10, 0x01, 0x20, 0x00, 0x52, 0x0b, 0x6b, 0x75,
	0x62, 0x65, 0x5f, 0x76, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x2e, 0x0a, 0x0c, 0x6b, 0x75, 0x62,
	0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0xba, 0xa4, 0xe1, 0xbc, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0c, 0x6b, 0x75, 0x62,
	0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x6d, 0x0a, 0x0d, 0x43, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1e, 0x0a, 0x1a, 0x43, 0x4c,
	0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53,
	0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14, 0x43, 0x4c,
	0x55, 0x53, 0x54, 0x45, 0x52, 0x5f, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x4f, 0x54, 0x48,
	0x45, 0x52, 0x10, 0x01, 0x12, 0x09, 0x0a, 0x05, 0x52, 0x45, 0x41, 0x44, 0x59, 0x10, 0x02, 0x12,
	0x0a, 0x0a, 0x06, 0x46, 0x41, 0x49, 0x4c, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x4f,
	0x46, 0x46, 0x4c, 0x49, 0x4e, 0x45, 0x10, 0x04, 0x22, 0x73, 0x0a, 0x0a, 0x4b, 0x75, 0x62, 0x65,
	0x56, 0x65, 0x6e, 0x64, 0x6f, 0x72, 0x12, 0x1b, 0x0a, 0x17, 0x4b, 0x55, 0x42, 0x45, 0x5f, 0x56,
	0x45, 0x4e, 0x44, 0x4f, 0x52, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45,
	0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x4b, 0x55, 0x42, 0x45, 0x5f, 0x56, 0x45, 0x4e, 0x44,
	0x4f, 0x52, 0x5f, 0x4f, 0x54, 0x48, 0x45, 0x52, 0x10, 0x01, 0x12, 0x07, 0x0a, 0x03, 0x41, 0x4b,
	0x53, 0x10, 0x02, 0x12, 0x07, 0x0a, 0x03, 0x45, 0x4b, 0x53, 0x10, 0x03, 0x12, 0x07, 0x0a, 0x03,
	0x49, 0x4b, 0x53, 0x10, 0x04, 0x12, 0x0d, 0x0a, 0x09, 0x4f, 0x50, 0x45, 0x4e, 0x53, 0x48, 0x49,
	0x46, 0x54, 0x10, 0x05, 0x12, 0x07, 0x0a, 0x03, 0x47, 0x4b, 0x45, 0x10, 0x06, 0x22, 0x84, 0x02,
	0x0a, 0x14, 0x52, 0x68, 0x65, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x23, 0x0a, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69,
	0x64, 0x18, 0x8d, 0xcb, 0xd9, 0x68, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8,
	0x01, 0x01, 0x52, 0x07, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x69, 0x64, 0x12, 0x61, 0x0a, 0x0b, 0x68,
	0x6f, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x91, 0xc5, 0xfc, 0x23, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x52, 0x68, 0x65, 0x6c, 0x48, 0x6f, 0x73, 0x74, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x48, 0x6f, 0x73, 0x74, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x42, 0x0d, 0xba, 0x48, 0x0a, 0xc8, 0x01, 0x01, 0x82, 0x01, 0x04, 0x10, 0x01, 0x20,
	0x00, 0x52, 0x0b, 0x68, 0x6f, 0x73, 0x74, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2a,
	0x0a, 0x0a, 0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0xba, 0xa4, 0xe1,
	0xbc, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x06, 0xba, 0x48, 0x03, 0xc8, 0x01, 0x01, 0x52, 0x0a,
	0x6f, 0x73, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x38, 0x0a, 0x0a, 0x48, 0x6f,
	0x73, 0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0f, 0x0a, 0x0b, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x44, 0x49, 0x53,
	0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x45, 0x4e, 0x41, 0x42, 0x4c,
	0x45, 0x44, 0x10, 0x02, 0x42, 0x4d, 0x0a, 0x0d, 0x61, 0x70, 0x69, 0x2e, 0x72, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x73, 0x50, 0x01, 0x5a, 0x3a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x6f, 0x6e, 0x79, 0x74, 0x68, 0x65, 0x6c, 0x65, 0x67, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x3b, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_resources_v1_resource_data_proto_rawDescOnce sync.Once
	file_resources_v1_resource_data_proto_rawDescData = file_resources_v1_resource_data_proto_rawDesc
)

func file_resources_v1_resource_data_proto_rawDescGZIP() []byte {
	file_resources_v1_resource_data_proto_rawDescOnce.Do(func() {
		file_resources_v1_resource_data_proto_rawDescData = protoimpl.X.CompressGZIP(file_resources_v1_resource_data_proto_rawDescData)
	})
	return file_resources_v1_resource_data_proto_rawDescData
}

var file_resources_v1_resource_data_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_resources_v1_resource_data_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_resources_v1_resource_data_proto_goTypes = []any{
	(K8SClusterResourceData_ClusterStatus)(0), // 0: resources.v1.K8sClusterResourceData.ClusterStatus
	(K8SClusterResourceData_KubeVendor)(0),    // 1: resources.v1.K8sClusterResourceData.KubeVendor
	(RhelHostResourceData_HostStatus)(0),      // 2: resources.v1.RhelHostResourceData.HostStatus
	(*K8SClusterResourceData)(nil),            // 3: resources.v1.K8sClusterResourceData
	(*RhelHostResourceData)(nil),              // 4: resources.v1.RhelHostResourceData
}
var file_resources_v1_resource_data_proto_depIdxs = []int32{
	0, // 0: resources.v1.K8sClusterResourceData.cluster_status:type_name -> resources.v1.K8sClusterResourceData.ClusterStatus
	1, // 1: resources.v1.K8sClusterResourceData.kube_vendor:type_name -> resources.v1.K8sClusterResourceData.KubeVendor
	2, // 2: resources.v1.RhelHostResourceData.host_status:type_name -> resources.v1.RhelHostResourceData.HostStatus
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_resources_v1_resource_data_proto_init() }
func file_resources_v1_resource_data_proto_init() {
	if File_resources_v1_resource_data_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_resources_v1_resource_data_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*K8SClusterResourceData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_resources_v1_resource_data_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*RhelHostResourceData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_resources_v1_resource_data_proto_rawDesc,
			NumEnums:      3,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_resources_v1_resource_data_proto_goTypes,
		DependencyIndexes: file_resources_v1_resource_data_proto_depIdxs,
		EnumInfos:         file_resources_v1_resource_data_proto_enumTypes,
		MessageInfos:      file_resources_v1_resource_data_proto_msgTypes,
	}.Build()
	File_resources_v1_resource_data_proto = out.File
	file_resources_v1_resource_data_proto_rawDesc = nil
	file_resources_v1_resource_data_proto_goTypes = nil
	file_resources_v1_resource_data_proto_depIdxs = nil
}
