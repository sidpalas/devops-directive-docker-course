// Copyright (c) 2012-2020 Ugorji Nwoke. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

package codec

// NAMING CONVENTION FOR TESTS
//
// function and variable/const names here fit a simple naming convention
//
// - each test is a doTestXXX(...). TestXXX calls doTestXXX.
// - testXXX are helper functions.
// - doTestXXX are test functions that take an extra arg of a handle.
// - testXXX variables and constants are only used in tests.
// - shared functions/vars/consts are testShared...
// - fnBenchmarkXX and fnTestXXX can be used as needed.
//
// - each TestXXX must only call testSetup once.
// - each test with a prefix __doTest is a dependent helper function,
//   which MUST not call testSetup itself.
//   doTestXXX or TestXXX may call it.

// if we are testing in parallel,
// then we don't want to share much state: testBytesFreeList, etc.

import (
	"bufio"
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"math/rand"
	"net"
	"net/rpc"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func init() {
	testPreInitFns = append(testPreInitFns, testInit)
}

const testRecoverPanicToErr = !debugging

// tests which check for errors will fail if testRecoverPanicToErr=false (debugging=true).
// Consequently, skip them.
var testSkipIfNotRecoverPanicToErrMsg = "tests checks for errors, and testRecoverPanicToErr=false"

func testPanicToErr(h errDecorator, err *error) {
	// Note: This method MUST be called directly from defer i.e. defer testPanicToErr ...
	// else it seems the recover is not fully handled
	if x := recover(); x != nil {
		panicValToErr(h, x, err)
	}
}

var testBytesFreeList bytesFreelist

type testCustomStringT string

// make these mapbyslice
type testMbsT []interface{}
type testMbsArr0T [0]interface{}
type testMbsArr4T [4]interface{}
type testMbsArr5T [5]interface{}
type testMbsCustStrT []testCustomStringT

func (testMbsT) MapBySlice()        {}
func (*testMbsArr0T) MapBySlice()   {}
func (*testMbsArr4T) MapBySlice()   {}
func (testMbsArr5T) MapBySlice()    {}
func (testMbsCustStrT) MapBySlice() {}

// type testSelferRecur struct{}

// func (s *testSelferRecur) CodecEncodeSelf(e *Encoder) {
// 	e.MustEncode(s)
// }
// func (s *testSelferRecur) CodecDecodeSelf(d *Decoder) {
// 	d.MustDecode(s)
// }

type testIntfMapI interface {
	GetIntfMapV() string
}

type testIntfMapT1 struct {
	IntfMapV string
}

func (x *testIntfMapT1) GetIntfMapV() string { return x.IntfMapV }

type testIntfMapT2 struct {
	IntfMapV string
}

func (x testIntfMapT2) GetIntfMapV() string { return x.IntfMapV }

type testMissingFieldsMap struct {
	m map[string]interface{}
}

func (mf *testMissingFieldsMap) CodecMissingField(field []byte, value interface{}) bool {
	if mf.m == nil {
		mf.m = map[string]interface{}{}
	}

	(mf.m)[string(field)] = value

	return true
}

func (mf *testMissingFieldsMap) CodecMissingFields() map[string]interface{} {
	return mf.m
}

var _ MissingFielder = (*testMissingFieldsMap)(nil)

var testErrWriterErr = errors.New("testErrWriterErr")

type testErrWriter struct{}

func (x *testErrWriter) Write(p []byte) (int, error) {
	return 0, testErrWriterErr
}

// ----

type testVerifyFlag uint8

const (
	_ testVerifyFlag = 1 << iota
	testVerifyMapTypeSame
	testVerifyMapTypeStrIntf
	testVerifyMapTypeIntfIntf
	// testVerifySliceIntf
	testVerifyForPython
	testVerifyDoNil
	testVerifyTimeAsInteger
)

func (f testVerifyFlag) isset(v testVerifyFlag) bool {
	return f&v == v
}

// const testSkipRPCTests = false

var (
	testTableNumPrimitives int
	testTableIdxTime       int
	testTableNumMaps       int

	// set this when running using bufio, etc
	testSkipRPCTests = false

	testSkipRPCTestsMsg = "testSkipRPCTests=true"
)

var (
	skipVerifyVal interface{} = &(struct{}{})

	testMapStrIntfTyp = reflect.TypeOf(map[string]interface{}(nil))

	// For Go Time, do not use a descriptive timezone.
	// It's unnecessary, and makes it harder to do a reflect.DeepEqual.
	// The Offset already tells what the offset should be, if not on UTC and unknown zone name.
	timeLoc        = time.FixedZone("", -8*60*60) // UTC-08:00 //time.UTC-8
	timeToCompare1 = time.Date(2012, 2, 2, 2, 2, 2, 2000, timeLoc).UTC()
	timeToCompare2 = time.Date(1900, 2, 2, 2, 2, 2, 2000, timeLoc).UTC()
	timeToCompare3 = time.Unix(0, 270).UTC() // use value that must be encoded as uint64 for nanoseconds (for cbor/msgpack comparison)
	//timeToCompare4 = time.Time{}.UTC() // does not work well with simple cbor time encoding (overflow)
	timeToCompare4 = time.Unix(-2013855848, 4223).UTC()

	table []interface{} // main items we encode
	// will encode a float32 as float64, or large int as uint

	testRpcServer = rpc.NewServer()
	testRpcInt    = new(TestRpcInt)
)

func init() {
	testRpcServer.Register(testRpcInt)
}

var wrapInt64Typ = reflect.TypeOf(wrapInt64(0))
var wrapBytesTyp = reflect.TypeOf(wrapBytes(nil))

var testUintToBytesTyp = reflect.TypeOf(testUintToBytes(0))
var testSelfExtTyp = reflect.TypeOf((*TestSelfExtImpl)(nil)).Elem()
var testSelfExt2Typ = reflect.TypeOf((*TestSelfExtImpl2)(nil)).Elem()

func testByteBuf(in []byte) *bytes.Buffer {
	return bytes.NewBuffer(in)
}

type TestABC struct {
	A, B, C string
}

func (x *TestABC) MarshalBinary() ([]byte, error) {
	return []byte(fmt.Sprintf("%s %s %s", x.A, x.B, x.C)), nil
}
func (x *TestABC) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%s %s %s", x.A, x.B, x.C)), nil
}
func (x *TestABC) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s %s %s"`, x.A, x.B, x.C)), nil
}

func (x *TestABC) UnmarshalBinary(data []byte) (err error) {
	ss := strings.Split(string(data), " ")
	x.A, x.B, x.C = ss[0], ss[1], ss[2]
	return
}
func (x *TestABC) UnmarshalText(data []byte) (err error) {
	return x.UnmarshalBinary(data)
}
func (x *TestABC) UnmarshalJSON(data []byte) (err error) {
	return x.UnmarshalBinary(data[1 : len(data)-1])
}

type TestABC2 struct {
	A, B, C string
}

func (x TestABC2) MarshalText() ([]byte, error) {
	return []byte(fmt.Sprintf("%s %s %s", x.A, x.B, x.C)), nil
}
func (x *TestABC2) UnmarshalText(data []byte) (err error) {
	ss := strings.Split(string(data), " ")
	x.A, x.B, x.C = ss[0], ss[1], ss[2]
	return
	// _, err = fmt.Sscanf(string(data), "%s %s %s", &x.A, &x.B, &x.C)
}

type TestSimplish struct {
	Ii int
	Ss string
	Ar [2]*TestSimplish
	Sl []*TestSimplish
	Mm map[string]*TestSimplish
}

type TestRpcABC struct {
	A, B, C string
}

type TestRpcInt struct {
	i int64
}

func (r *TestRpcInt) Update(n int, res *int) error {
	atomic.StoreInt64(&r.i, int64(n))
	*res = n
	return nil
}
func (r *TestRpcInt) Square(ignore int, res *int) error {
	i := int(atomic.LoadInt64(&r.i))
	*res = i * i
	return nil
}
func (r *TestRpcInt) Mult(n int, res *int) error {
	*res = int(atomic.LoadInt64(&r.i)) * n
	return nil
}
func (r *TestRpcInt) EchoStruct(arg TestRpcABC, res *string) error {
	*res = fmt.Sprintf("%#v", arg)
	return nil
}
func (r *TestRpcInt) Echo123(args []string, res *string) error {
	*res = fmt.Sprintf("%#v", args)
	return nil
}

type TestRawValue struct {
	R Raw
	I int
}

// ----

type testUnixNanoTimeExt struct {
	// keep timestamp here, so that do not incur interface-conversion costs
	// ts int64
}

func (x *testUnixNanoTimeExt) WriteExt(v interface{}) []byte {
	v2 := v.(*time.Time)
	bs := make([]byte, 8)
	bigenstd.PutUint64(bs, uint64(v2.UnixNano()))
	return bs
}
func (x *testUnixNanoTimeExt) ReadExt(v interface{}, bs []byte) {
	v2 := v.(*time.Time)
	ui := bigenstd.Uint64(bs)
	*v2 = time.Unix(0, int64(ui)).UTC()
}

type testUnixNanoTimeInterfaceExt struct{}

func (x testUnixNanoTimeInterfaceExt) ConvertExt(v interface{}) interface{} {
	v2 := v.(*time.Time) // structs are encoded by passing the ptr
	return v2.UTC().UnixNano()
}

func (x testUnixNanoTimeInterfaceExt) UpdateExt(dest interface{}, v interface{}) {
	tt := dest.(*time.Time)
	*tt = time.Unix(0, v.(int64)).UTC()
	// switch v2 := v.(type) {
	// case int64:
	// 	*tt = time.Unix(0, v2).UTC()
	// case uint64:
	// 	*tt = time.Unix(0, int64(v2)).UTC()
	// //case float64:
	// //case string:
	// default:
	// 	panic(fmt.Errorf("unsupported format for time conversion: expecting int64/uint64; got %T", v))
	// }
}

// ----

type wrapInt64Ext int64

func (x *wrapInt64Ext) WriteExt(v interface{}) []byte {
	v2 := uint64(int64(v.(wrapInt64)))
	bs := make([]byte, 8)
	bigenstd.PutUint64(bs, v2)
	return bs
}
func (x *wrapInt64Ext) ReadExt(v interface{}, bs []byte) {
	v2 := v.(*wrapInt64)
	ui := bigenstd.Uint64(bs)
	*v2 = wrapInt64(int64(ui))
}
func (x *wrapInt64Ext) ConvertExt(v interface{}) interface{} {
	return int64(v.(wrapInt64))
}
func (x *wrapInt64Ext) UpdateExt(dest interface{}, v interface{}) {
	v2 := dest.(*wrapInt64)
	*v2 = wrapInt64(v.(int64))
}

// ----

type wrapBytesExt struct{}

func (x *wrapBytesExt) WriteExt(v interface{}) []byte {
	return ([]byte)(v.(wrapBytes))
}
func (x *wrapBytesExt) ReadExt(v interface{}, bs []byte) {
	v2 := v.(*wrapBytes)
	*v2 = wrapBytes(bs)
}
func (x *wrapBytesExt) ConvertExt(v interface{}) interface{} {
	return ([]byte)(v.(wrapBytes))
}
func (x *wrapBytesExt) UpdateExt(dest interface{}, v interface{}) {
	v2 := dest.(*wrapBytes)
	// some formats (e.g. json) cannot nakedly determine []byte from string, so expect both
	switch v3 := v.(type) {
	case []byte:
		*v2 = wrapBytes(v3)
	case string:
		*v2 = wrapBytes([]byte(v3))
	default:
		panic(errors.New("UpdateExt for wrapBytesExt expects string or []byte"))
	}
	// *v2 = wrapBytes(v.([]byte))
}

// ----

// timeExt is an extension handler for time.Time, that uses binc model for encoding/decoding time.
// we used binc model, as that is the only custom time representation that we designed ourselves.
type timeBytesExt struct{}

func (x timeBytesExt) WriteExt(v interface{}) (bs []byte) {
	return bincEncodeTime(*(v.(*time.Time)))
	// return bincEncodeTime(v.(time.Time))
	// switch v2 := v.(type) {
	// case time.Time:
	// 	bs = bincEncodeTime(v2)
	// case *time.Time:
	// 	bs = bincEncodeTime(*v2)
	// default:
	// 	panic(fmt.Errorf("unsupported format for time conversion: expecting time.Time; got %T", v2))
	// }
	// return
}
func (x timeBytesExt) ReadExt(v interface{}, bs []byte) {
	tt, err := bincDecodeTime(bs)
	if err != nil {
		panic(err)
	}
	*(v.(*time.Time)) = tt
}

type timeInterfaceExt struct{}

func (x timeInterfaceExt) ConvertExt(v interface{}) interface{} {
	return timeBytesExt{}.WriteExt(v)
}
func (x timeInterfaceExt) UpdateExt(v interface{}, src interface{}) {
	timeBytesExt{}.ReadExt(v, src.([]byte))
}

// ----
type testUintToBytesExt struct{}

func (x testUintToBytesExt) WriteExt(v interface{}) (bs []byte) {
	z := uint32(v.(testUintToBytes))
	if z == 0 {
		return nil
	}
	return make([]byte, z)
}
func (x testUintToBytesExt) ReadExt(v interface{}, bs []byte) {
	*(v.(*testUintToBytes)) = testUintToBytes(len(bs))
}
func (x testUintToBytesExt) ConvertExt(v interface{}) interface{} {
	return x.WriteExt(v)
}
func (x testUintToBytesExt) UpdateExt(v interface{}, src interface{}) {
	x.ReadExt(v, src.([]byte))
}

// ----

func testSetupNoop() {}

func testSetup(t *testing.T, h *Handle) (fn func()) {
	return testSetupWithChecks(t, h, true)
}

// testSetup will ensure testInitAll is run, and then
// return a function that should be deferred to run at the end
// of the test.
//
// This function can track how much time a test took,
// or recover from panic's and fail the test appropriately.
func testSetupWithChecks(t *testing.T, h *Handle, allowParallel bool) (fn func()) {
	testOnce.Do(testInitAll)
	if allowParallel && testUseParallel {
		t.Parallel()
		if h != nil {
			*h = testHandleCopy(*h)
		}
	}
	// in case an error is seen, recover it here.
	if testRecoverPanicToErr {
		fnRecoverPanic := func() {
			if x := recover(); x != nil {
				var err error
				panicValToErr(errDecoratorDef{}, x, &err)
				t.Logf("recovered error: %v", err)
				t.FailNow()
			}
		}
		fn = fnRecoverPanic
	}
	if fn == nil {
		fn = testSetupNoop
	}
	return
}

func testBasicHandle(h Handle) *BasicHandle { return h.getBasicHandle() }

func testCodecEncode(ts interface{}, bsIn []byte, fn func([]byte) *bytes.Buffer, h Handle, useMust bool) (bs []byte, err error) {
	return testSharedCodecEncode(ts, bsIn, fn, h, testBasicHandle(h), useMust)
}

func testCodecDecode(bs []byte, ts interface{}, h Handle, useMust bool) (err error) {
	return testSharedCodecDecode(bs, ts, h, testBasicHandle(h), useMust)
}

func testCheckErr(t *testing.T, err error) {
	if err != nil {
		t.Logf("err: %v", err)
		t.FailNow()
	}
}

func testCheckEqual(t *testing.T, v1 interface{}, v2 interface{}, desc string) {
	t.Helper()
	if err := deepEqual(v1, v2); err != nil {
		t.Logf("Not Equal: %s: %v", desc, err)
		if testVerbose {
			t.Logf("\tv1: %v, v2: %v", v1, v2)
		}
		t.FailNow()
	}
}

func testInit() {
	gob.Register(new(TestStrucFlex))

	for _, v := range testHandles {
		bh := testBasicHandle(v)
		bh.clearInited() // so it is reinitialized next time around
		// pre-fill them first
		bh.EncodeOptions = testEncodeOptions
		bh.DecodeOptions = testDecodeOptions
		bh.RPCOptions = testRPCOptions
		// bh.InterfaceReset = true
		// bh.PreferArrayOverSlice = true
		// modify from flag'ish things
		bh.MaxInitLen = testMaxInitLen
	}

	var tTimeExt timeBytesExt
	var tBytesExt wrapBytesExt
	var tI64Ext wrapInt64Ext
	var tUintToBytesExt testUintToBytesExt

	// create legacy functions suitable for deprecated AddExt functionality,
	// and use on some places for testSimpleH e.g. for time.Time and wrapInt64
	var (
		myExtEncFn = func(x BytesExt, rv reflect.Value) (bs []byte, err error) {
			defer testPanicToErr(errDecoratorDef{}, &err)
			bs = x.WriteExt(rv.Interface())
			return
		}
		myExtDecFn = func(x BytesExt, rv reflect.Value, bs []byte) (err error) {
			defer testPanicToErr(errDecoratorDef{}, &err)
			x.ReadExt(rv.Interface(), bs)
			return
		}
		timeExtEncFn      = func(rv reflect.Value) (bs []byte, err error) { return myExtEncFn(tTimeExt, rv) }
		timeExtDecFn      = func(rv reflect.Value, bs []byte) (err error) { return myExtDecFn(tTimeExt, rv, bs) }
		wrapInt64ExtEncFn = func(rv reflect.Value) (bs []byte, err error) { return myExtEncFn(&tI64Ext, rv) }
		wrapInt64ExtDecFn = func(rv reflect.Value, bs []byte) (err error) { return myExtDecFn(&tI64Ext, rv, bs) }
	)

	chkErr := func(err error) {
		if err != nil {
			panic(err)
		}
	}

	// time.Time is a native type, so extensions will have no effect.
	// However, we add these here to ensure nothing happens.
	chkErr(testSimpleH.AddExt(timeTyp, 1, timeExtEncFn, timeExtDecFn))
	// testBincH.SetBytesExt(timeTyp, 1, timeExt{}) // time is builtin for binc
	chkErr(testMsgpackH.SetBytesExt(timeTyp, 1, timeBytesExt{}))
	chkErr(testCborH.SetInterfaceExt(timeTyp, 1, testUnixNanoTimeInterfaceExt{}))
	// testJsonH.SetInterfaceExt(timeTyp, 1, &testUnixNanoTimeExt{})

	// Add extensions for the testSelfExt
	chkErr(testSimpleH.SetBytesExt(testSelfExtTyp, 78, SelfExt))
	chkErr(testMsgpackH.SetBytesExt(testSelfExtTyp, 78, SelfExt))
	chkErr(testBincH.SetBytesExt(testSelfExtTyp, 78, SelfExt))
	chkErr(testJsonH.SetInterfaceExt(testSelfExtTyp, 78, SelfExt))
	chkErr(testCborH.SetInterfaceExt(testSelfExtTyp, 78, SelfExt))

	chkErr(testSimpleH.SetBytesExt(testSelfExt2Typ, 79, SelfExt))
	chkErr(testMsgpackH.SetBytesExt(testSelfExt2Typ, 79, SelfExt))
	chkErr(testBincH.SetBytesExt(testSelfExt2Typ, 79, SelfExt))
	chkErr(testJsonH.SetInterfaceExt(testSelfExt2Typ, 79, SelfExt))
	chkErr(testCborH.SetInterfaceExt(testSelfExt2Typ, 79, SelfExt))

	// Now, add extensions for the type wrapInt64 and wrapBytes,
	// so we can execute the Encode/Decode Ext paths.

	chkErr(testSimpleH.SetBytesExt(wrapBytesTyp, 32, &tBytesExt))
	chkErr(testMsgpackH.SetBytesExt(wrapBytesTyp, 32, &tBytesExt))
	chkErr(testBincH.SetBytesExt(wrapBytesTyp, 32, &tBytesExt))
	chkErr(testJsonH.SetInterfaceExt(wrapBytesTyp, 32, &tBytesExt))
	chkErr(testCborH.SetInterfaceExt(wrapBytesTyp, 32, &tBytesExt))

	chkErr(testSimpleH.SetBytesExt(testUintToBytesTyp, 33, &tUintToBytesExt))
	chkErr(testMsgpackH.SetBytesExt(testUintToBytesTyp, 33, &tUintToBytesExt))
	chkErr(testBincH.SetBytesExt(testUintToBytesTyp, 33, &tUintToBytesExt))
	chkErr(testJsonH.SetInterfaceExt(testUintToBytesTyp, 33, &tUintToBytesExt))
	chkErr(testCborH.SetInterfaceExt(testUintToBytesTyp, 33, &tUintToBytesExt))

	chkErr(testSimpleH.AddExt(wrapInt64Typ, 16, wrapInt64ExtEncFn, wrapInt64ExtDecFn))
	// chkErr(testSimpleH.SetBytesExt(wrapInt64Typ, 16, &tI64Ext))
	chkErr(testMsgpackH.SetBytesExt(wrapInt64Typ, 16, &tI64Ext))
	chkErr(testBincH.SetBytesExt(wrapInt64Typ, 16, &tI64Ext))
	chkErr(testJsonH.SetInterfaceExt(wrapInt64Typ, 16, &tI64Ext))
	chkErr(testCborH.SetInterfaceExt(wrapInt64Typ, 16, &tI64Ext))

	// primitives MUST be an even number, so it can be used as a mapBySlice also.
	primitives := []interface{}{
		int8(-8),
		int16(-1616),
		int32(-32323232),
		int64(-6464646464646464),
		uint8(192),
		uint16(1616),
		uint32(32323232),
		uint64(6464646464646464),
		byte(192),
		float32(-3232.0),
		float64(-6464646464.0),
		float32(3232.0),
		float64(6464.0),
		float64(6464646464.0),
		complex64(complex(160.0, 0)),
		complex128(complex(1616, 0)),
		false,
		true,
		"null",
		nil,
		"some&day>some<day",
		timeToCompare1,
		"",
		timeToCompare2,
		"bytestring",
		timeToCompare3,
		"none",
		timeToCompare4,
	}

	maps := []interface{}{
		map[string]bool{
			"true":  true,
			"false": false,
		},
		map[string]interface{}{
			"true":         "True",
			"false":        false,
			"uint16(1616)": uint16(1616),
		},
		//add a complex combo map in here. (map has list which has map)
		//note that after the first thing, everything else should be generic.
		map[string]interface{}{
			"list": []interface{}{
				int16(1616),
				int32(32323232),
				true,
				float32(-3232.0),
				map[string]interface{}{
					"TRUE":  true,
					"FALSE": false,
				},
				[]interface{}{true, false},
			},
			"int32": int32(32323232),
			"bool":  true,
			"LONG STRING": `
1234567890 1234567890 
1234567890 1234567890 
1234567890 1234567890 
ABCDEDFGHIJKLMNOPQRSTUVWXYZ 
abcdedfghijklmnopqrstuvwxyz 
ABCDEDFGHIJKLMNOPQRSTUVWXYZ 
abcdedfghijklmnopqrstuvwxyz 
"ABCDEDFGHIJKLMNOPQRSTUVWXYZ" 
'	a tab	'
\a\b\c\d\e
\b\f\n\r\t all literally
ugorji
`,
			"SHORT STRING": "1234567890",
		},
		map[interface{}]interface{}{
			true:       "true",
			uint8(138): false,
			false:      uint8(200),
		},
	}

	testTableNumPrimitives = len(primitives)
	testTableIdxTime = testTableNumPrimitives - 8
	testTableNumMaps = len(maps)

	table = []interface{}{}
	table = append(table, primitives...)
	table = append(table, primitives)
	table = append(table, testMbsT(primitives))
	table = append(table, maps...)
	table = append(table, newTestStrucFlex(0, testNumRepeatString, false, !testSkipIntf, testMapStringKeyOnly))
	// table = append(table, newTestStrucFlex(0, testNumRepeatString, true, !testSkipIntf, testMapStringKeyOnly))
}

func testTableVerify(f testVerifyFlag, h Handle) (av []interface{}) {
	av = make([]interface{}, len(table))
	lp := testTableNumPrimitives + 4
	// doNil := f & testVerifyDoNil == testVerifyDoNil
	// doPython := f & testVerifyForPython == testVerifyForPython
	switch {
	case f.isset(testVerifyForPython):
		for i, v := range table {
			if i == testTableNumPrimitives+1 || i > lp { // testTableNumPrimitives+1 is the mapBySlice
				av[i] = skipVerifyVal
				continue
			}
			av[i] = testVerifyVal(v, f, h)
		}
		// only do the python verify up to the maps, skipping the last 2 maps.
		av = av[:testTableNumPrimitives+2+testTableNumMaps-2]
	case f.isset(testVerifyDoNil):
		for i, v := range table {
			if i > lp {
				av[i] = skipVerifyVal
				continue
			}
			av[i] = testVerifyVal(v, f, h)
		}
	default:
		for i, v := range table {
			if i == lp {
				av[i] = skipVerifyVal
				continue
			}
			//av[i] = testVerifyVal(v, testVerifyMapTypeSame)
			switch v.(type) {
			case []interface{}:
				av[i] = testVerifyVal(v, f, h)
			case testMbsT:
				av[i] = testVerifyVal(v, f, h)
			case map[string]interface{}:
				av[i] = testVerifyVal(v, f, h)
			case map[interface{}]interface{}:
				av[i] = testVerifyVal(v, f, h)
			case time.Time:
				av[i] = testVerifyVal(v, f, h)
			default:
				av[i] = v
			}
		}
	}
	return
}

func testVerifyValInt(v int64, isMsgp bool) (v2 interface{}) {
	if isMsgp {
		if v >= 0 && v <= 127 {
			v2 = uint64(v)
		} else {
			v2 = int64(v)
		}
	} else if v >= 0 {
		v2 = uint64(v)
	} else {
		v2 = int64(v)
	}
	return
}

func testVerifyVal(v interface{}, f testVerifyFlag, h Handle) (v2 interface{}) {
	//for python msgpack,
	//  - all positive integers are unsigned 64-bit ints
	//  - all floats are float64
	_, isMsgp := h.(*MsgpackHandle)
	_, isCbor := h.(*CborHandle)
	switch iv := v.(type) {
	case int8:
		v2 = testVerifyValInt(int64(iv), isMsgp)
		// fmt.Printf(">>>> is msgp: %v, v: %T, %v ==> v2: %T, %v\n", isMsgp, v, v, v2, v2)
	case int16:
		v2 = testVerifyValInt(int64(iv), isMsgp)
	case int32:
		v2 = testVerifyValInt(int64(iv), isMsgp)
	case int64:
		v2 = testVerifyValInt(int64(iv), isMsgp)
	case uint8:
		v2 = uint64(iv)
	case uint16:
		v2 = uint64(iv)
	case uint32:
		v2 = uint64(iv)
	case uint64:
		v2 = uint64(iv)
	case float32:
		v2 = float64(iv)
	case float64:
		v2 = float64(iv)
	case complex64:
		v2 = float64(float32(real(iv)))
	case complex128:
		v2 = float64(real(iv))
	case []interface{}:
		m2 := make([]interface{}, len(iv))
		for j, vj := range iv {
			m2[j] = testVerifyVal(vj, f, h)
		}
		v2 = m2
	case testMbsT:
		m2 := make([]interface{}, len(iv))
		for j, vj := range iv {
			m2[j] = testVerifyVal(vj, f, h)
		}
		v2 = testMbsT(m2)
	case map[string]bool:
		switch {
		case f.isset(testVerifyMapTypeSame):
			m2 := make(map[string]bool)
			for kj, kv := range iv {
				m2[kj] = kv
			}
			v2 = m2
		case f.isset(testVerifyMapTypeStrIntf):
			m2 := make(map[string]interface{})
			for kj, kv := range iv {
				m2[kj] = kv
			}
			v2 = m2
		case f.isset(testVerifyMapTypeIntfIntf):
			m2 := make(map[interface{}]interface{})
			for kj, kv := range iv {
				m2[kj] = kv
			}
			v2 = m2
		}
	case map[string]interface{}:
		switch {
		case f.isset(testVerifyMapTypeSame):
			m2 := make(map[string]interface{})
			for kj, kv := range iv {
				m2[kj] = testVerifyVal(kv, f, h)
			}
			v2 = m2
		case f.isset(testVerifyMapTypeStrIntf):
			m2 := make(map[string]interface{})
			for kj, kv := range iv {
				m2[kj] = testVerifyVal(kv, f, h)
			}
			v2 = m2
		case f.isset(testVerifyMapTypeIntfIntf):
			m2 := make(map[interface{}]interface{})
			for kj, kv := range iv {
				m2[kj] = testVerifyVal(kv, f, h)
			}
			v2 = m2
		}
	case map[interface{}]interface{}:
		m2 := make(map[interface{}]interface{})
		for kj, kv := range iv {
			m2[testVerifyVal(kj, f, h)] = testVerifyVal(kv, f, h)
		}
		v2 = m2
	case time.Time:
		switch {
		case f.isset(testVerifyTimeAsInteger):
			if iv2 := iv.UnixNano(); iv2 >= 0 {
				v2 = uint64(iv2)
			} else {
				v2 = int64(iv2)
			}
		case isMsgp:
			v2 = iv.UTC()
		case isCbor:
			// fmt.Printf("%%%% cbor verifier\n")
			v2 = iv.UTC().Round(time.Microsecond)
		default:
			v2 = v
		}
	default:
		v2 = v
	}
	return
}

func testReleaseBytes(bs []byte) {
	if !testUseParallel {
		testBytesFreeList.put(bs)
	}
}

func testGetBytes() (bs []byte) {
	if !testUseParallel {
		bs = testBytesFreeList.get(64)
	}
	return
}

func testHandleCopy(h Handle) (h2 Handle) {
	switch v := h.(type) {
	case *JsonHandle:
		v2 := *v
		h2 = &v2
	case *CborHandle:
		v2 := *v
		h2 = &v2
	case *MsgpackHandle:
		v2 := *v
		h2 = &v2
	case *SimpleHandle:
		v2 := *v
		h2 = &v2
	case *BincHandle:
		v2 := *v
		h2 = &v2
	}
	return
}

func testMarshal(v interface{}, h Handle) (bs []byte, err error) {
	// return testCodecEncode(v, nil, testByteBuf, h)
	return testCodecEncode(v, testGetBytes(), testByteBuf, h, false)
}

func testUnmarshal(v interface{}, data []byte, h Handle) (err error) {
	return testCodecDecode(data, v, h, false)
}

func testMarshalErr(v interface{}, h Handle, t *testing.T, name string) (bs []byte) {
	t.Helper()
	bs, err := testCodecEncode(v, testGetBytes(), testByteBuf, h, true)
	if err != nil {
		t.Logf("%s: marshal failed: %v", name, err)
		if testVerbose {
			t.Logf("Error encoding %s: %v, Err: %v", name, v, err)
		}
		t.FailNow()
	}
	return
}

func testUnmarshalErr(v interface{}, data []byte, h Handle, t *testing.T, name string) {
	t.Helper()
	err := testCodecDecode(data, v, h, true)
	if err != nil {
		t.Logf("%s: unmarshal failed: %v", name, err)
		if testVerbose {
			t.Logf("Error Decoding into %s: %v, Err: %v", name, v, err)
		}
		t.FailNow()
	}
}

func testDeepEqualErr(v1, v2 interface{}, t *testing.T, name string) {
	t.Helper()
	if err := deepEqual(v1, v2); err == nil {
		if testVerbose {
			t.Logf("%s: values equal", name)
		}
	} else {
		t.Logf("%s: values not equal: %v", name, err)
		if testVerbose {
			t.Logf("%s: values not equal: %v. 1: %#v, 2: %#v", name, err, v1, v2)
		}
		t.FailNow()
	}
}

func testReadWriteCloser(c io.ReadWriteCloser) io.ReadWriteCloser {
	if testRpcBufsize <= 0 && rand.Int63()%2 == 0 {
		return c
	}
	return struct {
		io.Closer
		*bufio.Reader
		*bufio.Writer
	}{c, bufio.NewReaderSize(c, testRpcBufsize), bufio.NewWriterSize(c, testRpcBufsize)}
}

// testCodecTableOne allows us test for different variations based on arguments passed.
func testCodecTableOne(t *testing.T, testNil bool, h Handle,
	vs []interface{}, vsVerify []interface{}) {
	//if testNil, then just test for when a pointer to a nil interface{} is passed. It should work.
	//Current setup allows us test (at least manually) the nil interface or typed interface.
	if testVerbose {
		t.Logf("================ TestNil: %v: %v entries ================\n", testNil, len(vs))
	}

	if mh, ok := h.(*MsgpackHandle); ok {
		defer func(a, b bool) {
			mh.RawToString = a
			mh.PositiveIntUnsigned = b
		}(mh.RawToString, mh.PositiveIntUnsigned)
		mh.RawToString = true
		mh.PositiveIntUnsigned = false
	}

	bh := testBasicHandle(h)
	for i, v0 := range vs {
		if testVerbose {
			t.Logf("..............................................")
			t.Logf("         Testing: #%d:, %T, %#v\n", i, v0, v0)
		}
		// if a TestStrucFlex and we are doing a testNil,
		// ensure the fields which are not encodeable are set to nil appropriately
		// i.e. TestStrucFlex.{MstrUi64TSelf, mapMsu2wss}
		var mapMstrUi64TSelf map[stringUint64T]*stringUint64T
		var mapMsu2wss map[stringUint64T]wrapStringSlice
		// TestStrucFlex.{Msp2ss, Mip2ss} have pointer keys.
		// When we encode and the decode back into the same value,
		// the length of this map will effectively double, because
		// each pointer has equal underlying value, but are separate entries in the map.
		//
		// Best way to compare is to store them, and then compare them later if needed.
		var mapMsp2ss map[*string][]string
		var mapMip2ss map[*uint64][]string

		tsflex, _ := v0.(*TestStrucFlex)
		if tsflex != nil {
			mapMstrUi64TSelf = tsflex.MstrUi64TSelf
			mapMsu2wss = tsflex.Msu2wss
			mapMsp2ss = tsflex.Msp2ss
			mapMip2ss = tsflex.Mip2ss
			if testNil {
				tsflex.MstrUi64TSelf = nil
				tsflex.Msu2wss = nil
			}
		}
		b0 := testMarshalErr(v0, h, t, "v0")
		var b1 = b0
		if len(b0) > 1024 {
			b1 = b0[:1024]
		}
		bytesorstr := "string"
		if h.isBinary() {
			bytesorstr = "bytes"
			if len(b0) > 256 {
				b1 = b0[:256]
			}
		}
		if testVerbose {
			t.Logf("         Encoded %s: type: %T, len/cap: %v/%v, %v, %s\n", bytesorstr, v0, len(b0), cap(b0), b1, "...")
		}
		// TestStrucFlex has many fields which will encode differently if SignedInteger - so skip
		if _, ok := v0.(*TestStrucFlex); ok && bh.SignedInteger {
			continue
		}

		var v1 interface{}
		var err error

		if tsflex != nil {
			if testNil {
				tsflex.MstrUi64TSelf = mapMstrUi64TSelf
				tsflex.Msu2wss = mapMsu2wss
			}
			tsflex.Msp2ss = nil
			tsflex.Mip2ss = nil
		}

		if testNil {
			err = testUnmarshal(&v1, b0, h)
		} else if v0 != nil {
			v0rt := reflect.TypeOf(v0) // ptr
			if v0rt.Kind() == reflect.Ptr {
				err = testUnmarshal(v0, b0, h)
				v1 = v0
			} else {
				rv1 := reflect.New(v0rt)
				err = testUnmarshal(rv1.Interface(), b0, h)
				v1 = rv1.Elem().Interface()
				// v1 = reflect.Indirect(reflect.ValueOf(v1)).Interface()
			}
		}

		if tsflex != nil {
			// MARKER: consider compare tsflex.{Msp2ss, Mip2ss} to map{Msp2ss, Mip2ss}
			_, _ = mapMsp2ss, mapMip2ss
		}

		if testVerbose {
			t.Logf("         v1 returned: %T, %v %#v", v1, v1, v1)
		}
		// if v1 != nil {
		//	t.Logf("         v1 returned: %T, %#v", v1, v1)
		//	//we always indirect, because ptr to typed value may be passed (if not testNil)
		//	v1 = reflect.Indirect(reflect.ValueOf(v1)).Interface()
		// }
		if err != nil {
			t.Logf("-------- Error: %v", err)
			if testVerbose {
				t.Logf("-------- Partial return: %v", v1)
			}
			t.FailNow()
		}
		v0check := vsVerify[i]
		if v0check == skipVerifyVal || bh.SignedInteger {
			if testVerbose {
				t.Logf("        Nil Check skipped: Decoded: %T, %#v\n", v1, v1)
			}
			continue
		}
		if err = deepEqual(v0check, v1); err == nil {
			if testVerbose {
				t.Logf("++++++++ Before and After marshal matched\n")
			}
		} else {
			// t.Logf("-------- Before and After marshal do not match: Error: %v"+
			// 	" ====> GOLDEN: (%T) %#v, DECODED: (%T) %#v\n", err, v0check, v0check, v1, v1)
			t.Logf("-------- FAIL: Before and After marshal do not match: Error: %v", err)
			if testVerbose {
				t.Logf("    ....... GOLDEN:  (%T) %v %#v", v0check, v0check, v0check)
				t.Logf("    ....... DECODED: (%T) %v %#v", v1, v1, v1)
			}
			t.FailNow()
		}
		testReleaseBytes(b0)
	}
}

func doTestCodecTableOne(t *testing.T, h Handle) {
	// since this test modifies maps (and slices?), it should not be run in parallel,
	// else we may get "concurrent modification/range/set" errors.

	defer testSetupWithChecks(t, &h, false)()

	numPrim, numMap, idxTime, idxMap := testTableNumPrimitives, testTableNumMaps, testTableIdxTime, testTableNumPrimitives+2

	tableVerify := testTableVerify(testVerifyMapTypeSame, h)
	tableTestNilVerify := testTableVerify(testVerifyDoNil|testVerifyMapTypeStrIntf, h)
	switch v := h.(type) {
	case *MsgpackHandle:
		oldWriteExt := v.WriteExt
		v.WriteExt = true
		testCodecTableOne(t, false, h, table, tableVerify)
		v.WriteExt = oldWriteExt
	case *JsonHandle:
		//skip []interface{} containing time.Time, as it encodes as a number, but cannot decode back to time.Time.
		//As there is no real support for extension tags in json, this must be skipped.
		testCodecTableOne(t, false, h, table[:numPrim], tableVerify[:numPrim])
		testCodecTableOne(t, false, h, table[idxMap:], tableVerify[idxMap:])
	default:
		testCodecTableOne(t, false, h, table, tableVerify)
	}
	// func TestMsgpackAll(t *testing.T) {

	// //skip []interface{} containing time.Time
	// testCodecTableOne(t, false, h, table[:numPrim], tableVerify[:numPrim])
	// testCodecTableOne(t, false, h, table[numPrim+1:], tableVerify[numPrim+1:])
	// func TestMsgpackNilStringMap(t *testing.T) {
	var oldMapType reflect.Type
	v := testBasicHandle(h)

	oldMapType, v.MapType = v.MapType, testMapStrIntfTyp
	// defer func() { v.MapType = oldMapType }()
	//skip time.Time, []interface{} containing time.Time, last map, and newStruc
	testCodecTableOne(t, true, h, table[:idxTime], tableTestNilVerify[:idxTime])
	testCodecTableOne(t, true, h, table[idxMap:idxMap+numMap-1], tableTestNilVerify[idxMap:idxMap+numMap-1]) // failing one for msgpack
	v.MapType = oldMapType
	// func TestMsgpackNilIntf(t *testing.T) {

	//do last map and newStruc
	idx2 := idxMap + numMap - 1
	testCodecTableOne(t, true, h, table[idx2:], tableTestNilVerify[idx2:])
	//testCodecTableOne(t, true, h, table[17:18], tableTestNilVerify[17:18]) // do we need this?
}

func doTestCodecMiscOne(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	var err error
	bh := testBasicHandle(h)
	b := testMarshalErr(32, h, t, "32")
	// Cannot do this nil one, because faster type assertion decoding will panic
	// var i *int32
	// if err = testUnmarshal(b, i, nil); err == nil {
	// 	t.Logf("------- Expecting error because we cannot unmarshal to int32 nil ptr")
	// 	t.FailNow()
	// }
	var i2 int32
	testUnmarshalErr(&i2, b, h, t, "int32-ptr")
	if i2 != int32(32) {
		t.Logf("------- didn't unmarshal to 32: Received: %d", i2)
		t.FailNow()
	}

	// func TestMsgpackDecodePtr(t *testing.T) {
	ts := newTestStrucFlex(testDepth, testNumRepeatString, false, !testSkipIntf, testMapStringKeyOnly)
	testReleaseBytes(b)
	b = testMarshalErr(ts, h, t, "pointer-to-struct")
	if len(b) < 40 {
		t.Logf("------- Size must be > 40. Size: %d", len(b))
		t.FailNow()
	}
	var b1 = b
	if len(b1) > 256 {
		b1 = b1[:256]
	}
	if testVerbose {
		if h.isBinary() {
			t.Logf("------- b: size: %v, value: %v", len(b), b1)
		} else {
			t.Logf("------- b: size: %v, value: %s", len(b), b1)
		}
	}

	// ts2 := testEmptyTestStrucFlex()
	var ts2 = new(TestStrucFlex)
	// we initialize and start draining the chan, so that we can decode into it without it blocking due to no consumer
	ts2.Chstr = make(chan string, teststrucflexChanCap)
	go func() {
		for range ts2.Chstr {
		}
	}() // drain it

	testUnmarshalErr(ts2, b, h, t, "pointer-to-struct")
	if ts2.I64 != math.MaxInt64*2/3 {
		t.Logf("------- Unmarshal wrong. Expect I64 = 64. Got: %v", ts2.I64)
		t.FailNow()
	}
	close(ts2.Chstr)
	testReleaseBytes(b)

	// Note: These will not work with SliceElementReset or InterfaceReset=true, so handle that.
	defer func(a, b bool) {
		bh.SliceElementReset = a
		bh.InterfaceReset = b
	}(bh.SliceElementReset, bh.InterfaceReset)

	bh.SliceElementReset = false
	bh.InterfaceReset = false

	m := map[string]int{"A": 2, "B": 3}
	p := []interface{}{m}
	bs := testMarshalErr(p, h, t, "p")

	m2 := map[string]int{}
	p2 := []interface{}{m2}
	testUnmarshalErr(&p2, bs, h, t, "&p2")

	if m2["A"] != 2 || m2["B"] != 3 {
		t.Logf("FAIL: m2 not as expected: expecting: %v, got: %v", m, m2)
		t.FailNow()
	}

	testCheckEqual(t, p, p2, "p=p2")
	testCheckEqual(t, m, m2, "m=m2")
	if err = deepEqual(p, p2); err == nil {
		if testVerbose {
			t.Logf("p and p2 match")
		}
	} else {
		t.Logf("Not Equal: %v. p: %v, p2: %v", err, p, p2)
		t.FailNow()
	}
	if err = deepEqual(m, m2); err == nil {
		if testVerbose {
			t.Logf("m and m2 match")
		}
	} else {
		t.Logf("Not Equal: %v. m: %v, m2: %v", err, m, m2)
		t.FailNow()
	}
	testReleaseBytes(bs)

	// func TestMsgpackDecodeStructSubset(t *testing.T) {
	// test that we can decode a subset of the stream
	mm := map[string]interface{}{"A": 5, "B": 99, "C": 333}
	bs = testMarshalErr(mm, h, t, "mm")
	type ttt struct {
		A uint8
		C int32
	}
	var t2 ttt
	testUnmarshalErr(&t2, bs, h, t, "t2")
	t3 := ttt{5, 333}
	testCheckEqual(t, t2, t3, "t2=t3")
	testReleaseBytes(bs)

	// println(">>>>>")
	// test simple arrays, non-addressable arrays, slices
	type tarr struct {
		A int64
		B [3]int64
		C []byte
		D [3]byte
	}
	var tarr0 = tarr{1, [3]int64{2, 3, 4}, []byte{4, 5, 6}, [3]byte{7, 8, 9}}
	// test both pointer and non-pointer (value)
	for _, tarr1 := range []interface{}{tarr0, &tarr0} {
		bs = testMarshalErr(tarr1, h, t, "tarr1")
		if _, ok := h.(*JsonHandle); ok {
			if testVerbose {
				t.Logf("Marshal as: %s", bs)
			}
		}
		var tarr2 tarr
		testUnmarshalErr(&tarr2, bs, h, t, "tarr2")
		testCheckEqual(t, tarr0, tarr2, "tarr0=tarr2")
		testReleaseBytes(bs)
	}

	// test byte array, even if empty (msgpack only)
	if _, ok := h.(*MsgpackHandle); ok {
		type ystruct struct {
			Anarray []byte
		}
		var ya = ystruct{}
		testUnmarshalErr(&ya, []byte{0x91, 0x90}, h, t, "ya")
	}

	var tt1, tt2 time.Time
	tt2 = time.Now()
	bs = testMarshalErr(tt1, h, t, "zero-time-enc")
	testUnmarshalErr(&tt2, bs, h, t, "zero-time-dec")
	testDeepEqualErr(tt1, tt2, t, "zero-time-eq")
	testReleaseBytes(bs)

	// test encoding a slice of byte (but not []byte) and decoding into a []byte
	var sw = []wrapUint8{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J'}
	var bw []byte // ("ABCDEFGHIJ")
	bs = testMarshalErr(sw, h, t, "wrap-bytes-enc")
	testUnmarshalErr(&bw, bs, h, t, "wrap-bytes-dec")
	testDeepEqualErr(bw, []byte("ABCDEFGHIJ"), t, "wrap-bytes-eq")
	testReleaseBytes(bs)
}

func doTestCodecEmbeddedPointer(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	type Z int
	type A struct {
		AnInt int
	}
	type B struct {
		*Z
		*A
		MoreInt int
	}
	var z Z = 4
	x1 := &B{&z, &A{5}, 6}
	bs := testMarshalErr(x1, h, t, "x1")
	var x2 = new(B)
	testUnmarshalErr(x2, bs, h, t, "x2")
	testCheckEqual(t, x1, x2, "x1=x2")
	testReleaseBytes(bs)
}

func testCodecUnderlyingType(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	// Manual Test.
	// Run by hand, with accompanying printf.statements in fast-path.go
	// to ensure that the fast functions are called.
	type T1 map[string]string
	v := T1{"1": "1s", "2": "2s"}
	var bs []byte
	var err error
	NewEncoderBytes(&bs, h).MustEncode(v)
	if err != nil {
		t.Logf("Error during encode: %v", err)
		t.FailNow()
	}
	var v2 T1
	NewDecoderBytes(bs, h).MustDecode(&v2)
	if err != nil {
		t.Logf("Error during decode: %v", err)
		t.FailNow()
	}
}

func doTestCodecChan(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	// - send a slice []*int64 (sl1) into an chan (ch1) with cap > len(s1)
	// - encode ch1 as a stream array
	// - decode a chan (ch2), with cap > len(s1) from the stream array
	// - receive from ch2 into slice sl2
	// - compare sl1 and sl2

	{
		if testVerbose {
			t.Logf("*int64")
		}
		sl1 := make([]*int64, 4)
		for i := range sl1 {
			var j int64 = int64(i)
			sl1[i] = &j
		}
		ch1 := make(chan *int64, 4)
		for _, j := range sl1 {
			ch1 <- j
		}
		var bs []byte
		NewEncoderBytes(&bs, h).MustEncode(ch1)
		ch2 := make(chan *int64, 8)
		NewDecoderBytes(bs, h).MustDecode(&ch2)
		close(ch2)
		var sl2 []*int64
		for j := range ch2 {
			sl2 = append(sl2, j)
		}
		if err := deepEqual(sl1, sl2); err != nil {
			t.Logf("FAIL: Not Match: %v; len: %v, %v", err, len(sl1), len(sl2))
			if testVerbose {
				t.Logf("sl1: %#v, sl2: %#v", sl1, sl2)
			}
			t.FailNow()
		}
	}

	{
		if testVerbose {
			t.Logf("testBytesT []byte - input []byte")
		}
		type testBytesT []byte
		sl1 := make([]testBytesT, 4)
		for i := range sl1 {
			var j = []byte(strings.Repeat(strconv.FormatInt(int64(i), 10), i))
			sl1[i] = j
		}
		ch1 := make(chan testBytesT, 4)
		for _, j := range sl1 {
			ch1 <- j
		}
		var bs []byte
		NewEncoderBytes(&bs, h).MustEncode(ch1)
		ch2 := make(chan testBytesT, 8)
		NewDecoderBytes(bs, h).MustDecode(&ch2)
		close(ch2)
		var sl2 []testBytesT
		for j := range ch2 {
			// t.Logf(">>>> from chan: is nil? %v, %v", j == nil, j)
			sl2 = append(sl2, j)
		}
		if err := deepEqual(sl1, sl2); err != nil {
			t.Logf("FAIL: Not Match: %v; len: %v, %v", err, len(sl1), len(sl2))
			if testVerbose {
				t.Logf("sl1: %#v, sl2: %#v", sl1, sl2)
			}
			t.FailNow()
		}
	}
	{
		if testVerbose {
			t.Logf("testBytesT byte - input string/testBytesT")
		}
		type testBytesT byte
		sl1 := make([]testBytesT, 4)
		for i := range sl1 {
			var j = strconv.FormatInt(int64(i), 10)[0]
			sl1[i] = testBytesT(j)
		}
		ch1 := make(chan testBytesT, 4)
		for _, j := range sl1 {
			ch1 <- j
		}
		var bs []byte
		NewEncoderBytes(&bs, h).MustEncode(ch1)
		ch2 := make(chan testBytesT, 8)
		NewDecoderBytes(bs, h).MustDecode(&ch2)
		close(ch2)
		var sl2 []testBytesT
		for j := range ch2 {
			sl2 = append(sl2, j)
		}
		if err := deepEqual(sl1, sl2); err != nil {
			t.Logf("FAIL: Not Match: %v; len: %v, %v", err, len(sl1), len(sl2))
			t.FailNow()
		}
	}

	{
		if testVerbose {
			t.Logf("*[]byte")
		}
		sl1 := make([]byte, 4)
		for i := range sl1 {
			var j = strconv.FormatInt(int64(i), 10)[0]
			sl1[i] = byte(j)
		}
		ch1 := make(chan byte, 4)
		for _, j := range sl1 {
			ch1 <- j
		}
		var bs []byte
		NewEncoderBytes(&bs, h).MustEncode(ch1)
		ch2 := make(chan byte, 8)
		NewDecoderBytes(bs, h).MustDecode(&ch2)
		close(ch2)
		var sl2 []byte
		for j := range ch2 {
			sl2 = append(sl2, j)
		}
		if err := deepEqual(sl1, sl2); err != nil {
			t.Logf("FAIL: Not Match: %v; len: %v, %v", err, len(sl1), len(sl2))
			t.FailNow()
		}
	}
}

func doTestCodecRpcOne(t *testing.T, rr Rpc, h Handle, doRequest bool, exitSleepMs time.Duration) (port int) {
	defer testSetup(t, &h)()
	if testSkipRPCTests {
		t.Skip(testSkipRPCTestsMsg)
	}
	if !testRecoverPanicToErr {
		t.Skip(testSkipIfNotRecoverPanicToErrMsg)
	}

	if mh, ok := h.(*MsgpackHandle); ok && mh.SliceElementReset {
		t.Skipf("skipping ... MsgpackRpcSpec does not handle SliceElementReset - needs investigation")
	}

	if jsonH, ok := h.(*JsonHandle); ok && !jsonH.TermWhitespace {
		jsonH.TermWhitespace = true
		defer func() { jsonH.TermWhitespace = false }()
	}

	// srv := rpc.NewServer()
	// srv.Register(testRpcInt)
	srv := testRpcServer

	ln, err := net.Listen("tcp", "127.0.0.1:0") // listen on ipv4 localhost
	testCheckErr(t, err)
	port = (ln.Addr().(*net.TCPAddr)).Port
	if testVerbose {
		t.Logf("connFn: addr: %v, network: %v, port: %v", ln.Addr(), ln.Addr().Network(), port)
	}
	// var opts *DecoderOptions
	// opts := testDecOpts
	// opts.MapType = mapStrIntfTyp
	serverExitChan := make(chan bool, 1)
	var serverExitFlag uint64
	serverFn := func() {
		var conns []net.Conn
		var svrcodecs []rpc.ServerCodec
		defer func() {
			for i := range conns {
				svrcodecs[i].Close()
				conns[i].Close()
			}
			ln.Close()
			serverExitChan <- true
		}()
		for {
			conn1, err1 := ln.Accept()
			if atomic.LoadUint64(&serverExitFlag) == 1 {
				if conn1 != nil {
					conn1.Close()
				}
				break
			}
			if err1 != nil {
				// fmt.Printf("accept err1: %v\n", err1)
				if testVerbose {
					t.Logf("rpc error accepting connection: %v", err1)
				}
				continue
			}
			if conn1 != nil {
				sc := rr.ServerCodec(testReadWriteCloser(conn1), h)
				conns = append(conns, conn1)
				svrcodecs = append(svrcodecs, sc)
				go srv.ServeCodec(sc)
			}
		}
	}

	connFn := func() (bs net.Conn) {
		bs, err2 := net.Dial(ln.Addr().Network(), ln.Addr().String())
		testCheckErr(t, err2)
		return
	}

	exitFn := func() {
		if atomic.LoadUint64(&serverExitFlag) == 1 {
			return
		}
		atomic.StoreUint64(&serverExitFlag, 1)
		bs := connFn()
		defer bs.Close()
		<-serverExitChan
		// atomic.StoreUint64(&serverExitFlag, 0)
		// serverExitChan <- true
	}

	go serverFn()

	// runtime.Gosched()

	if exitSleepMs == 0 {
		defer exitFn()
	} else {
		go func() {
			time.Sleep(exitSleepMs)
			exitFn()
		}()
	}

	if doRequest {
		func() {
			bs := connFn()
			defer bs.Close()
			cc := rr.ClientCodec(testReadWriteCloser(bs), h)
			defer cc.Close()
			cl := rpc.NewClientWithCodec(cc)
			defer cl.Close()
			var up, sq, mult int
			var rstr string
			testCheckErr(t, cl.Call("TestRpcInt.Update", 5, &up))
			testCheckEqual(t, atomic.LoadInt64(&testRpcInt.i), int64(5), "testRpcInt.i=5")
			testCheckEqual(t, up, 5, "up=5")
			testCheckErr(t, cl.Call("TestRpcInt.Square", 1, &sq))
			testCheckEqual(t, sq, 25, "sq=25")
			testCheckErr(t, cl.Call("TestRpcInt.Mult", 20, &mult))
			testCheckEqual(t, mult, 100, "mult=100")
			testCheckErr(t, cl.Call("TestRpcInt.EchoStruct", TestRpcABC{"Aa", "Bb", "Cc"}, &rstr))
			testCheckEqual(t, rstr, fmt.Sprintf("%#v", TestRpcABC{"Aa", "Bb", "Cc"}), "rstr=")
			testCheckErr(t, cl.Call("TestRpcInt.Echo123", []string{"A1", "B2", "C3"}, &rstr))
			testCheckEqual(t, rstr, fmt.Sprintf("%#v", []string{"A1", "B2", "C3"}), "rstr=")
		}()
	}
	return
}

func doTestMapEncodeForCanonical(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	// println("doTestMapEncodeForCanonical")
	v1 := map[stringUint64T]interface{}{
		stringUint64T{"a", 1}: 1,
		stringUint64T{"b", 2}: "hello",
		stringUint64T{"c", 3}: map[string]interface{}{
			"c/a": 1,
			"c/b": "world",
			"c/c": []int{1, 2, 3, 4},
			"c/d": map[string]interface{}{
				"c/d/a": "fdisajfoidsajfopdjsaopfjdsapofda",
				"c/d/b": "fdsafjdposakfodpsakfopdsakfpodsakfpodksaopfkdsopafkdopsa",
				"c/d/c": "poir02  ir30qif4p03qir0pogjfpoaerfgjp ofke[padfk[ewapf kdp[afep[aw",
				"c/d/d": "fdsopafkd[sa f-32qor-=4qeof -afo-erfo r-eafo 4e-  o r4-qwo ag",
				"c/d/e": "kfep[a sfkr0[paf[a foe-[wq  ewpfao-q ro3-q ro-4qof4-qor 3-e orfkropzjbvoisdb",
				"c/d/f": "",
			},
			"c/e": map[int]string{
				1:     "1",
				22:    "22",
				333:   "333",
				4444:  "4444",
				55555: "55555",
			},
			"c/f": map[string]int{
				"1":     1,
				"22":    22,
				"333":   333,
				"4444":  4444,
				"55555": 55555,
			},
			"c/g": map[bool]int{
				false: 0,
				true:  1,
			},
			"c/t": map[time.Time]int64{
				time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC): time.Date(2019, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano(),
				time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC): time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano(),
				time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC): time.Date(2017, 1, 1, 0, 0, 0, 0, time.UTC).UnixNano(),
			},
		},
	}
	var v2 map[stringUint64T]interface{}
	var b1, b2, b3 []byte

	// encode v1 into b1, decode b1 into v2, encode v2 into b2, and compare b1 and b2.
	// OR
	// encode v1 into b1, decode b1 into v2, encode v2 into b2 and b3, and compare b2 and b3.
	//   e.g. when doing cbor indefinite, we may haveto use out-of-band encoding
	//   where each key is encoded as an indefinite length string, which makes it not the same
	//   order as the strings were lexicographically ordered before.

	var cborIndef bool
	if ch, ok := h.(*CborHandle); ok {
		cborIndef = ch.IndefiniteLength
	}

	// if ch, ok := h.(*BincHandle); ok && ch.AsSymbols != 2 {
	// 	defer func(u uint8) { ch.AsSymbols = u }(ch.AsSymbols)
	// 	ch.AsSymbols = 2
	// }

	bh := testBasicHandle(h)

	defer func(c, si bool) {
		bh.Canonical = c
		bh.SignedInteger = si
	}(bh.Canonical, bh.SignedInteger)
	bh.Canonical = true
	// bh.SignedInteger = true

	e1 := NewEncoderBytes(&b1, h)
	e1.MustEncode(v1)
	d1 := NewDecoderBytes(b1, h)
	d1.MustDecode(&v2)
	// testDeepEqualErr(v1, v2, t, "huh?")
	e2 := NewEncoderBytes(&b2, h)
	e2.MustEncode(v2)
	var b1t, b2t = b1, b2
	if cborIndef {
		e2 = NewEncoderBytes(&b3, h)
		e2.MustEncode(v2)
		b1t, b2t = b2, b3
	}

	if !bytes.Equal(b1t, b2t) {
		t.Logf("Unequal bytes of length: %v vs %v", len(b1t), len(b2t))
		if testVerbose {
			t.Logf("Unequal bytes: \n\t%v \n\tVS \n\t%v", b1t, b2t)
		}
		t.FailNow()
	}
}

func doTestStdEncIntf(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	args := [][2]interface{}{
		{&TestABC{"A", "BB", "CCC"}, new(TestABC)},
		{&TestABC2{"AAA", "BB", "C"}, new(TestABC2)},
	}
	for _, a := range args {
		var b []byte
		e := NewEncoderBytes(&b, h)
		e.MustEncode(a[0])
		d := NewDecoderBytes(b, h)
		d.MustDecode(a[1])
		if err := deepEqual(a[0], a[1]); err == nil {
			if testVerbose {
				t.Logf("++++ Objects match")
			}
		} else {
			t.Logf("---- FAIL: Objects do not match: y0: %v, y1: %v, err: %v", a[0], a[1], err)
			t.FailNow()
		}
	}
}

func doTestEncCircularRef(t *testing.T, h Handle) {
	if !testRecoverPanicToErr {
		t.Skip(testSkipIfNotRecoverPanicToErrMsg)
	}
	defer testSetup(t, &h)()
	type T1 struct {
		S string
		B bool
		T interface{}
	}
	type T2 struct {
		S string
		T *T1
	}
	type T3 struct {
		S string
		T *T2
	}
	t1 := T1{"t1", true, nil}
	t2 := T2{"t2", &t1}
	t3 := T3{"t3", &t2}
	t1.T = &t3

	var bs []byte
	var err error

	bh := testBasicHandle(h)
	if !bh.CheckCircularRef {
		bh.CheckCircularRef = true
		defer func() { bh.CheckCircularRef = false }()
	}
	err = NewEncoderBytes(&bs, h).Encode(&t3)
	if err == nil || !strings.Contains(err.Error(), "circular reference found") {
		t.Logf("expect circular reference error, got: %v", err)
		t.FailNow()
	}
	if x := err.Error(); strings.Contains(x, "circular") || strings.Contains(x, "cyclic") {
		if testVerbose {
			t.Logf("error detected as expected: %v", x)
		}
	} else {
		t.Logf("FAIL: error detected was not as expected: %v", x)
		t.FailNow()
	}
}

// TestAnonCycleT{1,2,3} types are used to test anonymous cycles.
// They are top-level, so that they can have circular references.
type (
	TestAnonCycleT1 struct {
		S string
		TestAnonCycleT2
	}
	TestAnonCycleT2 struct {
		S2 string
		TestAnonCycleT3
	}
	TestAnonCycleT3 struct {
		*TestAnonCycleT1
	}
)

func doTestAnonCycle(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	var x TestAnonCycleT1
	x.S = "hello"
	x.TestAnonCycleT2.S2 = "hello.2"
	x.TestAnonCycleT2.TestAnonCycleT3.TestAnonCycleT1 = &x

	// just check that you can get typeInfo for T1
	rt := reflect.TypeOf((*TestAnonCycleT1)(nil)).Elem()
	rtid := rt2id(rt)
	pti := testBasicHandle(h).getTypeInfo(rtid, rt)
	if testVerbose {
		t.Logf("[%s] pti: %v", h.Name(), pti)
	}
}

func doTestAllErrWriter(t *testing.T, hh ...Handle) {
	if !testRecoverPanicToErr {
		t.Skip(testSkipIfNotRecoverPanicToErrMsg)
	}
	defer testSetup(t, nil)()
	for _, h := range hh {
		if testUseParallel {
			h = testHandleCopy(h)
		}
		__doTestErrWriter(t, h)
	}
}

func __doTestErrWriter(t *testing.T, h Handle) {
	name := h.Name()
	var ew testErrWriter
	w := bufio.NewWriterSize(&ew, 4)
	enc := NewEncoder(w, h)
	for i := 0; i < 4; i++ {
		err := enc.Encode("ugorji")
		if ev, ok := err.(*codecError); ok {
			err = ev.Cause()
		}
		if err != testErrWriterErr {
			t.Logf("%s: expecting err: %v, received: %v", name, testErrWriterErr, err)
			t.FailNow()
		}
	}
}

func __doTestJsonLargeInteger(t *testing.T, v interface{}, ias uint8, jh *JsonHandle) {
	if testVerbose {
		t.Logf("Running TestJsonLargeInteger: v: %#v, ias: %c", v, ias)
	}
	oldIAS := jh.IntegerAsString
	defer func() { jh.IntegerAsString = oldIAS }()
	jh.IntegerAsString = ias

	var vu uint
	var vi int
	var vb bool
	var b []byte
	e := NewEncoderBytes(&b, jh)
	e.MustEncode(v)
	e.MustEncode(true)
	d := NewDecoderBytes(b, jh)
	// below, we validate that the json string or number was encoded,
	// then decode, and validate that the correct value was decoded.
	fnStrChk := func() {
		// check that output started with '"', and ended with '"true'
		// if !(len(b) >= 7 && b[0] == '"' && string(b[len(b)-7:]) == `" true `) {
		if !(len(b) >= 5 && b[0] == '"' && string(b[len(b)-5:]) == `"true`) {
			t.Logf("Expecting a JSON string, got: '%s'", b)
			t.FailNow()
		}
	}

	switch ias {
	case 'L':
		switch v2 := v.(type) {
		case int:
			v2n := int64(v2) // done to work with 32-bit OS
			if v2n > 1<<53 || (v2n < 0 && -v2n > 1<<53) {
				fnStrChk()
			}
		case uint:
			v2n := uint64(v2) // done to work with 32-bit OS
			if v2n > 1<<53 {
				fnStrChk()
			}
		}
	case 'A':
		fnStrChk()
	default:
		// check that output doesn't contain " at all
		for _, i := range b {
			if i == '"' {
				t.Logf("Expecting a JSON Number without quotation: got: %s", b)
				t.FailNow()
			}
		}
	}
	switch v2 := v.(type) {
	case int:
		d.MustDecode(&vi)
		d.MustDecode(&vb)
		// check that vb = true, and vi == v2
		if !(vb && vi == v2) {
			t.Logf("Expecting equal values from %s: got golden: %v, decoded: %v", b, v2, vi)
			t.FailNow()
		}
	case uint:
		d.MustDecode(&vu)
		d.MustDecode(&vb)
		// check that vb = true, and vi == v2
		if !(vb && vu == v2) {
			t.Logf("Expecting equal values from %s: got golden: %v, decoded: %v", b, v2, vu)
			t.FailNow()
		}
	}
}

func doTestRawValue(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	bh := testBasicHandle(h)
	if !bh.Raw {
		bh.Raw = true
		defer func() { bh.Raw = false }()
	}

	var i, i2 int
	var v, v2 TestRawValue
	var bs, bs2 []byte

	i = 1234 //1234567890
	v = TestRawValue{I: i}
	e := NewEncoderBytes(&bs, h)
	e.MustEncode(v.I)
	if testVerbose {
		t.Logf(">>> raw: %v, %s\n", bs, bs)
	}

	v.R = Raw(bs)
	e.ResetBytes(&bs2)
	e.MustEncode(v)

	if testVerbose {
		t.Logf(">>> bs2: %v, %s\n", bs2, bs2)
	}
	d := NewDecoderBytes(bs2, h)
	d.MustDecode(&v2)
	d.ResetBytes(v2.R)
	if testVerbose {
		t.Logf(">>> v2.R: %v, %s\n", ([]byte)(v2.R), ([]byte)(v2.R))
	}
	d.MustDecode(&i2)

	if testVerbose {
		t.Logf(">>> Encoded %v, decoded %v\n", i, i2)
	}
	// t.Logf("Encoded %v, decoded %v", i, i2)
	if i != i2 {
		t.Logf("Error: encoded %v, decoded %v", i, i2)
		t.FailNow()
	}
}

// Comprehensive testing that generates data encoded from python handle (cbor, msgpack),
// and validates that our code can read and write it out accordingly.
// We keep this unexported here, and put actual test in ext_dep_test.go.
// This way, it can be excluded by excluding file completely.
func doTestPythonGenStreams(t *testing.T, h Handle) {
	defer testSetup(t, &h)()

	// time0 := time.Now()
	// defer func() { xdebugf("python-gen-streams: %s: took: %v", h.Name(), time.Since(time0)) }()

	name := h.Name()
	if testVerbose {
		t.Logf("TestPythonGenStreams-%v", name)
	}
	tmpdir, err := ioutil.TempDir("", "golang-"+name+"-test")
	if err != nil {
		t.Logf("-------- Unable to create temp directory\n")
		t.FailNow()
	}
	defer os.RemoveAll(tmpdir)
	if testVerbose {
		t.Logf("tmpdir: %v", tmpdir)
	}
	cmd := exec.Command("python", "test.py", "testdata", tmpdir)
	//cmd.Stdin = strings.NewReader("some input")
	//cmd.Stdout = &out
	var cmdout []byte
	if cmdout, err = cmd.CombinedOutput(); err != nil {
		t.Logf("-------- Error running test.py testdata. Err: %v", err)
		t.Logf("         %v", string(cmdout))
		t.FailNow()
	}

	bh := testBasicHandle(h)

	defer func(tt reflect.Type) { bh.MapType = tt }(bh.MapType)
	defer func(b bool) { bh.RawToString = b }(bh.RawToString)

	// msgpack python needs Raw converted to string
	bh.RawToString = true

	oldMapType := bh.MapType
	tablePythonVerify := testTableVerify(testVerifyForPython|testVerifyTimeAsInteger|testVerifyMapTypeStrIntf, h)
	for i, v := range tablePythonVerify {
		// if v == uint64(0) && h is a *MsgpackHandle: v = int64(0)
		bh.MapType = oldMapType
		//load up the golden file based on number
		//decode it
		//compare to in-mem object
		//encode it again
		//compare to output stream
		if testVerbose {
			t.Logf("..............................................")
			t.Logf("         Testing: #%d: %T, %#v\n", i, v, v)
		}
		var bss []byte
		bss, err = ioutil.ReadFile(filepath.Join(tmpdir, strconv.Itoa(i)+"."+name+".golden"))
		if err != nil {
			t.Logf("-------- Error reading golden file: %d. Err: %v", i, err)
			t.FailNow()
			continue
		}
		bh.MapType = testMapStrIntfTyp

		var v1 interface{}
		if err = testUnmarshal(&v1, bss, h); err != nil {
			t.Logf("-------- Error decoding stream: %d: Err: %v", i, err)
			t.FailNow()
			continue
		}
		if v == skipVerifyVal {
			continue
		}
		//no need to indirect, because we pass a nil ptr, so we already have the value
		//if v1 != nil { v1 = reflect.Indirect(reflect.ValueOf(v1)).Interface() }
		if err = deepEqual(v, v1); err == nil {
			if testVerbose {
				t.Logf("++++++++ Objects match: %T, %v", v, v)
			}
		} else {
			t.Logf("-------- FAIL: Objects do not match: %v. Source: %T. Decoded: %T", err, v, v1)
			if testVerbose {
				t.Logf("--------   GOLDEN: %#v", v)
				// t.Logf("--------   DECODED: %#v <====> %#v", v1, reflect.Indirect(reflect.ValueOf(v1)).Interface())
				t.Logf("--------   DECODED: %#v <====> %#v", v1, reflect.Indirect(reflect.ValueOf(v1)).Interface())
			}
			t.FailNow()
		}
		bsb, err := testMarshal(v1, h)
		if err != nil {
			t.Logf("Error encoding to stream: %d: Err: %v", i, err)
			t.FailNow()
			continue
		}
		if err = deepEqual(bsb, bss); err == nil {
			if testVerbose {
				t.Logf("++++++++ Bytes match")
			}
		} else {
			xs := "--------"
			if reflect.ValueOf(v).Kind() == reflect.Map {
				xs = "        "
				if testVerbose {
					t.Logf("%s FAIL - bytes do not match, but it's a map (ok - dependent on ordering): %v", xs, err)
				}
			} else {
				t.Logf("%s FAIL - bytes do not match and is not a map (bad): %v", xs, err)
				t.FailNow()
			}
			if testVerbose {
				t.Logf("%s   FROM_FILE: %4d] %v", xs, len(bss), bss)
				t.Logf("%s     ENCODED: %4d] %v", xs, len(bsb), bsb)
			}
		}
		testReleaseBytes(bsb)
	}
	bh.MapType = oldMapType
}

// To test MsgpackSpecRpc, we test 3 scenarios:
//    - Go Client to Go RPC Service (contained within TestMsgpackRpcSpec)
//    - Go client to Python RPC Service (contained within doTestMsgpackRpcSpecGoClientToPythonSvc)
//    - Python Client to Go RPC Service (contained within doTestMsgpackRpcSpecPythonClientToGoSvc)
//
// This allows us test the different calling conventions
//    - Go Service requires only one argument
//    - Python Service allows multiple arguments

func doTestMsgpackRpcSpecGoClientToPythonSvc(t *testing.T, h Handle) {
	if testSkipRPCTests {
		t.Skip(testSkipRPCTestsMsg)
	}
	defer testSetup(t, &h)()

	// openPorts are between 6700 and 6800
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	openPort := strconv.FormatInt(6700+r.Int63n(99), 10)
	// openPort := "6792"
	cmd := exec.Command("python", "test.py", "rpc-server", openPort, "4")
	testCheckErr(t, cmd.Start())
	bs, err2 := net.Dial("tcp", ":"+openPort)
	maxSleepTime := 500 * time.Millisecond
	iterSleepTime := 5 * time.Millisecond
	for i := 0; i < int(maxSleepTime/iterSleepTime) && err2 != nil; i++ {
		time.Sleep(iterSleepTime) // time for python rpc server to start
		bs, err2 = net.Dial("tcp", ":"+openPort)
	}
	testCheckErr(t, err2)
	cc := MsgpackSpecRpc.ClientCodec(testReadWriteCloser(bs), h)
	cl := rpc.NewClientWithCodec(cc)
	defer cl.Close()
	var rstr string
	testCheckErr(t, cl.Call("EchoStruct", TestRpcABC{"Aa", "Bb", "Cc"}, &rstr))
	//testCheckEqual(t, rstr, "{'A': 'Aa', 'B': 'Bb', 'C': 'Cc'}")
	var mArgs MsgpackSpecRpcMultiArgs = []interface{}{"A1", "B2", "C3"}
	testCheckErr(t, cl.Call("Echo123", mArgs, &rstr))
	testCheckEqual(t, rstr, "1:A1 2:B2 3:C3", "rstr=")
	cmd.Process.Kill()
}

func doTestMsgpackRpcSpecPythonClientToGoSvc(t *testing.T, h Handle) {
	if testSkipRPCTests {
		t.Skip(testSkipRPCTestsMsg)
	}
	defer testSetup(t, &h)()
	// seems 10ms is not enough for test run; set high to 1 second, and client will close when done
	exitSleepDur := 1000 * time.Millisecond
	port := doTestCodecRpcOne(t, MsgpackSpecRpc, h, false, exitSleepDur)
	cmd := exec.Command("python", "test.py", "rpc-client-go-service", strconv.Itoa(port))
	var cmdout []byte
	var err error
	if cmdout, err = cmd.CombinedOutput(); err != nil {
		t.Logf("-------- Error running test.py rpc-client-go-service. Err: %v", err)
		t.Logf("         %v", string(cmdout))
		t.FailNow()
	}
	testCheckEqual(t, string(cmdout),
		fmt.Sprintf("%#v\n%#v\n", []string{"A1", "B2", "C3"}, TestRpcABC{"Aa", "Bb", "Cc"}), "cmdout=")
}

func doTestSwallowAndZero(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	v1 := newTestStrucFlex(testDepth, testNumRepeatString, false, false, testMapStringKeyOnly)
	var b1 []byte

	e1 := NewEncoderBytes(&b1, h)
	e1.MustEncode(v1)
	d1 := NewDecoderBytes(b1, h)
	d1.swallow()
	if d1.r().numread() != uint(len(b1)) {
		t.Logf("swallow didn't consume all encoded bytes: %v out of %v", d1.r().numread(), len(b1))
		t.FailNow()
	}
	setZero(v1)
	testDeepEqualErr(v1, &TestStrucFlex{}, t, "filled-and-zeroed")
}

func doTestRawExt(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	var b []byte
	var v RawExt // interface{}
	_, isJson := h.(*JsonHandle)
	_, isCbor := h.(*CborHandle)
	bh := testBasicHandle(h)
	// isValuer := isJson || isCbor
	// _ = isValuer
	for _, r := range []RawExt{
		{Tag: 99, Value: "9999", Data: []byte("9999")},
	} {
		e := NewEncoderBytes(&b, h)
		e.MustEncode(&r)
		// fmt.Printf(">>>> rawext: isnil? %v, %d - %v\n", b == nil, len(b), b)
		d := NewDecoderBytes(b, h)
		d.MustDecode(&v)
		var r2 = r
		switch {
		case isJson:
			r2.Tag = 0
			r2.Data = nil
		case isCbor:
			r2.Data = nil
		default:
			r2.Value = nil
		}
		testDeepEqualErr(v, r2, t, "rawext-default")
		// switch h.(type) {
		// case *JsonHandle:
		// 	testDeepEqualErr(r.Value, v, t, "rawext-json")
		// default:
		// 	var r2 = r
		// 	if isValuer {
		// 		r2.Data = nil
		// 	} else {
		// 		r2.Value = nil
		// 	}
		// 	testDeepEqualErr(v, r2, t, "rawext-default")
		// }
	}

	// Add testing for Raw also
	if b != nil {
		b = b[:0]
	}
	oldRawMode := bh.Raw
	defer func() { bh.Raw = oldRawMode }()
	bh.Raw = true

	var v2 Raw
	for _, s := range []string{
		"goodbye",
		"hello",
	} {
		e := NewEncoderBytes(&b, h)
		e.MustEncode(&s)
		// fmt.Printf(">>>> rawext: isnil? %v, %d - %v\n", b == nil, len(b), b)
		var r Raw = make([]byte, len(b))
		copy(r, b)
		d := NewDecoderBytes(b, h)
		d.MustDecode(&v2)
		testDeepEqualErr(v2, r, t, "raw-default")
	}

}

// func doTestTimeExt(t *testing.T, h Handle) {
// 	var t = time.Now()
// 	// add time ext to the handle
// }

func doTestMapStructKey(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	var b []byte
	var v interface{} // map[stringUint64T]wrapUint64Slice // interface{}
	bh := testBasicHandle(h)
	m := map[stringUint64T]wrapUint64Slice{
		stringUint64T{"55555", 55555}: []wrapUint64{12345},
		stringUint64T{"333", 333}:     []wrapUint64{123},
	}
	oldCanonical := bh.Canonical
	oldMapType := bh.MapType
	defer func() {
		bh.Canonical = oldCanonical
		bh.MapType = oldMapType
	}()

	bh.MapType = reflect.TypeOf((*map[stringUint64T]wrapUint64Slice)(nil)).Elem()
	for _, bv := range [2]bool{true, false} {
		b, v = nil, nil
		bh.Canonical = bv
		e := NewEncoderBytes(&b, h)
		e.MustEncode(m)
		d := NewDecoderBytes(b, h)
		d.MustDecode(&v)
		testDeepEqualErr(v, m, t, "map-structkey")
	}
}

func doTestDecodeNilMapValue(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	type Struct struct {
		Field map[uint16]map[uint32]struct{}
	}

	bh := testBasicHandle(h)
	defer func(t reflect.Type) {
		bh.MapType = t
	}(bh.MapType)

	// this test expects that nil doesn't result in deleting entries

	_, isJsonHandle := h.(*JsonHandle)

	toEncode := Struct{Field: map[uint16]map[uint32]struct{}{
		1: nil,
	}}

	bs := testMarshalErr(toEncode, h, t, "-")
	if isJsonHandle && testVerbose {
		t.Logf("json encoded: %s\n", bs)
	}

	var decoded Struct
	testUnmarshalErr(&decoded, bs, h, t, "-")
	if !reflect.DeepEqual(decoded, toEncode) {
		t.Logf("Decoded value %#v != %#v", decoded, toEncode)
		t.FailNow()
	}
	testReleaseBytes(bs)

	__doTestDecodeNilMapEntryValue(t, h)
}

func __doTestDecodeNilMapEntryValue(t *testing.T, h Handle) {
	type Entry struct{}
	type Entries struct {
		Map map[string]*Entry
	}

	c := Entries{
		Map: map[string]*Entry{
			"nil":   nil,
			"empty": &Entry{},
		},
	}

	bs, err := testMarshal(&c, h)
	if err != nil {
		t.Logf("failed to encode: %v", err)
		t.FailNow()
	}

	var f Entries
	err = testUnmarshal(&f, bs, h)
	if err != nil {
		t.Logf("failed to decode: %v", err)
		t.FailNow()
	}

	if !reflect.DeepEqual(c, f) {
		t.Logf("roundtrip encoding doesn't match\nexpected: %v\nfound:    %v\n\n"+
			"empty value: %#+v\nnil value:   %#+v",
			c, f, f.Map["empty"], f.Map["nil"])
		t.FailNow()
	}
	testReleaseBytes(bs)
}

func doTestEmbeddedFieldPrecedence(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	type Embedded struct {
		Field byte
	}
	type Struct struct {
		Field byte
		Embedded
	}
	toEncode := Struct{
		Field:    1,
		Embedded: Embedded{Field: 2},
	}
	_, isJsonHandle := h.(*JsonHandle)
	handle := testBasicHandle(h)
	oldMapType := handle.MapType
	defer func() { handle.MapType = oldMapType }()

	handle.MapType = reflect.TypeOf(map[interface{}]interface{}(nil))

	bs, err := testMarshal(toEncode, h)
	if err != nil {
		t.Logf("Error encoding: %v, Err: %v", toEncode, err)
		t.FailNow()
	}

	var decoded Struct
	err = testUnmarshal(&decoded, bs, h)
	if err != nil {
		t.Logf("Error decoding: %v", err)
		t.FailNow()
	}

	if decoded.Field != toEncode.Field {
		t.Logf("Decoded result %v != %v", decoded.Field, toEncode.Field) // hex to look at what was encoded
		if isJsonHandle {
			t.Logf("JSON encoded as: %s", bs) // hex to look at what was encoded
		}
		t.FailNow()
	}
	testReleaseBytes(bs)
}

func doTestLargeContainerLen(t *testing.T, h Handle) {
	defer testSetup(t, &h)()

	// This test can take a while if run multiple times in a loop, as it creates
	// large maps/slices. Use t.Short() appropriately to limit its execution time.

	// Note that this test does not make sense for formats which do not pre-record
	// the length, like json, or cbor with indefinite length.
	if _, ok := h.(*JsonHandle); ok {
		t.Skipf("skipping as json doesn't support prefixed lengths")
	}
	if c, ok := h.(*CborHandle); ok && c.IndefiniteLength {
		t.Skipf("skipping as cbor Indefinite Length doesn't use prefixed lengths")
	}

	bh := testBasicHandle(h)

	var sizes []int
	// sizes = []int{
	// 	0, 1,
	// 	math.MaxInt8, math.MaxInt8 + 4, math.MaxInt8 - 4,
	// 	math.MaxInt16, math.MaxInt16 + 4, math.MaxInt16 - 4,
	// 	math.MaxInt32, math.MaxInt32 - 4,
	// 	// math.MaxInt32 + 4, // bombs on 32-bit
	// 	// math.MaxInt64, math.MaxInt64 - 4, // bombs on 32-bit

	// 	math.MaxUint8, math.MaxUint8 + 4, math.MaxUint8 - 4,
	// 	math.MaxUint16, math.MaxUint16 + 4, math.MaxUint16 - 4,
	// 	// math.MaxUint32, math.MaxUint32 + 4, math.MaxUint32 - 4, // bombs on 32-bit
	// }
	sizes = []int{
		// ensure in ascending order (as creating mm below requires it)
		0,
		1,
		math.MaxInt8 + 4,
		math.MaxUint8 + 4,
	}
	if !testing.Short() {
		sizes = append(sizes, math.MaxUint16+4) // math.MaxInt16+4, math.MaxUint16+4
	}

	m := make(map[int][]struct{})
	for _, i := range sizes {
		m[i] = make([]struct{}, i)
	}

	bs := testMarshalErr(m, h, t, "-slices")
	var m2 = make(map[int][]struct{})
	testUnmarshalErr(m2, bs, h, t, "-slices")
	testDeepEqualErr(m, m2, t, "-slices")

	d, x := testSharedCodecDecoder(bs, h, bh)
	bs2 := d.d.nextValueBytes([]byte{})
	testSharedCodecDecoderAfter(d, x, bh)
	testDeepEqualErr(bs, bs2, t, "nextvaluebytes-slices")
	// if len(bs2) != 0 || len(bs2) != len(bs) { }
	testReleaseBytes(bs)

	// requires sizes to be in ascending order
	mm := make(map[int]struct{})
	for _, i := range sizes {
		for j := len(mm); j < i; j++ {
			mm[j] = struct{}{}
		}
		bs = testMarshalErr(mm, h, t, "-map")
		var mm2 = make(map[int]struct{})
		testUnmarshalErr(mm2, bs, h, t, "-map")
		testDeepEqualErr(mm, mm2, t, "-map")

		d, x = testSharedCodecDecoder(bs, h, bh)
		bs2 = d.d.nextValueBytes([]byte{})
		testSharedCodecDecoderAfter(d, x, bh)
		testDeepEqualErr(bs, bs2, t, "nextvaluebytes-map")
		testReleaseBytes(bs)
	}

	// do same tests for large strings (encoded as symbols or not)
	// skip if 32-bit or not using unsafe mode
	if safeMode || (32<<(^uint(0)>>63)) < 64 {
		return
	}

	// now, want to do tests for large strings, which
	// could be encoded as symbols.
	// to do this, we create a simple one-field struct,
	// use use flags to switch from symbols to non-symbols

	hbinc, okbinc := h.(*BincHandle)
	if okbinc {
		oldAsSymbols := hbinc.AsSymbols
		defer func() { hbinc.AsSymbols = oldAsSymbols }()
	}
	inOutLen := math.MaxUint16 * 3 / 2
	if testing.Short() {
		inOutLen = math.MaxUint8 * 2 // math.MaxUint16 / 16
	}
	var out = make([]byte, 0, inOutLen)
	var in []byte = make([]byte, inOutLen)
	for i := range in {
		in[i] = 'A'
	}

	e := NewEncoder(nil, h)

	sizes = []int{
		0, 1, 4, 8, 12, 16, 28, 32, 36,
		math.MaxInt8 - 4, math.MaxInt8, math.MaxInt8 + 4,
		math.MaxUint8, math.MaxUint8 + 4, math.MaxUint8 - 4,
	}
	if !testing.Short() {
		sizes = append(sizes,
			math.MaxInt16-4, math.MaxInt16, math.MaxInt16+4,
			math.MaxUint16, math.MaxUint16+4, math.MaxUint16-4)
	}

	for _, i := range sizes {
		var m1, m2 map[string]bool
		m1 = make(map[string]bool, 1)
		// var s1 = stringView(in[:i])
		var s1 = string(in[:i])
		// fmt.Printf("testcontainerlen: large string: i: %v, |%s|\n", i, s1)
		m1[s1] = true

		if okbinc {
			hbinc.AsSymbols = 2
		}
		out = out[:0]
		e.ResetBytes(&out)
		e.MustEncode(m1)
		// bs, _ = testMarshalErr(m1, h, t, "-")
		m2 = make(map[string]bool, 1)
		testUnmarshalErr(m2, out, h, t, "no-symbols-string")
		testDeepEqualErr(m1, m2, t, "no-symbols-string")

		d, x = testSharedCodecDecoder(out, h, bh)
		bs2 = d.d.nextValueBytes([]byte{})
		testSharedCodecDecoderAfter(d, x, bh)
		testDeepEqualErr(out, bs2, t, "nextvaluebytes-no-symbols-string")

		if okbinc {
			// now, do as symbols
			hbinc.AsSymbols = 1
			out = out[:0]
			e.ResetBytes(&out)
			e.MustEncode(m1)
			// bs, _ = testMarshalErr(m1, h, t, "-")
			m2 = make(map[string]bool, 1)
			testUnmarshalErr(m2, out, h, t, "symbols-string")
			testDeepEqualErr(m1, m2, t, "symbols-string")

			d, x = testSharedCodecDecoder(out, h, bh)
			bs2 = d.d.nextValueBytes([]byte{})
			testSharedCodecDecoderAfter(d, x, bh)
			testDeepEqualErr(out, bs2, t, "nextvaluebytes-symbols-string")
			hbinc.AsSymbols = 2
		}
	}

	// test out extensions with large output
	var xl, xl2 testUintToBytes
	xl = testUintToBytes(inOutLen)
	bs = testMarshalErr(xl, h, t, "-large-extension-bytes")
	testUnmarshalErr(&xl2, bs, h, t, "-large-extension-bytes")
	testDeepEqualErr(xl, xl2, t, "-large-extension-bytes")

	d, x = testSharedCodecDecoder(bs, h, bh)
	bs2 = d.d.nextValueBytes([]byte{})
	testSharedCodecDecoderAfter(d, x, bh)
	testDeepEqualErr(bs, bs2, t, "nextvaluebytes-large-extension-bytes")

	xl = testUintToBytes(0) // so it's WriteExt returns nil
	bs = testMarshalErr(xl, h, t, "-large-extension-bytes")
	testUnmarshalErr(&xl2, bs, h, t, "-large-extension-bytes")
	testDeepEqualErr(xl, xl2, t, "-large-extension-bytes")
}

func testRandomFillRV(v reflect.Value) {
	fneg := func() int64 {
		i := rand.Intn(1)
		if i == 1 {
			return 1
		}
		return -1
	}

	if v.Type() == timeTyp {
		v.Set(reflect.ValueOf(timeToCompare1))
		return
	}

	switch v.Kind() {
	case reflect.Invalid:
	case reflect.Ptr:
		if v.IsNil() {
			v.Set(reflect.New(v.Type().Elem()))
		}
		testRandomFillRV(v.Elem())
	case reflect.Interface:
		if v.IsNil() {
			v.Set(reflect.ValueOf("nothing"))
		} else {
			testRandomFillRV(v.Elem())
		}
	case reflect.Struct:
		for i, n := 0, v.NumField(); i < n; i++ {
			testRandomFillRV(v.Field(i))
		}
	case reflect.Slice:
		if v.IsNil() {
			v.Set(reflect.MakeSlice(v.Type(), 4, 4))
		}
		fallthrough
	case reflect.Array:
		for i, n := 0, v.Len(); i < n; i++ {
			testRandomFillRV(v.Index(i))
		}
	case reflect.Map:
		if v.IsNil() {
			v.Set(reflect.MakeMap(v.Type()))
		}
		if v.Len() == 0 {
			kt, vt := v.Type().Key(), v.Type().Elem()
			for i := 0; i < 4; i++ {
				k0 := reflect.New(kt).Elem()
				v0 := reflect.New(vt).Elem()
				testRandomFillRV(k0)
				testRandomFillRV(v0)
				v.SetMapIndex(k0, v0)
			}
		} else {
			for _, k := range v.MapKeys() {
				testRandomFillRV(v.MapIndex(k))
			}
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v.SetInt(fneg() * rand.Int63n(127))
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v.SetUint(uint64(rand.Int63n(255)))
	case reflect.Bool:
		v.SetBool(fneg() == 1)
	case reflect.Float32, reflect.Float64:
		v.SetFloat(float64(fneg()) * float64(rand.Float32()))
	case reflect.Complex64, reflect.Complex128:
		v.SetComplex(complex(float64(fneg())*float64(rand.Float32()), 0))
	case reflect.String:
		// ensure this string can test the extent of json string decoding
		v.SetString(strings.Repeat(strconv.FormatInt(rand.Int63n(99), 10), rand.Intn(8)) +
			"- ABC \x41=\x42 \u2318 - \r \b \f - \u2028 and \u2029 .")
	default:
		panic(fmt.Errorf("testRandomFillRV: unsupported type: %v", v.Kind()))
	}
}

func doTestTime(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	name := h.Name()
	// test time which uses the time.go implementation (ie Binc)
	var tt, tt2 time.Time
	// time in 1990
	tt = time.Unix(20*366*24*60*60, 1000*900).In(time.FixedZone("UGO", -5*60*60))
	// fmt.Printf("time tt: %v\n", tt)
	b := testMarshalErr(tt, h, t, "time-"+name)
	testUnmarshalErr(&tt2, b, h, t, "time-"+name)
	// per go documentation, test time with .Equal not ==
	if !tt2.Equal(tt) {
		t.Logf("%s: values not equal: 1: %v, 2: %v", name, tt2, tt)
		t.FailNow()
	}
	// testDeepEqualErr(tt.UTC(), tt2, t, "time-"+name)
	testReleaseBytes(b)
}

func doTestUintToInt(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	name := h.Name()
	var golden = [...]int64{
		0, 1, 22, 333, 4444, 55555, 666666,
		// msgpack ones
		24, 128,
		// standard ones
		math.MaxUint8, math.MaxUint8 + 4, math.MaxUint8 - 4,
		math.MaxUint16, math.MaxUint16 + 4, math.MaxUint16 - 4,
		math.MaxUint32, math.MaxUint32 + 4, math.MaxUint32 - 4,
		math.MaxInt8, math.MaxInt8 + 4, math.MaxInt8 - 4,
		math.MaxInt16, math.MaxInt16 + 4, math.MaxInt16 - 4,
		math.MaxInt32, math.MaxInt32 + 4, math.MaxInt32 - 4,
		math.MaxInt64, math.MaxInt64 - 4,
	}
	var i int64
	var ui, ui2 uint64
	var fi float64
	var b []byte
	for _, v := range golden {
		i = v
		ui = 0
		b = testMarshalErr(i, h, t, "int2uint-"+name)
		testUnmarshalErr(&ui, b, h, t, "int2uint-"+name)
		if ui != uint64(i) {
			t.Logf("%s: values not equal: %v, %v", name, ui, uint64(i))
			t.FailNow()
		}
		testReleaseBytes(b)

		ui = uint64(i)
		i = 0
		b = testMarshalErr(ui, h, t, "uint2int-"+name)
		testUnmarshalErr(&i, b, h, t, "uint2int-"+name)
		if i != int64(ui) {
			t.Logf("%s: values not equal: %v, %v", name, i, int64(ui))
			t.FailNow()
		}
		testReleaseBytes(b)

		if v == math.MaxInt64 {
			ui = uint64(-(v - 1))
		} else {
			ui = uint64(-v)
		}

		b = testMarshalErr(ui, h, t, "negint2uint-"+name)
		testUnmarshalErr(&ui2, b, h, t, "negint2uint-"+name)
		if ui2 != ui {
			t.Logf("%s: values not equal: %v, %v", name, ui2, ui)
			t.FailNow()
		}
		testReleaseBytes(b)

		fi = 0
		b = testMarshalErr(i, h, t, "int2float-"+name)
		testUnmarshalErr(&fi, b, h, t, "int2float-"+name)
		testReleaseBytes(b)
		if fi != float64(i) {
			t.Logf("%s: values not equal: %v, %v", name, fi, float64(i))
			t.FailNow()
		}
	}
}

func doTestDifferentMapOrSliceType(t *testing.T, h Handle) {
	if !testRecoverPanicToErr {
		t.Skip(testSkipIfNotRecoverPanicToErrMsg)
	}
	defer testSetup(t, &h)()

	if mh, ok := h.(*MsgpackHandle); ok {
		defer func(b bool) { mh.RawToString = b }(mh.RawToString)
		mh.RawToString = true
	}

	name := h.Name()

	// - maptype, slicetype: diff from map[string]intf, map[intf]intf or []intf, etc
	//   include map[interface{}]string where some keys are []byte.
	//   To test, take a sequence of []byte and string, and decode into []string and []interface.
	//   Also, decode into map[string]string, map[string]interface{}, map[interface{}]string

	bh := testBasicHandle(h)
	defer func(oldM, oldS reflect.Type) { bh.MapType, bh.SliceType = oldM, oldS }(bh.MapType, bh.SliceType)

	// test for arrays first
	fnArr := func() {
		defer func(b1, b2 bool) { bh.StringToRaw, _ = b1, b2 }(bh.StringToRaw, false)
		bh.StringToRaw = false
		vi := []interface{}{
			"hello 1",
			float64(111.0),
			"hello 3",
			float64(333.0),
			"hello 5",
		}

		var mi, mi2 map[string]float64
		mi = make(map[string]float64)

		mi[vi[0].(string)] = vi[1].(float64)
		mi[vi[2].(string)] = vi[3].(float64)

		var v4a, v4a2 testMbsArr4T
		copy(v4a[:], vi)

		b := testMarshalErr(v4a, h, t, "-")
		testUnmarshalErr(&mi2, b, h, t, "-")
		testDeepEqualErr(mi2, mi, t, "-")
		testUnmarshalErr(&v4a2, b, h, t, "-")
		testDeepEqualErr(v4a2, v4a, t, "-")
		testReleaseBytes(b)

		var v0a, v0a2 testMbsArr0T
		copy(v0a[:], vi)

		mi2 = nil
		b = testMarshalErr(v0a, h, t, "-")
		testUnmarshalErr(&mi2, b, h, t, "-")
		testDeepEqualErr(mi2, map[string]float64{}, t, "-")
		testUnmarshalErr(&v0a2, b, h, t, "-")
		testDeepEqualErr(v0a2, v0a, t, "-")
		testReleaseBytes(b)

		{
			var v5a testMbsArr5T
			copy(v5a[:], vi)
			b, err := testMarshal(v5a, h)
			testReleaseBytes(b)
			if err == nil || !strings.Contains(err.Error(), "mapBySlice requires even slice length") {
				t.Logf("mapBySlice for odd length array fail: expected mapBySlice error, got: %v", err)
				t.FailNow()
			}
		}
	}
	fnArr()

	var b []byte

	var vi = []interface{}{
		"hello 1",
		[]byte("hello 2"),
		"hello 3",
		[]byte("hello 4"),
		"hello 5",
	}

	var vs []string
	var v2i, v2s testMbsT
	var v2ss testMbsCustStrT

	// encode it as a map or as a slice
	for i, v := range vi {
		vv, ok := v.(string)
		if !ok {
			vv = string(v.([]byte))
		}
		vs = append(vs, vv)
		v2i = append(v2i, v, strconv.FormatInt(int64(i+1), 10))
		v2s = append(v2s, vv, strconv.FormatInt(int64(i+1), 10))
		v2ss = append(v2ss, testCustomStringT(vv), testCustomStringT(strconv.FormatInt(int64(i+1), 10)))
	}

	var v2d interface{}

	// encode vs as a list, and decode into a list and compare
	var goldSliceS = []string{"hello 1", "hello 2", "hello 3", "hello 4", "hello 5"}
	var goldSliceI = []interface{}{"hello 1", "hello 2", "hello 3", "hello 4", "hello 5"}
	var goldSlice = []interface{}{goldSliceS, goldSliceI}
	for j, g := range goldSlice {
		bh.SliceType = reflect.TypeOf(g)
		name := fmt.Sprintf("slice-%s-%v", name, j+1)
		b = testMarshalErr(vs, h, t, name)
		v2d = nil
		// v2d = reflect.New(bh.SliceType).Elem().Interface()
		testUnmarshalErr(&v2d, b, h, t, name)
		testDeepEqualErr(v2d, goldSlice[j], t, name)
		testReleaseBytes(b)
	}

	// to ensure that we do not use fast-path for map[intf]string, use a custom string type (for goldMapIS).
	// this will allow us to test out the path that sees a []byte where a map has an interface{} type,
	// and convert it to a string for the decoded map key.

	// encode v2i as a map, and decode into a map and compare
	var goldMapSS = map[string]string{"hello 1": "1", "hello 2": "2", "hello 3": "3", "hello 4": "4", "hello 5": "5"}
	var goldMapSI = map[string]interface{}{"hello 1": "1", "hello 2": "2", "hello 3": "3", "hello 4": "4", "hello 5": "5"}
	var goldMapIS = map[interface{}]testCustomStringT{"hello 1": "1", "hello 2": "2", "hello 3": "3", "hello 4": "4", "hello 5": "5"}
	var goldMap = []interface{}{goldMapSS, goldMapSI, goldMapIS}
	for j, g := range goldMap {
		bh.MapType = reflect.TypeOf(g)
		name := fmt.Sprintf("map-%s-%v", name, j+1)
		// for formats that clearly differentiate binary from string, use v2i
		// else use the v2s (with all strings, no []byte)
		v2d = nil
		// v2d = reflect.New(bh.MapType).Elem().Interface()
		switch h.(type) {
		case *MsgpackHandle, *BincHandle, *CborHandle:
			b = testMarshalErr(v2i, h, t, name)
			testUnmarshalErr(&v2d, b, h, t, name)
			testDeepEqualErr(v2d, goldMap[j], t, name)
			testReleaseBytes(b)
		default:
			b = testMarshalErr(v2s, h, t, name)
			testUnmarshalErr(&v2d, b, h, t, name)
			testDeepEqualErr(v2d, goldMap[j], t, name)
			testReleaseBytes(b)
			b = testMarshalErr(v2ss, h, t, name)
			v2d = nil
			testUnmarshalErr(&v2d, b, h, t, name)
			testDeepEqualErr(v2d, goldMap[j], t, name)
			testReleaseBytes(b)
		}
	}

	// encode []byte and decode into one with len < cap.
	// get the slices > decDefSliceCap*2, so things like json can go much higher
	{
		var bs1 = []byte("abcdefghijklmnopqrstuvwxyz")
		b := testMarshalErr(bs1, h, t, "enc-bytes")

		var bs2 = make([]byte, 32, 32)
		testUnmarshalErr(bs2, b, h, t, "dec-bytes")
		testDeepEqualErr(bs1, bs2[:len(bs1)], t, "cmp-enc-dec-bytes")

		testUnmarshalErr(&bs2, b, h, t, "dec-bytes-2")
		testDeepEqualErr(bs1, bs2, t, "cmp-enc-dec-bytes-2")

		bs2 = bs2[:2:4]
		testUnmarshalErr(&bs2, b, h, t, "dec-bytes-2")
		testDeepEqualErr(bs1, bs2, t, "cmp-enc-dec-bytes-2")

		bs2 = nil
		testUnmarshalErr(&bs2, b, h, t, "dec-bytes-2")
		testDeepEqualErr(bs1, bs2, t, "cmp-enc-dec-bytes-2")

		bs1 = []byte{}
		b = testMarshalErr(bs1, h, t, "enc-bytes")

		bs2 = nil
		testUnmarshalErr(&bs2, b, h, t, "dec-bytes-2")
		testDeepEqualErr(bs1, bs2, t, "cmp-enc-dec-bytes-2")

		type Ti32 int32
		v1 := []Ti32{
			9, 99, 999, 9999, 99999, 999999,
			9, 99, 999, 9999, 99999, 999999,
			9, 99, 999, 9999, 99999, 999999,
		}
		b = testMarshalErr(v1, h, t, "enc-Ti32")

		v2 := make([]Ti32, 20, 20)
		testUnmarshalErr(v2, b, h, t, "dec-Ti32")
		testDeepEqualErr(v1, v2[:len(v1)], t, "cmp-enc-dec-Ti32")

		testUnmarshalErr(&v2, b, h, t, "dec-Ti32-2")
		testDeepEqualErr(v1, v2, t, "cmp-enc-dec-Ti32-2")

		v2 = v2[:1:3]
		testUnmarshalErr(&v2, b, h, t, "dec-Ti32-2")
		testDeepEqualErr(v1, v2, t, "cmp-enc-dec-Ti32-2")

		v2 = nil
		testUnmarshalErr(&v2, b, h, t, "dec-Ti32-2")
		testDeepEqualErr(v1, v2, t, "cmp-enc-dec-Ti32-2")

		v1 = []Ti32{}
		b = testMarshalErr(v1, h, t, "enc-Ti32")

		v2 = nil
		testUnmarshalErr(&v2, b, h, t, "dec-Ti32-2")
		testDeepEqualErr(v1, v2, t, "cmp-enc-dec-Ti32-2")
	}
}

func doTestScalars(t *testing.T, h Handle) {
	defer testSetup(t, &h)()

	if mh, ok := h.(*MsgpackHandle); ok {
		defer func(b bool) { mh.RawToString = b }(mh.RawToString)
		mh.RawToString = true
	}

	// for each scalar:
	// - encode its ptr
	// - encode it (non-ptr)
	// - check that bytes are same
	// - make a copy (using reflect)
	// - check that same
	// - set zero on it
	// - check that its equal to 0 value
	// - decode into new
	// - compare to original

	bh := testBasicHandle(h)
	if !bh.Canonical {
		bh.Canonical = true
		defer func() { bh.Canonical = false }()
	}

	var bzero = testMarshalErr(nil, h, t, "nil-enc")

	vi := []interface{}{
		int(0),
		int8(0),
		int16(0),
		int32(0),
		int64(0),
		uint(0),
		uint8(0),
		uint16(0),
		uint32(0),
		uint64(0),
		uintptr(0),
		float32(0),
		float64(0),
		bool(false),
		time.Time{},
		string(""),
		[]byte(nil),
	}
	for _, v := range fastpathAv {
		vi = append(vi, reflect.Zero(v.rt).Interface())
	}
	for _, v := range vi {
		rv := reflect.New(reflect.TypeOf(v)).Elem()
		testRandomFillRV(rv)
		v = rv.Interface()

		rv2 := reflect.New(rv.Type())
		rv2.Elem().Set(rv)
		vp := rv2.Interface()

		var tname string
		switch rv.Kind() {
		case reflect.Map:
			tname = "map[" + rv.Type().Key().Name() + "]" + rv.Type().Elem().Name()
		case reflect.Slice:
			tname = "[]" + rv.Type().Elem().Name()
		default:
			tname = rv.Type().Name()
		}

		var b, b1, b2 []byte
		b1 = testMarshalErr(v, h, t, tname+"-enc")
		// store b1 into b, as b1 slice is reused for next marshal
		b = make([]byte, len(b1))
		copy(b, b1)
		b2 = testMarshalErr(vp, h, t, tname+"-enc-ptr")
		testDeepEqualErr(b1, b2, t, tname+"-enc-eq")

		// decode the nil value into rv2, and test that it is the zero value
		setZero(vp)
		testDeepEqualErr(rv2.Elem().Interface(), reflect.Zero(rv.Type()).Interface(), t, tname+"-enc-eq-zero-ref")

		testUnmarshalErr(vp, b, h, t, tname+"-dec")
		testDeepEqualErr(rv2.Elem().Interface(), v, t, tname+"-dec-eq")

		// test that we can decode an encoded nil into it
		testUnmarshalErr(vp, bzero, h, t, tname+"-dec-from-enc-nil")
		testDeepEqualErr(rv2.Elem().Interface(), reflect.Zero(rv.Type()).Interface(), t, tname+"-dec-from-enc-nil")
		testReleaseBytes(b1)
		testReleaseBytes(b2)
	}

	// test setZero for *Raw and reflect.Value
	var r0 Raw
	var r = Raw([]byte("hello"))
	setZero(&r)
	testDeepEqualErr(r, r0, t, "raw-zeroed")

	r = Raw([]byte("hello"))
	var rv = reflect.ValueOf(&r)
	setZero(rv)
	// note: we cannot test reflect.Value's because they might point to different pointers
	// and reflect.DeepEqual doesn't honor that.
	// testDeepEqualErr(rv, reflect.ValueOf(&r0), t, "raw-reflect-zeroed")
	testDeepEqualErr(rv.Interface(), &r0, t, "raw-reflect-zeroed")
	testReleaseBytes(bzero)
}

func doTestIntfMapping(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	name := h.Name()
	rti := reflect.TypeOf((*testIntfMapI)(nil)).Elem()
	defer func() { testBasicHandle(h).Intf2Impl(rti, nil) }()

	type T9 struct {
		I testIntfMapI
	}

	for i, v := range []testIntfMapI{
		// Use a valid string to test some extents of json string decoding
		&testIntfMapT1{"ABC \x41=\x42 \u2318 - \r \b \f - \u2028 and \u2029 ."},
		testIntfMapT2{"DEF"},
	} {
		if err := testBasicHandle(h).Intf2Impl(rti, reflect.TypeOf(v)); err != nil {
			t.Logf("Error mapping %v to %T", rti, v)
			t.FailNow()
		}
		var v1, v2 T9
		v1 = T9{v}
		b := testMarshalErr(v1, h, t, name+"-enc-"+strconv.Itoa(i))
		testUnmarshalErr(&v2, b, h, t, name+"-dec-"+strconv.Itoa(i))
		testDeepEqualErr(v1, v2, t, name+"-dec-eq-"+strconv.Itoa(i))
		testReleaseBytes(b)
	}
}

func doTestOmitempty(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	name := h.Name()
	if testBasicHandle(h).StructToArray {
		t.Skipf("skipping OmitEmpty test when StructToArray=true")
	}
	type T1 struct {
		A int  `codec:"a"`
		B *int `codec:"b,omitempty"`
		C int  `codec:"c,omitempty"`
	}
	type T2 struct {
		A int `codec:"a"`
	}
	var v1 T1
	var v2 T2
	b1 := testMarshalErr(v1, h, t, name+"-omitempty")
	b2 := testMarshalErr(v2, h, t, name+"-no-omitempty-trunc")
	testDeepEqualErr(b1, b2, t, name+"-omitempty-cmp")
	testReleaseBytes(b1)
	testReleaseBytes(b2)
}

func doTestMissingFields(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	name := h.Name()
	if codecgen {
		t.Skipf("skipping Missing Fields tests as it is not honored by codecgen")
	}
	if testBasicHandle(h).StructToArray {
		t.Skipf("skipping Missing Fields test when StructToArray=true")
	}
	// encode missingFielderT2, decode into missingFielderT1, encode it out again, decode into new missingFielderT2, compare
	v1 := missingFielderT2{S: "true seven eight", B: true, F: 777.0, I: -888}
	b1 := testMarshalErr(v1, h, t, name+"-missing-enc-2")

	var v2 missingFielderT1
	testUnmarshalErr(&v2, b1, h, t, name+"-missing-dec-1")

	b2 := testMarshalErr(&v2, h, t, name+"-missing-enc-1")

	var v3 missingFielderT2
	testUnmarshalErr(&v3, b2, h, t, name+"-missing-dec-2")

	testDeepEqualErr(v1, v3, t, name+"-missing-cmp-2")

	testReleaseBytes(b1)
	testReleaseBytes(b2)

	v4 := missingFielderT11{s1: "s111", S2: "S222"}
	b1 = testMarshalErr(v4, h, t, name+"-missing-enc-11")
	var m4 map[string]string
	testUnmarshalErr(&m4, b1, h, t, name+"-missing-dec-11")
	testDeepEqualErr(m4, map[string]string{"s1": "s111", "S2": "S222"}, t, name+"-missing-cmp-11")

	testReleaseBytes(b1)

	// test canonical interaction - with structs having some missing fields and some regular fields
	bh := testBasicHandle(h)

	defer func(c bool) {
		bh.Canonical = c
	}(bh.Canonical)
	bh.Canonical = true

	b1 = nil

	var s1 = struct {
		A int
		B int
		C int
	}{1, 2, 3}

	NewEncoderBytes(&b1, h).MustEncode(&s1)

	var s2 = struct {
		C int
		testMissingFieldsMap
	}{C: 3}
	s2.testMissingFieldsMap = testMissingFieldsMap{
		m: map[string]interface{}{
			"A": 1,
			"B": 2,
		},
	}

	for i := 0; i < 16; i++ {
		b2 = nil
		NewEncoderBytes(&b2, h).MustEncode(&s2)

		if !bytes.Equal(b1, b2) {
			t.Fatalf("bytes differed:'%s' vs '%s'", b1, b2)
		}
	}
}

func doTestMaxDepth(t *testing.T, h Handle) {
	if !testRecoverPanicToErr {
		t.Skip(testSkipIfNotRecoverPanicToErrMsg)
	}
	defer testSetup(t, &h)()
	name := h.Name()
	type T struct {
		I interface{} // value to encode
		M int16       // maxdepth
		S bool        // use swallow (decode into typed struct with only A1)
		E interface{} // error to find
	}
	type T1 struct {
		A1 *T1
	}
	var table []T
	var sfunc = func(n int) (s [1]interface{}, s1 *[1]interface{}) {
		s1 = &s
		for i := 0; i < n; i++ {
			var s0 [1]interface{}
			s1[0] = &s0
			s1 = &s0
		}

		return
		// var s []interface{}
		// s = append(s, []interface{})
		// s[0] = append(s[0], []interface{})
		// s[0][0] = append(s[0][0], []interface{})
		// s[0][0][0] = append(s[0][0][0], []interface{})
		// s[0][0][0][0] = append(s[0][0][0][0], []interface{})
		// return s
	}
	var mfunc = func(n int) (m map[string]interface{}, mlast map[string]interface{}) {
		m = make(map[string]interface{})
		mlast = make(map[string]interface{})
		m["A0"] = mlast
		for i := 1; i < n; i++ {
			m0 := make(map[string]interface{})
			mlast["A"+strconv.FormatInt(int64(i), 10)] = m0
			mlast = m0
		}

		return
	}
	s, s1 := sfunc(5)
	m, _ := mfunc(5)
	m99, _ := mfunc(99)

	s1[0] = m

	table = append(table, T{s, 0, false, nil})
	table = append(table, T{s, 256, false, nil})
	table = append(table, T{s, 7, false, errMaxDepthExceeded})
	table = append(table, T{s, 15, false, nil})
	table = append(table, T{m99, 15, true, errMaxDepthExceeded})
	table = append(table, T{m99, 215, true, nil})

	defer func(n int16) {
		testBasicHandle(h).MaxDepth = n
	}(testBasicHandle(h).MaxDepth)

	for i, v := range table {
		testBasicHandle(h).MaxDepth = v.M
		b1 := testMarshalErr(v.I, h, t, name+"-maxdepth-enc"+strconv.FormatInt(int64(i), 10))

		var err error
		v.S = false // MARKER: 20200925: swallow doesn't track depth anymore
		if v.S {
			var v2 T1
			err = testUnmarshal(&v2, b1, h)
		} else {
			var v2 interface{}
			err = testUnmarshal(&v2, b1, h)
		}
		var err0 interface{} = err
		if err1, ok := err.(*codecError); ok {
			err0 = err1.err
		}
		if err0 != v.E {
			t.Logf("Unexpected error testing max depth for depth %d: expected %v, received %v", v.M, v.E, err)
			t.FailNow()
		}

		// decode into something that just triggers swallow
		testReleaseBytes(b1)
	}
}

func doTestMultipleEncDec(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	name := h.Name()
	// encode a string multiple times.
	// decode it multiple times.
	// ensure we get the value each time
	var s1 = "ugorji"
	var s2 = "nwoke"
	var s11, s21 string
	var buf bytes.Buffer
	e := NewEncoder(&buf, h)
	e.MustEncode(s1)
	e.MustEncode(s2)
	d := NewDecoder(&buf, h)
	d.MustDecode(&s11)
	d.MustDecode(&s21)
	testDeepEqualErr(s1, s11, t, name+"-multiple-encode")
	testDeepEqualErr(s2, s21, t, name+"-multiple-encode")
}

func doTestSelfExt(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	name := h.Name()
	var ts TestSelfExtImpl
	ts.S = "ugorji"
	ts.I = 5678
	ts.B = true
	var ts2 TestSelfExtImpl

	bs := testMarshalErr(&ts, h, t, name)
	testUnmarshalErr(&ts2, bs, h, t, name)
	testDeepEqualErr(&ts, &ts2, t, name)
	testReleaseBytes(bs)
}

func doTestBytesEncodedAsArray(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	name := h.Name()
	// Need to test edge case where bytes are encoded as an array
	// (not using optimized []byte native format)

	// encode []int8 (or int32 or any numeric type) with all positive numbers
	// decode it into []byte
	var in = make([]int32, 128)
	var un = make([]uint8, 128)
	for i := range in {
		in[i] = int32(i)
		un[i] = uint8(i)
	}
	var out []byte
	bs := testMarshalErr(&in, h, t, name)
	testUnmarshalErr(&out, bs, h, t, name)

	testDeepEqualErr(un, out, t, name)
	testReleaseBytes(bs)
}

func doTestStrucEncDec(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	name := h.Name()

	{
		var ts1 = newTestStruc(2, testNumRepeatString, false, !testSkipIntf, testMapStringKeyOnly)
		var ts2 TestStruc
		bs := testMarshalErr(ts1, h, t, name)
		testUnmarshalErr(&ts2, bs, h, t, name)
		testDeepEqualErr(ts1, &ts2, t, name)
		testReleaseBytes(bs)
	}

	// Note: We cannot use TestStrucFlex because it has omitempty values,
	// Meaning that sometimes, encoded and decoded will not be same.
	// {
	// 	var ts1 = newTestStrucFlex(2, testNumRepeatString, false, !testSkipIntf, testMapStringKeyOnly)
	// 	var ts2 TestStrucFlex
	// 	bs := testMarshalErr(ts1, h, t, name)
	// 	fmt.Printf("\n\n%s\n\n", bs)
	// 	testUnmarshalErr(&ts2, bs, h, t, name)
	// 	testDeepEqualErr(ts1, &ts2, t, name)
	// }
}

func doTestStructKeyType(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	name := h.Name()

	mh, mok := h.(*MsgpackHandle)

	bch, bcok := h.(*BincHandle)

	bh := testBasicHandle(h)
	s2a := bh.StructToArray
	bh.StructToArray = false
	ir := bh.InterfaceReset
	bh.InterfaceReset = false
	var mfx bool
	if mok {
		mfx = mh.NoFixedNum
		mh.NoFixedNum = false
	}
	var bcsym uint8
	if bcok {
		bcsym = bch.AsSymbols
		bch.AsSymbols = 2 // MARKER: should be 0 but binc symbols do not work well
	}
	defer func() {
		bh.StructToArray = s2a
		bh.InterfaceReset = ir
		if mok {
			mh.NoFixedNum = mfx
		}
		if bcok {
			bch.AsSymbols = bcsym
		}
	}()

	var bs1, bs2 []byte

	var m = make(map[interface{}]interface{})

	fn := func(v interface{}) {
		v1 := reflect.New(reflect.TypeOf(v)).Elem().Interface()
		bs1 = testMarshalErr(v, h, t, "")
		testUnmarshalErr(&v1, bs1, h, t, "")
		testDeepEqualErr(v, v1, t, name+"")
		bs2 = testMarshalErr(m, h, t, "")
		// if _, ok := h.(*JsonHandle); ok {
		// 	xdebugf("bs1: %s, bs2: %s", bs1, bs2)
		// }
		// if bcok {
		// 	xdebug2f("-------------")
		// 	xdebugf("v: %#v, m: %#v", v, m)
		// 	xdebugf("bs1: %v", bs1)
		// 	xdebugf("bs2: %v", bs2)
		// 	xdebugf("bs1==bs2: %v", bytes.Equal(bs1, bs2))
		// 	testDeepEqualErr(bs1, bs2, t, name+"")
		// 	xdebug2f("-------------")
		// 	return
		// }
		testDeepEqualErr(bs1, bs2, t, name+"")
		testReleaseBytes(bs1)
		testReleaseBytes(bs2)
	}

	fnclr := func() {
		for k := range m {
			delete(m, k)
		}
	}

	m["F"] = 90
	fn(&testStrucKeyTypeT0{F: 90})
	fnclr()

	m["FFFF"] = 100
	fn(&testStrucKeyTypeT1{F: 100})
	fnclr()

	m[int64(-1)] = 200
	fn(&testStrucKeyTypeT2{F: 200})
	fnclr()

	m[int64(1)] = 300
	fn(&testStrucKeyTypeT3{F: 300})
	fnclr()

	m[float64(2.5)] = 400
	fn(&testStrucKeyTypeT4{F: 400})
	fnclr()
}

func doTestRawToStringToRawEtc(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	// name := h.Name()

	// Tests:
	// - RawToString
	// - StringToRaw
	// - MapValueReset
	// - DeleteOnMapValue (skipped - no longer supported)

	bh := testBasicHandle(h)

	r2s := bh.RawToString
	s2r := bh.StringToRaw
	can := bh.Canonical
	mvr := bh.MapValueReset

	mh, mok := h.(*MsgpackHandle)

	_, jok := h.(*JsonHandle)

	defer func() {
		bh.RawToString = r2s
		bh.StringToRaw = s2r
		bh.Canonical = can
		bh.MapValueReset = mvr
	}()

	bh.Canonical = false

	var bs1, bs2 []byte

	// encode: StringToRaw
	// decode: RawToString

	// compare encoded v1 to encoded v2, while setting StringToRaw to b
	fne := func(v1, v2 interface{}, b bool) {
		bh.StringToRaw = b
		bs1 = testMarshalErr(v1, h, t, "")
		// bs1 = []byte(string(bs1))
		bs2 = testMarshalErr(v2, h, t, "")
		testDeepEqualErr(bs1, bs2, t, "")
		testReleaseBytes(bs1)
		testReleaseBytes(bs2)
	}

	// encoded v1, decode naked and compare to v2
	fnd := func(v1, v2 interface{}, bs2r, br2s, bwext bool) {
		bh.RawToString = br2s
		bh.StringToRaw = bs2r
		if mok {
			mh.RawToString = bwext
		}
		bs1 = testMarshalErr(v1, h, t, "")
		var vn interface{}
		testUnmarshalErr(&vn, bs1, h, t, "")
		testDeepEqualErr(vn, v2, t, "")
		testReleaseBytes(bs1)
	}

	sv0 := "value"
	bv0 := []byte(sv0)

	sv1 := sv0
	bv1 := []byte(sv1)

	m1 := map[string]*string{"key": &sv1}
	m2 := map[string]*[]byte{"key": &bv1}

	// m3 := map[[3]byte]string{[3]byte{'k', 'e', 'y'}: sv0}
	m4 := map[[3]byte][]byte{[3]byte{'k', 'e', 'y'}: bv0}
	m5 := map[string][]byte{"key": bv0}
	// m6 := map[string]string{"key": sv0}

	m7 := map[interface{}]interface{}{"key": sv0}
	m8 := map[interface{}]interface{}{"key": bv0}

	// StringToRaw=true
	fne(m1, m4, true)

	// StringToRaw=false
	// compare encoded m2 to encoded m5
	fne(m2, m5, false)

	// json doesn't work well with StringToRaw and RawToString
	// when dealing with interfaces, because it cannot decipher
	// that a string should be treated as base64.
	if jok {
		goto MAP_VALUE_RESET
	}

	// if msgpack, always set WriteExt = RawToString

	// StringToRaw=true (RawToString=true)
	// encoded m1, decode naked and compare to m5
	fnd(m2, m7, true, true, true)

	// StringToRaw=true (RawToString=false)
	// encoded m1, decode naked and compare to m6
	fnd(m1, m8, true, false, false)

	// StringToRaw=false, RawToString=true
	// encode m1, decode naked, and compare to m6
	fnd(m2, m7, false, true, true)

MAP_VALUE_RESET:
	// set MapValueReset, and then decode i
	sv2 := "value-new"
	m9 := map[string]*string{"key": &sv2}

	bs1 = testMarshalErr(m1, h, t, "")

	bh.MapValueReset = false
	testUnmarshalErr(&m9, bs1, h, t, "")
	// if !(m9["key"] == m1["key"]
	testDeepEqualErr(sv2, "value", t, "")
	testDeepEqualErr(&sv2, m9["key"], t, "")

	sv2 = "value-new"
	m9["key"] = &sv2
	bh.MapValueReset = true
	testUnmarshalErr(&m9, bs1, h, t, "")
	testDeepEqualErr(sv2, "value-new", t, "")
	testDeepEqualErr("value", *(m9["key"]), t, "")

	// t1 = struct {
	// 	key string
	// }{ key: sv0 }
	// t2 := struct {
	// 	key []byte
	// }{ key: bv1 }
	testReleaseBytes(bs1)

}

// -----------------

func doTestJsonDecodeNonStringScalarInStringContext(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	var b = `{"s.true": "true", "b.true": true, "s.false": "false", "b.false": false, "s.10": "10", "i.10": 10, "i.-10": -10}`
	var golden = map[string]string{"s.true": "true", "b.true": "true", "s.false": "false", "b.false": "false", "s.10": "10", "i.10": "10", "i.-10": "-10"}

	var m map[string]string
	d := NewDecoderBytes([]byte(b), h)
	d.MustDecode(&m)
	if err := deepEqual(golden, m); err == nil {
		if testVerbose {
			t.Logf("++++ match: decoded: %#v", m)
		}
	} else {
		t.Logf("---- mismatch: %v ==> golden: %#v, decoded: %#v", err, golden, m)
		t.FailNow()
	}

	jh := h.(*JsonHandle)

	oldMapKeyAsString := jh.MapKeyAsString
	oldPreferFloat := jh.PreferFloat
	oldSignedInteger := jh.SignedInteger
	oldHTMLCharAsIs := jh.HTMLCharsAsIs
	oldMapType := jh.MapType

	defer func() {
		jh.MapKeyAsString = oldMapKeyAsString
		jh.PreferFloat = oldPreferFloat
		jh.SignedInteger = oldSignedInteger
		jh.HTMLCharsAsIs = oldHTMLCharAsIs
		jh.MapType = oldMapType
	}()

	jh.MapType = mapIntfIntfTyp
	jh.HTMLCharsAsIs = false
	jh.MapKeyAsString = true
	jh.PreferFloat = false
	jh.SignedInteger = false

	// Also test out decoding string values into naked interface
	b = `{"true": true, "false": false, "null": null, "7700000000000000000": 7700000000000000000}`
	const num = 7700000000000000000 // 77
	var golden2 = map[interface{}]interface{}{
		true:  true,
		false: false,
		nil:   nil,
		// uint64(num):uint64(num),
	}

	fn := func() {
		d.ResetBytes([]byte(b))
		var mf interface{}
		d.MustDecode(&mf)
		if err := deepEqual(golden2, mf); err != nil {
			t.Logf("---- mismatch: %v ==> golden: %#v, decoded: %#v", err, golden2, mf)
			t.FailNow()
		}
	}

	golden2[uint64(num)] = uint64(num)
	fn()
	delete(golden2, uint64(num))

	jh.SignedInteger = true
	golden2[int64(num)] = int64(num)
	fn()
	delete(golden2, int64(num))
	jh.SignedInteger = false

	jh.PreferFloat = true
	golden2[float64(num)] = float64(num)
	fn()
	delete(golden2, float64(num))
	jh.PreferFloat = false
}

func doTestJsonEncodeIndent(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	v := TestSimplish{
		Ii: -794,
		Ss: `A Man is
after the new line
	after new line and tab
`,
	}
	v2 := v
	v.Mm = make(map[string]*TestSimplish)
	for i := 0; i < len(v.Ar); i++ {
		v3 := v2
		v3.Ii += (i * 4)
		v3.Ss = fmt.Sprintf("%d - %s", v3.Ii, v3.Ss)
		if i%2 == 0 {
			v.Ar[i] = &v3
		}
		// v3 = v2
		v.Sl = append(v.Sl, &v3)
		v.Mm[strconv.FormatInt(int64(i), 10)] = &v3
	}
	jh := h.(*JsonHandle)

	oldcan := jh.Canonical
	oldIndent := jh.Indent
	oldS2A := jh.StructToArray
	defer func() {
		jh.Canonical = oldcan
		jh.Indent = oldIndent
		jh.StructToArray = oldS2A
	}()
	jh.Canonical = true
	jh.Indent = -1
	jh.StructToArray = false
	var bs []byte
	NewEncoderBytes(&bs, jh).MustEncode(&v)
	txt1Tab := string(bs)
	bs = nil
	jh.Indent = 120
	NewEncoderBytes(&bs, jh).MustEncode(&v)
	txtSpaces := string(bs)
	// fmt.Printf("\n-----------\n%s\n------------\n%s\n-------------\n", txt1Tab, txtSpaces)

	goldenResultTab := `{
	"Ar": [
		{
			"Ar": [
				null,
				null
			],
			"Ii": -794,
			"Mm": null,
			"Sl": null,
			"Ss": "-794 - A Man is\nafter the new line\n\tafter new line and tab\n"
		},
		null
	],
	"Ii": -794,
	"Mm": {
		"0": {
			"Ar": [
				null,
				null
			],
			"Ii": -794,
			"Mm": null,
			"Sl": null,
			"Ss": "-794 - A Man is\nafter the new line\n\tafter new line and tab\n"
		},
		"1": {
			"Ar": [
				null,
				null
			],
			"Ii": -790,
			"Mm": null,
			"Sl": null,
			"Ss": "-790 - A Man is\nafter the new line\n\tafter new line and tab\n"
		}
	},
	"Sl": [
		{
			"Ar": [
				null,
				null
			],
			"Ii": -794,
			"Mm": null,
			"Sl": null,
			"Ss": "-794 - A Man is\nafter the new line\n\tafter new line and tab\n"
		},
		{
			"Ar": [
				null,
				null
			],
			"Ii": -790,
			"Mm": null,
			"Sl": null,
			"Ss": "-790 - A Man is\nafter the new line\n\tafter new line and tab\n"
		}
	],
	"Ss": "A Man is\nafter the new line\n\tafter new line and tab\n"
}`

	if txt1Tab != goldenResultTab {
		t.Logf("decoded indented with tabs != expected: \nexpected: %s\nencoded: %s", goldenResultTab, txt1Tab)
		t.FailNow()
	}
	if txtSpaces != strings.Replace(goldenResultTab, "\t", strings.Repeat(" ", 120), -1) {
		t.Logf("decoded indented with spaces != expected: \nexpected: %s\nencoded: %s", goldenResultTab, txtSpaces)
		t.FailNow()
	}
}

func __doTestBufioDecReader(t *testing.T, bufsize int) {
	bufsizehalf := (bufsize + 1) / 2

	// try to read 85 bytes in chunks of 7 at a time.
	var s = strings.Repeat("01234'56789      ", 5)
	// fmt.Printf("s: %s\n", s)
	var r = strings.NewReader(s)
	var br bufioDecReader
	var blist bytesFreelist
	br.reset(r, bufsize, &blist)
	b, err := ioutil.ReadAll(br.r)
	if err != nil {
		panic(err)
	}
	var s2 = string(b)
	// fmt.Printf("s==s2: %v, len(s): %v, len(b): %v, len(s2): %v\n", s == s2, len(s), len(b), len(s2))
	if s != s2 {
		t.Logf("not equal: \ns:  %s\ns2: %s", s, s2)
		t.FailNow()
	}
	// Now, test search functions for skip, readTo and readUntil
	// readUntil ', readTo ', skip whitespace. 3 times in a loop, each time compare the token and/or outs
	// readUntil: see: 56789
	var out []byte
	var token byte
	br.reset(strings.NewReader(s), bufsizehalf, &blist)
	// println()
	for _, v2 := range [...]string{
		`01234'`,
		`56789      01234'`,
		`56789      01234'`,
		`56789      01234'`,
	} {
		out = br.readUntil('\'')
		testDeepEqualErr(string(out), v2[:len(v2)-1], t, "-")
		// fmt.Printf("readUntil: out: `%s`\n", out)
	}
	br.reset(strings.NewReader(s), bufsizehalf, &blist)
	// println()

	// 20200915: readTo is not longer available
	// for range [4]struct{}{} {
	// 	out = br.readTo(&numCharBitset)
	// 	testDeepEqualErr(string(out), `01234`, t, "-")
	// 	// fmt.Printf("readTo: out: `%s`\n", out)
	// 	out = br.readUntil('\'', true)
	// 	testDeepEqualErr(string(out), "'", t, "-")
	// 	// fmt.Printf("readUntil: out: `%s`\n", out)
	// 	out = br.readTo(&numCharBitset)
	// 	testDeepEqualErr(string(out), `56789`, t, "-")
	// 	// fmt.Printf("readTo: out: `%s`\n", out)
	// 	out = br.readUntil('0', true)
	// 	testDeepEqualErr(string(out), `      0`, t, "-")
	// 	// fmt.Printf("readUntil: out: `%s`\n", out)
	// 	br.unreadn1()
	// }
	br.reset(strings.NewReader(s), bufsizehalf, &blist)
	for range [4]struct{}{} {
		token = br.skipWhitespace() // br.skip(&whitespaceCharBitset)
		testDeepEqualErr(token, byte('0'), t, "-")
		out = br.readUntil(' ')
		testDeepEqualErr(string(out), `1234'56789`, t, "-")
	}
}

func doTestPreferArrayOverSlice(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	// encode a slice, decode it with PreferArrayOverSlice
	// if codecgen, skip the test (as codecgen doesn't work with PreferArrayOverSlice)
	if codecgen {
		t.Skip("skipping ... prefer array over slice is not supported in codecgen mode")
	}
	bh := testBasicHandle(h)
	paos := bh.PreferArrayOverSlice
	styp := bh.SliceType
	defer func() {
		bh.PreferArrayOverSlice = paos
		bh.SliceType = styp
	}()
	bh.PreferArrayOverSlice = true
	bh.SliceType = reflect.TypeOf(([]bool)(nil))

	s2 := [4]bool{true, false, true, false}
	s := s2[:]
	var v interface{}
	bs := testMarshalErr(s, h, t, t.Name())
	testUnmarshalErr(&v, bs, h, t, t.Name())
	testDeepEqualErr(s2, v, t, t.Name())
	testReleaseBytes(bs)
}

func doTestZeroCopyBytes(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	// jsonhandle and cborhandle with indefiniteLength do not support inline bytes, so skip them.
	if _, ok := h.(*JsonHandle); ok {
		t.Skipf("skipping ... zero copy bytes not supported by json handle")
	}
	if ch, ok := h.(*CborHandle); ok && ch.IndefiniteLength {
		t.Skipf("skipping ... zero copy bytes not supported by cbor handle with IndefiniteLength=true")
	}

	bh := testBasicHandle(h)
	zc := bh.ZeroCopy
	defer func() {
		bh.ZeroCopy = zc
	}()
	bh.ZeroCopy = true

	s := []byte("hello")
	var v []byte
	bs := testMarshalErr(s, h, t, t.Name())

	// Note: this test only works for decoding from []byte, so cannot use testUnmarshalErr
	NewDecoderBytes(bs, h).MustDecode(&v)
	// testUnmarshalErr(&v, bs, h, t, t.Name())

	// validate that bs and s points into the bs stream
	for i := range bs {
		if &bs[i] == &v[0] {
			return
		}
	}

	// if not match, then a failure happened.
	if len(bs) > 0 && len(v) > 0 {
		t.Logf("%s: ZeroCopy=true, but decoded (%p) is not slice of input: (%p)", h.Name(), &v[0], &bs[0])
	} else {
		t.Logf("%s: ZeroCopy=true, but decoded OR input slice is empty: %v, %v", h.Name(), v, bs)
	}
	testReleaseBytes(bs)
	t.FailNow()
}

func doTestNextValueBytes(t *testing.T, h Handle) {
	defer testSetup(t, &h)()

	bh := testBasicHandle(h)

	// - encode uint, int, float, bool, struct, map, slice, string - all separated by nil
	// - use nextvaluebytes to grab he's got each one, and decode it, and compare
	var inputs = []interface{}{
		uint64(7777),
		int64(9999),
		float64(12.25),
		true,
		false,
		map[string]uint64{"1": 1, "22": 22, "333": 333, "4444": 4444},
		[]string{"1", "22", "333", "4444"},
		// use *TestStruc, not *TestStrucFlex, as *TestStrucFlex is harder to compare with deep equal
		// Remember: *TestStruc was separated for this reason, affording comparing against other libraries
		newTestStruc(testDepth, testNumRepeatString, false, false, true),
		"1223334444",
	}
	var out []byte

	for i, v := range inputs {
		_ = i
		bs := testMarshalErr(v, h, t, "nextvaluebytes")
		out = append(out, bs...)
		bs2 := testMarshalErr(nil, h, t, "nextvaluebytes")
		out = append(out, bs2...)
		testReleaseBytes(bs)
		testReleaseBytes(bs2)
	}
	// out = append(out, []byte("----")...)

	var valueBytes = make([][]byte, len(inputs)*2)

	d, oldReadBufferSize := testSharedCodecDecoder(out, h, testBasicHandle(h))
	for i := 0; i < len(inputs)*2; i++ {
		valueBytes[i] = d.d.nextValueBytes([]byte{})
		// bs := d.d.nextValueBytes([]byte{})
		// valueBytes[i] = make([]byte, len(bs))
		// copy(valueBytes[i], bs)
	}
	if testUseIoEncDec >= 0 {
		bh.ReaderBufferSize = oldReadBufferSize
	}

	defer func(b bool) { bh.InterfaceReset = b }(bh.InterfaceReset)
	bh.InterfaceReset = false

	var result interface{}
	for i := 0; i < len(inputs); i++ {
		// result = reflect.New(reflect.TypeOf(inputs[i])).Elem().Interface()
		result = reflect.Zero(reflect.TypeOf(inputs[i])).Interface()
		testUnmarshalErr(&result, valueBytes[i*2], h, t, "nextvaluebytes")
		testDeepEqualErr(inputs[i], result, t, "nextvaluebytes-1")
		result = nil
		testUnmarshalErr(&result, valueBytes[(i*2)+1], h, t, "nextvaluebytes")
		testDeepEqualErr(nil, result, t, "nextvaluebytes-2")
	}
}

func doTestNumbers(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	__doTestIntegers(t, h)
	__doTestFloats(t, h)
	__doTestIntegerFloatConversions(t, h)
}

func __doTestIntegers(t *testing.T, h Handle) {
	// handle SignedInteger=true|false
	// decode into an interface{}

	bh := testBasicHandle(h)

	oldSignedInteger := bh.SignedInteger
	var oldPreferFloat bool
	var oldNoFixedNum bool
	jh, jok := h.(*JsonHandle)
	mh, mok := h.(*MsgpackHandle)
	if jok {
		oldPreferFloat = jh.PreferFloat
	}
	if mok {
		oldNoFixedNum = mh.NoFixedNum
		mh.NoFixedNum = true
	}

	defer func() {
		bh.SignedInteger = oldSignedInteger
		if jok {
			jh.PreferFloat = oldPreferFloat
		}
		if mok {
			mh.NoFixedNum = oldNoFixedNum
		}
	}()

	// var vi int64
	// var ui uint64
	var ii interface{}

	for _, v := range testUintsToParse {
		if jok {
			jh.PreferFloat = false
		}
		b := testMarshalErr(v, h, t, "test-integers")
		ii = nil
		bh.SignedInteger = true
		testUnmarshalErr(&ii, b, h, t, "test-integers")
		testDeepEqualErr(ii, int64(v), t, "test-integers-signed")
		ii = nil
		bh.SignedInteger = false
		testUnmarshalErr(&ii, b, h, t, "test-integers")
		testDeepEqualErr(ii, uint64(v), t, "test-integers-unsigned")
		ii = nil
		if jok {
			jh.PreferFloat = true
			testUnmarshalErr(&ii, b, h, t, "test-integers")
			testDeepEqualErr(ii, float64(v), t, "test-integers-float")
		}
		testReleaseBytes(b)
	}
}

var testUintsToParse = []uint64{
	// Note: use large integers, as some formats store small integers in an agnostic
	// way, where its not clear if signed or unsigned.
	2,
	2048,
	1<<63 - 4,
	1800000000e-2,
	18000000e+2,
	4.56e+4, // tests float32 exact parsing
	4.56e+16,
}

var testFloatsToParse = []float64{
	3,
	math.NaN(),
	math.Inf(1),
	math.Inf(-1),
	4.56e+4,  // tests float32 exact parsing
	4.56e+18, // float32 parsing - exp > 10
	4.56e+10,
	4.56e+30,
	1.01234567890123456789e+30,
	1.32e+5,
	1.32e-5,
	0e+01234567890123456789,
	1.7976931348623157e308,
	-1.7976931348623157e+308,
	1e308,
	1e-308,
	1.694649e-317,
	// 1e-4294967296,
	2.2250738585072012e-308,
	4.630813248087435e+307,
	1.00000000000000011102230246251565404236316680908203125,
	1.00000000000000033306690738754696212708950042724609375,
}

func __doTestFloats(t *testing.T, h Handle) {
	_, jok := h.(*JsonHandle)

	f64s := testFloatsToParse
	const unusedVal = 9999 // use this as a marker

	// marshall it, unmarshal it, compare to original
	// Note: JSON encodes NaN, inf, -inf as null, which is decoded as zero value (ie 0).
	for _, f64 := range f64s {
		{
			f := f64
			var w float64 = unusedVal
			b := testMarshalErr(f, h, t, "test-floats-enc")
			testUnmarshalErr(&w, b, h, t, "test-floats-dec")
			// we only check json for float64, as it doesn't differentiate
			if (jok && (math.IsNaN(f64) || math.IsInf(f64, 0)) && w != 0) ||
				(!jok && w != f && !math.IsNaN(float64(f))) {
				t.Logf("error testing float64: %v, decoded as: %v", f, w)
				t.FailNow()
			}
			var wi interface{}
			testUnmarshalErr(&wi, b, h, t, "test-floats-dec")
			if (jok && (math.IsNaN(f64) || math.IsInf(f64, 0)) && wi != nil) ||
				(!jok && wi.(float64) != f && !math.IsNaN(float64(f))) {
				t.Logf("error testing float64: %v, decoded as: %v", f, wi)
				t.FailNow()
			}
			testReleaseBytes(b)
		}
		{
			f := float32(f64)
			var w float32 = unusedVal
			b := testMarshalErr(f, h, t, "test-floats-enc")
			// xdebug2f("test float: of %v, encoded as: %s", f, b)
			testUnmarshalErr(&w, b, h, t, "test-floats-dec")
			if (jok && (math.IsNaN(f64) || math.IsInf(f64, 0)) && w != 0) ||
				(!jok && w != f && !math.IsNaN(float64(f))) {
				t.Logf("error testing float32: %v, decoded as: %v", f, w)
				t.FailNow()
			}
			testReleaseBytes(b)
		}
	}
}

func __doTestIntegerFloatConversions(t *testing.T, h Handle) {
	if !testRecoverPanicToErr {
		t.Skip(testSkipIfNotRecoverPanicToErrMsg)
	}
	type tI struct{ N int }
	type tU struct{ N uint }
	type tF struct{ N float64 }

	type elem struct {
		in  interface{}
		out interface{}
		err string
	}
	tests := []elem{
		// good
		{tI{5}, tF{5.0}, ""},
		{tU{5}, tF{5.0}, ""},
		{tF{5.0}, tU{5}, ""},
		{tF{5.0}, tI{5}, ""},
		{tF{-5.0}, tI{-5}, ""},
		// test negative number into unsigned integer
		{tI{-5}, tU{5}, "error"},
		{tF{-5.0}, tU{5}, "error"},
		// test fractional float into integer
		{tF{-5.7}, tU{5}, "error"},
		{tF{-5.7}, tI{5}, "error"},
	}
	for _, v := range tests {
		r := reflect.New(reflect.TypeOf(v.out))
		b := testMarshalErr(v.in, h, t, "")
		err := testUnmarshal(r.Interface(), b, h)
		if v.err == "" {
			testDeepEqualErr(err, nil, t, "")
			testDeepEqualErr(r.Elem().Interface(), v.out, t, "")
		} else if err == nil {
			t.Logf("expecting an error but didn't receive any, with in: %v, out: %v, expecting err matching: %v", v.in, v.out, v.err)
			t.FailNow()
		}
	}
}

func doTestStructFieldInfoToArray(t *testing.T, h Handle) {
	if !testRecoverPanicToErr {
		t.Skip(testSkipIfNotRecoverPanicToErrMsg)
	}
	defer testSetup(t, &h)()
	bh := testBasicHandle(h)

	defer func(b bool) { bh.CheckCircularRef = b }(bh.CheckCircularRef)
	bh.CheckCircularRef = true

	var vs = Sstructsmall{A: 99}
	var vb = Sstructbig{
		A:         77,
		B:         true,
		c:         "ccc 3 ccc",
		Ssmallptr: &vs,
		Ssmall:    vs,
	}
	vb.Sptr = &vb

	vba := SstructbigToArray{
		A:         vb.A,
		B:         vb.B,
		c:         vb.c,
		Ssmallptr: vb.Ssmallptr,
		Ssmall:    vb.Ssmall,
		Sptr:      vb.Sptr,
	}

	var b []byte
	var err error
	if !codecgen {
		// codecgen doesn't support CheckCircularRef, and these types are codecgen'ed
		b, err = testMarshal(&vba, h)
		testReleaseBytes(b)
		if err == nil || !strings.Contains(err.Error(), "circular reference found") {
			t.Logf("expect circular reference error, got: %v", err)
			t.FailNow()
		}
	}

	vb2 := vb
	vb.Sptr = nil // so we stop having the circular reference error
	vba.Sptr = &vb2

	var ss []interface{}

	// bh.CheckCircularRef = false
	b = testMarshalErr(&vba, h, t, "-")
	testUnmarshalErr(&ss, b, h, t, "-")
	testDeepEqualErr(ss[1], true, t, "-")
	testReleaseBytes(b)
}

func doTestDesc(t *testing.T, h Handle, m map[byte]string) {
	defer testSetup(t, &h)()
	for k, v := range m {
		if s := h.desc(k); s != v {
			t.Logf("error describing descriptor: '%q' i.e. 0x%x, expected '%s', got '%s'", k, k, v, s)
			t.FailNow()
		}
	}
}

func TestBufioDecReader(t *testing.T) {
	doTestBufioDecReader(t)
}

func doTestBufioDecReader(t *testing.T) {
	defer testSetup(t, nil)()
	__doTestBufioDecReader(t, 13)
	__doTestBufioDecReader(t, 3)
	__doTestBufioDecReader(t, 5)
	__doTestBufioDecReader(t, 127)
}

func TestAtomic(t *testing.T) {
	defer testSetup(t, nil)()
	// load, store, load, confirm
	if true {
		var a atomicTypeInfoSlice
		l := a.load()
		if l != nil {
			t.Logf("atomic fail: %T, expected load return nil, received: %v", a, l)
			t.FailNow()
		}
		l = append(l, rtid2ti{})
		a.store(l)
		l = a.load()
		if len(l) != 1 {
			t.Logf("atomic fail: %T, expected load to have length 1, received: %d", a, len(l))
			t.FailNow()
		}
	}
	if true {
		var a atomicRtidFnSlice
		l := a.load()
		if l != nil {
			t.Logf("atomic fail: %T, expected load return nil, received: %v", a, l)
			t.FailNow()
		}
		l = append(l, codecRtidFn{})
		a.store(l)
		l = a.load()
		if len(l) != 1 {
			t.Logf("atomic fail: %T, expected load to have length 1, received: %d", a, len(l))
			t.FailNow()
		}
	}
	if true {
		var a atomicClsErr
		l := a.load()
		if l.err != nil {
			t.Logf("atomic fail: %T, expected load return clsErr = nil, received: %v", a, l.err)
			t.FailNow()
		}
		l.err = io.EOF
		a.store(l)
		l = a.load()
		if l.err != io.EOF {
			t.Logf("atomic fail: %T, expected clsErr = io.EOF, received: %v", a, l.err)
			t.FailNow()
		}
	}
}

// -----------

func doTestJsonLargeInteger(t *testing.T, h Handle) {
	if !testRecoverPanicToErr {
		t.Skip(testSkipIfNotRecoverPanicToErrMsg)
	}
	defer testSetup(t, &h)()
	jh := h.(*JsonHandle)
	for _, i := range []uint8{'L', 'A', 0} {
		for _, j := range []interface{}{
			int64(1 << 60),
			-int64(1 << 60),
			0,
			1 << 20,
			-(1 << 20),
			uint64(1 << 60),
			uint(0),
			uint(1 << 20),
			int64(1840000e-2),
			uint64(1840000e+2),
		} {
			__doTestJsonLargeInteger(t, j, i, jh)
		}
	}

	oldIAS := jh.IntegerAsString
	defer func() { jh.IntegerAsString = oldIAS }()
	jh.IntegerAsString = 0

	type tt struct {
		s           string
		canI, canUi bool
		i           int64
		ui          uint64
	}

	var i int64
	var ui uint64
	var err error
	var d *Decoder = NewDecoderBytes(nil, jh)
	for _, v := range []tt{
		{"0", true, true, 0, 0},
		{"0000", true, true, 0, 0},
		{"0.00e+2", true, true, 0, 0},
		{"000e-2", true, true, 0, 0},
		{"0.00e-2", true, true, 0, 0},

		{"9223372036854775807", true, true, math.MaxInt64, math.MaxInt64},                             // maxint64
		{"92233720368547758.07e+2", true, true, math.MaxInt64, math.MaxInt64},                         // maxint64
		{"922337203685477580700e-2", true, true, math.MaxInt64, math.MaxInt64},                        // maxint64
		{"9223372.036854775807E+12", true, true, math.MaxInt64, math.MaxInt64},                        // maxint64
		{"9223372036854775807000000000000E-12", true, true, math.MaxInt64, math.MaxInt64},             // maxint64
		{"0.9223372036854775807E+19", true, true, math.MaxInt64, math.MaxInt64},                       // maxint64
		{"92233720368547758070000000000000000000E-19", true, true, math.MaxInt64, math.MaxInt64},      // maxint64
		{"0.000009223372036854775807E+24", true, true, math.MaxInt64, math.MaxInt64},                  // maxint64
		{"9223372036854775807000000000000000000000000E-24", true, true, math.MaxInt64, math.MaxInt64}, // maxint64

		{"-9223372036854775808", true, false, math.MinInt64, 0},                             // minint64
		{"-92233720368547758.08e+2", true, false, math.MinInt64, 0},                         // minint64
		{"-922337203685477580800E-2", true, false, math.MinInt64, 0},                        // minint64
		{"-9223372.036854775808e+12", true, false, math.MinInt64, 0},                        // minint64
		{"-9223372036854775808000000000000E-12", true, false, math.MinInt64, 0},             // minint64
		{"-0.9223372036854775808e+19", true, false, math.MinInt64, 0},                       // minint64
		{"-92233720368547758080000000000000000000E-19", true, false, math.MinInt64, 0},      // minint64
		{"-0.000009223372036854775808e+24", true, false, math.MinInt64, 0},                  // minint64
		{"-9223372036854775808000000000000000000000000E-24", true, false, math.MinInt64, 0}, // minint64

		{"18446744073709551615", false, true, 0, math.MaxUint64},                             // maxuint64
		{"18446744.073709551615E+12", false, true, 0, math.MaxUint64},                        // maxuint64
		{"18446744073709551615000000000000E-12", false, true, 0, math.MaxUint64},             // maxuint64
		{"0.000018446744073709551615E+24", false, true, 0, math.MaxUint64},                   // maxuint64
		{"18446744073709551615000000000000000000000000E-24", false, true, 0, math.MaxUint64}, // maxuint64

		// Add test for limit of uint64 where last digit is 0
		{"18446744073709551610", false, true, 0, math.MaxUint64 - 5},                             // maxuint64
		{"18446744.073709551610E+12", false, true, 0, math.MaxUint64 - 5},                        // maxuint64
		{"18446744073709551610000000000000E-12", false, true, 0, math.MaxUint64 - 5},             // maxuint64
		{"0.000018446744073709551610E+24", false, true, 0, math.MaxUint64 - 5},                   // maxuint64
		{"18446744073709551610000000000000000000000000E-24", false, true, 0, math.MaxUint64 - 5}, // maxuint64
		// {"", true, true},
	} {
		if v.s == "" {
			continue
		}
		d.ResetBytes([]byte(v.s))
		err = d.Decode(&ui)
		if (v.canUi && err != nil) || (!v.canUi && err == nil) || (v.canUi && err == nil && v.ui != ui) {
			t.Logf("Failing to decode %s (as unsigned): %v", v.s, err)
			t.FailNow()
		}

		d.ResetBytes([]byte(v.s))
		err = d.Decode(&i)
		if (v.canI && err != nil) || (!v.canI && err == nil) || (v.canI && err == nil && v.i != i) {
			t.Logf("Failing to decode %s (as signed): %v", v.s, err)
			t.FailNow()
		}
	}
}

func doTestJsonInvalidUnicode(t *testing.T, h Handle) {
	if !testRecoverPanicToErr {
		t.Skip(testSkipIfNotRecoverPanicToErrMsg)
	}
	defer testSetup(t, &h)()
	// t.Skipf("new json implementation does not handle bad unicode robustly")
	jh := h.(*JsonHandle)

	var err error

	// ---- test unmarshal ---
	var m = map[string]string{
		`"\udc49\u0430abc"`: "\uFFFDabc",
		`"\udc49\u0430"`:    "\uFFFD",
		`"\udc49abc"`:       "\uFFFDabc",
		`"\udc49"`:          "\uFFFD",
		`"\udZ49\u0430abc"`: "\uFFFD\u0430abc",
		`"\udcG9\u0430"`:    "\uFFFD\u0430",
		`"\uHc49abc"`:       "\uFFFDabc",
		`"\uKc49"`:          "\uFFFD",
	}

	for k, v := range m {
		// call testUnmarshal directly, so we can check for EOF
		// testUnmarshalErr(&s, []byte(k), jh, t, "-")
		var s string
		err = testUnmarshal(&s, []byte(k), jh)
		if err != nil {
			if err == io.EOF {
				continue
			}
			t.Logf("%s: unmarshal failed: %v", "-", err)
			t.FailNow()
		}

		if s != v {
			t.Logf("unmarshal: not equal: %q, %q", v, s)
			t.FailNow()
		}
	}

	// test some valid edge cases
	m = map[string]string{
		`"az\uD834\udD1E"`: "az𝄞",
		`"n\ud834\uDD1en"`: "n\U0001D11En", // "\uf09d849e", // "\UD834DD1E" // U+1DD1E g clef

		`"a\\\"\/\"\b\f\n\r\"\tz"`: "a\\\"/\"\b\f\n\r\"\tz",
	}
	for k, v := range m {
		var s string
		err = testUnmarshal(&s, []byte(k), jh)
		if err != nil {
			t.Logf("%s: unmarshal failed: %v", "-", err)
			t.FailNow()
		}

		if s != v {
			t.Logf("unmarshal: not equal: %q, %q", v, s)
			t.FailNow()
		}
	}

	// ---- test marshal ---
	var b = []byte{'"', 0xef, 0xbf, 0xbd} // this is " and unicode.ReplacementChar (as bytes)
	var m2 = map[string][]byte{
		string([]byte{0xef, 0xbf, 0xbd}):           append(b, '"'),
		string([]byte{0xef, 0xbf, 0xbd, 0x0, 0x0}): append(b, `\u0000\u0000"`...),

		"a\\\"/\"\b\f\n\r\"\tz": []byte(`"a\\\"/\"\b\f\n\r\"\tz"`),

		// our encoder doesn't support encoding using only ascii ... so need to use the utf-8 version
		// "n\U0001D11En": []byte(`"n\uD834\uDD1En"`),
		// "az𝄞": []byte(`"az\uD834\uDD1E"`),
		"n\U0001D11En": []byte(`"n𝄞n"`),
		"az𝄞":          []byte(`"az𝄞"`),

		string([]byte{129, 129}): []byte(`"\uFFFD\uFFFD"`),
	}

	for k, v := range m2 {
		b, err = testMarshal(k, jh)
		if err != nil {
			t.Logf("%s: marshal failed: %v", "-", err)
			t.FailNow()
		}

		if !bytes.Equal(b, v) {
			t.Logf("marshal: not equal: %q, %q", v, b)
			t.FailNow()
		}
		testReleaseBytes(b)
	}
}

func doTestJsonNumberParsing(t *testing.T, h Handle) {
	defer testSetup(t, &h)()
	parseInt64_reader := func(r readFloatResult) (v int64, fail bool) {
		u, fail := parseUint64_reader(r)
		if fail {
			return
		}
		if r.neg {
			v = -int64(u)
		} else {
			v = int64(u)
		}
		return
	}
	for _, f64 := range testFloatsToParse {
		// using large prec might make a non-exact float ..., while small prec might make lose some precision.
		// However, we can do a check, and only check if the (u)int64 and float64 can be converted to one another
		// without losing any precision. Then we can use any precision for the tests.
		var precs = [...]int{32, -1}
		var r readFloatResult
		var fail bool
		var fs uint64
		var fsi int64
		var fint, ffrac float64
		_ = ffrac
		var bv []byte
		for _, prec := range precs {
			f := f64
			if math.IsNaN(f64) || math.IsInf(f64, 0) {
				goto F32
			}
			bv = strconv.AppendFloat(nil, f, 'E', prec, 64)
			if fs, err := parseFloat64_custom(bv); err != nil || fs != f {
				t.Logf("float64 -> float64 error (prec: %v) parsing '%s', got %v, expected %v: %v", prec, bv, fs, f, err)
				t.FailNow()
			}
			// try decoding it a uint64 or int64
			fint, ffrac = math.Modf(f)
			// xdebug2f(">> fint: %v (ffrac: %v) as uint64: %v as float64: %v", fint, ffrac, uint64(fint), float64(uint64(fint)))
			bv = strconv.AppendFloat(bv[:0], fint, 'E', prec, 64)
			if f < 0 || fint != float64(uint64(fint)) {
				goto F64i
			}
			r = readFloat(bv, fi64u)
			fail = !r.ok
			if r.ok {
				fs, fail = parseUint64_reader(r)
			}
			if fail || fs != uint64(fint) {
				t.Logf("float64 -> uint64 error (prec: %v, fail: %v) parsing '%s', got %v, expected %v", prec, fail, bv, fs, uint64(fint))
				t.FailNow()
			}
		F64i:
			if fint != float64(int64(fint)) {
				goto F32
			}
			r = readFloat(bv, fi64u)
			fail = !r.ok
			if r.ok {
				fsi, fail = parseInt64_reader(r)
			}
			if fail || fsi != int64(fint) {
				t.Logf("float64 -> int64 error (prec: %v, fail: %v) parsing '%s', got %v, expected %v", prec, fail, bv, fsi, int64(fint))
				t.FailNow()
			}
		F32:
			f32 := float32(f64)
			f64n := float64(f32)
			if !(math.IsNaN(f64n) || math.IsInf(f64n, 0)) {
				bv := strconv.AppendFloat(nil, f64n, 'E', prec, 32)
				if fs, err := parseFloat32_custom(bv); err != nil || fs != f32 {
					t.Logf("float32 -> float32 error (prec: %v) parsing '%s', got %v, expected %v: %v", prec, bv, fs, f32, err)
					t.FailNow()
				}
			}
		}
	}

	for _, v := range testUintsToParse {
		const prec int = 32    // test with precisions of 32 so formatting doesn't truncate what doesn't fit into float
		v = uint64(float64(v)) // use a value that when converted back and forth, remains the same
		bv := strconv.AppendFloat(nil, float64(v), 'E', prec, 64)
		r := readFloat(bv, fi64u)
		if r.ok {
			if fs, fail := parseUint64_reader(r); fail || fs != v {
				t.Logf("uint64 -> uint64 error (prec: %v, fail: %v) parsing '%s', got %v, expected %v", prec, fail, bv, fs, v)
				t.FailNow()
			}
		} else {
			t.Logf("uint64 -> uint64 not ok (prec: %v) parsing '%s', expected %v", prec, bv, v)
			t.FailNow()
		}
		vi := int64(v)
		bv = strconv.AppendFloat(nil, float64(vi), 'E', prec, 64)
		r = readFloat(bv, fi64u)
		if r.ok {
			if fs, fail := parseInt64_reader(r); fail || fs != vi {
				t.Logf("int64 -> int64 error (prec: %v, fail: %v) parsing '%s', got %v, expected %v", prec, fail, bv, fs, vi)
				t.FailNow()
			}
		} else {
			t.Logf("int64 -> int64 not ok (prec: %v) parsing '%s', expected %v", prec, bv, vi)
			t.FailNow()
		}
	}
	{
		var f64 float64 = 64.0
		var f32 float32 = 32.0
		var i int = 11
		var u64 uint64 = 128

		for _, s := range []string{`""`, `null`} {
			b := []byte(s)
			testUnmarshalErr(&f64, b, h, t, "")
			testDeepEqualErr(f64, float64(0), t, "")
			testUnmarshalErr(&f32, b, h, t, "")
			testDeepEqualErr(f32, float32(0), t, "")
			testUnmarshalErr(&i, b, h, t, "")
			testDeepEqualErr(i, int(0), t, "")
			testUnmarshalErr(&u64, b, h, t, "")
			testDeepEqualErr(u64, uint64(0), t, "")
			testUnmarshalErr(&u64, b, h, t, "")
			testDeepEqualErr(u64, uint64(0), t, "")
		}
	}
}

func doTestMsgpackDecodeMapAndExtSizeMismatch(t *testing.T, h Handle) {
	if !testRecoverPanicToErr {
		t.Skip(testSkipIfNotRecoverPanicToErrMsg)
	}
	defer testSetup(t, &h)()
	fn := func(t *testing.T, b []byte, v interface{}) {
		if err := NewDecoderBytes(b, h).Decode(v); err != io.EOF && err != io.ErrUnexpectedEOF {
			t.Fatalf("expected EOF or ErrUnexpectedEOF, got %v", err)
		}
	}

	// a map claiming to have 0x10eeeeee KV pairs, but only has 1.
	var b = []byte{0xdf, 0x10, 0xee, 0xee, 0xee, 0x1, 0xa1, 0x1}
	var m1 map[int]string
	var m2 map[int][]byte
	fn(t, b, &m1)
	fn(t, b, &m2)

	// an extension claiming to have 0x7fffffff bytes, but only has 1.
	b = []byte{0xc9, 0x7f, 0xff, 0xff, 0xff, 0xda, 0x1}
	var a interface{}
	fn(t, b, &a)

	// b = []byte{0x00}
	// var s testSelferRecur
	// fn(t, b, &s)
}

func TestMapRangeIndex(t *testing.T) {
	defer testSetup(t, nil)()
	// t.Skip()
	type T struct {
		I int
		S string
		B bool
		M map[int]T
	}

	t1 := T{I: 1, B: true, S: "11", M: map[int]T{11: T{I: 11}}}
	t2 := T{I: 1, B: true, S: "12", M: map[int]T{12: T{I: 12}}}

	// ------

	var m1 = map[string]*T{
		"11": &t1,
		"12": &t2,
	}
	var m1c = make(map[string]T)
	for k, v := range m1 {
		m1c[k] = *v
	}

	fnrv := func(r1, r2 reflect.Value) reflect.Value {
		if r1.IsValid() {
			return r1
		}
		return r2
	}

	// var vx reflect.Value

	mt := reflect.TypeOf(m1)
	rvk := mapAddrLoopvarRV(mt.Key(), mt.Key().Kind())
	rvv := mapAddrLoopvarRV(mt.Elem(), mt.Elem().Kind())
	var it mapIter
	mapRange(&it, reflect.ValueOf(m1), rvk, rvv, true)
	for it.Next() {
		k := fnrv(it.Key(), rvk).Interface().(string)
		v := fnrv(it.Value(), rvv).Interface().(*T)
		testDeepEqualErr(m1[k], v, t, "map-key-eq-it-key")
		if _, ok := m1c[k]; ok {
			delete(m1c, k)
		} else {
			t.Logf("unexpected key in map: %v", k)
			t.FailNow()
		}
	}
	it.Done()
	testDeepEqualErr(len(m1c), 0, t, "all-keys-not-consumed")

	// ------

	var m2 = map[*T]T{
		&t1: t1,
		&t2: t2,
	}
	var m2c = make(map[*T]*T)
	for k := range m2 {
		m2c[k] = k
	}

	mt = reflect.TypeOf(m2)
	rvk = mapAddrLoopvarRV(mt.Key(), mt.Key().Kind())
	rvv = mapAddrLoopvarRV(mt.Elem(), mt.Elem().Kind())
	it = mapIter{} // zero it out first, before calling mapRange
	mapRange(&it, reflect.ValueOf(m2), rvk, rvv, true)
	for it.Next() {
		k := fnrv(it.Key(), rvk).Interface().(*T)
		v := fnrv(it.Value(), rvv).Interface().(T)
		testDeepEqualErr(m2[k], v, t, "map-key-eq-it-key")
		if _, ok := m2c[k]; ok {
			delete(m2c, k)
		} else {
			t.Logf("unexpected key in map: %v", k)
			t.FailNow()
		}
	}
	it.Done()
	testDeepEqualErr(len(m2c), 0, t, "all-keys-not-consumed")

	// ---- test mapGet

	fnTestMapIndex := func(mi ...interface{}) {
		for _, m0 := range mi {
			m := reflect.ValueOf(m0)
			mkt := m.Type().Key()
			mvt := m.Type().Elem()
			kfast := mapKeyFastKindFor(mkt.Kind())
			rvv := mapAddrLoopvarRV(mvt, mvt.Kind())
			visindirect := mapStoresElemIndirect(mvt.Size())
			visref := refBitset.isset(byte(mvt.Kind()))

			for _, k := range m.MapKeys() {
				mg := mapGet(m, k, rvv, kfast, visindirect, visref).Interface()
				testDeepEqualErr(m.MapIndex(k).Interface(), mg, t, "map-index-eq")
			}
		}
	}

	fnTestMapIndex(m1, m1c, m2, m2c)

	// var s string = "hello"
	// var tt = &T{I: 3}
	// ttTyp := reflect.TypeOf(tt)
	// _, _ = tt, ttTyp
	// mv := reflect.ValueOf(m)
	// it := mapRange(mv, reflect.ValueOf(&s).Elem(), reflect.ValueOf(&tt).Elem(), true) //ok
	// it := mapRange(mv, reflect.New(reflect.TypeOf(s)).Elem(), reflect.New(reflect.TypeOf(T{})).Elem(), true) // ok
	// it := mapRange(mv, reflect.New(reflect.TypeOf(s)).Elem(), reflect.New(ttTyp.Elem()), true) // !ok
	// it := mapRange(mv, reflect.New(reflect.TypeOf(s)).Elem(), reflect.New(reflect.TypeOf(T{})), true) !ok
	// it := mapRange(mv, reflect.New(reflect.TypeOf(s)).Elem(), reflect.New(reflect.TypeOf(T{})).Elem(), true) // ok

	// fmt.Printf("key: %#v\n", it.Key())
	// fmt.Printf("exp: %#v\n", mv.MapIndex(it.Key()))
	// fmt.Printf("val: %#v\n", it.Value())
	// testDeepEqualErr(mv.MapIndex(it.Key()), it.Value().Interface()
}

// ----------

func TestBincCodecsTable(t *testing.T) {
	doTestCodecTableOne(t, testBincH)
}

func TestBincCodecsMisc(t *testing.T) {
	doTestCodecMiscOne(t, testBincH)
}

func TestBincCodecsEmbeddedPointer(t *testing.T) {
	doTestCodecEmbeddedPointer(t, testBincH)
}

func TestBincStdEncIntf(t *testing.T) {
	doTestStdEncIntf(t, testBincH)
}

func TestSimpleCodecsTable(t *testing.T) {
	doTestCodecTableOne(t, testSimpleH)
}

func TestSimpleCodecsMisc(t *testing.T) {
	doTestCodecMiscOne(t, testSimpleH)
}

func TestSimpleCodecsEmbeddedPointer(t *testing.T) {
	doTestCodecEmbeddedPointer(t, testSimpleH)
}

func TestSimpleStdEncIntf(t *testing.T) {
	doTestStdEncIntf(t, testSimpleH)
}

func TestMsgpackCodecsTable(t *testing.T) {
	doTestCodecTableOne(t, testMsgpackH)
}

func TestMsgpackCodecsMisc(t *testing.T) {
	doTestCodecMiscOne(t, testMsgpackH)
}

func TestMsgpackCodecsEmbeddedPointer(t *testing.T) {
	doTestCodecEmbeddedPointer(t, testMsgpackH)
}

func TestMsgpackStdEncIntf(t *testing.T) {
	doTestStdEncIntf(t, testMsgpackH)
}

func TestCborCodecsTable(t *testing.T) {
	doTestCodecTableOne(t, testCborH)
}

func TestCborCodecsMisc(t *testing.T) {
	doTestCodecMiscOne(t, testCborH)
}

func TestCborCodecsEmbeddedPointer(t *testing.T) {
	doTestCodecEmbeddedPointer(t, testCborH)
}

func TestCborCodecChan(t *testing.T) {
	doTestCodecChan(t, testCborH)
}

func TestCborStdEncIntf(t *testing.T) {
	doTestStdEncIntf(t, testCborH)
}

func TestJsonCodecsTable(t *testing.T) {
	doTestCodecTableOne(t, testJsonH)
}

func TestJsonCodecsMisc(t *testing.T) {
	doTestCodecMiscOne(t, testJsonH)
}

func TestJsonCodecsEmbeddedPointer(t *testing.T) {
	doTestCodecEmbeddedPointer(t, testJsonH)
}

func TestJsonCodecChan(t *testing.T) {
	doTestCodecChan(t, testJsonH)
}

func TestJsonStdEncIntf(t *testing.T) {
	doTestStdEncIntf(t, testJsonH)
}

// ----- Raw ---------
func TestJsonRaw(t *testing.T) {
	doTestRawValue(t, testJsonH)
}
func TestBincRaw(t *testing.T) {
	doTestRawValue(t, testBincH)
}
func TestMsgpackRaw(t *testing.T) {
	doTestRawValue(t, testMsgpackH)
}
func TestSimpleRaw(t *testing.T) {
	doTestRawValue(t, testSimpleH)
}
func TestCborRaw(t *testing.T) {
	doTestRawValue(t, testCborH)
}

// ----- ALL (framework based) -----

func TestAllEncCircularRef(t *testing.T) {
	doTestEncCircularRef(t, testCborH)
}

func TestAllAnonCycle(t *testing.T) {
	doTestAnonCycle(t, testCborH)
}

func TestAllErrWriter(t *testing.T) {
	doTestAllErrWriter(t, testCborH, testJsonH)
}

// ----- RPC custom -----

func TestMsgpackRpcSpec(t *testing.T) {
	doTestCodecRpcOne(t, MsgpackSpecRpc, testMsgpackH, true, 0)
}

// ----- RPC -----

func TestBincRpcGo(t *testing.T) {
	doTestCodecRpcOne(t, GoRpc, testBincH, true, 0)
}

func TestSimpleRpcGo(t *testing.T) {
	doTestCodecRpcOne(t, GoRpc, testSimpleH, true, 0)
}

func TestMsgpackRpcGo(t *testing.T) {
	doTestCodecRpcOne(t, GoRpc, testMsgpackH, true, 0)
}

func TestCborRpcGo(t *testing.T) {
	doTestCodecRpcOne(t, GoRpc, testCborH, true, 0)
}

func TestJsonRpcGo(t *testing.T) {
	doTestCodecRpcOne(t, GoRpc, testJsonH, true, 0)
}

// ----- OTHERS -----

func TestBincMapEncodeForCanonical(t *testing.T) {
	t.Skipf("skipping ... needs investigation") // MARKER: testing fails??? Need to research
	doTestMapEncodeForCanonical(t, testBincH)
}

func TestSimpleMapEncodeForCanonical(t *testing.T) {
	doTestMapEncodeForCanonical(t, testSimpleH)
}

func TestMsgpackMapEncodeForCanonical(t *testing.T) {
	doTestMapEncodeForCanonical(t, testMsgpackH)
}

func TestCborMapEncodeForCanonical(t *testing.T) {
	doTestMapEncodeForCanonical(t, testCborH)
}

func TestJsonMapEncodeForCanonical(t *testing.T) {
	doTestMapEncodeForCanonical(t, testJsonH)
}

func TestBincUnderlyingType(t *testing.T) {
	testCodecUnderlyingType(t, testBincH)
}

func TestJsonSwallowAndZero(t *testing.T) {
	doTestSwallowAndZero(t, testJsonH)
}

func TestCborSwallowAndZero(t *testing.T) {
	doTestSwallowAndZero(t, testCborH)
}

func TestMsgpackSwallowAndZero(t *testing.T) {
	doTestSwallowAndZero(t, testMsgpackH)
}

func TestBincSwallowAndZero(t *testing.T) {
	doTestSwallowAndZero(t, testBincH)
}

func TestSimpleSwallowAndZero(t *testing.T) {
	doTestSwallowAndZero(t, testSimpleH)
}

func TestJsonRawExt(t *testing.T) {
	doTestRawExt(t, testJsonH)
}

func TestCborRawExt(t *testing.T) {
	doTestRawExt(t, testCborH)
}

func TestMsgpackRawExt(t *testing.T) {
	doTestRawExt(t, testMsgpackH)
}

func TestBincRawExt(t *testing.T) {
	doTestRawExt(t, testBincH)
}

func TestSimpleRawExt(t *testing.T) {
	doTestRawExt(t, testSimpleH)
}

func TestJsonMapStructKey(t *testing.T) {
	doTestMapStructKey(t, testJsonH)
}

func TestCborMapStructKey(t *testing.T) {
	doTestMapStructKey(t, testCborH)
}

func TestMsgpackMapStructKey(t *testing.T) {
	doTestMapStructKey(t, testMsgpackH)
}

func TestBincMapStructKey(t *testing.T) {
	doTestMapStructKey(t, testBincH)
}

func TestSimpleMapStructKey(t *testing.T) {
	doTestMapStructKey(t, testSimpleH)
}

func TestJsonDecodeNilMapValue(t *testing.T) {
	doTestDecodeNilMapValue(t, testJsonH)
}

func TestCborDecodeNilMapValue(t *testing.T) {
	doTestDecodeNilMapValue(t, testCborH)
}

func TestMsgpackDecodeNilMapValue(t *testing.T) {
	doTestDecodeNilMapValue(t, testMsgpackH)
}

func TestBincDecodeNilMapValue(t *testing.T) {
	doTestDecodeNilMapValue(t, testBincH)
}

func TestSimpleDecodeNilMapValue(t *testing.T) {
	doTestDecodeNilMapValue(t, testSimpleH)
}

func TestJsonEmbeddedFieldPrecedence(t *testing.T) {
	doTestEmbeddedFieldPrecedence(t, testJsonH)
}

func TestCborEmbeddedFieldPrecedence(t *testing.T) {
	doTestEmbeddedFieldPrecedence(t, testCborH)
}

func TestMsgpackEmbeddedFieldPrecedence(t *testing.T) {
	doTestEmbeddedFieldPrecedence(t, testMsgpackH)
}

func TestBincEmbeddedFieldPrecedence(t *testing.T) {
	doTestEmbeddedFieldPrecedence(t, testBincH)
}

func TestSimpleEmbeddedFieldPrecedence(t *testing.T) {
	doTestEmbeddedFieldPrecedence(t, testSimpleH)
}

func TestJsonLargeContainerLen(t *testing.T) {
	doTestLargeContainerLen(t, testJsonH)
}

func TestCborLargeContainerLen(t *testing.T) {
	doTestLargeContainerLen(t, testCborH)
}

func TestMsgpackLargeContainerLen(t *testing.T) {
	doTestLargeContainerLen(t, testMsgpackH)
}

func TestBincLargeContainerLen(t *testing.T) {
	doTestLargeContainerLen(t, testBincH)
}

func TestSimpleLargeContainerLen(t *testing.T) {
	doTestLargeContainerLen(t, testSimpleH)
}

func TestJsonTime(t *testing.T) {
	doTestTime(t, testJsonH)
}

func TestCborTime(t *testing.T) {
	doTestTime(t, testCborH)
}

func TestMsgpackTime(t *testing.T) {
	doTestTime(t, testMsgpackH)
}

func TestBincTime(t *testing.T) {
	doTestTime(t, testBincH)
}

func TestSimpleTime(t *testing.T) {
	doTestTime(t, testSimpleH)
}

func TestJsonUintToInt(t *testing.T) {
	doTestUintToInt(t, testJsonH)
}

func TestCborUintToInt(t *testing.T) {
	doTestUintToInt(t, testCborH)
}

func TestMsgpackUintToInt(t *testing.T) {
	doTestUintToInt(t, testMsgpackH)
}

func TestBincUintToInt(t *testing.T) {
	doTestUintToInt(t, testBincH)
}

func TestSimpleUintToInt(t *testing.T) {
	doTestUintToInt(t, testSimpleH)
}

func TestJsonDifferentMapOrSliceType(t *testing.T) {
	doTestDifferentMapOrSliceType(t, testJsonH)
}

func TestCborDifferentMapOrSliceType(t *testing.T) {
	doTestDifferentMapOrSliceType(t, testCborH)
}

func TestMsgpackDifferentMapOrSliceType(t *testing.T) {
	doTestDifferentMapOrSliceType(t, testMsgpackH)
}

func TestBincDifferentMapOrSliceType(t *testing.T) {
	doTestDifferentMapOrSliceType(t, testBincH)
}

func TestSimpleDifferentMapOrSliceType(t *testing.T) {
	doTestDifferentMapOrSliceType(t, testSimpleH)
}

func TestJsonScalars(t *testing.T) {
	doTestScalars(t, testJsonH)
}

func TestCborScalars(t *testing.T) {
	doTestScalars(t, testCborH)
}

func TestMsgpackScalars(t *testing.T) {
	doTestScalars(t, testMsgpackH)
}

func TestBincScalars(t *testing.T) {
	doTestScalars(t, testBincH)
}

func TestSimpleScalars(t *testing.T) {
	doTestScalars(t, testSimpleH)
}

func TestJsonOmitempty(t *testing.T) {
	doTestOmitempty(t, testJsonH)
}

func TestCborOmitempty(t *testing.T) {
	doTestOmitempty(t, testCborH)
}

func TestMsgpackOmitempty(t *testing.T) {
	doTestOmitempty(t, testMsgpackH)
}

func TestBincOmitempty(t *testing.T) {
	doTestOmitempty(t, testBincH)
}

func TestSimpleOmitempty(t *testing.T) {
	doTestOmitempty(t, testSimpleH)
}

func TestJsonIntfMapping(t *testing.T) {
	doTestIntfMapping(t, testJsonH)
}

func TestCborIntfMapping(t *testing.T) {
	doTestIntfMapping(t, testCborH)
}

func TestMsgpackIntfMapping(t *testing.T) {
	doTestIntfMapping(t, testMsgpackH)
}

func TestBincIntfMapping(t *testing.T) {
	doTestIntfMapping(t, testBincH)
}

func TestSimpleIntfMapping(t *testing.T) {
	doTestIntfMapping(t, testSimpleH)
}

func TestJsonMissingFields(t *testing.T) {
	doTestMissingFields(t, testJsonH)
}

func TestCborMissingFields(t *testing.T) {
	doTestMissingFields(t, testCborH)
}

func TestMsgpackMissingFields(t *testing.T) {
	doTestMissingFields(t, testMsgpackH)
}

func TestBincMissingFields(t *testing.T) {
	doTestMissingFields(t, testBincH)
}

func TestSimpleMissingFields(t *testing.T) {
	doTestMissingFields(t, testSimpleH)
}

func TestJsonMaxDepth(t *testing.T) {
	doTestMaxDepth(t, testJsonH)
}

func TestCborMaxDepth(t *testing.T) {
	doTestMaxDepth(t, testCborH)
}

func TestMsgpackMaxDepth(t *testing.T) {
	doTestMaxDepth(t, testMsgpackH)
}

func TestBincMaxDepth(t *testing.T) {
	doTestMaxDepth(t, testBincH)
}

func TestSimpleMaxDepth(t *testing.T) {
	doTestMaxDepth(t, testSimpleH)
}

func TestJsonSelfExt(t *testing.T) {
	doTestSelfExt(t, testJsonH)
}

func TestCborSelfExt(t *testing.T) {
	doTestSelfExt(t, testCborH)
}

func TestMsgpackSelfExt(t *testing.T) {
	doTestSelfExt(t, testMsgpackH)
}

func TestBincSelfExt(t *testing.T) {
	doTestSelfExt(t, testBincH)
}

func TestSimpleSelfExt(t *testing.T) {
	doTestSelfExt(t, testSimpleH)
}

func TestJsonBytesEncodedAsArray(t *testing.T) {
	doTestBytesEncodedAsArray(t, testJsonH)
}

func TestCborBytesEncodedAsArray(t *testing.T) {
	doTestBytesEncodedAsArray(t, testCborH)
}

func TestMsgpackBytesEncodedAsArray(t *testing.T) {
	doTestBytesEncodedAsArray(t, testMsgpackH)
}

func TestBincBytesEncodedAsArray(t *testing.T) {
	doTestBytesEncodedAsArray(t, testBincH)
}

func TestSimpleBytesEncodedAsArray(t *testing.T) {
	doTestBytesEncodedAsArray(t, testSimpleH)
}

func TestJsonStrucEncDec(t *testing.T) {
	doTestStrucEncDec(t, testJsonH)
}

func TestCborStrucEncDec(t *testing.T) {
	doTestStrucEncDec(t, testCborH)
}

func TestMsgpackStrucEncDec(t *testing.T) {
	doTestStrucEncDec(t, testMsgpackH)
}

func TestBincStrucEncDec(t *testing.T) {
	doTestStrucEncDec(t, testBincH)
}

func TestSimpleStrucEncDec(t *testing.T) {
	doTestStrucEncDec(t, testSimpleH)
}

func TestJsonRawToStringToRawEtc(t *testing.T) {
	doTestRawToStringToRawEtc(t, testJsonH)
}

func TestCborRawToStringToRawEtc(t *testing.T) {
	doTestRawToStringToRawEtc(t, testCborH)
}

func TestMsgpackRawToStringToRawEtc(t *testing.T) {
	doTestRawToStringToRawEtc(t, testMsgpackH)
}

func TestBincRawToStringToRawEtc(t *testing.T) {
	doTestRawToStringToRawEtc(t, testBincH)
}

func TestSimpleRawToStringToRawEtc(t *testing.T) {
	doTestRawToStringToRawEtc(t, testSimpleH)
}

func TestJsonStructKeyType(t *testing.T) {
	doTestStructKeyType(t, testJsonH)
}

func TestCborStructKeyType(t *testing.T) {
	doTestStructKeyType(t, testCborH)
}

func TestMsgpackStructKeyType(t *testing.T) {
	doTestStructKeyType(t, testMsgpackH)
}

func TestBincStructKeyType(t *testing.T) {
	doTestStructKeyType(t, testBincH)
}

func TestSimpleStructKeyType(t *testing.T) {
	doTestStructKeyType(t, testSimpleH)
}

func TestJsonPreferArrayOverSlice(t *testing.T) {
	doTestPreferArrayOverSlice(t, testJsonH)
}

func TestCborPreferArrayOverSlice(t *testing.T) {
	doTestPreferArrayOverSlice(t, testCborH)
}

func TestMsgpackPreferArrayOverSlice(t *testing.T) {
	doTestPreferArrayOverSlice(t, testMsgpackH)
}

func TestBincPreferArrayOverSlice(t *testing.T) {
	doTestPreferArrayOverSlice(t, testBincH)
}

func TestSimplePreferArrayOverSlice(t *testing.T) {
	doTestPreferArrayOverSlice(t, testSimpleH)
}

func TestJsonZeroCopyBytes(t *testing.T) {
	doTestZeroCopyBytes(t, testJsonH)
}

func TestCborZeroCopyBytes(t *testing.T) {
	doTestZeroCopyBytes(t, testCborH)
}

func TestMsgpackZeroCopyBytes(t *testing.T) {
	doTestZeroCopyBytes(t, testMsgpackH)
}

func TestBincZeroCopyBytes(t *testing.T) {
	doTestZeroCopyBytes(t, testBincH)
}

func TestSimpleZeroCopyBytes(t *testing.T) {
	doTestZeroCopyBytes(t, testSimpleH)
}

func TestJsonNextValueBytes(t *testing.T) {
	doTestNextValueBytes(t, testJsonH)
}

func TestCborNextValueBytes(t *testing.T) {
	// x := testCborH.IndefiniteLength
	// defer func() { testCborH.IndefiniteLength = x }()

	// xdebugf(">>>>> TestCborNextValueBytes: IndefiniteLength = false")
	// testCborH.IndefiniteLength = false
	// doTestNextValueBytes(t, testCborH)
	// xdebugf(">>>>> TestCborNextValueBytes: IndefiniteLength = true")
	// testCborH.IndefiniteLength = true
	doTestNextValueBytes(t, testCborH)
}

func TestMsgpackNextValueBytes(t *testing.T) {
	doTestNextValueBytes(t, testMsgpackH)
}

func TestBincNextValueBytes(t *testing.T) {
	doTestNextValueBytes(t, testBincH)
}

func TestSimpleNextValueBytes(t *testing.T) {
	doTestNextValueBytes(t, testSimpleH)
}

func TestJsonNumbers(t *testing.T) {
	doTestNumbers(t, testJsonH)
}

func TestCborNumbers(t *testing.T) {
	doTestNumbers(t, testCborH)
}

func TestMsgpackNumbers(t *testing.T) {
	doTestNumbers(t, testMsgpackH)
}

func TestBincNumbers(t *testing.T) {
	doTestNumbers(t, testBincH)
}

func TestSimpleNumbers(t *testing.T) {
	doTestNumbers(t, testSimpleH)
}

func TestJsonDesc(t *testing.T) {
	doTestDesc(t, testJsonH, map[byte]string{'"': `"`, '{': `{`, '}': `}`, '[': `[`, ']': `]`})
}

func TestCborDesc(t *testing.T) {
	m := make(map[byte]string)
	for k, v := range cbordescMajorNames {
		// if k == cborMajorSimpleOrFloat { m[k<<5] = "nil" }
		m[k<<5] = v
	}
	for k, v := range cbordescSimpleNames {
		m[k] = v
	}
	delete(m, cborMajorSimpleOrFloat<<5)
	doTestDesc(t, testCborH, m)
}

func TestMsgpackDesc(t *testing.T) {
	m := make(map[byte]string)
	for k, v := range mpdescNames {
		m[k] = v
	}
	m[mpPosFixNumMin] = "int"
	m[mpFixStrMin] = "string|bytes"
	m[mpFixArrayMin] = "array"
	m[mpFixMapMin] = "map"
	m[mpFixExt1] = "ext"

	doTestDesc(t, testMsgpackH, m)
}

func TestBincDesc(t *testing.T) {
	m := make(map[byte]string)
	for k, v := range bincdescVdNames {
		m[k<<4] = v
	}
	for k, v := range bincdescSpecialVsNames {
		m[k] = v
	}
	delete(m, bincVdSpecial<<4)
	doTestDesc(t, testBincH, m)
}

func TestSimpleDesc(t *testing.T) {
	doTestDesc(t, testSimpleH, simpledescNames)
}

func TestJsonStructFieldInfoToArray(t *testing.T) {
	doTestStructFieldInfoToArray(t, testJsonH)
}

func TestCborStructFieldInfoToArray(t *testing.T) {
	doTestStructFieldInfoToArray(t, testCborH)
}

func TestMsgpackStructFieldInfoToArray(t *testing.T) {
	doTestStructFieldInfoToArray(t, testMsgpackH)
}

func TestBincStructFieldInfoToArray(t *testing.T) {
	doTestStructFieldInfoToArray(t, testBincH)
}

func TestSimpleStructFieldInfoToArray(t *testing.T) {
	doTestStructFieldInfoToArray(t, testSimpleH)
}

// --------

func TestMultipleEncDec(t *testing.T) {
	doTestMultipleEncDec(t, testJsonH)
}

func TestJsonEncodeIndent(t *testing.T) {
	doTestJsonEncodeIndent(t, testJsonH)
}

func TestJsonDecodeNonStringScalarInStringContext(t *testing.T) {
	doTestJsonDecodeNonStringScalarInStringContext(t, testJsonH)
}

func TestJsonLargeInteger(t *testing.T) {
	doTestJsonLargeInteger(t, testJsonH)
}

func TestJsonInvalidUnicode(t *testing.T) {
	doTestJsonInvalidUnicode(t, testJsonH)
}

func TestJsonNumberParsing(t *testing.T) {
	doTestJsonNumberParsing(t, testJsonH)
}

func TestMsgpackDecodeMapAndExtSizeMismatch(t *testing.T) {
	doTestMsgpackDecodeMapAndExtSizeMismatch(t, testMsgpackH)
}
