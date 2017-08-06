// Code generated by protoc-gen-gopherjs. DO NOT EDIT.
// source: multi/multi2.proto

package multitest

import js "github.com/gopherjs/gopherjs/js"
import jspb "github.com/johanbrandhorst/protobuf/jspb"

// This is a compile-time assertion to ensure that this generated file
// is compatible with the jspb package it is being compiled against.
const _ = jspb.JspbPackageIsVersion1

type Multi2_Color int

const (
	Multi2_BLUE  Multi2_Color = 0
	Multi2_GREEN Multi2_Color = 1
	Multi2_RED   Multi2_Color = 2
)

var Multi2_Color_name = map[int]string{
	0: "BLUE",
	1: "GREEN",
	2: "RED",
}
var Multi2_Color_value = map[string]int{
	"BLUE":  0,
	"GREEN": 1,
	"RED":   2,
}

func (x Multi2_Color) String() string {
	return Multi2_Color_name[int(x)]
}

type Multi2 struct {
	*js.Object
}

// GetRequiredValue gets the RequiredValue of the Multi2.
func (m *Multi2) GetRequiredValue() (x int32) {
	if m == nil {
		return x
	}
	return int32(m.Call("getRequiredValue").Int())
}

// SetRequiredValue sets the RequiredValue of the Multi2.
func (m *Multi2) SetRequiredValue(v int32) {
	m.Call("setRequiredValue", v)
}

// GetColor gets the Color of the Multi2.
func (m *Multi2) GetColor() (x Multi2_Color) {
	if m == nil {
		return x
	}
	return Multi2_Color(m.Call("getColor").Int())
}

// SetColor sets the Color of the Multi2.
func (m *Multi2) SetColor(v Multi2_Color) {
	m.Call("setColor", v)
}

// New creates a new Multi2.
func (m *Multi2) New(requiredValue int32, color Multi2_Color) *Multi2 {
	m = &Multi2{
		Object: js.Global.Get("proto").Get("multitest").Get("Multi2").New([]interface{}{
			requiredValue,
			color,
		}),
	}

	return m
}

// Serialize marshals Multi2 to a slice of bytes.
func (m *Multi2) Serialize() []byte {
	return jspb.Serialize(m)
}

// Deserialize unmarshals a Multi2 from a slice of bytes.
func (m *Multi2) Deserialize(rawBytes []byte) (*Multi2, error) {
	obj, err := jspb.Deserialize(js.Global.Get("proto").Get("multitest").Get("Multi2"), rawBytes)
	if err != nil {
		return nil, err
	}

	return &Multi2{
		Object: obj,
	}, nil
}
