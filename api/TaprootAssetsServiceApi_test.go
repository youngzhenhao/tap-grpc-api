package api

import (
	"reflect"
	"testing"
)

// PASS
func TestGetInfo(t *testing.T) {
	want := "*taprpc.GetInfoResponse"
	got := reflect.TypeOf(GetInfo()).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// PASS
func TestListAssets(t *testing.T) {
	want := "*taprpc.ListAssetResponse"
	got := reflect.TypeOf(ListAssets(true, true, false)).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// PASS
func TestListBalances(t *testing.T) {
	want := "*taprpc.ListBalancesResponse"
	got := reflect.TypeOf(ListBalances(true)).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// PASS
func TestListGroups(t *testing.T) {
	want := "*taprpc.ListGroupsResponse"
	got := reflect.TypeOf(ListGroups()).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// PASS
func TestListTransfers(t *testing.T) {
	want := "*taprpc.ListTransfersResponse"
	got := reflect.TypeOf(ListTransfers()).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// PASS
func TestListUtxos(t *testing.T) {
	want := "*taprpc.ListUtxosResponse"
	got := reflect.TypeOf(ListUtxos(true)).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestNewAddr(t *testing.T) {
	want := "*taprpc.Addr"
	got := reflect.TypeOf(NewAddr("", 0).String())
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestSendAsset(t *testing.T) {
	want := "*taprpc.SendAssetResponse"
	got := reflect.TypeOf(SendAsset([]string{""}, 0)).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
