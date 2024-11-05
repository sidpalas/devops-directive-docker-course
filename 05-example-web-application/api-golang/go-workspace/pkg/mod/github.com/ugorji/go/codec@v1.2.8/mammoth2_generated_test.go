//go:build !codec.notmammoth
// +build !codec.notmammoth

// Copyright (c) 2012-2020 Ugorji Nwoke. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

// Code generated from mammoth2-test.go.tmpl - DO NOT EDIT.

package codec

// Increase codecoverage by covering all the codecgen paths, in fast-path and gen-helper.go....
//
// Note: even though this is built based on fast-path and gen-helper, we will run these tests
// in all modes, including notfastpath, etc.
//
// Add test file for creating a mammoth generated file as _mammoth_generated.go
//  - generate a second mammoth files in a different file: mammoth2_generated_test.go
//    mammoth-test.go.tmpl will do this
//  - run codecgen on it, into mammoth2_codecgen_generated_test.go (no build tags)
//  - as part of TestMammoth, run it also
//  - this will cover all the codecgen, gen-helper, etc in one full run
//  - check in mammoth* files into github also
//
// Now, add some types:
//  - some that implement BinaryMarshal, TextMarshal, JSONMarshal, and one that implements none of it
//  - create a wrapper type that includes TestMammoth2, with it in slices, and maps, and the custom types
//  - this wrapper object is what we work encode/decode (so that the codecgen methods are called)

// import "encoding/binary"

import "fmt"

type TestMammoth2 struct {
	FIntf       interface{}
	FptrIntf    *interface{}
	FString     string
	FptrString  *string
	FBytes      []byte
	FptrBytes   *[]byte
	FFloat32    float32
	FptrFloat32 *float32
	FFloat64    float64
	FptrFloat64 *float64
	FUint       uint
	FptrUint    *uint
	FUint8      uint8
	FptrUint8   *uint8
	FUint16     uint16
	FptrUint16  *uint16
	FUint32     uint32
	FptrUint32  *uint32
	FUint64     uint64
	FptrUint64  *uint64
	FUintptr    uintptr
	FptrUintptr *uintptr
	FInt        int
	FptrInt     *int
	FInt8       int8
	FptrInt8    *int8
	FInt16      int16
	FptrInt16   *int16
	FInt32      int32
	FptrInt32   *int32
	FInt64      int64
	FptrInt64   *int64
	FBool       bool
	FptrBool    *bool

	FSliceIntf       []interface{}
	FptrSliceIntf    *[]interface{}
	FSliceString     []string
	FptrSliceString  *[]string
	FSliceBytes      [][]byte
	FptrSliceBytes   *[][]byte
	FSliceFloat32    []float32
	FptrSliceFloat32 *[]float32
	FSliceFloat64    []float64
	FptrSliceFloat64 *[]float64
	FSliceUint8      []uint8
	FptrSliceUint8   *[]uint8
	FSliceUint64     []uint64
	FptrSliceUint64  *[]uint64
	FSliceInt        []int
	FptrSliceInt     *[]int
	FSliceInt32      []int32
	FptrSliceInt32   *[]int32
	FSliceInt64      []int64
	FptrSliceInt64   *[]int64
	FSliceBool       []bool
	FptrSliceBool    *[]bool

	FMapStringIntf       map[string]interface{}
	FptrMapStringIntf    *map[string]interface{}
	FMapStringString     map[string]string
	FptrMapStringString  *map[string]string
	FMapStringBytes      map[string][]byte
	FptrMapStringBytes   *map[string][]byte
	FMapStringUint8      map[string]uint8
	FptrMapStringUint8   *map[string]uint8
	FMapStringUint64     map[string]uint64
	FptrMapStringUint64  *map[string]uint64
	FMapStringInt        map[string]int
	FptrMapStringInt     *map[string]int
	FMapStringInt32      map[string]int32
	FptrMapStringInt32   *map[string]int32
	FMapStringFloat64    map[string]float64
	FptrMapStringFloat64 *map[string]float64
	FMapStringBool       map[string]bool
	FptrMapStringBool    *map[string]bool
	FMapUint8Intf        map[uint8]interface{}
	FptrMapUint8Intf     *map[uint8]interface{}
	FMapUint8String      map[uint8]string
	FptrMapUint8String   *map[uint8]string
	FMapUint8Bytes       map[uint8][]byte
	FptrMapUint8Bytes    *map[uint8][]byte
	FMapUint8Uint8       map[uint8]uint8
	FptrMapUint8Uint8    *map[uint8]uint8
	FMapUint8Uint64      map[uint8]uint64
	FptrMapUint8Uint64   *map[uint8]uint64
	FMapUint8Int         map[uint8]int
	FptrMapUint8Int      *map[uint8]int
	FMapUint8Int32       map[uint8]int32
	FptrMapUint8Int32    *map[uint8]int32
	FMapUint8Float64     map[uint8]float64
	FptrMapUint8Float64  *map[uint8]float64
	FMapUint8Bool        map[uint8]bool
	FptrMapUint8Bool     *map[uint8]bool
	FMapUint64Intf       map[uint64]interface{}
	FptrMapUint64Intf    *map[uint64]interface{}
	FMapUint64String     map[uint64]string
	FptrMapUint64String  *map[uint64]string
	FMapUint64Bytes      map[uint64][]byte
	FptrMapUint64Bytes   *map[uint64][]byte
	FMapUint64Uint8      map[uint64]uint8
	FptrMapUint64Uint8   *map[uint64]uint8
	FMapUint64Uint64     map[uint64]uint64
	FptrMapUint64Uint64  *map[uint64]uint64
	FMapUint64Int        map[uint64]int
	FptrMapUint64Int     *map[uint64]int
	FMapUint64Int32      map[uint64]int32
	FptrMapUint64Int32   *map[uint64]int32
	FMapUint64Float64    map[uint64]float64
	FptrMapUint64Float64 *map[uint64]float64
	FMapUint64Bool       map[uint64]bool
	FptrMapUint64Bool    *map[uint64]bool
	FMapIntIntf          map[int]interface{}
	FptrMapIntIntf       *map[int]interface{}
	FMapIntString        map[int]string
	FptrMapIntString     *map[int]string
	FMapIntBytes         map[int][]byte
	FptrMapIntBytes      *map[int][]byte
	FMapIntUint8         map[int]uint8
	FptrMapIntUint8      *map[int]uint8
	FMapIntUint64        map[int]uint64
	FptrMapIntUint64     *map[int]uint64
	FMapIntInt           map[int]int
	FptrMapIntInt        *map[int]int
	FMapIntInt32         map[int]int32
	FptrMapIntInt32      *map[int]int32
	FMapIntFloat64       map[int]float64
	FptrMapIntFloat64    *map[int]float64
	FMapIntBool          map[int]bool
	FptrMapIntBool       *map[int]bool
	FMapInt32Intf        map[int32]interface{}
	FptrMapInt32Intf     *map[int32]interface{}
	FMapInt32String      map[int32]string
	FptrMapInt32String   *map[int32]string
	FMapInt32Bytes       map[int32][]byte
	FptrMapInt32Bytes    *map[int32][]byte
	FMapInt32Uint8       map[int32]uint8
	FptrMapInt32Uint8    *map[int32]uint8
	FMapInt32Uint64      map[int32]uint64
	FptrMapInt32Uint64   *map[int32]uint64
	FMapInt32Int         map[int32]int
	FptrMapInt32Int      *map[int32]int
	FMapInt32Int32       map[int32]int32
	FptrMapInt32Int32    *map[int32]int32
	FMapInt32Float64     map[int32]float64
	FptrMapInt32Float64  *map[int32]float64
	FMapInt32Bool        map[int32]bool
	FptrMapInt32Bool     *map[int32]bool
}

// -----------

type testMammoth2Binary uint64

func (x testMammoth2Binary) MarshalBinary() (data []byte, err error) {
	data = make([]byte, 8)
	bigenstd.PutUint64(data, uint64(x))
	return
}
func (x *testMammoth2Binary) UnmarshalBinary(data []byte) (err error) {
	*x = testMammoth2Binary(bigenstd.Uint64(data))
	return
}

type testMammoth2Text uint64

func (x testMammoth2Text) MarshalText() (data []byte, err error) {
	data = []byte(fmt.Sprintf("%b", uint64(x)))
	return
}
func (x *testMammoth2Text) UnmarshalText(data []byte) (err error) {
	_, err = fmt.Sscanf(string(data), "%b", (*uint64)(x))
	return
}

type testMammoth2Json uint64

func (x testMammoth2Json) MarshalJSON() (data []byte, err error) {
	data = []byte(fmt.Sprintf("%v", uint64(x)))
	return
}
func (x *testMammoth2Json) UnmarshalJSON(data []byte) (err error) {
	_, err = fmt.Sscanf(string(data), "%v", (*uint64)(x))
	return
}

type testMammoth2Basic [4]uint64

type TestMammoth2Wrapper struct {
	V TestMammoth2
	T testMammoth2Text
	B testMammoth2Binary
	J testMammoth2Json
	C testMammoth2Basic
	M map[testMammoth2Basic]TestMammoth2
	L []TestMammoth2
	A [4]int64

	Tcomplex128 complex128
	Tcomplex64  complex64
	Tbytes      []uint8
	Tpbytes     *[]uint8
}
