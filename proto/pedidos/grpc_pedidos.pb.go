// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: proto/pedidos/grpc_pedidos.proto

package proto

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type GetProdutoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetProdutoRequest) Reset() {
	*x = GetProdutoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pedidos_grpc_pedidos_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProdutoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProdutoRequest) ProtoMessage() {}

func (x *GetProdutoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pedidos_grpc_pedidos_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProdutoRequest.ProtoReflect.Descriptor instead.
func (*GetProdutoRequest) Descriptor() ([]byte, []int) {
	return file_proto_pedidos_grpc_pedidos_proto_rawDescGZIP(), []int{0}
}

func (x *GetProdutoRequest) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Usuario struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Nome  string `protobuf:"bytes,2,opt,name=nome,proto3" json:"nome,omitempty"`
	Email string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Err   string `protobuf:"bytes,4,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *Usuario) Reset() {
	*x = Usuario{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pedidos_grpc_pedidos_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Usuario) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Usuario) ProtoMessage() {}

func (x *Usuario) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pedidos_grpc_pedidos_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Usuario.ProtoReflect.Descriptor instead.
func (*Usuario) Descriptor() ([]byte, []int) {
	return file_proto_pedidos_grpc_pedidos_proto_rawDescGZIP(), []int{1}
}

func (x *Usuario) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Usuario) GetNome() string {
	if x != nil {
		return x.Nome
	}
	return ""
}

func (x *Usuario) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Usuario) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

type Itens struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Descricao string  `protobuf:"bytes,2,opt,name=descricao,proto3" json:"descricao,omitempty"`
	Preco     float32 `protobuf:"fixed32,3,opt,name=preco,proto3" json:"preco,omitempty"`
	Qtde      uint32  `protobuf:"varint,4,opt,name=qtde,proto3" json:"qtde,omitempty"`
}

func (x *Itens) Reset() {
	*x = Itens{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pedidos_grpc_pedidos_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Itens) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Itens) ProtoMessage() {}

func (x *Itens) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pedidos_grpc_pedidos_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Itens.ProtoReflect.Descriptor instead.
func (*Itens) Descriptor() ([]byte, []int) {
	return file_proto_pedidos_grpc_pedidos_proto_rawDescGZIP(), []int{2}
}

func (x *Itens) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Itens) GetDescricao() string {
	if x != nil {
		return x.Descricao
	}
	return ""
}

func (x *Itens) GetPreco() float32 {
	if x != nil {
		return x.Preco
	}
	return 0
}

func (x *Itens) GetQtde() uint32 {
	if x != nil {
		return x.Qtde
	}
	return 0
}

type Pedido struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Data      string   `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Valor     float32  `protobuf:"fixed32,3,opt,name=valor,proto3" json:"valor,omitempty"`
	Pagamento string   `protobuf:"bytes,4,opt,name=pagamento,proto3" json:"pagamento,omitempty"`
	Usuario   *Usuario `protobuf:"bytes,5,opt,name=usuario,proto3" json:"usuario,omitempty"`
	Itens     []*Itens `protobuf:"bytes,6,rep,name=itens,proto3" json:"itens,omitempty"`
}

func (x *Pedido) Reset() {
	*x = Pedido{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pedidos_grpc_pedidos_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Pedido) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Pedido) ProtoMessage() {}

func (x *Pedido) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pedidos_grpc_pedidos_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Pedido.ProtoReflect.Descriptor instead.
func (*Pedido) Descriptor() ([]byte, []int) {
	return file_proto_pedidos_grpc_pedidos_proto_rawDescGZIP(), []int{3}
}

func (x *Pedido) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Pedido) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *Pedido) GetValor() float32 {
	if x != nil {
		return x.Valor
	}
	return 0
}

func (x *Pedido) GetPagamento() string {
	if x != nil {
		return x.Pagamento
	}
	return ""
}

func (x *Pedido) GetUsuario() *Usuario {
	if x != nil {
		return x.Usuario
	}
	return nil
}

func (x *Pedido) GetItens() []*Itens {
	if x != nil {
		return x.Itens
	}
	return nil
}

type GetProdutoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pedido *Pedido `protobuf:"bytes,1,opt,name=pedido,proto3" json:"pedido,omitempty"`
	Err    string  `protobuf:"bytes,2,opt,name=err,proto3" json:"err,omitempty"`
}

func (x *GetProdutoResponse) Reset() {
	*x = GetProdutoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pedidos_grpc_pedidos_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProdutoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProdutoResponse) ProtoMessage() {}

func (x *GetProdutoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pedidos_grpc_pedidos_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProdutoResponse.ProtoReflect.Descriptor instead.
func (*GetProdutoResponse) Descriptor() ([]byte, []int) {
	return file_proto_pedidos_grpc_pedidos_proto_rawDescGZIP(), []int{4}
}

func (x *GetProdutoResponse) GetPedido() *Pedido {
	if x != nil {
		return x.Pedido
	}
	return nil
}

func (x *GetProdutoResponse) GetErr() string {
	if x != nil {
		return x.Err
	}
	return ""
}

var File_proto_pedidos_grpc_pedidos_proto protoreflect.FileDescriptor

var file_proto_pedidos_grpc_pedidos_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x73, 0x2f,
	0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x55,
	0x0a, 0x07, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x6f, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x65, 0x72, 0x72, 0x22, 0x5f, 0x0a, 0x05, 0x49, 0x74, 0x65, 0x6e, 0x73, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x63, 0x61, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x63, 0x61, 0x6f, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x65, 0x63, 0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x70, 0x72, 0x65,
	0x63, 0x6f, 0x12, 0x12, 0x0a, 0x04, 0x71, 0x74, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0d,
	0x52, 0x04, 0x71, 0x74, 0x64, 0x65, 0x22, 0xae, 0x01, 0x0a, 0x06, 0x50, 0x65, 0x64, 0x69, 0x64,
	0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x02, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x6f, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x70,
	0x61, 0x67, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x70, 0x61, 0x67, 0x61, 0x6d, 0x65, 0x6e, 0x74, 0x6f, 0x12, 0x28, 0x0a, 0x07, 0x75, 0x73, 0x75,
	0x61, 0x72, 0x69, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x73, 0x75, 0x61, 0x72, 0x69, 0x6f, 0x52, 0x07, 0x75, 0x73, 0x75, 0x61,
	0x72, 0x69, 0x6f, 0x12, 0x22, 0x0a, 0x05, 0x69, 0x74, 0x65, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x49, 0x74, 0x65, 0x6e, 0x73,
	0x52, 0x05, 0x69, 0x74, 0x65, 0x6e, 0x73, 0x22, 0x4d, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x25, 0x0a,
	0x06, 0x70, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x50, 0x65, 0x64, 0x69, 0x64, 0x6f, 0x52, 0x06, 0x70, 0x65,
	0x64, 0x69, 0x64, 0x6f, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x72, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x65, 0x72, 0x72, 0x32, 0x56, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x74, 0x6f, 0x12, 0x41, 0x0a, 0x0a, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x74, 0x6f, 0x12, 0x18, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x74, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_pedidos_grpc_pedidos_proto_rawDescOnce sync.Once
	file_proto_pedidos_grpc_pedidos_proto_rawDescData = file_proto_pedidos_grpc_pedidos_proto_rawDesc
)

func file_proto_pedidos_grpc_pedidos_proto_rawDescGZIP() []byte {
	file_proto_pedidos_grpc_pedidos_proto_rawDescOnce.Do(func() {
		file_proto_pedidos_grpc_pedidos_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_pedidos_grpc_pedidos_proto_rawDescData)
	})
	return file_proto_pedidos_grpc_pedidos_proto_rawDescData
}

var file_proto_pedidos_grpc_pedidos_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_pedidos_grpc_pedidos_proto_goTypes = []interface{}{
	(*GetProdutoRequest)(nil),  // 0: proto.GetProdutoRequest
	(*Usuario)(nil),            // 1: proto.Usuario
	(*Itens)(nil),              // 2: proto.Itens
	(*Pedido)(nil),             // 3: proto.Pedido
	(*GetProdutoResponse)(nil), // 4: proto.GetProdutoResponse
}
var file_proto_pedidos_grpc_pedidos_proto_depIdxs = []int32{
	1, // 0: proto.Pedido.usuario:type_name -> proto.Usuario
	2, // 1: proto.Pedido.itens:type_name -> proto.Itens
	3, // 2: proto.GetProdutoResponse.pedido:type_name -> proto.Pedido
	0, // 3: proto.ServiceGetProduto.GetProduto:input_type -> proto.GetProdutoRequest
	4, // 4: proto.ServiceGetProduto.GetProduto:output_type -> proto.GetProdutoResponse
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_proto_pedidos_grpc_pedidos_proto_init() }
func file_proto_pedidos_grpc_pedidos_proto_init() {
	if File_proto_pedidos_grpc_pedidos_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_pedidos_grpc_pedidos_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProdutoRequest); i {
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
		file_proto_pedidos_grpc_pedidos_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Usuario); i {
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
		file_proto_pedidos_grpc_pedidos_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Itens); i {
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
		file_proto_pedidos_grpc_pedidos_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Pedido); i {
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
		file_proto_pedidos_grpc_pedidos_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProdutoResponse); i {
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
			RawDescriptor: file_proto_pedidos_grpc_pedidos_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_pedidos_grpc_pedidos_proto_goTypes,
		DependencyIndexes: file_proto_pedidos_grpc_pedidos_proto_depIdxs,
		MessageInfos:      file_proto_pedidos_grpc_pedidos_proto_msgTypes,
	}.Build()
	File_proto_pedidos_grpc_pedidos_proto = out.File
	file_proto_pedidos_grpc_pedidos_proto_rawDesc = nil
	file_proto_pedidos_grpc_pedidos_proto_goTypes = nil
	file_proto_pedidos_grpc_pedidos_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ServiceGetProdutoClient is the client API for ServiceGetProduto service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ServiceGetProdutoClient interface {
	GetProduto(ctx context.Context, in *GetProdutoRequest, opts ...grpc.CallOption) (*GetProdutoResponse, error)
}

type serviceGetProdutoClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceGetProdutoClient(cc grpc.ClientConnInterface) ServiceGetProdutoClient {
	return &serviceGetProdutoClient{cc}
}

func (c *serviceGetProdutoClient) GetProduto(ctx context.Context, in *GetProdutoRequest, opts ...grpc.CallOption) (*GetProdutoResponse, error) {
	out := new(GetProdutoResponse)
	err := c.cc.Invoke(ctx, "/proto.ServiceGetProduto/GetProduto", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ServiceGetProdutoServer is the server API for ServiceGetProduto service.
type ServiceGetProdutoServer interface {
	GetProduto(context.Context, *GetProdutoRequest) (*GetProdutoResponse, error)
}

// UnimplementedServiceGetProdutoServer can be embedded to have forward compatible implementations.
type UnimplementedServiceGetProdutoServer struct {
}

func (*UnimplementedServiceGetProdutoServer) GetProduto(context.Context, *GetProdutoRequest) (*GetProdutoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProduto not implemented")
}

func RegisterServiceGetProdutoServer(s *grpc.Server, srv ServiceGetProdutoServer) {
	s.RegisterService(&_ServiceGetProduto_serviceDesc, srv)
}

func _ServiceGetProduto_GetProduto_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProdutoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ServiceGetProdutoServer).GetProduto(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ServiceGetProduto/GetProduto",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ServiceGetProdutoServer).GetProduto(ctx, req.(*GetProdutoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ServiceGetProduto_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ServiceGetProduto",
	HandlerType: (*ServiceGetProdutoServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProduto",
			Handler:    _ServiceGetProduto_GetProduto_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/pedidos/grpc_pedidos.proto",
}
