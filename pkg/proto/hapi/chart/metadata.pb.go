// Code generated by protoc-gen-go.
// source: hapi/chart/metadata.proto
// DO NOT EDIT!

package chart

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Metadata_Engine int32

const (
	Metadata_UNKNOWN Metadata_Engine = 0
	Metadata_GOTPL   Metadata_Engine = 1
)

var Metadata_Engine_name = map[int32]string{
	0: "UNKNOWN",
	1: "GOTPL",
}
var Metadata_Engine_value = map[string]int32{
	"UNKNOWN": 0,
	"GOTPL":   1,
}

func (x Metadata_Engine) String() string {
	return proto.EnumName(Metadata_Engine_name, int32(x))
}
func (Metadata_Engine) EnumDescriptor() ([]byte, []int) { return fileDescriptor2, []int{1, 0} }

// Maintainer describes a Chart maintainer.
type Maintainer struct {
	// Name is a user name or organization name
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// Email is an optional email address to contact the named maintainer
	Email string `protobuf:"bytes,2,opt,name=email" json:"email,omitempty"`
}

func (m *Maintainer) Reset()                    { *m = Maintainer{} }
func (m *Maintainer) String() string            { return proto.CompactTextString(m) }
func (*Maintainer) ProtoMessage()               {}
func (*Maintainer) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{0} }

// 	Metadata for a Chart file. This models the structure of a Chart.yaml file.
//
// 	Spec: https://k8s.io/helm/blob/master/docs/design/chart_format.md#the-chart-file
type Metadata struct {
	// The name of the chart
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	// The URL to a relevant project page, git repo, or contact person
	Home string `protobuf:"bytes,2,opt,name=home" json:"home,omitempty"`
	// Source is the URL to the source code of this chart
	Sources []string `protobuf:"bytes,3,rep,name=sources" json:"sources,omitempty"`
	// A SemVer 2 conformant version string of the chart
	Version string `protobuf:"bytes,4,opt,name=version" json:"version,omitempty"`
	// A one-sentence description of the chart
	Description string `protobuf:"bytes,5,opt,name=description" json:"description,omitempty"`
	// A list of string keywords
	Keywords []string `protobuf:"bytes,6,rep,name=keywords" json:"keywords,omitempty"`
	// A list of name and URL/email address combinations for the maintainer(s)
	Maintainers []*Maintainer `protobuf:"bytes,7,rep,name=maintainers" json:"maintainers,omitempty"`
	// The name of the template engine to use. Defaults to 'gotpl'.
	Engine string `protobuf:"bytes,8,opt,name=engine" json:"engine,omitempty"`
	// The URL to an icon file.
	Icon string `protobuf:"bytes,9,opt,name=icon" json:"icon,omitempty"`
	// The API Version of this chart.
	ApiVersion string `protobuf:"bytes,10,opt,name=apiVersion" json:"apiVersion,omitempty"`
	// The condition to check to enable chart
	Condition string `protobuf:"bytes,11,opt,name=condition" json:"condition,omitempty"`
	// The tags to check to enable chart
	Tags string `protobuf:"bytes,12,opt,name=tags" json:"tags,omitempty"`
	// The version of the application enclosed inside of this chart.
	AppVersion string `protobuf:"bytes,13,opt,name=appVersion" json:"appVersion,omitempty"`
	// Whether or not this chart is deprecated
	Deprecated bool `protobuf:"varint,14,opt,name=deprecated" json:"deprecated,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor2, []int{1} }

func (m *Metadata) GetMaintainers() []*Maintainer {
	if m != nil {
		return m.Maintainers
	}
	return nil
}

func init() {
	proto.RegisterType((*Maintainer)(nil), "hapi.chart.Maintainer")
	proto.RegisterType((*Metadata)(nil), "hapi.chart.Metadata")
	proto.RegisterEnum("hapi.chart.Metadata_Engine", Metadata_Engine_name, Metadata_Engine_value)
}

func init() { proto.RegisterFile("hapi/chart/metadata.proto", fileDescriptor2) }

var fileDescriptor2 = []byte{
	// 339 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x92, 0x4f, 0x4b, 0xeb, 0x40,
	0x14, 0xc5, 0x5f, 0x5e, 0x9a, 0x7f, 0x37, 0xef, 0x3d, 0xca, 0xf0, 0x28, 0xa3, 0x88, 0x84, 0xae,
	0xb2, 0x4a, 0x41, 0x41, 0x5c, 0x0b, 0xe2, 0x42, 0xdb, 0x4a, 0xf0, 0x0f, 0xb8, 0x1b, 0x93, 0x4b,
	0x3b, 0x68, 0x66, 0xc2, 0xcc, 0xa8, 0xf8, 0xe5, 0xfc, 0x6c, 0x32, 0x93, 0xa4, 0xcd, 0xc2, 0xdd,
	0x3d, 0xe7, 0x97, 0x7b, 0x92, 0x73, 0x09, 0x1c, 0x6c, 0x59, 0xcb, 0x17, 0xd5, 0x96, 0x29, 0xb3,
	0x68, 0xd0, 0xb0, 0x9a, 0x19, 0x56, 0xb4, 0x4a, 0x1a, 0x49, 0xc0, 0xa2, 0xc2, 0xa1, 0xf9, 0x19,
	0xc0, 0x92, 0x71, 0x61, 0x18, 0x17, 0xa8, 0x08, 0x81, 0x89, 0x60, 0x0d, 0x52, 0x2f, 0xf3, 0xf2,
	0xa4, 0x74, 0x33, 0xf9, 0x0f, 0x01, 0x36, 0x8c, 0xbf, 0xd2, 0xdf, 0xce, 0xec, 0xc4, 0xfc, 0xcb,
	0x87, 0x78, 0xd9, 0xc7, 0xfe, 0xb8, 0x46, 0x60, 0xb2, 0x95, 0x0d, 0xf6, 0x5b, 0x6e, 0x26, 0x14,
	0x22, 0x2d, 0xdf, 0x54, 0x85, 0x9a, 0xfa, 0x99, 0x9f, 0x27, 0xe5, 0x20, 0x2d, 0x79, 0x47, 0xa5,
	0xb9, 0x14, 0x74, 0xe2, 0x16, 0x06, 0x49, 0x32, 0x48, 0x6b, 0xd4, 0x95, 0xe2, 0xad, 0xb1, 0x34,
	0x70, 0x74, 0x6c, 0x91, 0x43, 0x88, 0x5f, 0xf0, 0xf3, 0x43, 0xaa, 0x5a, 0xd3, 0xd0, 0xc5, 0xee,
	0x34, 0x39, 0x87, 0xb4, 0xd9, 0xd5, 0xd3, 0x34, 0xca, 0xfc, 0x3c, 0x3d, 0x99, 0x15, 0xfb, 0x03,
	0x14, 0xfb, 0xf6, 0xe5, 0xf8, 0x51, 0x32, 0x83, 0x10, 0xc5, 0x86, 0x0b, 0xa4, 0xb1, 0x7b, 0x65,
	0xaf, 0x6c, 0x2f, 0x5e, 0x49, 0x41, 0x93, 0xae, 0x97, 0x9d, 0xc9, 0x31, 0x00, 0x6b, 0xf9, 0x43,
	0x5f, 0x00, 0x1c, 0x19, 0x39, 0xe4, 0x08, 0x92, 0x4a, 0x8a, 0x9a, 0xbb, 0x06, 0xa9, 0xc3, 0x7b,
	0xc3, 0x26, 0x1a, 0xb6, 0xd1, 0xf4, 0x4f, 0x97, 0x68, 0xe7, 0x2e, 0xb1, 0x1d, 0x12, 0xff, 0x0e,
	0x89, 0x83, 0x63, 0x79, 0x8d, 0xad, 0xc2, 0x8a, 0x19, 0xac, 0xe9, 0xbf, 0xcc, 0xcb, 0xe3, 0x72,
	0xe4, 0xcc, 0x33, 0x08, 0x2f, 0xbb, 0xef, 0x4d, 0x21, 0xba, 0x5f, 0x5d, 0xaf, 0xd6, 0x8f, 0xab,
	0xe9, 0x2f, 0x92, 0x40, 0x70, 0xb5, 0xbe, 0xbb, 0xbd, 0x99, 0x7a, 0x17, 0xd1, 0x53, 0xe0, 0x0e,
	0xf0, 0x1c, 0xba, 0x9f, 0xe2, 0xf4, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x08, 0xf3, 0xcc, 0x66, 0x31,
	0x02, 0x00, 0x00,
}
