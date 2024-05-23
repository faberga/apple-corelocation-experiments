// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v3.19.6
// source: pb/BSSIDApple.proto

package pb

import (
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

type WifiDevice struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bssid    string               `protobuf:"bytes,1,opt,name=bssid,proto3" json:"bssid,omitempty"`
	Location *WifiDevice_Location `protobuf:"bytes,2,opt,name=location,proto3,oneof" json:"location,omitempty"`
}

func (x *WifiDevice) Reset() {
	*x = WifiDevice{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_BSSIDApple_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WifiDevice) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WifiDevice) ProtoMessage() {}

func (x *WifiDevice) ProtoReflect() protoreflect.Message {
	mi := &file_pb_BSSIDApple_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WifiDevice.ProtoReflect.Descriptor instead.
func (*WifiDevice) Descriptor() ([]byte, []int) {
	return file_pb_BSSIDApple_proto_rawDescGZIP(), []int{0}
}

func (x *WifiDevice) GetBssid() string {
	if x != nil {
		return x.Bssid
	}
	return ""
}

func (x *WifiDevice) GetLocation() *WifiDevice_Location {
	if x != nil {
		return x.Location
	}
	return nil
}

type AppleWLoc struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnknownValue0      *int64        `protobuf:"varint,1,opt,name=unknown_value0,json=unknownValue0,proto3,oneof" json:"unknown_value0,omitempty"`
	WifiDevices        []*WifiDevice `protobuf:"bytes,2,rep,name=wifi_devices,json=wifiDevices,proto3" json:"wifi_devices,omitempty"`
	UnknownValue1      *int32        `protobuf:"varint,3,opt,name=unknown_value1,json=unknownValue1,proto3,oneof" json:"unknown_value1,omitempty"`
	ReturnSingleResult *int32        `protobuf:"varint,4,opt,name=return_single_result,json=returnSingleResult,proto3,oneof" json:"return_single_result,omitempty"`
	APIName            *string       `protobuf:"bytes,5,opt,name=APIName,proto3,oneof" json:"APIName,omitempty"`
	UnknownValue2      *string       `protobuf:"bytes,6,opt,name=unknown_value2,json=unknownValue2,proto3,oneof" json:"unknown_value2,omitempty"`
}

func (x *AppleWLoc) Reset() {
	*x = AppleWLoc{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_BSSIDApple_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AppleWLoc) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AppleWLoc) ProtoMessage() {}

func (x *AppleWLoc) ProtoReflect() protoreflect.Message {
	mi := &file_pb_BSSIDApple_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AppleWLoc.ProtoReflect.Descriptor instead.
func (*AppleWLoc) Descriptor() ([]byte, []int) {
	return file_pb_BSSIDApple_proto_rawDescGZIP(), []int{1}
}

func (x *AppleWLoc) GetUnknownValue0() int64 {
	if x != nil && x.UnknownValue0 != nil {
		return *x.UnknownValue0
	}
	return 0
}

func (x *AppleWLoc) GetWifiDevices() []*WifiDevice {
	if x != nil {
		return x.WifiDevices
	}
	return nil
}

func (x *AppleWLoc) GetUnknownValue1() int32 {
	if x != nil && x.UnknownValue1 != nil {
		return *x.UnknownValue1
	}
	return 0
}

func (x *AppleWLoc) GetReturnSingleResult() int32 {
	if x != nil && x.ReturnSingleResult != nil {
		return *x.ReturnSingleResult
	}
	return 0
}

func (x *AppleWLoc) GetAPIName() string {
	if x != nil && x.APIName != nil {
		return *x.APIName
	}
	return ""
}

func (x *AppleWLoc) GetUnknownValue2() string {
	if x != nil && x.UnknownValue2 != nil {
		return *x.UnknownValue2
	}
	return ""
}

type WifiDevice_Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude           *int64 `protobuf:"varint,1,opt,name=latitude,proto3,oneof" json:"latitude,omitempty"`
	Longitude          *int64 `protobuf:"varint,2,opt,name=longitude,proto3,oneof" json:"longitude,omitempty"`
	HorizontalAccuracy *int64 `protobuf:"varint,3,opt,name=horizontal_accuracy,json=horizontalAccuracy,proto3,oneof" json:"horizontal_accuracy,omitempty"`
	UnknownValue4      *int64 `protobuf:"varint,4,opt,name=unknown_value4,json=unknownValue4,proto3,oneof" json:"unknown_value4,omitempty"`
	Altitude           *int64 `protobuf:"varint,5,opt,name=altitude,proto3,oneof" json:"altitude,omitempty"`
	VerticalAccuracy   *int64 `protobuf:"varint,6,opt,name=vertical_accuracy,json=verticalAccuracy,proto3,oneof" json:"vertical_accuracy,omitempty"`
	// optional int64 speed = 7;
	// optional int64 course = 8;
	// optional int64 timestamp = 9;
	// optional int64 unknown_context = 10;
	MotionActivityType       *int64 `protobuf:"varint,11,opt,name=motion_activity_type,json=motionActivityType,proto3,oneof" json:"motion_activity_type,omitempty"`
	MotionActivityConfidence *int64 `protobuf:"varint,12,opt,name=motion_activity_confidence,json=motionActivityConfidence,proto3,oneof" json:"motion_activity_confidence,omitempty"`
}

func (x *WifiDevice_Location) Reset() {
	*x = WifiDevice_Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pb_BSSIDApple_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WifiDevice_Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WifiDevice_Location) ProtoMessage() {}

func (x *WifiDevice_Location) ProtoReflect() protoreflect.Message {
	mi := &file_pb_BSSIDApple_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WifiDevice_Location.ProtoReflect.Descriptor instead.
func (*WifiDevice_Location) Descriptor() ([]byte, []int) {
	return file_pb_BSSIDApple_proto_rawDescGZIP(), []int{0, 0}
}

func (x *WifiDevice_Location) GetLatitude() int64 {
	if x != nil && x.Latitude != nil {
		return *x.Latitude
	}
	return 0
}

func (x *WifiDevice_Location) GetLongitude() int64 {
	if x != nil && x.Longitude != nil {
		return *x.Longitude
	}
	return 0
}

func (x *WifiDevice_Location) GetHorizontalAccuracy() int64 {
	if x != nil && x.HorizontalAccuracy != nil {
		return *x.HorizontalAccuracy
	}
	return 0
}

func (x *WifiDevice_Location) GetUnknownValue4() int64 {
	if x != nil && x.UnknownValue4 != nil {
		return *x.UnknownValue4
	}
	return 0
}

func (x *WifiDevice_Location) GetAltitude() int64 {
	if x != nil && x.Altitude != nil {
		return *x.Altitude
	}
	return 0
}

func (x *WifiDevice_Location) GetVerticalAccuracy() int64 {
	if x != nil && x.VerticalAccuracy != nil {
		return *x.VerticalAccuracy
	}
	return 0
}

func (x *WifiDevice_Location) GetMotionActivityType() int64 {
	if x != nil && x.MotionActivityType != nil {
		return *x.MotionActivityType
	}
	return 0
}

func (x *WifiDevice_Location) GetMotionActivityConfidence() int64 {
	if x != nil && x.MotionActivityConfidence != nil {
		return *x.MotionActivityConfidence
	}
	return 0
}

var File_pb_BSSIDApple_proto protoreflect.FileDescriptor

var file_pb_BSSIDApple_proto_rawDesc = []byte{
	0x0a, 0x13, 0x70, 0x62, 0x2f, 0x42, 0x53, 0x53, 0x49, 0x44, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x05, 0x0a, 0x0a, 0x57, 0x69, 0x66, 0x69, 0x44, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x73, 0x73, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x73, 0x73, 0x69, 0x64, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x57,
	0x69, 0x66, 0x69, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x88, 0x01,
	0x01, 0x1a, 0x9e, 0x04, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1f,
	0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x00, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x88, 0x01, 0x01, 0x12,
	0x21, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x34, 0x0a, 0x13, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c,
	0x5f, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x48,
	0x02, 0x52, 0x12, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x41, 0x63, 0x63,
	0x75, 0x72, 0x61, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x75, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x34, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x03, 0x52, 0x0d, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x56, 0x61, 0x6c, 0x75, 0x65,
	0x34, 0x88, 0x01, 0x01, 0x12, 0x1f, 0x0a, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x48, 0x04, 0x52, 0x08, 0x61, 0x6c, 0x74, 0x69, 0x74, 0x75,
	0x64, 0x65, 0x88, 0x01, 0x01, 0x12, 0x30, 0x0a, 0x11, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61,
	0x6c, 0x5f, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03,
	0x48, 0x05, 0x52, 0x10, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61, 0x6c, 0x41, 0x63, 0x63, 0x75,
	0x72, 0x61, 0x63, 0x79, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x14, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x03, 0x48, 0x06, 0x52, 0x12, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x41,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x54, 0x79, 0x70, 0x65, 0x88, 0x01, 0x01, 0x12, 0x41,
	0x0a, 0x1a, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74,
	0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x0c, 0x20, 0x01,
	0x28, 0x03, 0x48, 0x07, 0x52, 0x18, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x41, 0x63, 0x74, 0x69,
	0x76, 0x69, 0x74, 0x79, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e, 0x63, 0x65, 0x88, 0x01,
	0x01, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65, 0x42, 0x0c,
	0x0a, 0x0a, 0x5f, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x42, 0x16, 0x0a, 0x14,
	0x5f, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x6f, 0x6e, 0x74, 0x61, 0x6c, 0x5f, 0x61, 0x63, 0x63, 0x75,
	0x72, 0x61, 0x63, 0x79, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x34, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x61, 0x6c, 0x74, 0x69,
	0x74, 0x75, 0x64, 0x65, 0x42, 0x14, 0x0a, 0x12, 0x5f, 0x76, 0x65, 0x72, 0x74, 0x69, 0x63, 0x61,
	0x6c, 0x5f, 0x61, 0x63, 0x63, 0x75, 0x72, 0x61, 0x63, 0x79, 0x42, 0x17, 0x0a, 0x15, 0x5f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x74,
	0x79, 0x70, 0x65, 0x42, 0x1d, 0x0a, 0x1b, 0x5f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61,
	0x63, 0x74, 0x69, 0x76, 0x69, 0x74, 0x79, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x64, 0x65, 0x6e,
	0x63, 0x65, 0x42, 0x0b, 0x0a, 0x09, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22,
	0xf3, 0x02, 0x0a, 0x09, 0x41, 0x70, 0x70, 0x6c, 0x65, 0x57, 0x4c, 0x6f, 0x63, 0x12, 0x2a, 0x0a,
	0x0e, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x30, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x48, 0x00, 0x52, 0x0d, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x30, 0x88, 0x01, 0x01, 0x12, 0x2e, 0x0a, 0x0c, 0x77, 0x69, 0x66,
	0x69, 0x5f, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x0b, 0x2e, 0x57, 0x69, 0x66, 0x69, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x52, 0x0b, 0x77, 0x69,
	0x66, 0x69, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x73, 0x12, 0x2a, 0x0a, 0x0e, 0x75, 0x6e, 0x6b,
	0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x48, 0x01, 0x52, 0x0d, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x56, 0x61, 0x6c, 0x75,
	0x65, 0x31, 0x88, 0x01, 0x01, 0x12, 0x35, 0x0a, 0x14, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f,
	0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x05, 0x48, 0x02, 0x52, 0x12, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x53, 0x69, 0x6e,
	0x67, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x88, 0x01, 0x01, 0x12, 0x1d, 0x0a, 0x07,
	0x41, 0x50, 0x49, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x48, 0x03, 0x52,
	0x07, 0x41, 0x50, 0x49, 0x4e, 0x61, 0x6d, 0x65, 0x88, 0x01, 0x01, 0x12, 0x2a, 0x0a, 0x0e, 0x75,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x48, 0x04, 0x52, 0x0d, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x56, 0x61,
	0x6c, 0x75, 0x65, 0x32, 0x88, 0x01, 0x01, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x75, 0x6e, 0x6b, 0x6e,
	0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x30, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x75,
	0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x31, 0x42, 0x17, 0x0a,
	0x15, 0x5f, 0x72, 0x65, 0x74, 0x75, 0x72, 0x6e, 0x5f, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x5f,
	0x72, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x42, 0x0a, 0x0a, 0x08, 0x5f, 0x41, 0x50, 0x49, 0x4e, 0x61,
	0x6d, 0x65, 0x42, 0x11, 0x0a, 0x0f, 0x5f, 0x75, 0x6e, 0x6b, 0x6e, 0x6f, 0x77, 0x6e, 0x5f, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x32, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pb_BSSIDApple_proto_rawDescOnce sync.Once
	file_pb_BSSIDApple_proto_rawDescData = file_pb_BSSIDApple_proto_rawDesc
)

func file_pb_BSSIDApple_proto_rawDescGZIP() []byte {
	file_pb_BSSIDApple_proto_rawDescOnce.Do(func() {
		file_pb_BSSIDApple_proto_rawDescData = protoimpl.X.CompressGZIP(file_pb_BSSIDApple_proto_rawDescData)
	})
	return file_pb_BSSIDApple_proto_rawDescData
}

var file_pb_BSSIDApple_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_pb_BSSIDApple_proto_goTypes = []interface{}{
	(*WifiDevice)(nil),          // 0: WifiDevice
	(*AppleWLoc)(nil),           // 1: AppleWLoc
	(*WifiDevice_Location)(nil), // 2: WifiDevice.Location
}
var file_pb_BSSIDApple_proto_depIdxs = []int32{
	2, // 0: WifiDevice.location:type_name -> WifiDevice.Location
	0, // 1: AppleWLoc.wifi_devices:type_name -> WifiDevice
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_pb_BSSIDApple_proto_init() }
func file_pb_BSSIDApple_proto_init() {
	if File_pb_BSSIDApple_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pb_BSSIDApple_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WifiDevice); i {
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
		file_pb_BSSIDApple_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AppleWLoc); i {
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
		file_pb_BSSIDApple_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WifiDevice_Location); i {
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
	file_pb_BSSIDApple_proto_msgTypes[0].OneofWrappers = []interface{}{}
	file_pb_BSSIDApple_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_pb_BSSIDApple_proto_msgTypes[2].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_BSSIDApple_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_BSSIDApple_proto_goTypes,
		DependencyIndexes: file_pb_BSSIDApple_proto_depIdxs,
		MessageInfos:      file_pb_BSSIDApple_proto_msgTypes,
	}.Build()
	File_pb_BSSIDApple_proto = out.File
	file_pb_BSSIDApple_proto_rawDesc = nil
	file_pb_BSSIDApple_proto_goTypes = nil
	file_pb_BSSIDApple_proto_depIdxs = nil
}
