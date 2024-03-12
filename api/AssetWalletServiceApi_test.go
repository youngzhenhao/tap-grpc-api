package api

import (
	"reflect"
	"testing"
)

func TestAnchorVirtualPsbts(t *testing.T) {
	want := "*assetwalletrpc.AnchorVirtualPsbtsResponse"
	got := reflect.TypeOf(AnchorVirtualPsbts(nil)).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestFundVirtualPsbt(t *testing.T) {
	want := "*assetwalletrpc.FundVirtualPsbtResponse"
	got := reflect.TypeOf(FundVirtualPsbt(false)).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// === RUN   TestNextInternalKey
// 2024/03/11 15:29:03 assetwalletrpc NextInternalKey Error: rpc error: code = Unknown desc = key family must be set to a non-zero value
// PASS
func TestNextInternalKey(t *testing.T) {
	want := "*assetwalletrpc.NextInternalKeyResponse"
	got := reflect.TypeOf(NextInternalKey(1)).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// PASS
func TestNextScriptKey(t *testing.T) {
	want := "*assetwalletrpc.NextScriptKeyResponse"
	got := reflect.TypeOf(NextScriptKey(1)).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// === RUN   TestQueryInternalKey
// 2024/03/08 16:20:18 assetwalletrpc QueryInternalKey Error: rpc error: code = Unknown desc = unknown request
func TestQueryInternalKey(t *testing.T) {
	want := "*assetwalletrpc.QueryInternalKeyResponse"
	got := reflect.TypeOf(QueryInternalKey("03eb1540863194fdbb08d3b99d69321b5e3e51be6425148ab305d66d441496d77f")).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// === RUN   TestProveAssetOwnership
// 2024/03/11 09:38:27 assetwalletrpc ProveAssetOwnership Error: rpc error: code = Unknown desc = invalid script key: unknown script key length: 66
func TestProveAssetOwnership(t *testing.T) {
	want := "*assetwalletrpc.ProveAssetOwnershipResponse"
	got := reflect.TypeOf(ProveAssetOwnership("e361fa40e96e0dd9c47e7367bb556d8acf1c4ea07421850c538057e5a4ebdfd6", "02daf0c9d98dc8e7c6aceeabe6c271cfc88d3733da6c03e19146eb25a10844d418")).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// === RUN   TestVerifyAssetOwnership
// 2024/03/11 09:40:37 assetwalletrpc VerifyAssetOwnership Error: rpc error: code = Unknown desc = a valid proof must be specified
func TestVerifyAssetOwnership(t *testing.T) {
	want := "*assetwalletrpc.VerifyAssetOwnershipResponse"
	got := reflect.TypeOf(VerifyAssetOwnership("").String())
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestQueryScriptKey(t *testing.T) {
	want := "*assetwalletrpc.QueryScriptKeyResponse"
	got := reflect.TypeOf(QueryScriptKey("")).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestRemoveUTXOLease(t *testing.T) {
	want := "*assetwalletrpc.RemoveUTXOLeaseResponse"
	got := reflect.TypeOf(RemoveUTXOLease()).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestSignVirtualPsbt(t *testing.T) {
	want := "*assetwalletrpc.SignVirtualPsbtResponse"
	got := reflect.TypeOf(SignVirtualPsbt("")).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
