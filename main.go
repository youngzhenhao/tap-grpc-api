package main

import (
	"fmt"
	"tap-api/api"
)

func main() {

	//tmp()
	//fmt.Print(api.StructToJsonString(api.Info()))
	//fmt.Print(api.StructToJsonString(api.ListFederationServers()))
	//fmt.Printf("%v\n", api.StructToJsonString(api.ListBatches()))
	//fmt.Printf("\nlen:%v\n", api.StructToJsonString(len(api.ListBatches().Batches)))
	//fmt.Printf("%v\n", api.ListBatches().Batches[6].Assets[1].GetName())
	//fmt.Printf("%v\n", hex.EncodeToString(api.ListBatches().Batches[6].Assets[0].GetAssetMeta().GetData()))
	//fmt.Printf("%v\n", api.StructToJsonString(api.NextInternalKey(1)))
	//fmt.Printf("%v\n", api.StructToJsonString(api.NextScriptKey(1)))

	fmt.Printf("%v\n", api.StructToJsonString(api.NewAddr("4302729c975340536905d8fe209c5606620e2fd00c01891fcdf726c5b1ea7eee", 3)))

}

func tmp() {

}
