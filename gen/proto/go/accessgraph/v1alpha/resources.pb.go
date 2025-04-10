//
// Teleport
// Copyright (C) 2023  Gravitational, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        (unknown)
// source: accessgraph/v1alpha/resources.proto

package accessgraphv1alpha

import (
	v14 "github.com/gravitational/teleport/api/gen/proto/go/teleport/accessgraph/v1"
	v1 "github.com/gravitational/teleport/api/gen/proto/go/teleport/accesslist/v1"
	v11 "github.com/gravitational/teleport/api/gen/proto/go/teleport/crownjewel/v1"
	v12 "github.com/gravitational/teleport/api/gen/proto/go/teleport/dbobject/v1"
	v13 "github.com/gravitational/teleport/api/gen/proto/go/teleport/devicetrust/v1"
	types "github.com/gravitational/teleport/api/types"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// ResourceList is a list of resources to send to the access graph.
type ResourceList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Resources     []*ResourceEntry       `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceList) Reset() {
	*x = ResourceList{}
	mi := &file_accessgraph_v1alpha_resources_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceList) ProtoMessage() {}

func (x *ResourceList) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_resources_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceList.ProtoReflect.Descriptor instead.
func (*ResourceList) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_resources_proto_rawDescGZIP(), []int{0}
}

func (x *ResourceList) GetResources() []*ResourceEntry {
	if x != nil {
		return x.Resources
	}
	return nil
}

// ResourceHeaderList is a list of resource headers to send to the access graph.
type ResourceHeaderList struct {
	state         protoimpl.MessageState  `protogen:"open.v1"`
	Resources     []*types.ResourceHeader `protobuf:"bytes,1,rep,name=resources,proto3" json:"resources,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceHeaderList) Reset() {
	*x = ResourceHeaderList{}
	mi := &file_accessgraph_v1alpha_resources_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceHeaderList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceHeaderList) ProtoMessage() {}

func (x *ResourceHeaderList) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_resources_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceHeaderList.ProtoReflect.Descriptor instead.
func (*ResourceHeaderList) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_resources_proto_rawDescGZIP(), []int{1}
}

func (x *ResourceHeaderList) GetResources() []*types.ResourceHeader {
	if x != nil {
		return x.Resources
	}
	return nil
}

// AccessListsMembers is the request to declare users as members of access lists.
type AccessListsMembers struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// members is the list of members to add to access lists.
	Members       []*v1.Member `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AccessListsMembers) Reset() {
	*x = AccessListsMembers{}
	mi := &file_accessgraph_v1alpha_resources_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AccessListsMembers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessListsMembers) ProtoMessage() {}

func (x *AccessListsMembers) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_resources_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessListsMembers.ProtoReflect.Descriptor instead.
func (*AccessListsMembers) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_resources_proto_rawDescGZIP(), []int{2}
}

func (x *AccessListsMembers) GetMembers() []*v1.Member {
	if x != nil {
		return x.Members
	}
	return nil
}

// ExcludeAccessListsMembers is the request to exclude users from access lists.
type ExcludeAccessListsMembers struct {
	state         protoimpl.MessageState     `protogen:"open.v1"`
	Members       []*ExcludeAccessListMember `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExcludeAccessListsMembers) Reset() {
	*x = ExcludeAccessListsMembers{}
	mi := &file_accessgraph_v1alpha_resources_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExcludeAccessListsMembers) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExcludeAccessListsMembers) ProtoMessage() {}

func (x *ExcludeAccessListsMembers) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_resources_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExcludeAccessListsMembers.ProtoReflect.Descriptor instead.
func (*ExcludeAccessListsMembers) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_resources_proto_rawDescGZIP(), []int{3}
}

func (x *ExcludeAccessListsMembers) GetMembers() []*ExcludeAccessListMember {
	if x != nil {
		return x.Members
	}
	return nil
}

// ExcludeAccessListMember is the request to exclude a user from an access list.
type ExcludeAccessListMember struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	AccessList    string                 `protobuf:"bytes,1,opt,name=access_list,json=accessList,proto3" json:"access_list,omitempty"`
	Username      string                 `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ExcludeAccessListMember) Reset() {
	*x = ExcludeAccessListMember{}
	mi := &file_accessgraph_v1alpha_resources_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ExcludeAccessListMember) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExcludeAccessListMember) ProtoMessage() {}

func (x *ExcludeAccessListMember) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_resources_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExcludeAccessListMember.ProtoReflect.Descriptor instead.
func (*ExcludeAccessListMember) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_resources_proto_rawDescGZIP(), []int{4}
}

func (x *ExcludeAccessListMember) GetAccessList() string {
	if x != nil {
		return x.AccessList
	}
	return ""
}

func (x *ExcludeAccessListMember) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

// ResourceEntry is a wrapper for the supported resource types.
type ResourceEntry struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Types that are valid to be assigned to Resource:
	//
	//	*ResourceEntry_User
	//	*ResourceEntry_Role
	//	*ResourceEntry_Server
	//	*ResourceEntry_AccessRequest
	//	*ResourceEntry_KubernetesServer
	//	*ResourceEntry_AppServer
	//	*ResourceEntry_DatabaseServer
	//	*ResourceEntry_WindowsDesktop
	//	*ResourceEntry_AccessList
	//	*ResourceEntry_CrownJewel
	//	*ResourceEntry_DatabaseObject
	//	*ResourceEntry_Device
	//	*ResourceEntry_PrivateKey
	//	*ResourceEntry_AuthorizedKey
	Resource      isResourceEntry_Resource `protobuf_oneof:"resource"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ResourceEntry) Reset() {
	*x = ResourceEntry{}
	mi := &file_accessgraph_v1alpha_resources_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ResourceEntry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResourceEntry) ProtoMessage() {}

func (x *ResourceEntry) ProtoReflect() protoreflect.Message {
	mi := &file_accessgraph_v1alpha_resources_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResourceEntry.ProtoReflect.Descriptor instead.
func (*ResourceEntry) Descriptor() ([]byte, []int) {
	return file_accessgraph_v1alpha_resources_proto_rawDescGZIP(), []int{5}
}

func (x *ResourceEntry) GetResource() isResourceEntry_Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *ResourceEntry) GetUser() *types.UserV2 {
	if x != nil {
		if x, ok := x.Resource.(*ResourceEntry_User); ok {
			return x.User
		}
	}
	return nil
}

func (x *ResourceEntry) GetRole() *types.RoleV6 {
	if x != nil {
		if x, ok := x.Resource.(*ResourceEntry_Role); ok {
			return x.Role
		}
	}
	return nil
}

func (x *ResourceEntry) GetServer() *types.ServerV2 {
	if x != nil {
		if x, ok := x.Resource.(*ResourceEntry_Server); ok {
			return x.Server
		}
	}
	return nil
}

func (x *ResourceEntry) GetAccessRequest() *types.AccessRequestV3 {
	if x != nil {
		if x, ok := x.Resource.(*ResourceEntry_AccessRequest); ok {
			return x.AccessRequest
		}
	}
	return nil
}

func (x *ResourceEntry) GetKubernetesServer() *types.KubernetesServerV3 {
	if x != nil {
		if x, ok := x.Resource.(*ResourceEntry_KubernetesServer); ok {
			return x.KubernetesServer
		}
	}
	return nil
}

func (x *ResourceEntry) GetAppServer() *types.AppServerV3 {
	if x != nil {
		if x, ok := x.Resource.(*ResourceEntry_AppServer); ok {
			return x.AppServer
		}
	}
	return nil
}

func (x *ResourceEntry) GetDatabaseServer() *types.DatabaseServerV3 {
	if x != nil {
		if x, ok := x.Resource.(*ResourceEntry_DatabaseServer); ok {
			return x.DatabaseServer
		}
	}
	return nil
}

func (x *ResourceEntry) GetWindowsDesktop() *types.WindowsDesktopV3 {
	if x != nil {
		if x, ok := x.Resource.(*ResourceEntry_WindowsDesktop); ok {
			return x.WindowsDesktop
		}
	}
	return nil
}

func (x *ResourceEntry) GetAccessList() *v1.AccessList {
	if x != nil {
		if x, ok := x.Resource.(*ResourceEntry_AccessList); ok {
			return x.AccessList
		}
	}
	return nil
}

func (x *ResourceEntry) GetCrownJewel() *v11.CrownJewel {
	if x != nil {
		if x, ok := x.Resource.(*ResourceEntry_CrownJewel); ok {
			return x.CrownJewel
		}
	}
	return nil
}

func (x *ResourceEntry) GetDatabaseObject() *v12.DatabaseObject {
	if x != nil {
		if x, ok := x.Resource.(*ResourceEntry_DatabaseObject); ok {
			return x.DatabaseObject
		}
	}
	return nil
}

func (x *ResourceEntry) GetDevice() *v13.Device {
	if x != nil {
		if x, ok := x.Resource.(*ResourceEntry_Device); ok {
			return x.Device
		}
	}
	return nil
}

func (x *ResourceEntry) GetPrivateKey() *v14.PrivateKey {
	if x != nil {
		if x, ok := x.Resource.(*ResourceEntry_PrivateKey); ok {
			return x.PrivateKey
		}
	}
	return nil
}

func (x *ResourceEntry) GetAuthorizedKey() *v14.AuthorizedKey {
	if x != nil {
		if x, ok := x.Resource.(*ResourceEntry_AuthorizedKey); ok {
			return x.AuthorizedKey
		}
	}
	return nil
}

type isResourceEntry_Resource interface {
	isResourceEntry_Resource()
}

type ResourceEntry_User struct {
	// user is a user resource
	User *types.UserV2 `protobuf:"bytes,1,opt,name=user,proto3,oneof"`
}

type ResourceEntry_Role struct {
	// role is a role resource
	Role *types.RoleV6 `protobuf:"bytes,2,opt,name=role,proto3,oneof"`
}

type ResourceEntry_Server struct {
	// server is a node/server resource
	Server *types.ServerV2 `protobuf:"bytes,3,opt,name=server,proto3,oneof"`
}

type ResourceEntry_AccessRequest struct {
	// access_request is a resource for access requests
	AccessRequest *types.AccessRequestV3 `protobuf:"bytes,4,opt,name=access_request,json=accessRequest,proto3,oneof"`
}

type ResourceEntry_KubernetesServer struct {
	// kubernetes_server is a kubernetes server resource
	KubernetesServer *types.KubernetesServerV3 `protobuf:"bytes,5,opt,name=kubernetes_server,json=kubernetesServer,proto3,oneof"`
}

type ResourceEntry_AppServer struct {
	// app_server is an application server resource
	AppServer *types.AppServerV3 `protobuf:"bytes,6,opt,name=app_server,json=appServer,proto3,oneof"`
}

type ResourceEntry_DatabaseServer struct {
	// database_server is a database server resource
	DatabaseServer *types.DatabaseServerV3 `protobuf:"bytes,7,opt,name=database_server,json=databaseServer,proto3,oneof"`
}

type ResourceEntry_WindowsDesktop struct {
	// windows_desktop is a resource for Windows desktop host.
	WindowsDesktop *types.WindowsDesktopV3 `protobuf:"bytes,8,opt,name=windows_desktop,json=windowsDesktop,proto3,oneof"`
}

type ResourceEntry_AccessList struct {
	// access_list is a resource for access lists.
	AccessList *v1.AccessList `protobuf:"bytes,9,opt,name=access_list,json=accessList,proto3,oneof"`
}

type ResourceEntry_CrownJewel struct {
	// crown_jewel is a resource for crown jewels.
	CrownJewel *v11.CrownJewel `protobuf:"bytes,10,opt,name=crown_jewel,json=crownJewel,proto3,oneof"`
}

type ResourceEntry_DatabaseObject struct {
	// database_object is a resource for database objects.
	DatabaseObject *v12.DatabaseObject `protobuf:"bytes,11,opt,name=database_object,json=databaseObject,proto3,oneof"`
}

type ResourceEntry_Device struct {
	// device is a device trust resource.
	Device *v13.Device `protobuf:"bytes,12,opt,name=device,proto3,oneof"`
}

type ResourceEntry_PrivateKey struct {
	// private_key represents a private key resource found in user's laptops.
	PrivateKey *v14.PrivateKey `protobuf:"bytes,13,opt,name=private_key,json=privateKey,proto3,oneof"`
}

type ResourceEntry_AuthorizedKey struct {
	// authorized_key represents a authorized key for a server.
	AuthorizedKey *v14.AuthorizedKey `protobuf:"bytes,14,opt,name=authorized_key,json=authorizedKey,proto3,oneof"`
}

func (*ResourceEntry_User) isResourceEntry_Resource() {}

func (*ResourceEntry_Role) isResourceEntry_Resource() {}

func (*ResourceEntry_Server) isResourceEntry_Resource() {}

func (*ResourceEntry_AccessRequest) isResourceEntry_Resource() {}

func (*ResourceEntry_KubernetesServer) isResourceEntry_Resource() {}

func (*ResourceEntry_AppServer) isResourceEntry_Resource() {}

func (*ResourceEntry_DatabaseServer) isResourceEntry_Resource() {}

func (*ResourceEntry_WindowsDesktop) isResourceEntry_Resource() {}

func (*ResourceEntry_AccessList) isResourceEntry_Resource() {}

func (*ResourceEntry_CrownJewel) isResourceEntry_Resource() {}

func (*ResourceEntry_DatabaseObject) isResourceEntry_Resource() {}

func (*ResourceEntry_Device) isResourceEntry_Resource() {}

func (*ResourceEntry_PrivateKey) isResourceEntry_Resource() {}

func (*ResourceEntry_AuthorizedKey) isResourceEntry_Resource() {}

var File_accessgraph_v1alpha_resources_proto protoreflect.FileDescriptor

const file_accessgraph_v1alpha_resources_proto_rawDesc = "" +
	"\n" +
	"#accessgraph/v1alpha/resources.proto\x12\x13accessgraph.v1alpha\x1a-teleport/access_graph/v1/authorized_key.proto\x1a*teleport/access_graph/v1/private_key.proto\x1a'teleport/accesslist/v1/accesslist.proto\x1a'teleport/crownjewel/v1/crownjewel.proto\x1a#teleport/dbobject/v1/dbobject.proto\x1a$teleport/devicetrust/v1/device.proto\x1a!teleport/legacy/types/types.proto\"P\n" +
	"\fResourceList\x12@\n" +
	"\tresources\x18\x01 \x03(\v2\".accessgraph.v1alpha.ResourceEntryR\tresources\"I\n" +
	"\x12ResourceHeaderList\x123\n" +
	"\tresources\x18\x01 \x03(\v2\x15.types.ResourceHeaderR\tresources\"N\n" +
	"\x12AccessListsMembers\x128\n" +
	"\amembers\x18\x01 \x03(\v2\x1e.teleport.accesslist.v1.MemberR\amembers\"c\n" +
	"\x19ExcludeAccessListsMembers\x12F\n" +
	"\amembers\x18\x01 \x03(\v2,.accessgraph.v1alpha.ExcludeAccessListMemberR\amembers\"V\n" +
	"\x17ExcludeAccessListMember\x12\x1f\n" +
	"\vaccess_list\x18\x01 \x01(\tR\n" +
	"accessList\x12\x1a\n" +
	"\busername\x18\x02 \x01(\tR\busername\"\x8d\a\n" +
	"\rResourceEntry\x12#\n" +
	"\x04user\x18\x01 \x01(\v2\r.types.UserV2H\x00R\x04user\x12#\n" +
	"\x04role\x18\x02 \x01(\v2\r.types.RoleV6H\x00R\x04role\x12)\n" +
	"\x06server\x18\x03 \x01(\v2\x0f.types.ServerV2H\x00R\x06server\x12?\n" +
	"\x0eaccess_request\x18\x04 \x01(\v2\x16.types.AccessRequestV3H\x00R\raccessRequest\x12H\n" +
	"\x11kubernetes_server\x18\x05 \x01(\v2\x19.types.KubernetesServerV3H\x00R\x10kubernetesServer\x123\n" +
	"\n" +
	"app_server\x18\x06 \x01(\v2\x12.types.AppServerV3H\x00R\tappServer\x12B\n" +
	"\x0fdatabase_server\x18\a \x01(\v2\x17.types.DatabaseServerV3H\x00R\x0edatabaseServer\x12B\n" +
	"\x0fwindows_desktop\x18\b \x01(\v2\x17.types.WindowsDesktopV3H\x00R\x0ewindowsDesktop\x12E\n" +
	"\vaccess_list\x18\t \x01(\v2\".teleport.accesslist.v1.AccessListH\x00R\n" +
	"accessList\x12E\n" +
	"\vcrown_jewel\x18\n" +
	" \x01(\v2\".teleport.crownjewel.v1.CrownJewelH\x00R\n" +
	"crownJewel\x12O\n" +
	"\x0fdatabase_object\x18\v \x01(\v2$.teleport.dbobject.v1.DatabaseObjectH\x00R\x0edatabaseObject\x129\n" +
	"\x06device\x18\f \x01(\v2\x1f.teleport.devicetrust.v1.DeviceH\x00R\x06device\x12G\n" +
	"\vprivate_key\x18\r \x01(\v2$.teleport.access_graph.v1.PrivateKeyH\x00R\n" +
	"privateKey\x12P\n" +
	"\x0eauthorized_key\x18\x0e \x01(\v2'.teleport.access_graph.v1.AuthorizedKeyH\x00R\rauthorizedKeyB\n" +
	"\n" +
	"\bresourceBWZUgithub.com/gravitational/teleport/gen/proto/go/accessgraph/v1alpha;accessgraphv1alphab\x06proto3"

var (
	file_accessgraph_v1alpha_resources_proto_rawDescOnce sync.Once
	file_accessgraph_v1alpha_resources_proto_rawDescData []byte
)

func file_accessgraph_v1alpha_resources_proto_rawDescGZIP() []byte {
	file_accessgraph_v1alpha_resources_proto_rawDescOnce.Do(func() {
		file_accessgraph_v1alpha_resources_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_accessgraph_v1alpha_resources_proto_rawDesc), len(file_accessgraph_v1alpha_resources_proto_rawDesc)))
	})
	return file_accessgraph_v1alpha_resources_proto_rawDescData
}

var file_accessgraph_v1alpha_resources_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_accessgraph_v1alpha_resources_proto_goTypes = []any{
	(*ResourceList)(nil),              // 0: accessgraph.v1alpha.ResourceList
	(*ResourceHeaderList)(nil),        // 1: accessgraph.v1alpha.ResourceHeaderList
	(*AccessListsMembers)(nil),        // 2: accessgraph.v1alpha.AccessListsMembers
	(*ExcludeAccessListsMembers)(nil), // 3: accessgraph.v1alpha.ExcludeAccessListsMembers
	(*ExcludeAccessListMember)(nil),   // 4: accessgraph.v1alpha.ExcludeAccessListMember
	(*ResourceEntry)(nil),             // 5: accessgraph.v1alpha.ResourceEntry
	(*types.ResourceHeader)(nil),      // 6: types.ResourceHeader
	(*v1.Member)(nil),                 // 7: teleport.accesslist.v1.Member
	(*types.UserV2)(nil),              // 8: types.UserV2
	(*types.RoleV6)(nil),              // 9: types.RoleV6
	(*types.ServerV2)(nil),            // 10: types.ServerV2
	(*types.AccessRequestV3)(nil),     // 11: types.AccessRequestV3
	(*types.KubernetesServerV3)(nil),  // 12: types.KubernetesServerV3
	(*types.AppServerV3)(nil),         // 13: types.AppServerV3
	(*types.DatabaseServerV3)(nil),    // 14: types.DatabaseServerV3
	(*types.WindowsDesktopV3)(nil),    // 15: types.WindowsDesktopV3
	(*v1.AccessList)(nil),             // 16: teleport.accesslist.v1.AccessList
	(*v11.CrownJewel)(nil),            // 17: teleport.crownjewel.v1.CrownJewel
	(*v12.DatabaseObject)(nil),        // 18: teleport.dbobject.v1.DatabaseObject
	(*v13.Device)(nil),                // 19: teleport.devicetrust.v1.Device
	(*v14.PrivateKey)(nil),            // 20: teleport.access_graph.v1.PrivateKey
	(*v14.AuthorizedKey)(nil),         // 21: teleport.access_graph.v1.AuthorizedKey
}
var file_accessgraph_v1alpha_resources_proto_depIdxs = []int32{
	5,  // 0: accessgraph.v1alpha.ResourceList.resources:type_name -> accessgraph.v1alpha.ResourceEntry
	6,  // 1: accessgraph.v1alpha.ResourceHeaderList.resources:type_name -> types.ResourceHeader
	7,  // 2: accessgraph.v1alpha.AccessListsMembers.members:type_name -> teleport.accesslist.v1.Member
	4,  // 3: accessgraph.v1alpha.ExcludeAccessListsMembers.members:type_name -> accessgraph.v1alpha.ExcludeAccessListMember
	8,  // 4: accessgraph.v1alpha.ResourceEntry.user:type_name -> types.UserV2
	9,  // 5: accessgraph.v1alpha.ResourceEntry.role:type_name -> types.RoleV6
	10, // 6: accessgraph.v1alpha.ResourceEntry.server:type_name -> types.ServerV2
	11, // 7: accessgraph.v1alpha.ResourceEntry.access_request:type_name -> types.AccessRequestV3
	12, // 8: accessgraph.v1alpha.ResourceEntry.kubernetes_server:type_name -> types.KubernetesServerV3
	13, // 9: accessgraph.v1alpha.ResourceEntry.app_server:type_name -> types.AppServerV3
	14, // 10: accessgraph.v1alpha.ResourceEntry.database_server:type_name -> types.DatabaseServerV3
	15, // 11: accessgraph.v1alpha.ResourceEntry.windows_desktop:type_name -> types.WindowsDesktopV3
	16, // 12: accessgraph.v1alpha.ResourceEntry.access_list:type_name -> teleport.accesslist.v1.AccessList
	17, // 13: accessgraph.v1alpha.ResourceEntry.crown_jewel:type_name -> teleport.crownjewel.v1.CrownJewel
	18, // 14: accessgraph.v1alpha.ResourceEntry.database_object:type_name -> teleport.dbobject.v1.DatabaseObject
	19, // 15: accessgraph.v1alpha.ResourceEntry.device:type_name -> teleport.devicetrust.v1.Device
	20, // 16: accessgraph.v1alpha.ResourceEntry.private_key:type_name -> teleport.access_graph.v1.PrivateKey
	21, // 17: accessgraph.v1alpha.ResourceEntry.authorized_key:type_name -> teleport.access_graph.v1.AuthorizedKey
	18, // [18:18] is the sub-list for method output_type
	18, // [18:18] is the sub-list for method input_type
	18, // [18:18] is the sub-list for extension type_name
	18, // [18:18] is the sub-list for extension extendee
	0,  // [0:18] is the sub-list for field type_name
}

func init() { file_accessgraph_v1alpha_resources_proto_init() }
func file_accessgraph_v1alpha_resources_proto_init() {
	if File_accessgraph_v1alpha_resources_proto != nil {
		return
	}
	file_accessgraph_v1alpha_resources_proto_msgTypes[5].OneofWrappers = []any{
		(*ResourceEntry_User)(nil),
		(*ResourceEntry_Role)(nil),
		(*ResourceEntry_Server)(nil),
		(*ResourceEntry_AccessRequest)(nil),
		(*ResourceEntry_KubernetesServer)(nil),
		(*ResourceEntry_AppServer)(nil),
		(*ResourceEntry_DatabaseServer)(nil),
		(*ResourceEntry_WindowsDesktop)(nil),
		(*ResourceEntry_AccessList)(nil),
		(*ResourceEntry_CrownJewel)(nil),
		(*ResourceEntry_DatabaseObject)(nil),
		(*ResourceEntry_Device)(nil),
		(*ResourceEntry_PrivateKey)(nil),
		(*ResourceEntry_AuthorizedKey)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_accessgraph_v1alpha_resources_proto_rawDesc), len(file_accessgraph_v1alpha_resources_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_accessgraph_v1alpha_resources_proto_goTypes,
		DependencyIndexes: file_accessgraph_v1alpha_resources_proto_depIdxs,
		MessageInfos:      file_accessgraph_v1alpha_resources_proto_msgTypes,
	}.Build()
	File_accessgraph_v1alpha_resources_proto = out.File
	file_accessgraph_v1alpha_resources_proto_goTypes = nil
	file_accessgraph_v1alpha_resources_proto_depIdxs = nil
}
