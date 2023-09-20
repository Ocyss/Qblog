package rpc

import (
	"testing"
)

func Log(t *testing.T, resp any, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func Logs(t *testing.T, resp any, err error) {
	t.Helper()
	if err != nil {
		t.Log(err)
	} else {
		t.Log(resp)
	}
}
