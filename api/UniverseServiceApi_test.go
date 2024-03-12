package api

import (
	"reflect"
	"testing"
)

// PASS
func TestInfo(t *testing.T) {
	want := "*universerpc.InfoResponse"
	got := reflect.TypeOf(Info()).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// PASS
func TestListFederationServers(t *testing.T) {
	want := "*universerpc.ListFederationServersResponse"
	got := reflect.TypeOf(ListFederationServers()).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
