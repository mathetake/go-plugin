// Protocol Buffers - Google's data interchange format
// Copyright 2008 Google Inc.  All rights reserved.
// Copyright 2022 Teppei Fukuda.  All rights reserved.
// https://developers.google.com/protocol-buffers/

// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.0.1
// 	protoc               v3.21.5
// source: types/known/emptypb/empty.proto

package emptypb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// A generic empty message that you can re-use to avoid defining duplicated
// empty messages in your APIs. A typical example is to use it as the request
// or the response type of an API method. For instance:
//
//     service Foo {
//       rpc Bar(google.protobuf.Empty) returns (google.protobuf.Empty);
//     }
//
type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}
