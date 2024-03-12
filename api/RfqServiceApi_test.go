package api

import (
	"reflect"
	"testing"
)

func TestAddAssetBuyOrder(t *testing.T) {
	want := "*rfqrpc.AddAssetBuyOrderResponse"
	got := reflect.TypeOf(AddAssetBuyOrder()).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestAddAssetSellOffer(t *testing.T) {
	want := "*rfqrpc.AddAssetSellOfferResponse"
	got := reflect.TypeOf(AddAssetSellOffer()).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

// === RUN   TestQueryRfqAcceptedQuotes
// 2024/03/08 16:20:39 rfqrpc QueryRfqAcceptedQuotes Error: rpc error: code = Unknown desc = unknown request
func TestQueryRfqAcceptedQuotes(t *testing.T) {
	want := "*rfqrpc.QueryRfqAcceptedQuotesResponse"
	got := reflect.TypeOf(QueryRfqAcceptedQuotes()).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}

func TestSubscribeRfqEventNtfns(t *testing.T) {
	want := "*rfqrpc.SubscribeRfqEventNtfnsResponse"
	got := reflect.TypeOf(SubscribeRfqEventNtfns()).String()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("expected:%v, got:%v", want, got)
	}
}
