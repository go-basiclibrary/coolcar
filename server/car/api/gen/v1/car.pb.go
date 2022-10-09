// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.6
// source: car.proto

package carpb

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

type CarStatus int32

const (
	CarStatus_CS_NOT_SPECIFIED CarStatus = 0
	CarStatus_LOCKED           CarStatus = 1 //上锁
	CarStatus_UNLOCKING        CarStatus = 2 //开锁中
	CarStatus_UNLOCKED         CarStatus = 3 //已开锁
	CarStatus_LOCKING          CarStatus = 4
)

// Enum value maps for CarStatus.
var (
	CarStatus_name = map[int32]string{
		0: "CS_NOT_SPECIFIED",
		1: "LOCKED",
		2: "UNLOCKING",
		3: "UNLOCKED",
		4: "LOCKING",
	}
	CarStatus_value = map[string]int32{
		"CS_NOT_SPECIFIED": 0,
		"LOCKED":           1,
		"UNLOCKING":        2,
		"UNLOCKED":         3,
		"LOCKING":          4,
	}
)

func (x CarStatus) Enum() *CarStatus {
	p := new(CarStatus)
	*p = x
	return p
}

func (x CarStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (CarStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_car_proto_enumTypes[0].Descriptor()
}

func (CarStatus) Type() protoreflect.EnumType {
	return &file_car_proto_enumTypes[0]
}

func (x CarStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use CarStatus.Descriptor instead.
func (CarStatus) EnumDescriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{0}
}

type CarEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Car *Car   `protobuf:"bytes,2,opt,name=car,proto3" json:"car,omitempty"`
}

func (x *CarEntity) Reset() {
	*x = CarEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CarEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CarEntity) ProtoMessage() {}

func (x *CarEntity) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CarEntity.ProtoReflect.Descriptor instead.
func (*CarEntity) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{0}
}

func (x *CarEntity) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CarEntity) GetCar() *Car {
	if x != nil {
		return x.Car
	}
	return nil
}

type Driver struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	AvatarUrl string `protobuf:"bytes,2,opt,name=avatar_url,json=avatarUrl,proto3" json:"avatar_url,omitempty"`
}

func (x *Driver) Reset() {
	*x = Driver{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Driver) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Driver) ProtoMessage() {}

func (x *Driver) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Driver.ProtoReflect.Descriptor instead.
func (*Driver) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{1}
}

func (x *Driver) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Driver) GetAvatarUrl() string {
	if x != nil {
		return x.AvatarUrl
	}
	return ""
}

type Location struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Latitude  float64 `protobuf:"fixed64,1,opt,name=latitude,proto3" json:"latitude,omitempty"`
	Longitude float64 `protobuf:"fixed64,2,opt,name=longitude,proto3" json:"longitude,omitempty"`
}

func (x *Location) Reset() {
	*x = Location{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Location) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Location) ProtoMessage() {}

func (x *Location) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Location.ProtoReflect.Descriptor instead.
func (*Location) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{2}
}

func (x *Location) GetLatitude() float64 {
	if x != nil {
		return x.Latitude
	}
	return 0
}

func (x *Location) GetLongitude() float64 {
	if x != nil {
		return x.Longitude
	}
	return 0
}

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status   CarStatus `protobuf:"varint,1,opt,name=status,proto3,enum=car.v1.CarStatus" json:"status,omitempty"`
	Driver   *Driver   `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver,omitempty"`
	Position *Location `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
	TripId   string    `protobuf:"bytes,4,opt,name=trip_id,json=tripId,proto3" json:"trip_id,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{3}
}

func (x *Car) GetStatus() CarStatus {
	if x != nil {
		return x.Status
	}
	return CarStatus_CS_NOT_SPECIFIED
}

func (x *Car) GetDriver() *Driver {
	if x != nil {
		return x.Driver
	}
	return nil
}

func (x *Car) GetPosition() *Location {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Car) GetTripId() string {
	if x != nil {
		return x.TripId
	}
	return ""
}

type CreateCarReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *CreateCarReq) Reset() {
	*x = CreateCarReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCarReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCarReq) ProtoMessage() {}

func (x *CreateCarReq) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCarReq.ProtoReflect.Descriptor instead.
func (*CreateCarReq) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{4}
}

type GetCarReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCarReq) Reset() {
	*x = GetCarReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarReq) ProtoMessage() {}

func (x *GetCarReq) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarReq.ProtoReflect.Descriptor instead.
func (*GetCarReq) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{5}
}

func (x *GetCarReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCarsReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCarsReq) Reset() {
	*x = GetCarsReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarsReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarsReq) ProtoMessage() {}

func (x *GetCarsReq) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarsReq.ProtoReflect.Descriptor instead.
func (*GetCarsReq) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{6}
}

type GetCarsRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cars []*CarEntity `protobuf:"bytes,1,rep,name=cars,proto3" json:"cars,omitempty"`
}

func (x *GetCarsRsp) Reset() {
	*x = GetCarsRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarsRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarsRsp) ProtoMessage() {}

func (x *GetCarsRsp) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarsRsp.ProtoReflect.Descriptor instead.
func (*GetCarsRsp) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{7}
}

func (x *GetCarsRsp) GetCars() []*CarEntity {
	if x != nil {
		return x.Cars
	}
	return nil
}

type LockCarReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *LockCarReq) Reset() {
	*x = LockCarReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockCarReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockCarReq) ProtoMessage() {}

func (x *LockCarReq) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockCarReq.ProtoReflect.Descriptor instead.
func (*LockCarReq) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{8}
}

func (x *LockCarReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type LockCarRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *LockCarRsp) Reset() {
	*x = LockCarRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LockCarRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LockCarRsp) ProtoMessage() {}

func (x *LockCarRsp) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LockCarRsp.ProtoReflect.Descriptor instead.
func (*LockCarRsp) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{9}
}

type UnlockCarReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string  `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Driver *Driver `protobuf:"bytes,2,opt,name=driver,proto3" json:"driver,omitempty"`
	TripId string  `protobuf:"bytes,3,opt,name=trip_id,json=tripId,proto3" json:"trip_id,omitempty"`
}

func (x *UnlockCarReq) Reset() {
	*x = UnlockCarReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockCarReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockCarReq) ProtoMessage() {}

func (x *UnlockCarReq) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockCarReq.ProtoReflect.Descriptor instead.
func (*UnlockCarReq) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{10}
}

func (x *UnlockCarReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UnlockCarReq) GetDriver() *Driver {
	if x != nil {
		return x.Driver
	}
	return nil
}

func (x *UnlockCarReq) GetTripId() string {
	if x != nil {
		return x.TripId
	}
	return ""
}

type UnlockCarRsp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UnlockCarRsp) Reset() {
	*x = UnlockCarRsp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnlockCarRsp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnlockCarRsp) ProtoMessage() {}

func (x *UnlockCarRsp) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnlockCarRsp.ProtoReflect.Descriptor instead.
func (*UnlockCarRsp) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{11}
}

type UpdateCarReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status   CarStatus `protobuf:"varint,2,opt,name=status,proto3,enum=car.v1.CarStatus" json:"status,omitempty"`
	Position *Location `protobuf:"bytes,3,opt,name=position,proto3" json:"position,omitempty"`
}

func (x *UpdateCarReq) Reset() {
	*x = UpdateCarReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCarReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCarReq) ProtoMessage() {}

func (x *UpdateCarReq) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCarReq.ProtoReflect.Descriptor instead.
func (*UpdateCarReq) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{12}
}

func (x *UpdateCarReq) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UpdateCarReq) GetStatus() CarStatus {
	if x != nil {
		return x.Status
	}
	return CarStatus_CS_NOT_SPECIFIED
}

func (x *UpdateCarReq) GetPosition() *Location {
	if x != nil {
		return x.Position
	}
	return nil
}

type UpdateCarResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *UpdateCarResponse) Reset() {
	*x = UpdateCarResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCarResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCarResponse) ProtoMessage() {}

func (x *UpdateCarResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCarResponse.ProtoReflect.Descriptor instead.
func (*UpdateCarResponse) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{13}
}

var File_car_proto protoreflect.FileDescriptor

var file_car_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x61, 0x72,
	0x2e, 0x76, 0x31, 0x22, 0x3a, 0x0a, 0x09, 0x43, 0x61, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x1d, 0x0a, 0x03, 0x63, 0x61, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e,
	0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x03, 0x63, 0x61, 0x72, 0x22,
	0x37, 0x0a, 0x06, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x76, 0x61,
	0x74, 0x61, 0x72, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61,
	0x76, 0x61, 0x74, 0x61, 0x72, 0x55, 0x72, 0x6c, 0x22, 0x44, 0x0a, 0x08, 0x4c, 0x6f, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x08, 0x6c, 0x61, 0x74, 0x69, 0x74, 0x75, 0x64, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x09, 0x6c, 0x6f, 0x6e, 0x67, 0x69, 0x74, 0x75, 0x64, 0x65, 0x22, 0x9f,
	0x01, 0x0a, 0x03, 0x43, 0x61, 0x72, 0x12, 0x29, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x26, 0x0a, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x0e, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65,
	0x72, 0x52, 0x06, 0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x6f, 0x73,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x61,
	0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70,
	0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x70, 0x5f,
	0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x69, 0x70, 0x49, 0x64,
	0x22, 0x0e, 0x0a, 0x0c, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71,
	0x22, 0x1b, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0c, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x73, 0x52, 0x65, 0x71, 0x22, 0x33, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x43, 0x61, 0x72, 0x73, 0x52, 0x73, 0x70, 0x12, 0x25, 0x0a, 0x04, 0x63, 0x61, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31,
	0x2e, 0x43, 0x61, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x63, 0x61, 0x72, 0x73,
	0x22, 0x1c, 0x0a, 0x0a, 0x4c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x0c,
	0x0a, 0x0a, 0x4c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x52, 0x73, 0x70, 0x22, 0x5f, 0x0a, 0x0c,
	0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x26, 0x0a, 0x06,
	0x64, 0x72, 0x69, 0x76, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x63,
	0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x44, 0x72, 0x69, 0x76, 0x65, 0x72, 0x52, 0x06, 0x64, 0x72,
	0x69, 0x76, 0x65, 0x72, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x72, 0x69, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74, 0x72, 0x69, 0x70, 0x49, 0x64, 0x22, 0x0e, 0x0a,
	0x0c, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x52, 0x73, 0x70, 0x22, 0x77, 0x0a,
	0x0c, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x29, 0x0a,
	0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x11, 0x2e,
	0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x2c, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x61, 0x72,
	0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x08, 0x70, 0x6f,
	0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x13, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2a, 0x57, 0x0a, 0x09, 0x43,
	0x61, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x14, 0x0a, 0x10, 0x43, 0x53, 0x5f, 0x4e,
	0x4f, 0x54, 0x5f, 0x53, 0x50, 0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a,
	0x0a, 0x06, 0x4c, 0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x55, 0x4e,
	0x4c, 0x4f, 0x43, 0x4b, 0x49, 0x4e, 0x47, 0x10, 0x02, 0x12, 0x0c, 0x0a, 0x08, 0x55, 0x4e, 0x4c,
	0x4f, 0x43, 0x4b, 0x45, 0x44, 0x10, 0x03, 0x12, 0x0b, 0x0a, 0x07, 0x4c, 0x4f, 0x43, 0x4b, 0x49,
	0x4e, 0x47, 0x10, 0x04, 0x32, 0xc9, 0x02, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72,
	0x12, 0x14, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x28, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x12, 0x11, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x61, 0x72, 0x12, 0x31, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x73, 0x12, 0x12,
	0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x73, 0x52, 0x73, 0x70, 0x12, 0x31, 0x0a, 0x07, 0x4c, 0x6f, 0x63, 0x6b, 0x43, 0x61,
	0x72, 0x12, 0x12, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x6f, 0x63, 0x6b, 0x43,
	0x61, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x12, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x4c,
	0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x52, 0x73, 0x70, 0x12, 0x37, 0x0a, 0x09, 0x55, 0x6e, 0x6c,
	0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x12, 0x14, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e,
	0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x63,
	0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x6c, 0x6f, 0x63, 0x6b, 0x43, 0x61, 0x72, 0x52,
	0x73, 0x70, 0x12, 0x3c, 0x0a, 0x09, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x12,
	0x14, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x19, 0x2e, 0x63, 0x61, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x1e, 0x5a, 0x1c, 0x63, 0x6f, 0x6f, 0x6c, 0x63, 0x61, 0x72, 0x2f, 0x63, 0x61, 0x72, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x61, 0x72, 0x70, 0x62,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_car_proto_rawDescOnce sync.Once
	file_car_proto_rawDescData = file_car_proto_rawDesc
)

func file_car_proto_rawDescGZIP() []byte {
	file_car_proto_rawDescOnce.Do(func() {
		file_car_proto_rawDescData = protoimpl.X.CompressGZIP(file_car_proto_rawDescData)
	})
	return file_car_proto_rawDescData
}

var file_car_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_car_proto_msgTypes = make([]protoimpl.MessageInfo, 14)
var file_car_proto_goTypes = []interface{}{
	(CarStatus)(0),            // 0: car.v1.CarStatus
	(*CarEntity)(nil),         // 1: car.v1.CarEntity
	(*Driver)(nil),            // 2: car.v1.Driver
	(*Location)(nil),          // 3: car.v1.Location
	(*Car)(nil),               // 4: car.v1.Car
	(*CreateCarReq)(nil),      // 5: car.v1.CreateCarReq
	(*GetCarReq)(nil),         // 6: car.v1.GetCarReq
	(*GetCarsReq)(nil),        // 7: car.v1.GetCarsReq
	(*GetCarsRsp)(nil),        // 8: car.v1.GetCarsRsp
	(*LockCarReq)(nil),        // 9: car.v1.LockCarReq
	(*LockCarRsp)(nil),        // 10: car.v1.LockCarRsp
	(*UnlockCarReq)(nil),      // 11: car.v1.UnlockCarReq
	(*UnlockCarRsp)(nil),      // 12: car.v1.UnlockCarRsp
	(*UpdateCarReq)(nil),      // 13: car.v1.UpdateCarReq
	(*UpdateCarResponse)(nil), // 14: car.v1.UpdateCarResponse
}
var file_car_proto_depIdxs = []int32{
	4,  // 0: car.v1.CarEntity.car:type_name -> car.v1.Car
	0,  // 1: car.v1.Car.status:type_name -> car.v1.CarStatus
	2,  // 2: car.v1.Car.driver:type_name -> car.v1.Driver
	3,  // 3: car.v1.Car.position:type_name -> car.v1.Location
	1,  // 4: car.v1.GetCarsRsp.cars:type_name -> car.v1.CarEntity
	2,  // 5: car.v1.UnlockCarReq.driver:type_name -> car.v1.Driver
	0,  // 6: car.v1.UpdateCarReq.status:type_name -> car.v1.CarStatus
	3,  // 7: car.v1.UpdateCarReq.position:type_name -> car.v1.Location
	5,  // 8: car.v1.CarService.CreateCar:input_type -> car.v1.CreateCarReq
	6,  // 9: car.v1.CarService.GetCar:input_type -> car.v1.GetCarReq
	7,  // 10: car.v1.CarService.GetCars:input_type -> car.v1.GetCarsReq
	9,  // 11: car.v1.CarService.LockCar:input_type -> car.v1.LockCarReq
	11, // 12: car.v1.CarService.UnlockCar:input_type -> car.v1.UnlockCarReq
	13, // 13: car.v1.CarService.UpdateCar:input_type -> car.v1.UpdateCarReq
	1,  // 14: car.v1.CarService.CreateCar:output_type -> car.v1.CarEntity
	4,  // 15: car.v1.CarService.GetCar:output_type -> car.v1.Car
	8,  // 16: car.v1.CarService.GetCars:output_type -> car.v1.GetCarsRsp
	10, // 17: car.v1.CarService.LockCar:output_type -> car.v1.LockCarRsp
	12, // 18: car.v1.CarService.UnlockCar:output_type -> car.v1.UnlockCarRsp
	14, // 19: car.v1.CarService.UpdateCar:output_type -> car.v1.UpdateCarResponse
	14, // [14:20] is the sub-list for method output_type
	8,  // [8:14] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_car_proto_init() }
func file_car_proto_init() {
	if File_car_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_car_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CarEntity); i {
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
		file_car_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Driver); i {
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
		file_car_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Location); i {
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
		file_car_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Car); i {
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
		file_car_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCarReq); i {
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
		file_car_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarReq); i {
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
		file_car_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarsReq); i {
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
		file_car_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarsRsp); i {
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
		file_car_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockCarReq); i {
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
		file_car_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LockCarRsp); i {
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
		file_car_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockCarReq); i {
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
		file_car_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnlockCarRsp); i {
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
		file_car_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCarReq); i {
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
		file_car_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCarResponse); i {
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
			RawDescriptor: file_car_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   14,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_car_proto_goTypes,
		DependencyIndexes: file_car_proto_depIdxs,
		EnumInfos:         file_car_proto_enumTypes,
		MessageInfos:      file_car_proto_msgTypes,
	}.Build()
	File_car_proto = out.File
	file_car_proto_rawDesc = nil
	file_car_proto_goTypes = nil
	file_car_proto_depIdxs = nil
}
