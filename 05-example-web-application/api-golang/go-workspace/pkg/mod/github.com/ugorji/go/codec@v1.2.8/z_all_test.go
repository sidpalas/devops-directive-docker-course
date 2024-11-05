// Copyright (c) 2012-2020 Ugorji Nwoke. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

//go:build alltests && go1.9
// +build alltests,go1.9

package codec

// Run this using:
//   go test -tags=alltests -run=Suite -coverprofile=cov.out
//   go tool cover -html=cov.out
//
// Because build tags are a build time parameter, we will have to test out the
// different tags separately.
// Tags: x codecgen codec.safe codec.notfastpath appengine
//
// These tags should be added to alltests, e.g.
//   go test '-tags=alltests x codecgen' -run=Suite -coverprofile=cov.out
//
// To run all tests before submitting code, run:
//    a=( "" "codec.safe" "codecgen" "codec.notfastpath" "codecgen codec.notfastpath" "codecgen codec.safe" "codec.safe codec.notfastpath" )
//    for i in "${a[@]}"; do echo ">>>> TAGS: $i"; go test "-tags=alltests $i" -run=Suite; done
//
// This suite of tests requires support for subtests and suites,
// and conseqneutly will only run on go1.7 and above.

// find . -name "codec_test.go" | xargs grep -e '^func Test' | \
//     cut -d '(' -f 1 | cut -d ' ' -f 2 | \
//     while read f; do echo "t.Run(\"$f\", $f)"; done

import (
	"testing"
	"time"
)

// func TestMain(m *testing.M) {
// 	println("calling TestMain")
// 	// set some parameters
// 	exitcode := m.Run()
// 	os.Exit(exitcode)
// }

type testTimeTracker struct {
	t time.Time
}

func (tt *testTimeTracker) Elapsed() (d time.Duration) {
	if !tt.t.IsZero() {
		d = time.Since(tt.t)
	}
	tt.t = time.Now()
	return
}

func testGroupResetFlags() {
	testRpcBufsize = 2048
	testUseIoEncDec = -1
	testUseReset = false
	testUseParallel = false
	testMaxInitLen = 0
	testUseIoWrapper = false
	testNumRepeatString = 8
	testDepth = 0
	testDecodeOptions = DecodeOptions{}
	testEncodeOptions = EncodeOptions{}

	testJsonH.Indent = 0 // -1?
	testJsonH.HTMLCharsAsIs = false
	testJsonH.MapKeyAsString = false
	testJsonH.PreferFloat = false

	testCborH.IndefiniteLength = false
	testCborH.TimeRFC3339 = false
	testCborH.SkipUnexpectedTags = false
	testBincH.AsSymbols = 0 // 2? AsSymbolNone

	testMsgpackH.WriteExt = false
	testMsgpackH.NoFixedNum = false
	testMsgpackH.PositiveIntUnsigned = false

	testSimpleH.EncZeroValuesAsNil = false
}

func testCodecGroup(t *testing.T) {
	// <setup code>
	testJsonGroup(t)
	testBincGroup(t)
	testCborGroup(t)
	testMsgpackGroup(t)
	testSimpleGroup(t)
	// testSimpleMammothGroup(t)
	// testRpcGroup(t)
	testNonHandlesGroup(t)

	testCodecGroupV(t)
	// <tear-down code>
}

func testCodecGroupV(t *testing.T) {
	testJsonGroupV(t)
	testBincGroupV(t)
	testCborGroupV(t)
	testMsgpackGroupV(t)
	testSimpleGroupV(t)
}

func testJsonGroup(t *testing.T) {
	t.Run("TestJsonCodecsEmbeddedPointer", TestJsonCodecsEmbeddedPointer)
	t.Run("TestJsonCodecChan", TestJsonCodecChan)
	t.Run("TestJsonStdEncIntf", TestJsonStdEncIntf)
	t.Run("TestJsonMammoth", TestJsonMammoth)
	t.Run("TestJsonRaw", TestJsonRaw)
	t.Run("TestJsonRpcGo", TestJsonRpcGo)
	t.Run("TestJsonLargeInteger", TestJsonLargeInteger)
	t.Run("TestJsonDecodeNonStringScalarInStringContext", TestJsonDecodeNonStringScalarInStringContext)
	t.Run("TestJsonEncodeIndent", TestJsonEncodeIndent)

	t.Run("TestJsonRawExt", TestJsonRawExt)
	t.Run("TestJsonMapStructKey", TestJsonMapStructKey)
	t.Run("TestJsonDecodeNilMapValue", TestJsonDecodeNilMapValue)
	t.Run("TestJsonEmbeddedFieldPrecedence", TestJsonEmbeddedFieldPrecedence)
	t.Run("TestJsonLargeContainerLen", TestJsonLargeContainerLen)
	t.Run("TestJsonMammothMapsAndSlices", TestJsonMammothMapsAndSlices)
	t.Run("TestJsonTime", TestJsonTime)
	t.Run("TestJsonUintToInt", TestJsonUintToInt)
	t.Run("TestJsonDifferentMapOrSliceType", TestJsonDifferentMapOrSliceType)
	t.Run("TestJsonScalars", TestJsonScalars)
	t.Run("TestJsonOmitempty", TestJsonOmitempty)
	t.Run("TestJsonIntfMapping", TestJsonIntfMapping)
	t.Run("TestJsonMissingFields", TestJsonMissingFields)
	t.Run("TestJsonMaxDepth", TestJsonMaxDepth)
	t.Run("TestJsonSelfExt", TestJsonSelfExt)
	t.Run("TestJsonBytesEncodedAsArray", TestJsonBytesEncodedAsArray)
	t.Run("TestJsonMapEncodeForCanonical", TestJsonMapEncodeForCanonical)
	t.Run("TestJsonRawToStringToRawEtc", TestJsonRawToStringToRawEtc)
	t.Run("TestJsonStructKeyType", TestJsonStructKeyType)
	t.Run("TestJsonPreferArrayOverSlice", TestJsonPreferArrayOverSlice)
	t.Run("TestJsonZeroCopyBytes", TestJsonZeroCopyBytes)
	t.Run("TestJsonNumbers", TestJsonNumbers)
	t.Run("TestJsonDesc", TestJsonDesc)
	t.Run("TestJsonStructFieldInfoToArray", TestJsonStructFieldInfoToArray)

	t.Run("TestJsonInvalidUnicode", TestJsonInvalidUnicode)
	t.Run("TestJsonNumberParsing", TestJsonNumberParsing)
}

func testJsonGroupV(t *testing.T) {
	t.Run("TestJsonCodecsTable", TestJsonCodecsTable)
	t.Run("TestJsonCodecsMisc", TestJsonCodecsMisc)
	t.Run("TestJsonSwallowAndZero", TestJsonSwallowAndZero)
	t.Run("TestJsonNextValueBytes", TestJsonNextValueBytes)
	t.Run("TestJsonStrucEncDec", TestJsonStrucEncDec)
}

func testBincGroup(t *testing.T) {
	t.Run("TestBincCodecsEmbeddedPointer", TestBincCodecsEmbeddedPointer)
	t.Run("TestBincStdEncIntf", TestBincStdEncIntf)
	t.Run("TestBincMammoth", TestBincMammoth)
	t.Run("TestBincRaw", TestBincRaw)
	t.Run("TestBincRpcGo", TestBincRpcGo)
	t.Run("TestBincUnderlyingType", TestBincUnderlyingType)

	t.Run("TestBincRawExt", TestBincRawExt)
	t.Run("TestBincMapStructKey", TestBincMapStructKey)
	t.Run("TestBincDecodeNilMapValue", TestBincDecodeNilMapValue)
	t.Run("TestBincEmbeddedFieldPrecedence", TestBincEmbeddedFieldPrecedence)
	t.Run("TestBincLargeContainerLen", TestBincLargeContainerLen)
	t.Run("TestBincMammothMapsAndSlices", TestBincMammothMapsAndSlices)
	t.Run("TestBincTime", TestBincTime)
	t.Run("TestBincUintToInt", TestBincUintToInt)
	t.Run("TestBincDifferentMapOrSliceType", TestBincDifferentMapOrSliceType)
	t.Run("TestBincScalars", TestBincScalars)
	t.Run("TestBincOmitempty", TestBincOmitempty)
	t.Run("TestBincIntfMapping", TestBincIntfMapping)
	t.Run("TestBincMissingFields", TestBincMissingFields)
	t.Run("TestBincMaxDepth", TestBincMaxDepth)
	t.Run("TestBincSelfExt", TestBincSelfExt)
	t.Run("TestBincBytesEncodedAsArray", TestBincBytesEncodedAsArray)
	t.Run("TestBincMapEncodeForCanonical", TestBincMapEncodeForCanonical)
	t.Run("TestBincRawToStringToRawEtc", TestBincRawToStringToRawEtc)
	t.Run("TestBincStructKeyType", TestBincStructKeyType)
	t.Run("TestBincPreferArrayOverSlice", TestBincPreferArrayOverSlice)
	t.Run("TestBincZeroCopyBytes", TestBincZeroCopyBytes)
	t.Run("TestBincNumbers", TestBincNumbers)
	t.Run("TestBincDesc", TestBincDesc)
	t.Run("TestBincStructFieldInfoToArray", TestBincStructFieldInfoToArray)
}

func testBincGroupV(t *testing.T) {
	t.Run("TestBincCodecsTable", TestBincCodecsTable)
	t.Run("TestBincCodecsMisc", TestBincCodecsMisc)
	t.Run("TestBincSwallowAndZero", TestBincSwallowAndZero)
	t.Run("TestBincNextValueBytes", TestBincNextValueBytes)
	t.Run("TestBincStrucEncDec", TestBincStrucEncDec)
}

func testCborGroup(t *testing.T) {
	t.Run("TestCborCodecsEmbeddedPointer", TestCborCodecsEmbeddedPointer)
	t.Run("TestCborCodecChan", TestCborCodecChan)
	t.Run("TestCborStdEncIntf", TestCborStdEncIntf)
	t.Run("TestCborMammoth", TestCborMammoth)
	t.Run("TestCborRaw", TestCborRaw)
	t.Run("TestCborRpcGo", TestCborRpcGo)

	t.Run("TestCborRawExt", TestCborRawExt)
	t.Run("TestCborMapStructKey", TestCborMapStructKey)
	t.Run("TestCborDecodeNilMapValue", TestCborDecodeNilMapValue)
	t.Run("TestCborEmbeddedFieldPrecedence", TestCborEmbeddedFieldPrecedence)
	t.Run("TestCborLargeContainerLen", TestCborLargeContainerLen)
	t.Run("TestCborMammothMapsAndSlices", TestCborMammothMapsAndSlices)
	t.Run("TestCborTime", TestCborTime)
	t.Run("TestCborUintToInt", TestCborUintToInt)
	t.Run("TestCborDifferentMapOrSliceType", TestCborDifferentMapOrSliceType)
	t.Run("TestCborScalars", TestCborScalars)
	t.Run("TestCborOmitempty", TestCborOmitempty)
	t.Run("TestCborIntfMapping", TestCborIntfMapping)
	t.Run("TestCborMissingFields", TestCborMissingFields)
	t.Run("TestCborMaxDepth", TestCborMaxDepth)
	t.Run("TestCborSelfExt", TestCborSelfExt)
	t.Run("TestCborBytesEncodedAsArray", TestCborBytesEncodedAsArray)
	t.Run("TestCborMapEncodeForCanonical", TestCborMapEncodeForCanonical)
	t.Run("TestCborRawToStringToRawEtc", TestCborRawToStringToRawEtc)
	t.Run("TestCborStructKeyType", TestCborStructKeyType)
	t.Run("TestCborPreferArrayOverSlice", TestCborPreferArrayOverSlice)
	t.Run("TestCborZeroCopyBytes", TestCborZeroCopyBytes)
	t.Run("TestCborNumbers", TestCborNumbers)
	t.Run("TestCborDesc", TestCborDesc)
	t.Run("TestCborStructFieldInfoToArray", TestCborStructFieldInfoToArray)

	t.Run("TestCborHalfFloat", TestCborHalfFloat)
	t.Run("TestCborSkipTags", TestCborSkipTags)
}

func testCborGroupV(t *testing.T) {
	t.Run("TestCborCodecsTable", TestCborCodecsTable)
	t.Run("TestCborCodecsMisc", TestCborCodecsMisc)
	t.Run("TestCborSwallowAndZero", TestCborSwallowAndZero)
	t.Run("TestCborNextValueBytes", TestCborNextValueBytes)
	t.Run("TestCborStrucEncDec", TestCborStrucEncDec)
}

func testMsgpackGroup(t *testing.T) {
	t.Run("TestMsgpackCodecsEmbeddedPointer", TestMsgpackCodecsEmbeddedPointer)
	t.Run("TestMsgpackStdEncIntf", TestMsgpackStdEncIntf)
	t.Run("TestMsgpackMammoth", TestMsgpackMammoth)
	t.Run("TestMsgpackRaw", TestMsgpackRaw)
	t.Run("TestMsgpackRpcGo", TestMsgpackRpcGo)
	t.Run("TestMsgpackRpcSpec", TestMsgpackRpcSpec)

	t.Run("TestMsgpackRawExt", TestMsgpackRawExt)
	t.Run("TestMsgpackMapStructKey", TestMsgpackMapStructKey)
	t.Run("TestMsgpackDecodeNilMapValue", TestMsgpackDecodeNilMapValue)
	t.Run("TestMsgpackEmbeddedFieldPrecedence", TestMsgpackEmbeddedFieldPrecedence)
	t.Run("TestMsgpackLargeContainerLen", TestMsgpackLargeContainerLen)
	t.Run("TestMsgpackMammothMapsAndSlices", TestMsgpackMammothMapsAndSlices)
	t.Run("TestMsgpackTime", TestMsgpackTime)
	t.Run("TestMsgpackUintToInt", TestMsgpackUintToInt)
	t.Run("TestMsgpackDifferentMapOrSliceType", TestMsgpackDifferentMapOrSliceType)
	t.Run("TestMsgpackScalars", TestMsgpackScalars)
	t.Run("TestMsgpackOmitempty", TestMsgpackOmitempty)
	t.Run("TestMsgpackIntfMapping", TestMsgpackIntfMapping)
	t.Run("TestMsgpackMissingFields", TestMsgpackMissingFields)
	t.Run("TestMsgpackMaxDepth", TestMsgpackMaxDepth)
	t.Run("TestMsgpackSelfExt", TestMsgpackSelfExt)
	t.Run("TestMsgpackBytesEncodedAsArray", TestMsgpackBytesEncodedAsArray)
	t.Run("TestMsgpackMapEncodeForCanonical", TestMsgpackMapEncodeForCanonical)
	t.Run("TestMsgpackRawToStringToRawEtc", TestMsgpackRawToStringToRawEtc)
	t.Run("TestMsgpackStructKeyType", TestMsgpackStructKeyType)
	t.Run("TestMsgpackPreferArrayOverSlice", TestMsgpackPreferArrayOverSlice)
	t.Run("TestMsgpackZeroCopyBytes", TestMsgpackZeroCopyBytes)
	t.Run("TestMsgpackNumbers", TestMsgpackNumbers)
	t.Run("TestMsgpackDesc", TestMsgpackDesc)
	t.Run("TestMsgpackStructFieldInfoToArray", TestMsgpackStructFieldInfoToArray)

	t.Run("TestMsgpackDecodeMapAndExtSizeMismatch", TestMsgpackDecodeMapAndExtSizeMismatch)
}

func testMsgpackGroupV(t *testing.T) {
	t.Run("TestMsgpackCodecsTable", TestMsgpackCodecsTable)
	t.Run("TestMsgpackCodecsMisc", TestMsgpackCodecsMisc)
	t.Run("TestMsgpackSwallowAndZero", TestMsgpackSwallowAndZero)
	t.Run("TestMsgpackNextValueBytes", TestMsgpackNextValueBytes)
	t.Run("TestMsgpackStrucEncDec", TestMsgpackStrucEncDec)
}

func testSimpleGroup(t *testing.T) {
	t.Run("TestSimpleCodecsEmbeddedPointer", TestSimpleCodecsEmbeddedPointer)
	t.Run("TestSimpleStdEncIntf", TestSimpleStdEncIntf)
	t.Run("TestSimpleMammoth", TestSimpleMammoth)
	t.Run("TestSimpleRaw", TestSimpleRaw)
	t.Run("TestSimpleRpcGo", TestSimpleRpcGo)

	t.Run("TestSimpleRawExt", TestSimpleRawExt)
	t.Run("TestSimpleMapStructKey", TestSimpleMapStructKey)
	t.Run("TestSimpleDecodeNilMapValue", TestSimpleDecodeNilMapValue)
	t.Run("TestSimpleEmbeddedFieldPrecedence", TestSimpleEmbeddedFieldPrecedence)
	t.Run("TestSimpleLargeContainerLen", TestSimpleLargeContainerLen)
	t.Run("TestSimpleMammothMapsAndSlices", TestSimpleMammothMapsAndSlices)
	t.Run("TestSimpleTime", TestSimpleTime)
	t.Run("TestSimpleUintToInt", TestSimpleUintToInt)
	t.Run("TestSimpleDifferentMapOrSliceType", TestSimpleDifferentMapOrSliceType)
	t.Run("TestSimpleScalars", TestSimpleScalars)
	t.Run("TestSimpleOmitempty", TestSimpleOmitempty)
	t.Run("TestSimpleIntfMapping", TestSimpleIntfMapping)
	t.Run("TestSimpleMissingFields", TestSimpleMissingFields)
	t.Run("TestSimpleMaxDepth", TestSimpleMaxDepth)
	t.Run("TestSimpleSelfExt", TestSimpleSelfExt)
	t.Run("TestSimpleBytesEncodedAsArray", TestSimpleBytesEncodedAsArray)
	t.Run("TestSimpleMapEncodeForCanonical", TestSimpleMapEncodeForCanonical)
	t.Run("TestSimpleRawToStringToRawEtc", TestSimpleRawToStringToRawEtc)
	t.Run("TestSimpleStructKeyType", TestSimpleStructKeyType)
	t.Run("TestSimplePreferArrayOverSlice", TestSimplePreferArrayOverSlice)
	t.Run("TestSimpleZeroCopyBytes", TestSimpleZeroCopyBytes)
	t.Run("TestSimpleNumbers", TestSimpleNumbers)
	t.Run("TestSimpleDesc", TestSimpleDesc)
	t.Run("TestSimpleStructFieldInfoToArray", TestSimpleStructFieldInfoToArray)
}

func testSimpleGroupV(t *testing.T) {
	t.Run("TestSimpleCodecsTable", TestSimpleCodecsTable)
	t.Run("TestSimpleCodecsMisc", TestSimpleCodecsMisc)
	t.Run("TestSimpleSwallowAndZero", TestSimpleSwallowAndZero)
	t.Run("TestSimpleNextValueBytes", TestSimpleNextValueBytes)
	t.Run("TestSimpleStrucEncDec", TestSimpleStrucEncDec)
}

func testSimpleMammothGroup(t *testing.T) {
	t.Run("TestSimpleMammothMapsAndSlices", TestSimpleMammothMapsAndSlices)
}

func testRpcGroup(t *testing.T) {
	t.Run("TestBincRpcGo", TestBincRpcGo)
	t.Run("TestSimpleRpcGo", TestSimpleRpcGo)
	t.Run("TestMsgpackRpcGo", TestMsgpackRpcGo)
	t.Run("TestCborRpcGo", TestCborRpcGo)
	t.Run("TestJsonRpcGo", TestJsonRpcGo)
	t.Run("TestMsgpackRpcSpec", TestMsgpackRpcSpec)
}

func testNonHandlesGroup(t *testing.T) {
	// grep "func Test" codec_test.go | grep -v -E '(Cbor|Json|Simple|Msgpack|Binc)'
	t.Run("TestBufioDecReader", TestBufioDecReader)
	t.Run("TestAtomic", TestAtomic)
	t.Run("TestAllEncCircularRef", TestAllEncCircularRef)
	t.Run("TestAllAnonCycle", TestAllAnonCycle)
	t.Run("TestMultipleEncDec", TestMultipleEncDec)
	t.Run("TestAllErrWriter", TestAllErrWriter)
	t.Run("TestMapRangeIndex", TestMapRangeIndex)
}

func TestCodecSuite(t *testing.T) {
	var tt testTimeTracker
	tt.Elapsed()

	fnRun := func(s string, f func(t *testing.T)) {
		t.Run(s, f)
		// xdebugf("%s: %v", s, tt.Elapsed())
	}

	testGroupResetFlags()

	testReinit() // so flag.Parse() is called first, and never called again

	fnRun("optionsFalse", testCodecGroup)

	testUseIoEncDec = 0
	testUseReset = true
	testUseParallel = true

	testDecodeOptions.ZeroCopy = true
	testDecodeOptions.InternString = true
	testDecodeOptions.MapValueReset = true

	// testDecodeOptions.ErrorIfNoField = true // error, as expected fields not there
	// testDecodeOptions.ErrorIfNoArrayExpand = true // no error, but no error case either
	// testDecodeOptions.PreferArrayOverSlice = true // error??? because slice != array.
	testDecodeOptions.SignedInteger = true // error as deepEqual compares int64 to uint64
	testDecodeOptions.SliceElementReset = true
	testDecodeOptions.InterfaceReset = true
	testDecodeOptions.RawToString = true
	testDecodeOptions.PreferPointerForStructOrArray = true

	testEncodeOptions.StructToArray = true
	testEncodeOptions.Canonical = true
	testEncodeOptions.CheckCircularRef = true
	testEncodeOptions.RecursiveEmptyCheck = true
	testEncodeOptions.OptimumSize = true

	// testEncodeOptions.Raw = true
	// testEncodeOptions.StringToRaw = true

	testJsonH.HTMLCharsAsIs = true
	// testJsonH.MapKeyAsString = true
	// testJsonH.PreferFloat = true

	testCborH.IndefiniteLength = true
	testCborH.TimeRFC3339 = true
	testCborH.SkipUnexpectedTags = true

	testMsgpackH.WriteExt = true
	testMsgpackH.NoFixedNum = true
	testMsgpackH.PositiveIntUnsigned = true

	// testSimpleH.EncZeroValuesAsNil = true

	testReinit()
	fnRun("optionsTrue", testCodecGroup)

	testGroupResetFlags()

	// ---
	testDepth = 4
	if testing.Short() {
		testDepth = 2
	}
	testReinit()
	fnRun("optionsTrue-deepstruct", testCodecGroupV)
	testDepth = 0

	// ---
	// testEncodeOptions.AsSymbols = AsSymbolAll
	testUseIoWrapper = true
	testReinit()
	fnRun("optionsTrue-ioWrapper", testCodecGroupV)
	testUseIoWrapper = false

	// testUseIoEncDec = -1

	// ---
	// make buffer small enough so that we have to re-fill multiple times
	// and also such that writing a quoted struct name e.g. "LongFieldNameXYZ"
	// will require a re-fill, and test out bufioEncWriter.writeqstr well.
	// Due to last requirement, we prefer 16 to 128.
	testSkipRPCTests = true
	testUseIoEncDec = 16
	// testDecodeOptions.ReaderBufferSize = 128
	// testEncodeOptions.WriterBufferSize = 128
	testReinit()
	fnRun("optionsTrue-bufio", testCodecGroupV)
	// testDecodeOptions.ReaderBufferSize = 0
	// testEncodeOptions.WriterBufferSize = 0
	testSkipRPCTests = false
	testUseIoEncDec = -1

	// ---
	testNumRepeatString = 32
	testReinit()
	fnRun("optionsTrue-largestrings", testCodecGroupV)
	testNumRepeatString = 8

	testGroupResetFlags()

	// ---
	fnJsonReset := func(ml int, d int8, hca, mkas bool) func() {
		return func() {
			testMaxInitLen = ml
			testJsonH.Indent = d
			testJsonH.HTMLCharsAsIs = hca
			testJsonH.MapKeyAsString = mkas
		}
	}(testMaxInitLen, testJsonH.Indent, testJsonH.HTMLCharsAsIs, testJsonH.MapKeyAsString)

	testMaxInitLen = 10
	testJsonH.MapKeyAsString = true

	testJsonH.Indent = 8
	testJsonH.HTMLCharsAsIs = true
	testReinit()
	fnRun("json-spaces-htmlcharsasis-initLen10", testJsonGroup)

	testJsonH.Indent = -1
	testJsonH.HTMLCharsAsIs = false
	testReinit()
	fnRun("json-tabs-initLen10", testJsonGroup)

	fnJsonReset()

	// ---
	oldSymbols := testBincH.AsSymbols

	testBincH.AsSymbols = 2 // AsSymbolNone
	testReinit()
	fnRun("binc-no-symbols", testBincGroup)

	testBincH.AsSymbols = 1 // AsSymbolAll
	testReinit()
	fnRun("binc-all-symbols", testBincGroup)

	testBincH.AsSymbols = oldSymbols

	// ---
	oldEncZeroValuesAsNil := testSimpleH.EncZeroValuesAsNil
	testSimpleH.EncZeroValuesAsNil = !testSimpleH.EncZeroValuesAsNil
	testReinit()
	fnRun("simple-enczeroasnil", testSimpleMammothGroup) // testSimpleGroup
	testSimpleH.EncZeroValuesAsNil = oldEncZeroValuesAsNil

	// ---
	testUseIoEncDec = 16
	testRPCOptions.RPCNoBuffer = false
	testReinit()
	testRpcBufsize = 0
	fnRun("rpc-buf-0", testRpcGroup)
	testRpcBufsize = 0
	fnRun("rpc-buf-00", testRpcGroup)
	testRpcBufsize = 0
	fnRun("rpc-buf-000", testRpcGroup)
	testRpcBufsize = 16
	fnRun("rpc-buf-16", testRpcGroup)
	testRpcBufsize = 2048
	fnRun("rpc-buf-2048", testRpcGroup)

	testRPCOptions.RPCNoBuffer = true
	testRpcBufsize = 0
	fnRun("rpc-buf-0-rpcNoBuffer", testRpcGroup)
	testRpcBufsize = 0
	fnRun("rpc-buf-00-rpcNoBuffer", testRpcGroup)
	testRpcBufsize = 2048
	fnRun("rpc-buf-2048-rpcNoBuffer", testRpcGroup)

	testGroupResetFlags()
}

// func TestCodecSuite(t *testing.T) {
// 	testReinit() // so flag.Parse() is called first, and never called again
// 	testDecodeOptions, testEncodeOptions = DecodeOptions{}, EncodeOptions{}
// 	testGroupResetFlags()
// 	testReinit()
// 	t.Run("optionsFalse", func(t *testing.T) {
// 		t.Run("TestJsonMammothMapsAndSlices", TestJsonMammothMapsAndSlices)
// 	})
// }
