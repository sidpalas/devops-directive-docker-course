// Copyright (c) 2012-2020 Ugorji Nwoke. All rights reserved.
// Use of this source code is governed by a MIT license found in the LICENSE file.

//go:build x
// +build x

package codec

// These tests are used to verify msgpack and cbor implementations against their python libraries.
// If you have the library installed, you can enable the tests back by running: go test -tags=x .
// Look at test.py for how to setup your environment.
//
// Also, name all tests here with suffix X, so that we can filter to run only these
// using: go test -tags x -run '.*X$'

import (
	"testing"
)

func TestMsgpackPythonGenStreamsX(t *testing.T) {
	doTestPythonGenStreams(t, testMsgpackH)
}

func TestCborPythonGenStreamsX(t *testing.T) {
	doTestPythonGenStreams(t, testCborH)
}

func TestMsgpackRpcSpecGoClientToPythonSvcX(t *testing.T) {
	doTestMsgpackRpcSpecGoClientToPythonSvc(t, testMsgpackH)
}

func TestMsgpackRpcSpecPythonClientToGoSvcX(t *testing.T) {
	doTestMsgpackRpcSpecPythonClientToGoSvc(t, testMsgpackH)
}
