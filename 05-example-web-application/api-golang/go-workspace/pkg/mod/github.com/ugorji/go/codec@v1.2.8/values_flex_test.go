// comment this out // // + build testing

// Copyright (c) 2012-2020 Ugorji Nwoke. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

package codec

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

const teststrucflexChanCap = 64

// This file contains values used by tests alone.
// This is where we may try out different things,
// that other engines may not support or may barf upon
// e.g. custom extensions for wrapped types, maps with non-string keys, etc.

// some funky types to test codecgen

type codecgenA struct {
	ZZ []byte
}
type codecgenB struct {
	AA codecgenA
}
type codecgenC struct {
	_struct struct{} `codec:",omitempty"`
	BB      codecgenB
}

type TestCodecgenG struct {
	TestCodecgenG int
}
type codecgenH struct {
	TestCodecgenG
}
type codecgenI struct {
	codecgenH
}

type codecgenK struct {
	X int
	Y string
}
type codecgenL struct {
	X int
	Y uint32
}
type codecgenM struct {
	codecgenK
	codecgenL
}

// some types to test struct keytype

type testStrucKeyTypeT0 struct {
	_struct struct{}
	F       int
}
type testStrucKeyTypeT1 struct {
	_struct struct{} `codec:",string"`
	F       int      `codec:"FFFF"`
}
type testStrucKeyTypeT2 struct {
	_struct struct{} `codec:",int"`
	F       int      `codec:"-1"`
}
type testStrucKeyTypeT3 struct {
	_struct struct{} `codec:",uint"`
	F       int      `codec:"1"`
}
type testStrucKeyTypeT4 struct {
	_struct struct{} `codec:",float"`
	F       int      `codec:"2.5"`
}

// Some unused types just stored here

type Bbool bool
type Aarray [1]string
type Sstring string
type Sstructsmall struct {
	A int
}

type Sstructbig struct {
	_struct struct{}
	A       int
	B       bool
	c       string
	// Sval Sstruct
	Ssmallptr *Sstructsmall
	Ssmall    Sstructsmall
	Sptr      *Sstructbig
}

type SstructbigToArray struct {
	_struct struct{} `codec:",toarray"`
	A       int
	B       bool
	c       string
	// Sval Sstruct
	Ssmallptr *Sstructsmall
	Ssmall    Sstructsmall
	Sptr      *Sstructbig
}

// small struct for testing that codecgen works for unexported types
type tLowerFirstLetter struct {
	I int
	u uint64
	S string
	b []byte
}

// Some wrapped types which are used as extensions
type wrapInt64 int64
type wrapUint8 uint8
type wrapBytes []byte

// some types that define how to marshal themselves
type testMarshalAsJSON bool
type testMarshalAsBinary []byte
type testMarshalAsText string

// some types used for extensions
type testUintToBytes uint32

func (x testMarshalAsJSON) MarshalJSON() (data []byte, err error) {
	if x {
		data = []byte("true")
	} else {
		data = []byte("false")
	}
	return
}
func (x *testMarshalAsJSON) UnmarshalJSON(data []byte) (err error) {
	switch string(data) {
	case "true":
		*x = true
	case "false":
		*x = false
	default:
		err = fmt.Errorf("testMarshalAsJSON failed to decode as bool: %s", data)
	}
	return
}

func (x testMarshalAsBinary) MarshalBinary() (data []byte, err error) {
	data = []byte(x)
	return
}
func (x *testMarshalAsBinary) UnmarshalBinary(data []byte) (err error) {
	*x = data
	return
}

func (x testMarshalAsText) MarshalText() (text []byte, err error) {
	text = []byte(x)
	return
}
func (x *testMarshalAsText) UnmarshalText(text []byte) (err error) {
	*x = testMarshalAsText(string(text))
	return
}

type AnonInTestStrucIntf struct {
	Islice []interface{}
	Ms     map[string]interface{}
	Nintf  interface{} //don't set this, so we can test for nil
	T      time.Time
	Tptr   *time.Time
}

type missingFielderT1 struct {
	S string
	B bool
	f float64
	i int64
}

func (t *missingFielderT1) CodecMissingField(field []byte, value interface{}) bool {
	switch string(field) {
	case "F":
		t.f = value.(float64)
	case "I":
		t.i = value.(int64)
	default:
		return false
	}
	return true
}

func (t *missingFielderT1) CodecMissingFields() map[string]interface{} {
	return map[string]interface{}{"F": t.f, "I": t.i}
}

type missingFielderT11 struct {
	s1 string
	S2 string
}

func (t *missingFielderT11) CodecMissingField(field []byte, value interface{}) bool {
	if "s1" == string(field) {
		t.s1 = value.(string)
		return true
	}
	return false
}

// missingFielderT11 implements CodecMissingFields on the value (not pointer)
func (t missingFielderT11) CodecMissingFields() map[string]interface{} {
	return map[string]interface{}{"s1": t.s1}
}

type missingFielderT2 struct {
	S string
	B bool
	F float64
	I int64
}

type testSelfExtHelper struct {
	S string
	I int64
	B bool
}

type TestSelfExtImpl struct {
	testSelfExtHelper
}

type TestSelfExtImpl2 struct {
	M string
	O bool
}

type TestTwoNakedInterfaces struct {
	A interface{}
	B interface{}
}

var testWRepeated512 wrapBytes
var testStrucTime = time.Date(2012, 2, 2, 2, 2, 2, 2000, time.UTC).UTC()

func init() {
	var testARepeated512 [512]byte
	for i := range testARepeated512 {
		testARepeated512[i] = 'A'
	}
	testWRepeated512 = wrapBytes(testARepeated512[:])
}

type TestStrucFlex struct {
	_struct struct{} `codec:",omitempty"` //set omitempty for every field
	TestStrucCommon

	Chstr chan string

	Mis     map[int]string
	Mbu64   map[bool]struct{}
	Mu8e    map[byte]struct{}
	Mu8u64  map[byte]stringUint64T
	Msp2ss  map[*string][]string
	Mip2ss  map[*uint64][]string
	Ms2misu map[string]map[uint64]stringUint64T
	Miwu64s map[int]wrapUint64Slice
	Mfwss   map[float64]wrapStringSlice
	Mf32wss map[float32]wrapStringSlice
	Mui2wss map[uint64]wrapStringSlice

	// DecodeNaked bombs because stringUint64T is decoded as a map,
	// and a map cannot be the key type of a map.
	// Ensure this is set to nil if decoding into a nil interface{}.
	Msu2wss map[stringUint64T]wrapStringSlice

	Ci64       wrapInt64
	Swrapbytes []wrapBytes
	Swrapuint8 []wrapUint8

	ArrStrUi64T [4]stringUint64T

	Ui64array      [4]uint64
	Ui64slicearray []*[4]uint64

	SintfAarray []interface{}

	// Ensure this is set to nil if decoding into a nil interface{}.
	MstrUi64TSelf map[stringUint64T]*stringUint64T

	Ttime    time.Time
	Ttimeptr *time.Time

	// make this a ptr, so that it could be set or not.
	// for comparison (e.g. with msgp), give it a struct tag (so it is not inlined),
	// make this one omitempty (so it is excluded if nil).
	*AnonInTestStrucIntf `json:",omitempty"`

	M          map[interface{}]interface{} `json:"-"`
	Msu        map[wrapString]interface{}
	Mtsptr     map[string]*TestStrucFlex
	Mts        map[string]TestStrucFlex
	Its        []*TestStrucFlex
	Nteststruc *TestStrucFlex

	MarJ testMarshalAsJSON
	MarT testMarshalAsText
	MarB testMarshalAsBinary

	XuintToBytes testUintToBytes

	Ffunc       func() error // expect this to be skipped/ignored
	Bboolignore bool         `codec:"-"` // expect this to be skipped/ignored

	Cmplx64  complex64
	Cmplx128 complex128
}

func newTestStrucFlex(depth, n int, bench, useInterface, useStringKeyOnly bool) (ts *TestStrucFlex) {
	ts = &TestStrucFlex{
		Chstr: make(chan string, teststrucflexChanCap),

		Miwu64s: map[int]wrapUint64Slice{
			5: []wrapUint64{1, 2, 3, 4, 5},
			3: []wrapUint64{1, 2, 3},
		},

		Mf32wss: map[float32]wrapStringSlice{
			5.0: []wrapString{"1.0", "2.0", "3.0", "4.0", "5.0"},
			3.0: []wrapString{"1.0", "2.0", "3.0"},
		},

		Mui2wss: map[uint64]wrapStringSlice{
			5: []wrapString{"1.0", "2.0", "3.0", "4.0", "5.0"},
			3: []wrapString{"1.0", "2.0", "3.0"},
		},

		Mfwss: map[float64]wrapStringSlice{
			5.0: []wrapString{"1.0", "2.0", "3.0", "4.0", "5.0"},
			3.0: []wrapString{"1.0", "2.0", "3.0"},
		},

		// DecodeNaked bombs here, because the stringUint64T is decoded as a map,
		// and a map cannot be the key type of a map.
		// Ensure this is set to nil if decoding into a nil interface{}.
		Msu2wss: map[stringUint64T]wrapStringSlice{
			stringUint64T{"5", 5}: []wrapString{"1", "2", "3", "4", "5"},
			stringUint64T{"3", 3}: []wrapString{"1", "2", "3"},
		},

		Mis: map[int]string{
			1:   "one",
			22:  "twenty two",
			-44: "minus forty four",
		},

		Ms2misu: map[string]map[uint64]stringUint64T{
			"1":   {1: {"11", 11}},
			"22":  {1: {"2222", 2222}},
			"333": {1: {"333333", 333333}},
		},

		Mbu64:  map[bool]struct{}{false: {}, true: {}},
		Mu8e:   map[byte]struct{}{1: {}, 2: {}, 3: {}, 4: {}},
		Mu8u64: make(map[byte]stringUint64T),
		Mip2ss: make(map[*uint64][]string),
		Msp2ss: make(map[*string][]string),
		M:      make(map[interface{}]interface{}),
		Msu:    make(map[wrapString]interface{}),

		Ci64: -22,
		Swrapbytes: []wrapBytes{ // lengths of 1, 2, 4, 8, 16, 32, 64, 128, 256,
			testWRepeated512[:1],
			testWRepeated512[:2],
			testWRepeated512[:4],
			testWRepeated512[:8],
			testWRepeated512[:16],
			testWRepeated512[:32],
			testWRepeated512[:64],
			testWRepeated512[:128],
			testWRepeated512[:256],
			testWRepeated512[:512],
		},
		Swrapuint8: []wrapUint8{
			'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J',
		},
		Ui64array:     [4]uint64{4, 16, 64, 256},
		ArrStrUi64T:   [4]stringUint64T{{"4", 4}, {"3", 3}, {"2", 2}, {"1", 1}},
		SintfAarray:   []interface{}{Aarray{"s"}},
		MstrUi64TSelf: make(map[stringUint64T]*stringUint64T, numStrUi64T),

		Ttime:    testStrucTime,
		Ttimeptr: &testStrucTime,

		MarJ: true,
		MarT: "hello string",
		MarB: []byte("hello bytes"),

		XuintToBytes: 16,

		Cmplx64:  complex(16, 0),
		Cmplx128: complex(1616, 0),
	}

	var strslice []string
	for i := uint64(0); i < numStrUi64T; i++ {
		s := strings.Repeat(strconv.FormatUint(i, 10), 4)
		ss := stringUint64T{S: s, U: i}
		strslice = append(strslice, s)
		// Ensure this is set to nil if decoding into a nil interface{}.
		ts.MstrUi64TSelf[ss] = &ss
		ts.Mu8u64[s[0]] = ss
		ts.Mip2ss[&i] = strslice
		ts.Msp2ss[&s] = strslice
		// add some other values of maps and pointers into M
		ts.M[s] = strslice
		// cannot use this, as converting stringUint64T to interface{} returns map,
		// DecodeNaked does this, causing "hash of unhashable value" as some maps cannot be map keys
		// ts.M[ss] = &ss
		ts.Msu[wrapString(s)] = &ss
		s = s + "-map"
		ss.S = s
		ts.M[s] = map[string]string{s: s}
		ts.Msu[wrapString(s)] = map[string]string{s: s}
	}

	numChanSend := cap(ts.Chstr) / 4 // 8
	for i := 0; i < numChanSend; i++ {
		ts.Chstr <- strings.Repeat("A", i+1)
	}

	ts.Ui64slicearray = []*[4]uint64{&ts.Ui64array, &ts.Ui64array}

	if useInterface {
		ts.AnonInTestStrucIntf = &AnonInTestStrucIntf{
			Islice: []interface{}{strRpt(n, "true"), true, strRpt(n, "no"), false, uint64(288), float64(0.4)},
			Ms: map[string]interface{}{
				strRpt(n, "true"):     strRpt(n, "true"),
				strRpt(n, "int64(9)"): false,
			},
			T:    testStrucTime,
			Tptr: &testStrucTime,
		}
	}

	populateTestStrucCommon(&ts.TestStrucCommon, n, bench, useInterface, useStringKeyOnly)
	if depth > 0 {
		depth--
		if ts.Mtsptr == nil {
			ts.Mtsptr = make(map[string]*TestStrucFlex)
		}
		if ts.Mts == nil {
			ts.Mts = make(map[string]TestStrucFlex)
		}
		ts.Mtsptr["0"] = newTestStrucFlex(depth, n, bench, useInterface, useStringKeyOnly)
		ts.Mts["0"] = *(ts.Mtsptr["0"])
		ts.Its = append(ts.Its, ts.Mtsptr["0"])
	}
	return
}
