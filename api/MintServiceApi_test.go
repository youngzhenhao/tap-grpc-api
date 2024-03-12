package api

import (
	"reflect"
	"testing"
)

// === RUN   TestCancelBatch
// 2024/03/11 14:36:19 mintrpc CancelBatch Error: rpc error: code = Unknown desc = unable to cancel batch: multiple caretakers not supported
func TestCancelBatch(t *testing.T) {
	want := "*mintrpc.CancelBatchResponse"
	got := reflect.TypeOf(CancelBatch()).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// PASS
func TestFinalizeBatch(t *testing.T) {
	want := "*mintrpc.FinalizeBatchResponse"
	got := reflect.TypeOf(FinalizeBatch(false, 0)).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// PASS
func TestListBatches(t *testing.T) {
	want := "*mintrpc.ListBatchResponse"
	got := reflect.TypeOf(ListBatches()).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// PASS
func TestMintAsset(t *testing.T) {
	want := "*mintrpc.MintAssetResponse"
	got := reflect.TypeOf(MintAsset(
		false,
		false,
		"grpcTestMintAsset2",
		// "74657374206d696e74206173736574",
		"74657374206d696e742061737365742074776f",
		false,
		1000,
		true,
		false,
		"",
		"",
		false)).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
