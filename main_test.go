package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

// Test for Sendrequest for Short Endpoint
func Test_SendRequestShortTrue(t *testing.T) {
	var expected = []byte(`{"uid":"187ef4436122d1cc2f40dc2b92f0eba0"}`)
	expected = append(expected, 10)
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// Send response to be tested
		rw.Write(expected)

	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	api := API{server.Client(), "https://ionaapp.com/assignment-magic/dk", "/short/ab"}
	got, _ := api.sendRequest()

	cmp := bytes.Compare(got, expected)
	if cmp != 0 {
		t.Errorf("response doest not match")
	}
}

// Test for Sendrequest for Long Endpoint

func Test_SendRequestLong(t *testing.T) {
	// Start a local HTTP server
	var expected = []byte(`{"uid":"53f1ce1310367adfe34e70e39c454d88"}`)
	expected = append(expected, 10)
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		// Send response to be tested
		rw.Write(expected)

	}))
	// Close the server when test finishes
	defer server.Close()

	// Use Client & URL from our local test server
	api := API{server.Client(), "https://ionaapp.com/assignment-magic/dk", "/long/ab2"}
	got, _ := api.sendRequest()

	cmp := bytes.Compare(got, expected)
	if cmp != 0 {
		t.Errorf("response doest not match")
	}
}

// Test for vaild UID
func TestUId(t *testing.T) {
	name := "187ef4436122d1cc2f40dc2b92f0eba0"
	want := true
	got := isValidUID(name)
	if want != got {
		t.Errorf("got UID is invaild, expected to be vaild")
	}
}

// Test for Invaild UID
func TestUUIdfalse(t *testing.T) {
	name := "187ef4436122d1cc2f40dc2b92f0eba02"
	want := false
	got := isValidUID(name)
	if want != got {
		t.Errorf("got UID is vaild, expected to be invaild ")
	}
}

// Test for vaildParam for long
func Test_ParamforlongTrue(t *testing.T) {
	name := "ab2"
	want := true
	got := isvaildparamforlong(name)
	if want != got {
		t.Errorf("got param  invaild, expected to be vaild")
	}
}

// Test for InvaildParam for long
func Test_ParamforlongFalse(t *testing.T) {
	name := "ab"
	want := false
	got := isvaildparamforlong(name)
	if want != got {
		t.Errorf("got param nvaild, expected to be invaild")
	}
}

// Test for vaildParam for short
func Test_ParamforShortTrue(t *testing.T) {
	name := "ab"
	want := true
	got := isvaildparamforshort(name)
	if want != got {
		t.Errorf("got param  invaild, expected to be vaild")
	}
}

// Test for InvaildParam for short
func Test_ParamforShortFalse(t *testing.T) {
	name := "ab2"
	want := false
	got := isvaildparamforshort(name)
	if want != got {
		t.Errorf("got param  vaild, expected to be invaild")
	}
}
