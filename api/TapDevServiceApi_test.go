package api

import (
	"reflect"
	"testing"
)

func TestImportProof(t *testing.T) {
	want := "*rfqrpc.ImportProofResponse"
	got := reflect.TypeOf(ImportProof("", "")).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
